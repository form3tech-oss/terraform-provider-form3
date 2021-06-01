package form3

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
)

func TestAccSepactLiquidityAssociation_basic(t *testing.T) {
	directOrgID := uuid.New().String()
indirectOrgID := uuid.New().String()
	verifyOrgDoesNotExist(t, directOrgID)
	verifyOrgDoesNotExist(t, indirectOrgID)

	createData := sepactLiquidityAssociationData{
		ParentOrganisationID: os.Getenv("FORM3_ORGANISATION_ID"),
		ReachableBIC:         generateTestBicWithLength(11),
		DirectParticipant: participant{
			Path:           "form3_sepact_liquidity_association.direct_participant_association",
			OrganisationID: directOrgID,
			AssociationID:  uuid.New().String(),
			BIC:            generateTestBicWithLength(8),
			IBAN:           generateRandomIban(),
			Name:           "Direct Participant",
		},
		IndirectParticipant: participant{
			Path:           "form3_sepact_liquidity_association.indirect_participant_association",
			OrganisationID: indirectOrgID,
			AssociationID:  uuid.New().String(),
			BIC:            generateTestBicWithLength(8),
			IBAN:           generateRandomIban(),
			Name:           "Indirect Participant",
		},
	}

	updateData := createData
	updateData.DirectParticipant.Name = "Updated Direct Participant"
	updateData.DirectParticipant.BIC = generateTestBicWithLength(8)
	updateData.DirectParticipant.IBAN = generateRandomIban()
	updateData.ReachableBIC = generateTestBicWithLength(11)
	updateData.IndirectParticipant.Name = "Updated Indirect Participant"
	updateData.IndirectParticipant.BIC = generateTestBicWithLength(8)
	updateData.IndirectParticipant.IBAN = generateRandomIban()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepactLiquidityAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSepactLiquidityAssociationConfig(createData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepactLiquidityAssociationExists(createData.DirectParticipant.Path),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "association_id", createData.DirectParticipant.AssociationID),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "organisation_id", createData.DirectParticipant.OrganisationID),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "name", createData.DirectParticipant.Name),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "settlement_bic", createData.DirectParticipant.BIC),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "settlement_iban", createData.DirectParticipant.IBAN),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "address_city", "London"),
					resource.TestCheckResourceAttr(createData.DirectParticipant.Path, "address_country", "United Kingdom"),

					testAccCheckSepactLiquidityAssociationExists(createData.IndirectParticipant.Path),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "association_id", createData.IndirectParticipant.AssociationID),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "organisation_id", createData.IndirectParticipant.OrganisationID),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "name", createData.IndirectParticipant.Name),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "reachable_bic", createData.ReachableBIC),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "settlement_bic", createData.IndirectParticipant.BIC),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "settlement_iban", createData.IndirectParticipant.IBAN),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "address_city", "London"),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "address_country", "United Kingdom"),
					resource.TestCheckResourceAttr(createData.IndirectParticipant.Path, "direct_participant_id", createData.DirectParticipant.AssociationID),
				),
			},
			{
				Config: testSepactLiquidityAssociationConfig(updateData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepactLiquidityAssociationExists(updateData.DirectParticipant.Path),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "association_id", updateData.DirectParticipant.AssociationID),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "organisation_id", updateData.DirectParticipant.OrganisationID),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "name", updateData.DirectParticipant.Name),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "settlement_bic", updateData.DirectParticipant.BIC),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "settlement_iban", updateData.DirectParticipant.IBAN),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "address_city", "London"),
					resource.TestCheckResourceAttr(updateData.DirectParticipant.Path, "address_country", "United Kingdom"),

					testAccCheckSepactLiquidityAssociationExists(updateData.IndirectParticipant.Path),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "association_id", updateData.IndirectParticipant.AssociationID),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "organisation_id", updateData.IndirectParticipant.OrganisationID),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "name", updateData.IndirectParticipant.Name),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "reachable_bic", updateData.ReachableBIC),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "settlement_bic", updateData.IndirectParticipant.BIC),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "settlement_iban", updateData.IndirectParticipant.IBAN),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "address_city", "London"),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "address_country", "United Kingdom"),
					resource.TestCheckResourceAttr(updateData.IndirectParticipant.Path, "direct_participant_id", updateData.DirectParticipant.AssociationID),
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

type participant struct{ Path, OrganisationID, AssociationID, BIC, IBAN, Name string }
type sepactLiquidityAssociationData struct {
	ParentOrganisationID                   string
	ReachableBIC                           string
	DirectParticipant, IndirectParticipant participant
}

func testSepactLiquidityAssociationConfig(data sepactLiquidityAssociationData) string {
	var sb strings.Builder

	err := template.Must(template.New("sepact-liquidity.tf").Parse(`
resource "form3_organisation" "direct_participant" {
	organisation_id        = "{{ .DirectParticipant.OrganisationID }}"
	parent_organisation_id = "{{ .ParentOrganisationID }}"
	name 		           = "terraform direct participant organisation"
}

resource "form3_organisation" "indirect_participant" {
	organisation_id        = "{{ .IndirectParticipant.OrganisationID }}"
	parent_organisation_id = "{{ .ParentOrganisationID }}"
	name 		           = "terraform indirect participant organisation"
}

resource "form3_sepact_liquidity_association" "direct_participant_association" {
	organisation_id         = "${form3_organisation.direct_participant.organisation_id}"
	association_id          = "{{ .DirectParticipant.AssociationID }}"
	name                    = "{{ .DirectParticipant.Name }}"
    settlement_bic          = "{{ .DirectParticipant.BIC }}"
    settlement_iban         = "{{ .DirectParticipant.IBAN }}"
	address_street          = "Harp Ln"
	address_building_number = "7"
	address_city            = "London"
	address_country         = "United Kingdom"
}

resource "form3_sepact_liquidity_association" "indirect_participant_association" {
	organisation_id         = "${form3_organisation.indirect_participant.organisation_id}"
	association_id          = "{{ .IndirectParticipant.AssociationID }}"
	name                    = "{{ .IndirectParticipant.Name }}"
    reachable_bic           = "{{ .ReachableBIC }}"
    settlement_bic          = "{{ .IndirectParticipant.BIC }}"
    settlement_iban         = "{{ .IndirectParticipant.IBAN }}"
	address_street          = "Harp Ln"
	address_building_number = "7"
	address_city            = "London"
	address_country         = "United Kingdom"
	direct_participant_id   = "${form3_sepact_liquidity_association.direct_participant_association.association_id}"
}
`)).Execute(&sb, data)
	if err != nil {
		log.Fatal(err)
	}

	return sb.String()
}
