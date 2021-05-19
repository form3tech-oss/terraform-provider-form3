package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/organisations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccOrganisation_basic(t *testing.T) {
	var organisationResponse organisations.GetUnitsIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganisationDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3OrganisationConfig(organisationID, parentOrganisationID, testOrgName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganisationExists("form3_organisation.organisation", &organisationResponse),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "name", testOrgName),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "parent_organisation_id", parentOrganisationID),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "organisation_id", organisationID),
				),
			},
			{
				Config: getTestForm3OrganisationConfigUpdate(organisationID, parentOrganisationID, testOrgName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganisationExists("form3_organisation.organisation", &organisationResponse),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "name", fmt.Sprintf("%s-updated", testOrgName)),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "parent_organisation_id", parentOrganisationID),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "organisation_id", organisationID),
				),
			},
		},
	})
}

func TestAccOrganisation_importBasic(t *testing.T) {

	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	resourceName := "form3_organisation.organisation"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganisationDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3OrganisationConfig(organisationID, parentOrganisationID, testOrgName),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckOrganisationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_organisation" {
			continue
		}

		response, err := client.OrganisationClient.Organisations.GetUnitsID(organisations.NewGetUnitsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckOrganisationExists(resourceKey string, organisation *organisations.GetUnitsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.OrganisationClient.Organisations.GetUnitsID(organisations.NewGetUnitsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		organisation = foundRecord

		return nil
	}
}

func getTestForm3OrganisationConfig(orgID, parOrgID, orgName string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
	}`, orgID, parOrgID, orgName)
}

func getTestForm3OrganisationConfigUpdate(orgID, parOrgID, orgName string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s-updated"
	}`, orgID, parOrgID, orgName)
}
