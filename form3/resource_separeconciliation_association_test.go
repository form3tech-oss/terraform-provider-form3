package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	uuid "github.com/satori/go.uuid"
)

func TestAccSepaReconciliationAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	associationId := uuid.NewV4().String()

	name := uuid.NewV4().String()
	bic := uuid.NewV4().String()
	iban := uuid.NewV4().String()
	address_street := uuid.NewV4().String()
	address_building_number := uuid.NewV4().String()
	address_city := uuid.NewV4().String()
	address_country := uuid.NewV4().String()

	assoc_path := "form3_separeconciliation_association.association"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaReconciliationAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testForm3SepaReconciliationAssociationConfigA,
					organisationId,
					parentOrganisationId,
					associationId,
					name,
					bic,
					iban,
					address_street,
					address_building_number,
					address_city,
					address_country,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaReconciliationAssociationExists(assoc_path),
					resource.TestCheckResourceAttr(assoc_path, "association_id", associationId),
					resource.TestCheckResourceAttr(assoc_path, "organisation_id", organisationId),
					resource.TestCheckResourceAttr(assoc_path, "name", name),
					resource.TestCheckResourceAttr(assoc_path, "bic", bic),
					resource.TestCheckResourceAttr(assoc_path, "iban", iban),
					resource.TestCheckResourceAttr(assoc_path, "address_street", address_street),
					resource.TestCheckResourceAttr(assoc_path, "address_building_number", address_building_number),
					resource.TestCheckResourceAttr(assoc_path, "address_city", address_city),
					resource.TestCheckResourceAttr(assoc_path, "address_country", address_country),
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
			return fmt.Errorf("no sepa reconciliation Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetLhvID(associations.NewGetLhvIDParams().
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
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_separeconciliation_association" "association" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
	association_id          = "%s"
	name                    = "%s
	bic                     = "%s
	iban                    = "%s
	address_street          = "%s
	address_building_number = "%s
	address_city            = "%s
	address_country         = "%s
}`
