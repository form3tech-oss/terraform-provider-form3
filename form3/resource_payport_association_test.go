package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestAccStarlingPayportAssociation_basic(t *testing.T) {
	var payportResponse associations.GetPayportIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStarlingPayportAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3StarlingPayportAssociationConfigA, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStarlingPayportAssociationExists("form3_payport_association.association", &payportResponse),
					resource.TestCheckResourceAttr(
						"form3_payport_association.association", "participant_id", "08008802"),
					resource.TestCheckResourceAttr(
						"form3_payport_association.association", "participant_id", "08008802"),
				),
			},
		},
	})
}

func testAccCheckStarlingPayportAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_payport_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetPayportID(associations.NewGetPayportIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckStarlingPayportAssociationExists(resourceKey string, association *associations.GetPayportIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetPayportID(associations.NewGetPayportIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		association = foundRecord

		return nil
	}
}

const testForm3StarlingPayportAssociationConfigA = `
resource "form3_payport_association" "association" {
	organisation_id                  = "%s"
	payport_association_id           = "6e89d652-0344-4ab4-8118-f3ac30397ee8"
	participant_id	                 = "08008802"
  customer_sending_fps_institution = "444443"
  sponsor_bank_id                  = "111113"
  sponsor_account_number           = "22222223"
}`
