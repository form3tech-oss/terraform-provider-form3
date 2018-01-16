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

func TestAccBankID_basic(t *testing.T) {
	var bankIDResponse accounts.GetBankidsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	bankResourceId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBankIDDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BankIDConfigA, organisationId, parentOrganisationId, bankResourceId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBankIDExists("form3_bank_id.bank_id", &bankIDResponse),
					resource.TestCheckResourceAttr(
						"form3_bank_id.bank_id", "bank_id", "999999"),
					resource.TestCheckResourceAttr(
						"form3_bank_id.bank_id", "bank_id_code", "GBDSC"),
					resource.TestCheckResourceAttr(
						"form3_bank_id.bank_id", "country", "GB"),
				),
			},
		},
	})
}

func TestAccBankID_importBasic(t *testing.T) {

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	bankResourceId := uuid.NewV4().String()

	resourceName := "form3_bank_id.bank_id"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganisationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testForm3BankIDConfigA, organisationId, parentOrganisationId, bankResourceId),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckBankIDDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_bankID" {
			continue
		}

		response, err := client.AccountClient.Accounts.GetBankidsID(accounts.NewGetBankidsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckBankIDExists(resourceKey string, bankIDResponse *accounts.GetBankidsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AccountClient.Accounts.GetBankidsID(accounts.NewGetBankidsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		bankIDResponse = foundRecord

		return nil
	}
}

const testForm3BankIDConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_bank_id" "bank_id" {
	organisation_id  = "${form3_organisation.organisation.organisation_id}"
  bank_resource_id = "%s"
	bank_id       	 = "999999"
  bank_id_code     = "GBDSC"
  country          = "GB"
}
`
