package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
  "github.com/ewilde/go-form3/client/accounts"
)

func TestAccBankID_basic(t *testing.T) {
	var bankIDResponse accounts.GetBankidsIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBankIDDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BankIDConfigA, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBankIDExists("form3_bank_id.bank_id", &bankIDResponse),
					resource.TestCheckResourceAttr(
						"form3_bank_id.bank_id", "bank_id", "400309"),
					resource.TestCheckResourceAttr(
						"form3_bank_id.bank_id", "bank_id_code", "GBDSC"),
					resource.TestCheckResourceAttr(
						"form3_bank_id.bank_id", "country", "GB"),
				),
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
resource "form3_bank_id" "bank_id" {
	organisation_id  = "%s"
  bank_resource_id = "2f606e19-d32e-4dcb-9237-7348f71d7da0"
	bank_id       	 = "400309"
  bank_id_code     = "GBDSC"
  country          = "GB"
}
`
