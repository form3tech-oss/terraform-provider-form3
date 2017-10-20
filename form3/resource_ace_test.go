package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/ace"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestAccAce_basic(t *testing.T) {
	var aceResponse ace.GetRolesRoleIDAcesAceIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAceDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3AceConfigA, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAceExists("form3_ace.ace", &aceResponse),
					resource.TestCheckResourceAttr(
						"form3_ace.ace", "role_id", "f7bd2404-275a-4b95-8203-b5e8f027bfaf"),
					resource.TestCheckResourceAttr(
						"form3_ace.ace", "action", "CREATE"),
					resource.TestCheckResourceAttr(
						"form3_ace.ace", "record_type", "account"),
				),
			},
		},
	})
}

func testAccCheckAceDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_ace" {
			continue
		}

		response, err := client.SecurityClient.Ace.GetRolesRoleIDAcesAceID(ace.NewGetRolesRoleIDAcesAceIDParams().
			WithAceID(strfmt.UUID(rs.Primary.ID)).
			WithRoleID(strfmt.UUID(rs.Primary.Attributes["role_id"])))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckAceExists(resourceKey string, aceResponse *ace.GetRolesRoleIDAcesAceIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.SecurityClient.Ace.GetRolesRoleIDAcesAceID(ace.NewGetRolesRoleIDAcesAceIDParams().
			WithAceID(strfmt.UUID(rs.Primary.ID)).
			WithRoleID(strfmt.UUID(rs.Primary.Attributes["role_id"])))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		aceResponse = foundRecord

		return nil
	}
}

const testForm3AceConfigA = `
resource "form3_ace" "ace" {
	organisation_id = "${form3_role.role.organisation_id}"
  ace_id      = "bfcd20ce-bf9a-4f1f-ad2e-637b2f568c8e"
	role_id 		= "${form3_role.role.role_id}"
	action   		= "CREATE"
  record_type = "account"
}

resource "form3_role" "role" {
	organisation_id = "%s"
	role_id 		= "f7bd2404-275a-4b95-8203-b5e8f027bfaf"
	name     		= "terraform-role"
}
`
