package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/accounts"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"os"
	"testing"
)

func TestAccAccountConfigurationBasic(t *testing.T) {
	var accountResponse accounts.GetAccountconfigurationsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	accountConfigurationId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3AccountConfigurationConfigA, organisationId, parentOrganisationId, accountConfigurationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountConfigurationExists("form3_account_configuration.configuration", &accountResponse),
					resource.TestCheckResourceAttr(
						"form3_account_configuration.configuration", "account_generation_enabled", "true"),
				),
			},
		},
	})
}

func TestAccAccountConfigurationImportBasic(t *testing.T) {

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	accountConfigurationId := uuid.NewV4().String()
	resourceName := "form3_account_configuration.configuration"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountConfigurationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testForm3AccountConfigurationConfigA, organisationId, parentOrganisationId, accountConfigurationId),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAccountConfigurationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_account_configuration" {
			continue
		}

		response, err := client.AccountClient.Accounts.GetAccountconfigurationsID(accounts.NewGetAccountconfigurationsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckAccountConfigurationExists(resourceKey string, configuration *accounts.GetAccountconfigurationsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AccountClient.Accounts.GetAccountconfigurationsID(accounts.NewGetAccountconfigurationsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		configuration = foundRecord

		return nil
	}
}

const testForm3AccountConfigurationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_account_configuration" "configuration" {
	organisation_id            = "${form3_organisation.organisation.organisation_id}"
	account_configuration_id   = "%s"
	account_generation_enabled = true
  account_generation_configuration = [
        {
            country               = "US"
            valid_account_ranges  = [
                {
                   minimum = "8400000000"
                   maximum = "8409999999"
                }
            ]
        }
    ]
}`
