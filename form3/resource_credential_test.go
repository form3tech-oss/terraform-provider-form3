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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCredential_basic(t *testing.T) {
	var credentialResponse models.Credential
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	roleID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCredentialDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3CredentialConfigA, organisationId, roleID, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialExists("form3_credential.credential", &credentialResponse)),
			},
		},
	})
}

func testAccCheckCredentialDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_credential" {
			continue
		}

		_, err := client.SecurityClient.Users.GetUsersUserIDCredentials(users.NewGetUsersUserIDCredentialsParams().
			WithUserID(strfmt.UUID(strfmt.UUID(rs.Primary.Attributes["user_id"]))))

		if err == nil {
			return fmt.Errorf("error listing credentials: %s", form3.JsonErrorPrettyPrint(err))
		}
	}

	return nil
}

func testAccCheckCredentialExists(resourceKey string, credential *models.Credential) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		log.Printf("[INFO] Checking that credential with client id: %s exists", rs.Primary.ID)
		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		credentialList, err := client.SecurityClient.Users.GetUsersUserIDCredentials(users.NewGetUsersUserIDCredentialsParams().
			WithUserID(strfmt.UUID(rs.Primary.Attributes["user_id"])))

		if err != nil {
			return err
		}

		found := false
		for _, element := range credentialList.Payload.Data {
			log.Printf("[DEBUG] Checking that credential with client id: %s exists comparing with %s", rs.Primary.ID, element.ClientID.String())
			if element.ClientID.String() == rs.Primary.ID {
				credential = element
				found = true
				if len(credential.ClientID.String()) != 36 {
					return fmt.Errorf("expected credential %s to have a length of 36 got %d", credential.ClientID.String(),
						len(credential.ClientID.String()))
				}
			}
		}

		if !found {
			return fmt.Errorf("credential record not found")
		}

		return nil
	}
}

const testForm3CredentialConfigA = `
resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "%s"
	name     		= "terraform-role"
}
resource "form3_credential" "credential" {
	user_id 		= "${form3_user.user.user_id}"
}
resource "form3_user" "user" {
	organisation_id = "%s"
	user_id 		= "${uuid()}"
	user_name 		= "terraform-user"
	email 			= "terraform-user@form3.tech"
	roles 			= ["${form3_role.role.role_id}"]
  	lifecycle {
		ignore_changes = ["user_id"]
  	}
}`
