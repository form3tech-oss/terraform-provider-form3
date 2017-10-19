package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/roles"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestAccRole_basic(t *testing.T) {
	var roleResponse roles.GetRolesRoleIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3RoleConfigA, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleExists("form3_role.role", &roleResponse),
					resource.TestCheckResourceAttr(
						"form3_role.role", "name", "terraform-role"),
				),
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

const testForm3RoleConfigA = `
resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "933a03db-08fa-42af-8e4b-9a088b178b3a"
	name     		= "terraform-role"
}`
