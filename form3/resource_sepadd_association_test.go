package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSepaDDAssociation_basic(t *testing.T) {
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	associationID := uuid.New().String()
	bic := generateTestBic()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaDDAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3SepaDDAssociationConfig(organisationID, parentOrganisationID, testOrgName, associationID, bic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaDDAssociationExists("form3_sepadd_association.association"),
					resource.TestCheckResourceAttr("form3_sepadd_association.association", "association_id", associationID),
					resource.TestCheckResourceAttr("form3_sepadd_association.association", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_sepadd_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepadd_association.association", "business_user", "PR344567"),
					resource.TestCheckResourceAttr("form3_sepadd_association.association", "receiver_business_user", "PR344568"),
					resource.TestCheckResourceAttr("form3_sepadd_association.association", "local_instrument", "CORE"),
					resource.TestCheckResourceAttr("form3_sepadd_association.association", "allow_submissions", "true"),
				),
			},
		},
	})
}

func testAccCheckSepaDDAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_sepadd_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSepaddID(associations.NewGetSepaddIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("sepa dd record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSepaDDAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no sepadd Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepaddID(associations.NewGetSepaddIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("sepa dd record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

func getTestForm3SepaDDAssociationConfig(orgID, parOrgID, orgName, assocID, bic string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		           = "%s"
	}

	resource "form3_sepadd_association" "association" {
		organisation_id        = "${form3_organisation.organisation.organisation_id}"
		association_id         = "%s"
		bic                    = "%s"
		business_user          = "PR344567"
		receiver_business_user = "PR344568"
		local_instrument       = "CORE"
		allow_submissions      = true
	}`, orgID, parOrgID, orgName, assocID, bic)
}
