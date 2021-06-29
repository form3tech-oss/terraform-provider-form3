package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccMandateManagement_basic(t *testing.T) {
	var mandateManagementResponse associations.GetMandatemanagementIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	mandateManagementId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testMandateManagementDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3MandateManagementConfig(organisationID, parentOrganisationID, testOrgName, mandateManagementId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMandateManagementExists("form3_mandate_management_association.mandate_management", &mandateManagementResponse),
					resource.TestCheckResourceAttr("form3_mandate_management_association.mandate_management", "payment_scheme", string(models.PaymentSchemeBACS)),
				),
			},
		},
	})
}

func TestAccMandateManagement_importBasic(t *testing.T) {
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	mandateManagementId := uuid.New().String()

	resourceName := "form3_mandate_management_association.mandate_management"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3MandateManagementConfig(organisationID, parentOrganisationID, testOrgName, mandateManagementId),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckMandateManagementExists(resourceKey string, mandateManagement *associations.GetMandatemanagementIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetMandatemanagementID(associations.NewGetMandatemanagementIDParams().WithID(strfmt.UUID(rs.Primary.ID)))
		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		mandateManagement = foundRecord

		return nil
	}
}

func testMandateManagementDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_mandate_management_association" {
			continue
		}
		response, err := client.AssociationClient.Associations.GetMandatemanagementID(associations.NewGetMandatemanagementIDParams().WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}
	return nil
}

func getTestForm3MandateManagementConfig(orgID, parOrgID, orgName, mandateManagementID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		           = "%s"
	}

	resource "form3_mandate_management_association" "mandate_management" {
		organisation_id       	= "${form3_organisation.organisation.organisation_id}"
		association_id		    = "%s"
		payment_scheme    	    = "%s"
	}`, orgID, parOrgID, orgName, mandateManagementID, models.PaymentSchemeBACS)
}
