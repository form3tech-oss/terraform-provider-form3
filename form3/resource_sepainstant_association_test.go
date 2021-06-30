package form3

import (
	"fmt"
	"os"
	"strconv"
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
	defer verifyOrgDoesNotExist(t, organisationId)
	defer verifyOrgDoesNotExist(t, sponsoredOrganisationId)

	bic := generateTestBic()
	bic2 := generateTestBic()
	bicSponsored := generateTestBic()
	reachableBics := []string{generateTestBicWithLength(11)}
	businessUserDN := fmt.Sprintf("cn=%s,ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu", strings.ToLower(bic))
	businessUserDN2 := fmt.Sprintf("cn=%s,ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu", strings.ToLower(bic2))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaInstantAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SepaInstantAssociationConfigA,
					organisationId, parentOrganisationId, testOrgName, associationId,
					sponsoredOrganisationId, sponsoredAssociationId,
					bic, bicSponsored, strings.Join(reachableBics, "\",\"")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "business_user_dn", businessUserDN),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "transport_profile_id", "TEST_PROFILE_1"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "simulator_only", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "disable_outbound_payments", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_admission_decision", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_check", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "clearing_system", "auto"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "sponsor_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "association_id", sponsoredAssociationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "organisation_id", sponsoredOrganisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "business_user_dn", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "transport_profile_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "bic", bicSponsored),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "sponsor_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "reachable_bics.#", strconv.Itoa(len(reachableBics))),
				),
			},
			{
				Config: fmt.Sprintf(testForm3SepaInstantAssociationUpdatedConfig,
					organisationId, parentOrganisationId, testOrgName, associationId,
					sponsoredOrganisationId, sponsoredAssociationId,
					bic2, bicSponsored, strings.Join([]string{generateTestBicWithLength(11), generateTestBicWithLength(11)}, "\",\"")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "business_user_dn", businessUserDN2),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "transport_profile_id", "TEST_PROFILE_2"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "bic", bic2),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "simulator_only", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "disable_outbound_payments", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_admission_decision", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_check", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "clearing_system", "tips"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "sponsor_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "reachable_bics.#", "2"),
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
	organisation_name 		 = "%s"
	association_id           = "%s"
	organisation_sponsor_id  = "%s"
	sponsored_association_id = "%s"
	bic					     = "%s"
	bic_sponsored			 = "%s"
    reachable_bics           = ["%s"]
}

resource "form3_organisation" "organisation" {
	organisation_id        = "${local.organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "${local.organisation_name}"
}

resource "form3_sepainstant_association" "association" {
	organisation_id      = "${form3_organisation.organisation.organisation_id}"
	association_id       = "${local.association_id}"
  	business_user_dn     = "cn=${lower(local.bic)},ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu"
  	transport_profile_id = "TEST_PROFILE_1"
	bic                  = "${local.bic}"
	simulator_only       = true
    disable_outbound_payments = true
}

resource "form3_organisation" "organisation_sponsored" {
	organisation_id        = "${local.organisation_sponsor_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "${local.organisation_name}"
}

resource "form3_sepainstant_association" "association_sponsored" {
    organisation_id      = "${form3_organisation.organisation_sponsored.organisation_id}"
    association_id       = "${local.sponsored_association_id}"
    business_user_dn     = ""
    transport_profile_id = ""
    bic                  = "${local.bic_sponsored}"
    simulator_only       = true
    sponsor_id           = "${form3_sepainstant_association.association.association_id}"
    reachable_bics       = "${local.reachable_bics}"

    depends_on = [
      "form3_sepainstant_association.association"
    ]
}
`

const testForm3SepaInstantAssociationUpdatedConfig = `
locals {
	organisation_id          = "%s"
	parent_organisation_id   = "%s"
	organisation_name 		 = "%s"
	association_id           = "%s"
	organisation_sponsor_id  = "%s"
	sponsored_association_id = "%s"
	bic						 = "%s"
	bic_sponsored			 = "%s"
    reachable_bics           = ["%s"]
}

resource "form3_organisation" "organisation" {
	organisation_id        = "${local.organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "${local.organisation_name}"
}

resource "form3_sepainstant_association" "association" {
	organisation_id                    = "${form3_organisation.organisation.organisation_id}"
	association_id                     = "${local.association_id}"
  	business_user_dn                   = "cn=${lower(local.bic)},ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu"
  	transport_profile_id               = "TEST_PROFILE_2"
	bic                                = "${local.bic}"
	simulator_only                     = true
	disable_outbound_payments          = true
    enable_customer_admission_decision = true
	enable_customer_check              = true
	clearing_system                    = "tips"
}

resource "form3_organisation" "organisation_sponsored" {
	organisation_id        = "${local.organisation_sponsor_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "terraform-organisation"
}

resource "form3_sepainstant_association" "association_sponsored" {
    organisation_id      = "${form3_organisation.organisation_sponsored.organisation_id}"
    association_id       = "${local.sponsored_association_id}"
    business_user_dn     = ""
    transport_profile_id = ""
    bic                  = "${local.bic_sponsored}"
    simulator_only       = true
    sponsor_id           = "${form3_sepainstant_association.association.association_id}"
    reachable_bics       = "${local.reachable_bics}"

    depends_on = [
      "form3_sepainstant_association.association"
    ]
}
`
