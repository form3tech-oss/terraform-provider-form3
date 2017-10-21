package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/organisations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestAccOrganisation_basic(t *testing.T) {
	var organisationResponse organisations.GetUnitsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := parentOrganisationId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganisationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3OrganisationConfigA, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganisationExists("form3_organisation.organisation", &organisationResponse),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "name", "terraform-organisation"),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "parent_organisation_id", parentOrganisationId),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "organisation_id", "ea586d19-ed46-4574-938f-539d2cc03efe"),
				),
			},
			{
				Config: fmt.Sprintf(testForm3OrganisationConfigAUpdate, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganisationExists("form3_organisation.organisation", &organisationResponse),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "name", "terraform-organisation-updated"),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "parent_organisation_id", parentOrganisationId),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "organisation_id", "ea586d19-ed46-4574-938f-539d2cc03efe"),
				),
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

const testForm3OrganisationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "ea586d19-ed46-4574-938f-539d2cc03efe"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}`

const testForm3OrganisationConfigAUpdate = `
resource "form3_organisation" "organisation" {
	organisation_id        = "ea586d19-ed46-4574-938f-539d2cc03efe"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation-updated"
}`
