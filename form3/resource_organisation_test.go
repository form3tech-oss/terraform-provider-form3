package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/organisations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/satori/go.uuid"
	"os"
	"testing"
)

func TestAccOrganisation_basic(t *testing.T) {
	var organisationResponse organisations.GetUnitsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganisationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3OrganisationConfigA, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganisationExists("form3_organisation.organisation", &organisationResponse),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "name", "terraform-organisation"),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "parent_organisation_id", parentOrganisationId),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "organisation_id", organisationId),
				),
			},
			{
				Config: fmt.Sprintf(testForm3OrganisationConfigAUpdate, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganisationExists("form3_organisation.organisation", &organisationResponse),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "name", "terraform-organisation-updated"),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "parent_organisation_id", parentOrganisationId),
					resource.TestCheckResourceAttr(
						"form3_organisation.organisation", "organisation_id", organisationId),
				),
			},
		},
	})
}

func TestAccOrganisation_importBasic(t *testing.T) {

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()

	resourceName := "form3_organisation.organisation"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganisationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testForm3OrganisationConfigA, organisationId, parentOrganisationId),
			},

			resource.TestStep{
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

const testForm3OrganisationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}`

const testForm3OrganisationConfigAUpdate = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation-updated"
}`
