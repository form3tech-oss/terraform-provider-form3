package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGocardlessAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)
	associationId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGocardlessAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3GocardlessAssociationConfigA, organisationId, parentOrganisationId, associationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGocardlessAssociationExists("form3_gocardless_association.association"),
					resource.TestCheckResourceAttr("form3_gocardless_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_gocardless_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_gocardless_association.association", "schemes.0", "BACS"),
				),
			},
			{
				Config: fmt.Sprintf(testForm3GocardlessAssociationConfigB, organisationId, parentOrganisationId, associationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGocardlessAssociationExists("form3_gocardless_association.association"),
					resource.TestCheckResourceAttr("form3_gocardless_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_gocardless_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_gocardless_association.association", "schemes.0", "BACS"),
					resource.TestCheckResourceAttr("form3_gocardless_association.association", "schemes.1", "SEPADD"),
				),
			},
		},
	})
}

func testAccCheckGocardlessAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_gocardless_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetGocardlessID(associations.NewGetGocardlessIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("gocardless record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckGocardlessAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no gocardless Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetGocardlessID(associations.NewGetGocardlessIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("gocardless record not found, expected %s but found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3GocardlessAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_gocardless_association" "association" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
	association_id  = "%s"
	schemes         = ["BACS"]
}`

const testForm3GocardlessAssociationConfigB = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_gocardless_association" "association" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
	association_id  = "%s"
	schemes         = ["BACS", "SEPADD"]
}`
