package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"os"
	"testing"
)

func TestAccConfirmationOfPayeeAssociation_basic(t *testing.T) {
	var response associations.GetConfirmationOfPayeeIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConfirmationOfPayeeAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3ConfirmationOfPayeeAssociationConfigA, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfirmationOfPayeeAssociationExists("form3_confirmation_of_payee_association.association", &response),
					resource.TestCheckResourceAttr("form3_confirmation_of_payee_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_confirmation_of_payee_association.association", "association_id", "ad5e20e5-800d-4143-9936-ca1007da3a03"),
					resource.TestCheckResourceAttr("form3_confirmation_of_payee_association.association", "open_banking_organisation_id", "43bdb305-e4b8-4eeb-9b26-b9cb0cad1ae5"),
					resource.TestCheckResourceAttr("form3_confirmation_of_payee_association.association", "open_banking_public_key_id", "cb9c35c3-de6c-4376-bb23-a5ca2cbdb142"),
					resource.TestCheckResourceAttr("form3_confirmation_of_payee_association.association", "signing_key_id", "336145dc-587e-47b9-a0c4-61ae25e1f35e"),
					resource.TestCheckResourceAttr("form3_confirmation_of_payee_association.association", "signing_dn", "SIGNER"),
				),
			},
		},
	})
}

func testAccCheckConfirmationOfPayeeAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_confirmation_of_payee_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckConfirmationOfPayeeAssociationExists(resourceKey string, association *associations.GetConfirmationOfPayeeIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		association = foundRecord

		return nil
	}
}

const testForm3ConfirmationOfPayeeAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_confirmation_of_payee_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "ad5e20e5-800d-4143-9936-ca1007da3a03"
  open_banking_organisation_id     = "43bdb305-e4b8-4eeb-9b26-b9cb0cad1ae5"
  open_banking_public_key_id       = "cb9c35c3-de6c-4376-bb23-a5ca2cbdb142"
  signing_key_id                   = "336145dc-587e-47b9-a0c4-61ae25e1f35e"
  signing_dn                       = "SIGNER"
}`
