package form3

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	uuid "github.com/satori/go.uuid"
)

func TestAccSepaSctAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	associationId := uuid.NewV4().String()
	bic := fmt.Sprintf("TESTBIC%d", randomNumber(1000, 9999))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaSctAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SepaSctAssociationConfigA, organisationId, parentOrganisationId, associationId, bic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaSctAssociationExists("form3_sepasct_association.association"),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "business_user", "PR344567"),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "receiver_business_user", "PR344568"),
				),
			},
		},
	})
}

func testAccCheckSepaSctAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_sepasct_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSepasctID(associations.NewGetSepasctIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("sepa sct record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSepaSctAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no sepasct Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepasctID(associations.NewGetSepasctIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("sepa sct record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

func randomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

const testForm3SepaSctAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_sepasct_association" "association" {
	organisation_id        = "${form3_organisation.organisation.organisation_id}"
	association_id         = "%s"
	bic                    = "%s"
  business_user          = "PR344567"
  receiver_business_user = "PR344568"
}`
