package form3

import (
	"fmt"
	"os"
	"strings"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSepaInstantAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	sponsoredOrganisationId := uuid.New().String()
	associationId := uuid.New().String()
	sponsoredAssociationId := uuid.New().String()

	bic := generateTestBic()
	bicCN := strings.ToLower(bic)
	bicSponsored := generateTestBic()
	businessUserDN := fmt.Sprintf("cn=%s,ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu", bicCN)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaInstantAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SepaInstantAssociationConfigA, organisationId, parentOrganisationId, associationId, sponsoredOrganisationId, sponsoredAssociationId, bicCN, bic, bicSponsored),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "business_user_dn", businessUserDN),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "transport_profile_id", "TEST_PROFILE_1"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "simulator_only", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "disable_outbound_payments", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "sponsor_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "association_id", sponsoredAssociationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "organisation_id", sponsoredOrganisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "business_user_dn", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "transport_profile_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "bic", bicSponsored),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "sponsor_id", associationId),
				),
			},
			{
				Config: fmt.Sprintf(testForm3SepaInstantAssociationUpdatedConfig, organisationId, parentOrganisationId, associationId, bicCN, bic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "business_user_dn", businessUserDN),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "transport_profile_id", "TEST_PROFILE_1"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "simulator_only", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "disable_outbound_payments", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "sponsor_id", ""),
				),
			},
		},
	})
}

func testAccCheckSepaInstantAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_sepainstant_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("sepa instant record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSepaInstantAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no bacs Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("sepa instant record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3SepaInstantAssociationConfigA = `
locals {
	organisation_id          = "%s"
	parent_organisation_id   = "%s"
	association_id           = "%s"
	organisation_sponsor_id  = "%s"
	sponsored_association_id = "%s"
	bic_cn					 = "%s"
	bic					     = "%s"
	bic_sponsored			 = "%s"
}

resource "form3_organisation" "organisation" {
	organisation_id        = "${local.organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_sepainstant_association" "association" {
	organisation_id      = "${form3_organisation.organisation.organisation_id}"
	association_id       = "${local.association_id}"
  	business_user_dn     = "cn=${local.bic_cn},ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu"
  	transport_profile_id = "TEST_PROFILE_1"
	bic                  = "${local.bic}"
	simulator_only       = true
}

resource "form3_organisation" "organisation_sponsored" {
	organisation_id        = "${local.organisation_sponsor_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_sepainstant_association" "association_sponsored" {
  organisation_id      = "${form3_organisation.organisation_sponsored.organisation_id}"
  association_id       = "${local.sponsored_association_id}"
  business_user_dn     = ""
  transport_profile_id = ""
  bic                  = "${local.bic_sponsored}"
  simulator_only       = true
  sponsor_id           = "${form3_sepainstant_association.association.association_id}"

  depends_on = [
    "form3_sepainstant_association.association"
  ]
}
`

const testForm3SepaInstantAssociationUpdatedConfig = `
locals {
	organisation_id          = "%s"
	parent_organisation_id   = "%s"
	association_id           = "%s"
	bic_cn					 = "%s"
	bic						 = "%s"
}

resource "form3_organisation" "organisation" {
	organisation_id        = "${local.organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_sepainstant_association" "association" {
	organisation_id           = "${form3_organisation.organisation.organisation_id}"
	association_id            = "${local.association_id}"
  	business_user_dn          = "cn=${local.bic_cn},ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu"
  	transport_profile_id      = "TEST_PROFILE_1"
	bic                       = "${local.bic}"
	simulator_only            = true
	disable_outbound_payments = true
}
`
