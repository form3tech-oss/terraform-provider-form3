package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccUser_basic(t *testing.T) {
	var userResponse users.GetUsersUserIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	userId := uuid.New().String()
	roleId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3UserConfigA, organisationId, roleId, organisationId, userId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("form3_user.user", &userResponse),
					resource.TestCheckResourceAttr(
						"form3_user.user", "user_name", "terraform-user"),
					resource.TestCheckResourceAttr(
						"form3_user.user", "email", "terraform-user@form3.tech"),
					resource.TestCheckResourceAttr(
						"form3_user.user", "roles.0", roleId),
				),
			},
			{
				Config: fmt.Sprintf(testForm3UserConfigAUpdate, organisationId, roleId, organisationId, userId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("form3_user.user", &userResponse),
					resource.TestCheckResourceAttr(
						"form3_user.user", "email", "dude@form3.tech"),
					resource.TestCheckResourceAttr(
						"form3_user.user", "roles.0", roleId),
				),
			},
		},
	})
}

func TestAccUser_importBasic(t *testing.T) {

	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	userId := uuid.New().String()
	roleId := uuid.New().String()

	resourceName := "form3_user.user"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3UserConfigA, organisationId, roleId, organisationId, userId),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckUserDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_user" {
			continue
		}

		response, err := client.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckUserExists(resourceKey string, user *users.GetUsersUserIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		user = foundRecord

		return nil
	}
}

const testForm3UserConfigA = `
resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "%s"
	name     		= "terraform-role"
}

resource "form3_user" "user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 	= "terraform-user"
 	email 			= "terraform-user@form3.tech"
	roles 			= ["${form3_role.role.role_id}"]
}`

const testForm3UserConfigAUpdate = `
resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "%s"
	name     		= "terraform-role"
}

resource "form3_user" "user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 	= "terraform-user"
  email			  = "dude@form3.tech"
	roles 			= ["${form3_role.role.role_id}"]
}`
