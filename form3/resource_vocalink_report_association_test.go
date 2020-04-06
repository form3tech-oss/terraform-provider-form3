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

func TestAccVocalinkReportAssociation_basic(t *testing.T) {
	var response associations.GetVocalinkreportIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.NewV4().String()
	associationID := uuid.NewV4().String()

	bacsMemberKeyID := uuid.NewV4().String()
	bacsMemberCertificateID := uuid.NewV4().String()
	fpsMemberKeyID := uuid.NewV4().String()
	fpsMemberCertificateID := uuid.NewV4().String()
	bacsServiceUserKeyID := uuid.NewV4().String()
	bacsServiceUserCertificateID := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVocalinkReportAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3VocalinkReportAssociationConfigA,
					organisationID,
					parentOrganisationID,
					associationID,
					bacsMemberKeyID,
					bacsMemberCertificateID,
					fpsMemberKeyID,
					fpsMemberCertificateID,
					bacsServiceUserKeyID,
					bacsServiceUserCertificateID,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocalinkReportAssociationExists("form3_vocalink_report_association.association", &response),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "association_id", associationID),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_member_key_id", bacsMemberKeyID),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_member_certificate_id", bacsMemberCertificateID),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "fps_member_key_id", fpsMemberKeyID),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "fps_member_certificate_id", fpsMemberCertificateID),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_service_user_key_id", bacsServiceUserKeyID),
					resource.TestCheckResourceAttr("form3_vocalink_report_association.association", "bacs_service_user_certificate_id", bacsServiceUserCertificateID),
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

locals {
	association_id							= "%s"
	bacs_member_key_id						= "%s"
	bacs_member_certificate_id				= "%s"
	fps_member_key_id						= "%s"
	fps_member_certificate_id				= "%s"
	bacs_service_user_key_id				= "%s"
	bacs_service_user_certificate_id        = "%s"
}

resource "form3_vocalink_report_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "${local.association_id}"
  bacs_member_key_id               = "${local.bacs_member_key_id}"
  bacs_member_certificate_id       = "${local.bacs_member_certificate_id}"
  fps_member_key_id                = "${local.fps_member_key_id}"
  fps_member_certificate_id        = "${local.fps_member_certificate_id}"
  bacs_service_user_number         = "998877"
  bacs_service_user_key_id         = "${local.bacs_service_user_key_id}"
  bacs_service_user_certificate_id = "${local.bacs_service_user_certificate_id}"
}`
