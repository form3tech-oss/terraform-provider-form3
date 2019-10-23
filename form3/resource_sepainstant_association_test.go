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

func TestAccSepaInstantAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	associationId := uuid.NewV4().String()
	sponsoredAssociationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaInstantAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SepaInstantAssociationConfigA, organisationId, parentOrganisationId, associationId, sponsoredAssociationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "business_user_dn", "cn=testbic8,ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "transport_profile_id", "TEST_PROFILE_1"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "bic", "TESTBIC8"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "simulator_only", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "sponsor_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "association_id", sponsoredAssociationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "sponsor_id", associationId),
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
	sponsored_association_id = "%s"
}
resource "form3_organisation" "organisation" {
	organisation_id        = "${local.organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		           = "terraform-organisation"
}

resource "form3_sepainstant_association" "association" {
	organisation_id      = "${form3_organisation.organisation.organisation_id}"
	association_id       = "${local.association_id}"
  	business_user_dn     = "cn=testbic8,ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu"
  	transport_profile_id = "TEST_PROFILE_1"
	bic                  = "TESTBIC8"
	simulator_only       = true
}

resource "form3_sepainstant_association" "association_sponsored" {
  organisation_id      = "${form3_organisation.organisation.organisation_id}"
  association_id       = "${local.sponsored_association_id}"
  business_user_dn     = ""
  transport_profile_id = ""
  bic                  = "TESTBIC8"
  simulator_only       = true
  sponsor_id           = "${form3_sepainstant_association.association.association_id}"	

  depends_on = [
    "form3_sepainstant_association.association"
  ]
}
`
