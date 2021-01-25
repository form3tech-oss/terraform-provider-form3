package form3

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
)

func TestAccSepactLiquidityAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	type participant struct{ path, organisationId, associationId, bic, iban string }
	directParticipant := participant{
		path:           "form3_sepact_liquidity_association.direct_participant_association",
		organisationId: uuid.New().String(),
		associationId:  uuid.New().String(),
		bic:            generateTestBicWithLength(8),
		iban:           generateRandomIban(),
	}
	indirectParticipant := participant{
		path:           "form3_sepact_liquidity_association.indirect_participant_association",
		organisationId: uuid.New().String(),
		associationId:  uuid.New().String(),
		bic:            generateTestBicWithLength(8),
		iban:           generateRandomIban(),
	}
	reachableBic := generateTestBicWithLength(11)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepactLiquidityAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SepactLiquidityAssociationConfig,
					parentOrganisationId,
					directParticipant.organisationId,
					indirectParticipant.organisationId,
					directParticipant.associationId,
					indirectParticipant.associationId,
					directParticipant.bic,
					indirectParticipant.bic,
					directParticipant.iban,
					indirectParticipant.iban,
					reachableBic,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepactLiquidityAssociationExists(directParticipant.path),
					resource.TestCheckResourceAttr(directParticipant.path, "association_id", directParticipant.associationId),
					resource.TestCheckResourceAttr(directParticipant.path, "organisation_id", directParticipant.organisationId),
					resource.TestCheckResourceAttr(directParticipant.path, "name", "Direct Participant"),
					resource.TestCheckResourceAttr(directParticipant.path, "settlement_bic", directParticipant.bic),
					resource.TestCheckResourceAttr(directParticipant.path, "settlement_iban", directParticipant.iban),
					resource.TestCheckResourceAttr(directParticipant.path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(directParticipant.path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(directParticipant.path, "address_city", "London"),
					resource.TestCheckResourceAttr(directParticipant.path, "address_country", "United Kingdom"),

					testAccCheckSepactLiquidityAssociationExists(indirectParticipant.path),
					resource.TestCheckResourceAttr(indirectParticipant.path, "association_id", indirectParticipant.associationId),
					resource.TestCheckResourceAttr(indirectParticipant.path, "organisation_id", indirectParticipant.organisationId),
					resource.TestCheckResourceAttr(indirectParticipant.path, "name", "Indirect Participant"),
					resource.TestCheckResourceAttr(indirectParticipant.path, "reachable_bic", reachableBic),
					resource.TestCheckResourceAttr(indirectParticipant.path, "settlement_bic", indirectParticipant.bic),
					resource.TestCheckResourceAttr(indirectParticipant.path, "settlement_iban", indirectParticipant.iban),
					resource.TestCheckResourceAttr(indirectParticipant.path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(indirectParticipant.path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(indirectParticipant.path, "address_city", "London"),
					resource.TestCheckResourceAttr(indirectParticipant.path, "address_country", "United Kingdom"),
					resource.TestCheckResourceAttr(indirectParticipant.path, "direct_participant_id", directParticipant.associationId),
				),
			},
		},
	})
}

func testAccCheckSepactLiquidityAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_sepact_liquidity_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSepactLiquidityAssociationID(
			associations.NewGetSepactLiquidityAssociationIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.ID)),
		)

		if err == nil {
			return fmt.Errorf("sepact-liquidity association `%s` still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSepactLiquidityAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no sepact-liquidity association ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepactLiquidityAssociationID(
			associations.NewGetSepactLiquidityAssociationIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.ID)),
		)
		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("sepact-liquidity association not found, expected `%s` but found `%s`", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3SepactLiquidityAssociationConfig = `
locals {
    parent_organisation_id               = "%s"
	direct_participant_organisation_id   = "%s"
	indirect_participant_organisation_id = "%s"
	direct_participant_association_id    = "%s"
	indirect_participant_association_id  = "%s"
    direct_participant_bic               = "%s"
    indirect_participant_bic             = "%s"
    direct_participant_iban              = "%s"
    indirect_participant_iban            = "%s"
	reachable_bic                        = "%s"
}

resource "form3_organisation" "direct_participant" {
	organisation_id        = "${local.direct_participant_organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "terraform direct participant organisation"
}

resource "form3_organisation" "indirect_participant" {
	organisation_id        = "${local.indirect_participant_organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "terraform indirect participant organisation"
}

resource "form3_sepact_liquidity_association" "direct_participant_association" {
	organisation_id         = "${form3_organisation.direct_participant.organisation_id}"
	association_id          = "${local.direct_participant_association_id}"
	name                    = "Direct Participant"
    settlement_bic          = "${local.direct_participant_bic}"
    settlement_iban         = "${local.direct_participant_iban}"
	address_street          = "Harp Ln"
	address_building_number = "7"
	address_city            = "London"
	address_country         = "United Kingdom"
}

resource "form3_sepact_liquidity_association" "indirect_participant_association" {
	organisation_id         = "${form3_organisation.indirect_participant.organisation_id}"
	association_id          = "${local.indirect_participant_association_id}"
	name                    = "Indirect Participant"
    reachable_bic           = "${local.reachable_bic}"
    settlement_bic          = "${local.indirect_participant_bic}"
    settlement_iban         = "${local.indirect_participant_iban}"
	address_street          = "Harp Ln"
	address_building_number = "7"
	address_city            = "London"
	address_country         = "United Kingdom"
	direct_participant_id   = "${form3_sepact_liquidity_association.direct_participant_association.association_id}"
}
`
