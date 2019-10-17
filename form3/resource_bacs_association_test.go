package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	uuid "github.com/satori/go.uuid"
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
				Config: fmt.Sprintf(testForm3BacsAssociationConfigWithCerts, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_number", "112238"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_number", "12345678"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "sorting_code", "123456"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_type", "1"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "association_id", "ad5e20e5-800d-4143-9936-ca1007da3a03"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "input_key_id", "8f77e1ba-944e-44f3-a845-f99ba80af63c"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "input_certificate_id", "23d4fa5d-ef38-48de-b9e4-22f45004bb50"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "messaging_key_id", "ce3b888b-2328-49ed-9c04-cda0035f8fd0"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "messaging_certificate_id", "f79162cb-cbde-4152-b2f0-3bde47da3332"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "output_key_id", "a7984808-a3bb-4951-827c-d4d15d01ac0b"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "output_certificate_id", "6dd9ca5d-b64a-4b59-a287-ad4ea82acb4f"),
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
				Config: fmt.Sprintf(testForm3BacsAssociationConfigZeroAccountType, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_number", "112233"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_number", "87654321"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "sorting_code", "654321"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_type", "0"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "association_id", "ba2283f5-e194-4e12-ac8d-ae9bb08eeddb"),
				),
			},
		},
	})
}

func TestAccBacsAssociation_withBankIdAndCentre(t *testing.T) {
	var bacsResponse associations.GetBacsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBacsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BacsAssociationConfigWithBankIdAndCentre, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "bank_code", "1234"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "centre_number", "42"),
				),
			},
		},
	})
}

func TestAccBacsAssociation_withTestFileSubmissionFlag(t *testing.T) {
	var bacsResponse associations.GetBacsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBacsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BacsAssociationConfigWithTestFileFlag, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "test_file_submission", "true"),
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

const testForm3BacsAssociationConfigWithCerts = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "ad5e20e5-800d-4143-9936-ca1007da3a03"
	service_user_number              = "112238",
    account_number                   = "12345678",
    sorting_code                     = "123456",
    account_type                     = 1

    input_key_id                     = "8f77e1ba-944e-44f3-a845-f99ba80af63c"
    input_certificate_id             = "23d4fa5d-ef38-48de-b9e4-22f45004bb50"

    messaging_key_id                 = "ce3b888b-2328-49ed-9c04-cda0035f8fd0"
    messaging_certificate_id         = "f79162cb-cbde-4152-b2f0-3bde47da3332"

    output_key_id                    = "a7984808-a3bb-4951-827c-d4d15d01ac0b"
    output_certificate_id            = "6dd9ca5d-b64a-4b59-a287-ad4ea82acb4f"
}`

const testForm3BacsAssociationConfigZeroAccountType = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "ba2283f5-e194-4e12-ac8d-ae9bb08eeddb"
	service_user_number              = "112233",
    account_number                   = "87654321",
    sorting_code                     = "654321",
    account_type                     = 0
}`

const testForm3BacsAssociationConfigWithBankIdAndCentre = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "ba2283f5-e194-4e12-ac8d-ae9bb08eeeee"
	service_user_number              = "112233",
    account_number                   = "87654321",
    sorting_code                     = "654321",
    account_type                     = 0,
    bank_code                        = "1234",
    centre_number                    = "42"
}`

const testForm3BacsAssociationConfigWithTestFileFlag = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "ba2283f5-e194-4e12-ac8d-ae9bb08eeeee"
	service_user_number              = "112233",
    account_number                   = "87654321",
    sorting_code                     = "654321",
    account_type                     = 0,
    bank_code                        = "1234",
    centre_number                    = "42",
    test_file_submission             = true
}`
