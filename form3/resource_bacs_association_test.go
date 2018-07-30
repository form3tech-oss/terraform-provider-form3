package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"os"
	"testing"
)

func TestAccBacsAssociation_basic(t *testing.T) {
	var bacsResponse associations.GetBacsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBacsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BacsAssociationConfigA, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_number", "112238"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_number", "12345678"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "sorting_code", "123456"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_type", "1"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "association_id", "ad5e20e5-800d-4143-9936-ca1007da3a03"),
				),
			},
		},
	})
}

func TestAccBacsAssociation_zeroAccountType(t *testing.T) {
	var bacsResponse associations.GetBacsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBacsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BacsAssociationConfigB, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_number", "112238"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_number", "12345678"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "sorting_code", "123456"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_type", "1"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "association_id", "ad5e20e5-800d-4143-9936-ca1007da3a03"),
				),
			},
		},
	})
}

func testAccCheckBacsAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_bacs_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("bacs record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckBacsAssociationExists(resourceKey string, association *associations.GetBacsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("bacs not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no bacs Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("bacs record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		association = foundRecord

		return nil
	}
}

const testForm3BacsAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "ad5e20e5-800d-4143-9936-ca1007da3a03"
	service_user_number              = "112238",
  account_number                   = "12345678",
  sorting_code                     = "123456",
  account_type                     = 1
}`

const testForm3BacsAssociationConfigB = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "ba2283f5-e194-4e12-ac8d-ae9bb08eeddb"
	service_user_number              = "112233",
  account_number                   = "87654321",
  sorting_code                     = "654321",
  account_type                     = "0"
}`
