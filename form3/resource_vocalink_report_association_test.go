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

func TestAccVocalinkReportAssociation_basic(t *testing.T) {
	var response associations.GetVocalinkreportIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	associationId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVocalinkReportAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3VocalinkReportAssociationConfigA, organisationId, parentOrganisationId, associationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocalinkReportAssociationExists("form3_vocalink_report_association.association", &response),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_member_key_id", "43bdb305-e4b8-4eeb-9b26-b9cb0cad1ae5"),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_member_certificate_id", "cb9c35c3-de6c-4376-bb23-a5ca2cbdb142"),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "fps_member_key_id", "336145dc-587e-47b9-a0c4-61ae25e1f35e"),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "fps_member_certificate_id", "7a7a0f00-c507-49c1-ba89-30fd8521173b"),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_service_user_key_id", "33b06551-e9f0-4e84-b1a0-ab23f6ea5611"),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_service_user_certificate_id", "f9962c0c-992a-4ce7-a6d5-8e8b6222ad65"),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_service_user_number", "998877"),
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
	association_id                   = "%s"
	bacs_member_key_id               = "43bdb305-e4b8-4eeb-9b26-b9cb0cad1ae5"
	bacs_member_certificate_id       = "cb9c35c3-de6c-4376-bb23-a5ca2cbdb142"
	fps_member_key_id                = "336145dc-587e-47b9-a0c4-61ae25e1f35e"
	fps_member_certificate_id        = "7a7a0f00-c507-49c1-ba89-30fd8521173b"
	bacs_service_user_number         = "998877"
	bacs_service_user_key_id         = "33b06551-e9f0-4e84-b1a0-ab23f6ea5611"
	bacs_service_user_certificate_id = "f9962c0c-992a-4ce7-a6d5-8e8b6222ad65"
}`
