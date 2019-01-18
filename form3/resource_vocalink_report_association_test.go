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

func TestAccVocalinkReportAssociation_basic(t *testing.T) {
	var response associations.GetVocalinkreportIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVocalinkReportAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3VocalinkReportAssociationConfigA, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocalinkReportAssociationExists("form3_vocalink_report_association.association", &response),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "association_id", "ad5e20e5-800d-4143-9936-ca1007da3a03"),
				),
			},
		},
	})
}

func testAccCheckVocalinkReportAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_vocalink_report_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckVocalinkReportAssociationExists(resourceKey string, association *associations.GetVocalinkreportIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
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

const testForm3VocalinkReportAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_vocalink_report_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "ad5e20e5-800d-4143-9936-ca1007da3a03"
}`
