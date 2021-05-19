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

func TestAccSwiftAssociation_basic(t *testing.T) {
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	associationID := uuid.New().String()
	bic := generateTestBic()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSwiftAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3SwiftAssociationConfig(organisationID, parentOrganisationID, testOrgName, associationID, bic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSwiftAssociationExists("form3_swift_association.association"),
					resource.TestCheckResourceAttr("form3_swift_association.association", "association_id", associationID),
					resource.TestCheckResourceAttr("form3_swift_association.association", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_swift_association.association", "bic", bic),
				),
			},
		},
	})
}

func testAccCheckSwiftAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_swift_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSwiftID(associations.NewGetSwiftIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("swift record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSwiftAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no swift Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSwiftID(associations.NewGetSwiftIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("swift record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}
		return nil
	}
}

func getTestForm3SwiftAssociationConfig(orgID, parOrgID, orgName, assocID, bic string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		           = "%s"
	}

	resource "form3_swift_association" "association" {
		organisation_id        = "${form3_organisation.organisation.organisation_id}"
		association_id         = "%s"
		bic                    = "%s"
	}`, orgID, parOrgID, orgName, assocID, bic)
}
