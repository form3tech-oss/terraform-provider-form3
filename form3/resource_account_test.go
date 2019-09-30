package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	uuid "github.com/satori/go.uuid"
)

func TestAccAccount_basic(t *testing.T) {
	var accountResponse accounts.GetAccountsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	accountId := uuid.NewV4().String()
	bankResourceId := uuid.NewV4().String()
	bicId := uuid.NewV4().String()
	bic := "NWABCD13"
	accountNumber := randomAccountNumber()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3AccountConfigA, organisationId, parentOrganisationId, accountId, accountNumber, bic, bankResourceId, bicId, bic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountExists("form3_account.account", &accountResponse),
					resource.TestCheckResourceAttr("form3_account.account", "account_id", accountId),
					resource.TestCheckResourceAttr("form3_account.account", "account_number", strconv.Itoa(accountNumber)),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id", "401005"),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id_code", "GBDSC"),
					resource.TestCheckResourceAttr("form3_account.account", "bic", "NWABCD13"),
					resource.TestCheckResourceAttr("form3_account.account", "country", "GB"),
				),
			},
		},
	})
}

func TestAccAccount_basic_with_iban(t *testing.T) {
	var accountResponse accounts.GetAccountsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	accountId := uuid.NewV4().String()
	bankResourceId := uuid.NewV4().String()
	bicId := uuid.NewV4().String()
	bic := "NWABCD13"
	accountNumber := randomAccountNumber()
	iban := "GB65FTHR40000166854176"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3AccountConfigWithIban, organisationId, parentOrganisationId, accountId, accountNumber, iban, bic, bankResourceId, bicId, bic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountExists("form3_account.account", &accountResponse),
					resource.TestCheckResourceAttr("form3_account.account", "account_id", accountId),
					resource.TestCheckResourceAttr("form3_account.account", "account_number", strconv.Itoa(accountNumber)),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id", "401005"),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id_code", "GBDSC"),
					resource.TestCheckResourceAttr("form3_account.account", "bic", "NWABCD13"),
					resource.TestCheckResourceAttr("form3_account.account", "country", "GB"),
					resource.TestCheckResourceAttr("form3_account.account", "iban", iban),
				),
			},
		},
	})
}

func randomAccountNumber() int {
	accountNumber := rand.Intn(99999999)
	for accountNumber < 10000000 {
		accountNumber = rand.Intn(99999999)
	}
	return accountNumber
}

func TestAccAccount_importBasic(t *testing.T) {

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	accountId := uuid.NewV4().String()
	bankResourceId := uuid.NewV4().String()
	bicId := uuid.NewV4().String()
	bic := "NWABCD14"
	accountNumber := randomAccountNumber()

	resourceName := "form3_account.account"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3AccountConfigA, organisationId, parentOrganisationId, accountId, accountNumber, bic, bankResourceId, bicId, bic),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAccount_import_with_iban(t *testing.T) {

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	accountId := uuid.NewV4().String()
	bankResourceId := uuid.NewV4().String()
	bicId := uuid.NewV4().String()
	bic := "NWABCD14"
	accountNumber := randomAccountNumber()
	iban := "GB65FTHR40000166854176"

	resourceName := "form3_account.account"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3AccountConfigWithIban, organisationId, parentOrganisationId, accountId, accountNumber, iban, bic, bankResourceId, bicId, bic),
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
  account_number   = "%d"
  bank_id          = "401005"
  bank_id_code     = "GBDSC"
  bic              = "%s"
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
	bic       	    = "%s"
}
`

const testForm3AccountConfigWithIban = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_account" "account" {
	organisation_id  = "${form3_organisation.organisation.organisation_id}"
  account_id       = "%s"
  account_number   = "%d"
  iban             = "%s"
  bank_id          = "401005"
  bank_id_code     = "GBDSC"
  bic              = "%s"
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
	bic       	    = "%s"
}
`
