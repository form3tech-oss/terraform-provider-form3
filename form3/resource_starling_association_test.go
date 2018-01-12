//package form3
//
//import (
//	"fmt"
//	"github.com/ewilde/go-form3"
//	"github.com/ewilde/go-form3/client/associations"
//	"github.com/go-openapi/strfmt"
//	"github.com/hashicorp/terraform/helper/resource"
//	"github.com/hashicorp/terraform/terraform"
//	"os"
//	"testing"
//)
//
//func TestAccStarlingAssociation_basic(t *testing.T) {
//	var starlingResponse associations.GetStarlingIDOK
//	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
//
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckStarlingAssociationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: fmt.Sprintf(testForm3StarlingAssociationConfigA, organisationId),
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckStarlingAssociationExists("form3_starling_association.association", &starlingResponse),
//					resource.TestCheckResourceAttr(
//						"form3_starling_association.association", "starling_account_name", "account-1"),
//				),
//			},
//		},
//	})
//}
//
//func testAccCheckStarlingAssociationDestroy(state *terraform.State) error {
//	client := testAccProvider.Meta().(*form3.AuthenticatedClient)
//
//	for _, rs := range state.RootModule().Resources {
//		if rs.Type != "form3_starling_association" {
//			continue
//		}
//
//		response, err := client.AssociationClient.Associations.GetStarlingID(associations.NewGetStarlingIDParams().
//			WithID(strfmt.UUID(rs.Primary.ID)))
//
//		if err == nil {
//			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
//		}
//	}
//
//	return nil
//}
//
//func testAccCheckStarlingAssociationExists(resourceKey string, association *associations.GetStarlingIDOK) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[resourceKey]
//
//		if !ok {
//			return fmt.Errorf("not found: %s", resourceKey)
//		}
//
//		if rs.Primary.ID == "" {
//			return fmt.Errorf("no Record ID is set")
//		}
//
//		client := testAccProvider.Meta().(*form3.AuthenticatedClient)
//
//		foundRecord, err := client.AssociationClient.Associations.GetStarlingID(associations.NewGetStarlingIDParams().
//			WithID(strfmt.UUID(rs.Primary.ID)))
//
//		if err != nil {
//			return err
//		}
//
//		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
//			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
//		}
//
//		association = foundRecord
//
//		return nil
//	}
//}
//
//const testForm3StarlingAssociationConfigA = `
//resource "form3_organisation" "organisation" {
//	organisation_id        = "${uuid()}"
//	parent_organisation_id = "%s"
//	name 		               = "terraform-organisation"
//
//  lifecycle {
//    ignore_changes = ["organisation_id"]
//  }
//}
//
//resource "form3_starling_association" "association" {
//	organisation_id       = "${form3_organisation.organisation.organisation_id}"
//	association_id        = "0b2fc31e-b778-448b-977d-1e7f828a81eb"
//	starling_account_name	= "account-1"
//}`
