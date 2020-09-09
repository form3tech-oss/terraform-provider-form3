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

func TestAccProductsAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	associationId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProductsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3ProductsAssociationConfigA, organisationId, parentOrganisationId, associationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProductsAssociationExists("form3_products_association.association"),
					resource.TestCheckResourceAttr("form3_products_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_products_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_products_association.association", "product", "INTERNATIONAL_SERVICES"),
					resource.TestCheckResourceAttr("form3_products_association.association", "provider", "EBURY"),
				),
			},
		},
	})
}

func testAccCheckProductsAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_products_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetProductsID(associations.NewGetProductsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("product record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckProductsAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no product Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetProductsID(associations.NewGetProductsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("product record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3ProductsAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}
resource "form3_products_association" "association" {
	organisation_id        = "${form3_organisation.organisation.organisation_id}"
	association_id         = "%s"
	product                = "INTERNATIONAL_SERVICES"
	provider               = "EBURY"
}`
