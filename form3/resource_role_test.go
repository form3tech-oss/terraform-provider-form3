package form3

import (
	"fmt"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"

	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/roles"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccRole_basic(t *testing.T) {
	var roleResponse roles.GetRolesRoleIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	roleId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3RoleConfigA, organisationId, roleId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleExists("form3_role.role", &roleResponse),
					resource.TestCheckResourceAttr(
						"form3_role.role", "name", "terraform-role"),
				),
			},
		},
	})
}

func TestAccRole_destroyWithoutUser(t *testing.T) {
	var userResponse users.GetUsersUserIDOK

	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	roleId := uuid.New().String()
	userId := uuid.New().String()

	const configA = `
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
}
`

	const configB = `
resource "form3_user" "user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 		= "terraform-user"
 	email 			= "terraform-user@form3.tech"
}`

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			if err := testAccCheckRoleDestroy(state); err != nil {
				return err
			}
			return testAccCheckUserDestroy(state)
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(configA, organisationId, roleId, organisationId, userId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("form3_user.user", &userResponse),
					resource.TestCheckResourceAttr(
						"form3_user.user", "email", "terraform-user@form3.tech"),
					resource.TestCheckResourceAttr(
						"form3_user.user", "roles.0", roleId),
				),
			},
			{
				Config: fmt.Sprintf(configB, organisationId, userId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("form3_user.user", &userResponse),
					resource.TestCheckResourceAttr(
						"form3_user.user", "email", "terraform-user@form3.tech"),
					testAccCheckRoleDestroy,
					resource.TestCheckResourceAttr(
						"form3_user.user", "roles", "[]"),
				),
			},
		},
	})
}

func TestAccRole_importBasic(t *testing.T) {

	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	roleId := uuid.New().String()

	resourceName := "form3_role.role"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRoleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testForm3RoleConfigA, organisationId, roleId),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckRoleDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_role" {
			continue
		}

		response, err := client.SecurityClient.Roles.GetRolesRoleID(roles.NewGetRolesRoleIDParams().WithRoleID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckRoleExists(resourceKey string, role *roles.GetRolesRoleIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.SecurityClient.Roles.GetRolesRoleID(roles.NewGetRolesRoleIDParams().WithRoleID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		role = foundRecord

		return nil
	}
}

func testAccCheckRoleDoesntExist(resourceKey string, roleId uuid.UUID) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.SecurityClient.Roles.GetRolesRoleID(roles.NewGetRolesRoleIDParams().WithRoleID(strfmt.UUID(roleId.String())))
		if err != nil {
			return err
		}
		if foundRecord != nil {
			return fmt.Errorf("record was not deleted")
		}
		return nil
	}
}

const testForm3RoleConfigA = `
resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "%s"
	name     		= "terraform-role"
}`
