package form3

import (
	"fmt"
	"log"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"math/rand"
)

func TestAccCredentialSso_single(t *testing.T) {
	log.SetOutput(os.Stdout)
	organisationID := os.Getenv("FORM3_ORGANISATION_ID")
	ssoUserId := newSsoUserId()
	userID := uuid.New().String()
	roleID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCredentialSsoDestroy,
		Steps: []resource.TestStep{
			// Given I have created an user
			// When i add sso credential to this user
			// Then I can see credential was added
			{
				Config: fmt.Sprintf(testForm3CredentialSsoConfigSingle, organisationID, roleID, organisationID, userID, ssoUserId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_single", "user_id", userID),
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_single", "sso_user_id", ssoUserId),
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_single", "organisation_id", organisationID),
				),
			},
		},
	})
}

func TestAccCredentialSso_multiple_for_same_user(t *testing.T) {
	log.SetOutput(os.Stdout)
	organisationID := os.Getenv("FORM3_ORGANISATION_ID")
	ssoUserId1 := newSsoUserId()
	ssoUserId2 := newSsoUserId()
	userID := uuid.New().String()
	roleID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCredentialSsoDestroy,
		Steps: []resource.TestStep{
			// Given I have created an user
			// When i add a sso credential to this user for ssoUserId1
			// When i add a second sso credential to this user for ssoUserId2
			// Then I can see both credentials added
			{
				Config: fmt.Sprintf(testForm3CredentialSsoConfigMulti, organisationID, roleID, organisationID, userID, ssoUserId1, ssoUserId2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_multi_one", "user_id", userID),
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_multi_one", "sso_user_id", ssoUserId1),
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_multi_one", "organisation_id", organisationID),

					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_multi_two", "user_id", userID),
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_multi_two", "sso_user_id", ssoUserId2),
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_multi_two", "organisation_id", organisationID),
				),
			},
		},
	})
}

func TestAccCredentialSso_preexisting_state(t *testing.T) {
	log.SetOutput(os.Stdout)
	organisationID := os.Getenv("FORM3_ORGANISATION_ID")
	ssoUserId := newSsoUserId()
	userID := uuid.New().String()
	roleID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3UserOnly, organisationID, roleID, organisationID, userID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("form3_user.sso_test_user", "user_id", userID),
				),
			},
			{
				PreConfig: func() {
					createPreexistingSsoUserManuallyInSecurityApi(t, strfmt.UUID(userID), strfmt.UUID(organisationID), ssoUserId)
				},
				Config: fmt.Sprintf(testForm3CredentialSsoConfigSingle, organisationID, roleID, organisationID, userID, ssoUserId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_single", "user_id", userID),
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_single", "sso_user_id", ssoUserId),
					resource.TestCheckResourceAttr("form3_credential_sso.test_sso_single", "organisation_id", organisationID),
				),
			},
		},
	})
}

func createPreexistingSsoUserManuallyInSecurityApi(t *testing.T, userID strfmt.UUID, organisationID strfmt.UUID, ssoUserId string) {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	params := users.NewPostUsersUserIDCredentialsSsoParams().WithUserID(userID).
		WithSsoUserCreation(&models.SsoUserCreation{Data: &models.SsoUser{
			Attributes: &models.SsoUserAttributes{
				SsoID:  ssoUserId,
				UserID: userID,
			},
			ID:             form3.NewUUIDValue(),
			OrganisationID: organisationID,
		}})
	_, err := client.SecurityClient.Users.PostUsersUserIDCredentialsSso(params)
	if err != nil {
		t.Fatalf("failed to create sso user: %v", err)
	}
}

func testAccCheckCredentialSsoDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_credential_sso" {
			continue
		}

		response, err := client.SecurityClient.Users.GetUsersUserIDCredentialsSsoSsoUserID(users.NewGetUsersUserIDCredentialsSsoSsoUserIDParams().
			WithUserID(strfmt.UUID(rs.Primary.Attributes["user_id"])).WithSsoUserID(rs.Primary.Attributes["sso_user_id"]))

		if err == nil {
			return fmt.Errorf("sso credential %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func newSsoUserId() string {
	return fmt.Sprintf("testSsoUserId_%d", rand.Intn(100))
}

const testForm3CredentialSsoConfigSingle = `
resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "%s"
	name     		= "terraform-role"
}

resource "form3_user" "sso_test_user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 	    = "terraform-user"
	email 			= "terraform-user@form3.tech"
	roles 			= ["${form3_role.role.role_id}"]
}

resource "form3_credential_sso" "test_sso_single" {
	user_id 		= "${form3_user.sso_test_user.user_id}"
	organisation_id = "${form3_user.sso_test_user.organisation_id}"
	sso_user_id    = "%s"
}`

const testForm3UserOnly = `
resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "%s"
	name     		= "terraform-role"
}

resource "form3_user" "sso_test_user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 	    = "terraform-user"
	email 			= "terraform-user@form3.tech"
	roles 			= ["${form3_role.role.role_id}"]
}`

const testForm3CredentialSsoConfigMulti = `
resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "%s"
	name     		= "terraform-role"
}
resource "form3_user" "sso_test_user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 	    = "terraform-user"
	email 			= "terraform-user@form3.tech"
	roles 			= ["${form3_role.role.role_id}"]
}

resource "form3_credential_sso" "test_sso_multi_one" {
	user_id 		= "${form3_user.sso_test_user.user_id}"
	organisation_id = "${form3_user.sso_test_user.organisation_id}"
	sso_user_id   = "%s"
}

resource "form3_credential_sso" "test_sso_multi_two" {
	user_id 		= "${form3_user.sso_test_user.user_id}"
	organisation_id = "${form3_user.sso_test_user.organisation_id}"
	sso_user_id   = "%s"

depends_on = [ "form3_credential_sso.test_sso_multi_one" ]
}`
