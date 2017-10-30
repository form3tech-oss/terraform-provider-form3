package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/accounts"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestAccBic_basic(t *testing.T) {
	var bicResponse accounts.GetBicsIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBicDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BicConfigA, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBicExists("form3_bic.bic", &bicResponse),
					resource.TestCheckResourceAttr(
						"form3_bic.bic", "bic", "NWBKGB09"),
				),
			},
		},
	})
}

func testAccCheckBicDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_bic" {
			continue
		}

		response, err := client.AccountClient.Accounts.GetBicsID(accounts.NewGetBicsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckBicExists(resourceKey string, bicResponse *accounts.GetBicsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AccountClient.Accounts.GetBicsID(accounts.NewGetBicsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		bicResponse = foundRecord

		return nil
	}
}

const testForm3BicConfigA = `
resource "form3_bic" "bic" {
	organisation_id = "%s"
  bic_id          = "8fe21b3a-f113-4318-abf1-2201882a5ee8"
	bic       	    = "NWBKGB09"
}
`
