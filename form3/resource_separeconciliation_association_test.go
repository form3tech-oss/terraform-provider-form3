package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"os"
	"testing"
)

func TestAccSepaReconciliationAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	sponsoredOrganisationId := uuid.NewV4().String()
	associationId := uuid.NewV4().String()
	sponsoredAssociationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaReconciliationAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SepaReconciliationAssociationConfigA, organisationId, parentOrganisationId, associationId, sponsoredOrganisationId, sponsoredAssociationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaReconciliationAssociationExists("form3_separeconciliation_association.association"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association", "bic", "TESTBIC8"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association", "sponsor_id", ""),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association", "name", "john"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association", "country", "DE"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association", "iban", "fakeiban"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association_sponsored", "association_id", sponsoredAssociationId),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association_sponsored", "organisation_id", sponsoredOrganisationId),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association_sponsored", "bic", "TESTBICCXXX"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association_sponsored", "name", "doe"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association_sponsored", "country", "FR"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association_sponsored", "iban", "fakeiban"),
					resource.TestCheckResourceAttr("form3_separeconciliation_association.association_sponsored", "sponsor_id", associationId),
				),
			},
		},
	})
}

func testAccCheckSepaReconciliationAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_separeconciliation_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSepareconciliationID(associations.NewGetSepareconciliationIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("sepa reconciliation record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSepaReconciliationAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepareconciliationID(associations.NewGetSepareconciliationIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("sepa reconciliation record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3SepaReconciliationAssociationConfigA = `
locals {
  organisation_id          = "%s"
  parent_organisation_id   = "%s"
  association_id           = "%s"
  organisation_sponsor_id  = "%s"
  sponsored_association_id = "%s"
}

resource "form3_organisation" "organisation" {
  organisation_id        = "${local.organisation_id}"
  parent_organisation_id = "${local.parent_organisation_id}"
  name 		           = "terraform-organisation"
}

resource "form3_separeconciliation_association" "association" {
  organisation_id      = "${form3_organisation.organisation.organisation_id}"
  association_id       = "${local.association_id}"
  bic                  = "TESTBIC8"
  name                 = "john"
  country              = "DE"
  iban                 = "fakeiban"
}

resource "form3_organisation" "organisation_sponsored" {
  organisation_id        = "${local.organisation_sponsor_id}"
  parent_organisation_id = "${local.parent_organisation_id}"
  name 		           = "terraform-organisation"
}

resource "form3_separeconciliation_association" "association_sponsored" {
  organisation_id      = "${form3_organisation.organisation_sponsored.organisation_id}"
  association_id       = "${local.sponsored_association_id}"
  bic                  = "TESTBICCXXX"
  name                 = "doe"
  country              = "FR"
  iban                 = "fakeiban"
  sponsor_id           = "${form3_separeconciliation_association.association.association_id}"
}
`
