package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSepaReconciliationAssociation_basic(t *testing.T) {
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	sponsorOrganisationID := uuid.New().String()
	sponsoredOrganisationID := uuid.New().String()
	sponsorAssociationID := uuid.New().String()
	sponsoredAssociationID := uuid.New().String()

	sponsor_assoc_path := "form3_separeconciliation_association.sponsor_association"
	sponsored_assoc_path := "form3_separeconciliation_association.sponsored_association"

	sponsorAssociationTechnicalBic := generateTestBic()
	sponsorAssociationReconciliationBic := generateTestBic()
	sponsorAssociationReconciliationIban := generateRandomIban()
	sponsoredAssociationTechnicalBic := generateTestBic()
	sponsoredAssociationReconciliationBic := generateTestBic()
	sponsoredAssociationReconciliationIban := generateRandomIban()

	config := fmt.Sprintf(
		testForm3SepaReconciliationAssociationConfigA,
		parentOrganisationID,
		sponsorOrganisationID,
		sponsoredOrganisationID,
		sponsorAssociationID,
		sponsoredAssociationID,
		sponsorAssociationTechnicalBic,
		sponsorAssociationReconciliationBic,
		sponsorAssociationReconciliationIban,
		sponsoredAssociationTechnicalBic,
		sponsoredAssociationReconciliationBic,
		sponsoredAssociationReconciliationIban,
	)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaReconciliationAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaReconciliationAssociationExists(sponsor_assoc_path),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "association_id", sponsorAssociationID),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "organisation_id", sponsorOrganisationID),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "name", "Sponsor company"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "technical_bic", sponsorAssociationTechnicalBic),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "reconciliation_bic", sponsorAssociationReconciliationBic),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "reconciliation_iban", sponsorAssociationReconciliationIban),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "address_city", "London"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "address_country", "United Kingdom"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "sponsor_id", ""),

					testAccCheckSepaReconciliationAssociationExists(sponsored_assoc_path),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "association_id", sponsoredAssociationID),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "organisation_id", sponsoredOrganisationID),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "name", "Sponsored company"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "technical_bic", sponsoredAssociationTechnicalBic),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "reconciliation_bic", sponsoredAssociationReconciliationBic),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "reconciliation_iban", sponsoredAssociationReconciliationIban),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "address_city", "London"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "address_country", "United Kingdom"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "sponsor_id", sponsorAssociationID),
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

		response, err := client.AssociationClient.Associations.GetSepaReconciliationID(associations.NewGetSepaReconciliationIDParams().
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
			return fmt.Errorf("no sepa reconciliation Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepaReconciliationID(associations.NewGetSepaReconciliationIDParams().
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
	parent_organisation_id    					= "%s"

	sponsor_organisation_id   					= "%s"
	sponsored_organisation_id 					= "%s"

	sponsor_association_id    					= "%s"
	sponsored_association_id  					= "%s"

	sponsor_association_technical_bic 			= "%s"
	sponsor_association_reconciliation_bic 		= "%s"
	sponsor_association_reconciliation_iban 		= "%s"

	sponsored_association_technical_bic 			= "%s"
	sponsored_association_reconciliation_bic 	= "%s"
	sponsored_association_reconciliation_iban 	= "%s"
}

resource "form3_organisation" "sponsor" {
	organisation_id        = "${local.sponsor_organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		               = "terraform-sponsor-organisation"
}

resource "form3_organisation" "sponsored" {
	organisation_id        = "${local.sponsored_organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		               = "terraform-sponsored-organisation"
}

resource "form3_separeconciliation_association" "sponsor_association" {
	organisation_id         = "${form3_organisation.sponsor.organisation_id}"
	association_id          = "${local.sponsor_association_id}"
	name                    = "Sponsor company"
	technical_bic           = "${local.sponsor_association_technical_bic}"
	reconciliation_bic      = "${local.sponsor_association_reconciliation_bic}"
	reconciliation_iban     = "${local.sponsor_association_reconciliation_iban}"
	address_street          = "Harp Ln"
	address_building_number = "7"
	address_city            = "London"
	address_country         = "United Kingdom"
}

resource "form3_separeconciliation_association" "sponsored_association" {
	organisation_id         = "${form3_organisation.sponsored.organisation_id}"
	association_id          = "${local.sponsored_association_id}"
	name                    = "Sponsored company"
	technical_bic           = "${local.sponsored_association_technical_bic}"
	reconciliation_bic      = "${local.sponsored_association_reconciliation_bic}"
	reconciliation_iban     = "${local.sponsored_association_reconciliation_iban}"
	address_street          = "Harp Ln"
	address_building_number = "7"
	address_city            = "London"
	address_country         = "United Kingdom"
	sponsor_id              = "${form3_separeconciliation_association.sponsor_association.association_id}"
}
`
