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

func TestAccLhvAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)
	associationId := uuid.New().String()
	clientCode := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLhvAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3LhvAssociationConfigA, organisationId, parentOrganisationId, associationId, clientCode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLhvAssociationExists("form3_lhv_association.association"),
					resource.TestCheckResourceAttr("form3_lhv_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_lhv_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_lhv_association.association", "name", "terraform-association"),
					resource.TestCheckResourceAttr("form3_lhv_association.association", "client_code", clientCode),
					resource.TestCheckResourceAttr("form3_lhv_association.association", "client_country", "GB"),
				),
			},
		},
	})
}

func testAccCheckLhvAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_lhv_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetLhvAssociationID(
			associations.NewGetLhvAssociationIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.ID)),
		)

		if err == nil {
			return fmt.Errorf("lhv record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckLhvAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no lhv Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetLhvAssociationID(
			associations.NewGetLhvAssociationIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.ID)),
		)

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("lhv record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3LhvAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		              = "terraform-provider-form3-test-organisation"
}

resource "form3_lhv_association" "association" {
	organisation_id  = "${form3_organisation.organisation.organisation_id}"
    name             = "terraform-association"
	association_id   = "%s"
	client_code      = "%s"
    client_country   = "GB"
}`
