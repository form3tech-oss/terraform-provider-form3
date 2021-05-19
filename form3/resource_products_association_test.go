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

func TestAccProductsAssociation_basic(t *testing.T) {
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	associationID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProductsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3ProductsAssociationConfig(organisationID, parentOrganisationID, testOrgName, associationID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProductsAssociationExists("form3_products_association.association"),
					resource.TestCheckResourceAttr("form3_products_association.association", "association_id", associationID),
					resource.TestCheckResourceAttr("form3_products_association.association", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_products_association.association", "product", "INTERNATIONAL_SERVICES"),
					resource.TestCheckResourceAttr("form3_products_association.association", "product_provider", "EBURY"),
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

func getTestForm3ProductsAssociationConfig(orgID, parOrgID, orgName, assocID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
	}
	resource "form3_products_association" "association" {
		organisation_id        = "${form3_organisation.organisation.organisation_id}"
		association_id         = "%s"
		product                = "INTERNATIONAL_SERVICES"
		product_provider       = "EBURY"
	}`, orgID, parOrgID, orgName, assocID)
}
