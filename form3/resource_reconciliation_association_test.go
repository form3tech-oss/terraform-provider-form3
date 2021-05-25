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

func TestAccReconciliationAssociation_basic(t *testing.T) {
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	associationID := uuid.New().String()
	name := "test organisation"

	// hardcoded value
	// https://github.com/hashicorp/terraform-plugin-sdk/v2/issues/196
	bankID := "QWKEHG33"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckReconciliationAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3ReconciliationAssociationConfig(organisationID, parentOrganisationID, testOrgName, associationID, name, bankID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReconciliationAssociationExists("form3_reconciliation_association.association"),
					resource.TestCheckResourceAttr("form3_reconciliation_association.association", "association_id", associationID),
					resource.TestCheckResourceAttr("form3_reconciliation_association.association", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_reconciliation_association.association", "name", name),
					resource.TestCheckResourceAttr("form3_reconciliation_association.association", "bank_ids.#", "1"),
					resource.TestCheckResourceAttr("form3_reconciliation_association.association", "bank_ids.3361187273", bankID),
					resource.TestCheckResourceAttr("form3_reconciliation_association.association", "scheme_type", "SEPAINSTANT"),
				),
			},
		},
	})
}

func testAccCheckReconciliationAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_reconciliation_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetReconciliationAssociationID(
			associations.NewGetReconciliationAssociationIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.ID)),
		)

		if err == nil {
			return fmt.Errorf("reconciliation association `%s` still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckReconciliationAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no reconciliation association ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetReconciliationAssociationID(
			associations.NewGetReconciliationAssociationIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.ID)),
		)
		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("reconciliation association not found, expected `%s` but found `%s`", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

func getTestForm3ReconciliationAssociationConfig(orgID, parOrgID, orgName, assocID, assocName, bankIDs string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		           = "%s"
	}

	resource "form3_reconciliation_association" "association" {
		organisation_id = "${form3_organisation.organisation.organisation_id}"
		association_id  = "%s"
		name            = "%s"
		bank_ids        = [ "%s" ]
		scheme_type     = "SEPAINSTANT"
	}`, orgID, parOrgID, orgName, assocID, assocName, bankIDs)
}
