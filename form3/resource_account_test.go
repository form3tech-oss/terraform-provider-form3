package form3

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"

	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAccount_basic(t *testing.T) {
	var before accounts.GetAccountsIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	accountID := uuid.New().String()
	bankResourceID := uuid.New().String()
	bicID := uuid.New().String()
	bic := generateTestBic()
	accountNumber := randomAccountNumber()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: getForm3AccountTFConfig(organisationID, parentOrganisationID, testOrgName, accountID, bic, bankResourceID, bicID, accountNumber),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountExists("form3_account.account", &before),
					resource.TestCheckResourceAttr("form3_account.account", "account_id", accountID),
					resource.TestCheckResourceAttr("form3_account.account", "account_number", strconv.Itoa(accountNumber)),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id", "401005"),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id_code", "GBDSC"),
					resource.TestCheckResourceAttr("form3_account.account", "bic", bic),
					resource.TestCheckResourceAttr("form3_account.account", "country", "GB"),
					resource.TestCheckResourceAttrSet("form3_account.account", "iban"),
				),
			},
		},
	})
}

func generateRandomIban() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return fmt.Sprintf("GB22ABCD192837%08d", r1.Intn(100000000))
}

func TestAccAccount_basic_with_iban(t *testing.T) {
	var accountResponse accounts.GetAccountsIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	accountID := uuid.New().String()
	bankResourceID := uuid.New().String()
	bicID := uuid.New().String()
	bic := generateTestBic()
	iban := generateRandomIban()
	accountNumber := randomAccountNumber()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: getForm3AccountTFConfigWithIban(organisationID, parentOrganisationID, testOrgName, accountID, bic, bankResourceID, bicID, iban, accountNumber),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountExists("form3_account.account", &accountResponse),
					resource.TestCheckResourceAttr("form3_account.account", "account_id", accountID),
					resource.TestCheckResourceAttr("form3_account.account", "account_number", strconv.Itoa(accountNumber)),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id", "401005"),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id_code", "GBDSC"),
					resource.TestCheckResourceAttr("form3_account.account", "bic", bic),
					resource.TestCheckResourceAttr("form3_account.account", "country", "GB"),
					resource.TestCheckResourceAttr("form3_account.account", "iban", iban),
				),
			},
		},
	})
}

func TestAccAccount_basic_with_iban_without_account_number(t *testing.T) {
	var accountResponse accounts.GetAccountsIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	accountID := uuid.New().String()
	bankResourceID := uuid.New().String()
	bicID := uuid.New().String()
	bic := generateTestBic()
	iban := generateRandomIban()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3AccountConfigWithIbanWithoutAccountNumber(organisationID, parentOrganisationID, testOrgName, accountID, bic, bankResourceID, bicID, iban),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountExists("form3_account.account", &accountResponse),
					resource.TestCheckResourceAttr("form3_account.account", "account_id", accountID),
					resource.TestCheckResourceAttr("form3_account.account", "account_number", ""),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id", "401005"),
					resource.TestCheckResourceAttr("form3_account.account", "bank_id_code", "GBDSC"),
					resource.TestCheckResourceAttr("form3_account.account", "bic", bic),
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

	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	accountID := uuid.New().String()
	bankResourceID := uuid.New().String()
	bicID := uuid.New().String()
	bic := generateTestBic()
	accountNumber := randomAccountNumber()

	resourceName := "form3_account.account"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: getForm3AccountTFConfig(organisationID, parentOrganisationID, testOrgName, accountID, bic, bankResourceID, bicID, accountNumber),
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
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	accountID := uuid.New().String()
	bankResourceID := uuid.New().String()
	bicID := uuid.New().String()
	bic := generateTestBic()
	accountNumber := randomAccountNumber()
	iban := fmt.Sprintf("GB65FTHR400001%d", accountNumber)

	resourceName := "form3_account.account"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: getForm3AccountTFConfigWithIban(organisationID, parentOrganisationID, testOrgName, accountID, bic, bankResourceID, bicID, iban, accountNumber),
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

		*accountResponse = *foundRecord
		return nil
	}
}

func getForm3AccountTFConfig(organisationID, parentOrganisationID, orgName, accountID, bic, bankResourceID, bicID string, accountNumber int) string {

	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
	}
	
	resource "form3_account_configuration" "customer_backoffice_configuration" {
	  organisation_id             = "${form3_organisation.organisation.organisation_id}"
	  account_configuration_id    = "${uuid()}"
	  account_generation_enabled  = true
	
	  lifecycle {
		ignore_changes = ["account_configuration_id"]
	  }
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
	`, organisationID, parentOrganisationID, orgName, accountID, accountNumber, bic, bankResourceID, bicID, bic)
}

func getForm3AccountTFConfigWithIban(organisationID, parentOrganisationID, orgName, accountID, bic, bankResourceID, bicID, iban string, accountNumber int) string {

	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
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
	`, organisationID, parentOrganisationID, orgName, accountID, accountNumber, iban, bic, bankResourceID, bicID, bic)
}

func getTestForm3AccountConfigWithIbanWithoutAccountNumber(organisationID, parentOrganisationID, orgName, accountID, bic, bankResourceID, bicID, iban string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
	}
	
	resource "form3_account_configuration" "customer_backoffice_configuration" {
	  organisation_id             = "${form3_organisation.organisation.organisation_id}"
	  account_configuration_id    = "${uuid()}"
	  account_generation_enabled  = true
	
	  lifecycle {
		ignore_changes = ["account_configuration_id"]
	  }
	}
	
	resource "form3_account" "account" {
	  organisation_id  = "${form3_organisation.organisation.organisation_id}"
	  account_id       = "%s"
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
	  bank_id          = "401005"
	  bank_id_code     = "GBDSC"
	  country          = "GB"
	}
	
	resource "form3_bic" "bic" {
		organisation_id = "${form3_organisation.organisation.organisation_id}"
		bic_id          = "%s"
		bic       	    = "%s"
	}`, organisationID, parentOrganisationID, orgName, accountID, iban, bic, bankResourceID, bicID, bic)
}
