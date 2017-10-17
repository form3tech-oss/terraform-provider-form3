package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/users"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestAccUser_basic(t *testing.T) {
	var userResponse users.GetUsersUserIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3UserConfigA, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("form3_user.user", &userResponse),
					resource.TestCheckResourceAttr(
						"form3_user.user", "user_name", "terraform-user"),
					resource.TestCheckResourceAttr(
						"form3_user.user", "email", "terraform-user@form3.tech"),
				),
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

		response, err := client.ApiClients.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(strfmt.UUID(rs.Primary.ID)))

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

		foundRecord, err := client.ApiClients.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(strfmt.UUID(rs.Primary.ID)))

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
resource "form3_user" "user" {
	organisation_id = "%s"
	user_id = "44247ebb-fe01-44ab-887d-7f344481712f"
	user_name = "terraform-user"
  	email = "terraform-user@form3.tech"
}`
