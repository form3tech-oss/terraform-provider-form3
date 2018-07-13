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

func TestAccAccount_basic(t *testing.T) {
	var accountResponse accounts.GetAccountsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	accountId := uuid.NewV4().String()
	bankResourceId := uuid.NewV4().String()
	bicId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3AccountConfigA, organisationId, parentOrganisationId, accountId, bankResourceId, bicId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountExists("form3_account.account", &accountResponse),
					resource.TestCheckResourceAttr("form3_account.account", "account_id", accountId),
					resource.TestCheckResourceAttr("form3_account.account", "account_number", "12345678"),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id", "401005"),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id_code", "GBDSC"),
					resource.TestCheckResourceAttr("form3_account.account", "bic", "NWABCD13"),
					resource.TestCheckResourceAttr("form3_account.account", "country", "GB"),
				),
			},
		},
	})
}

func TestAccAccount_importBasic(t *testing.T) {

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	accountId := uuid.NewV4().String()
	bankResourceId := uuid.NewV4().String()
	bicId := uuid.NewV4().String()

	resourceName := "form3_account.account"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3AccountConfigA, organisationId, parentOrganisationId, accountId, bankResourceId, bicId),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAccountDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_account" {
			continue
		}

		response, err := client.AccountClient.Accounts.GetAccountsID(accounts.NewGetAccountsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckAccountExists(resourceKey string, accountResponse *accounts.GetAccountsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AccountClient.Accounts.GetAccountsID(accounts.NewGetAccountsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		accountResponse = foundRecord

		return nil
	}
}

const testForm3AccountConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_account" "account" {
	organisation_id  = "${form3_organisation.organisation.organisation_id}"
  account_id       = "%s"
	account_number 	 = "12345678"
  bank_id          = "401005"
  bank_id_code     = "GBDSC"
  bic              = "NWABCD13"
  country          = "GB"
  depends_on       = ["form3_bank_id.bank_id", "form3_bic.bic"]
}

resource "form3_bank_id" "bank_id" {
	organisation_id  = "${form3_organisation.organisation.organisation_id}"
  bank_resource_id = "%s"
	bank_id       	 = "401005"
  bank_id_code     = "GBDSC"
  country          = "GB" 
}

resource "form3_bic" "bic" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
  bic_id          = "%s"
	bic       	    = "NWABCD13"
}
`
