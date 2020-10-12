package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccEburyAssociationAccount_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	associationId := uuid.New().String()
	associationAccountId := uuid.New().String()

	accountNumber := "GB1158E1BBDC7E5648C5A19594442A413ECB"
	accountNumberCode := "IBAN"
	accountWithBic := "SSMMKKDD"
	accountWithBankId := "200605"
	accountWithBankIdCode := "GBDSC"
	accountLabels := []string{"funding_account"}
	currency := "GBP"
	country := "GB"
	iban := "GB1158E1BBDC7E5648C5A19594442A413ECB"

	accountNumberUpdate := "GB1158E1BBDC7E5648C5A19594442A413ECC"
	accountNumberCodeUpdate := "BBAN"
	accountWithBicUpdate := "SSMMKKDE"
	accountWithBankIdUpdate := "200606"
	accountWithBankIdCodeUpdate := "GBDSD"
	accountLabelsUpdate := []string{"funding_account"}
	currencyUpdate := "EUR"
	countryUpdate := "FR"
	ibanUpdate := "GB1158E1BBDC7E5648C5A19594442A413ECC"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEburyAssociationAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3EburyAssociationAccountConfig, organisationId, parentOrganisationId, associationId, associationAccountId,
					accountNumber, accountNumberCode, accountWithBic, accountWithBankId, accountWithBankIdCode, accountLabels[0], currency, country, iban),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEburyAssociationAccountExists("form3_ebury_association_account.association_account"),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "association_account_id", associationAccountId),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_number", accountNumber),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_number_code", accountNumberCode),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_with_bic", accountWithBic),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_with_bank_id", accountWithBankId),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_with_bank_id_code", accountWithBankIdCode),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_labels.0", accountLabels[0]),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "currency", currency),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "country", country),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "iban", iban),
				),
			},
			{
				Config: fmt.Sprintf(testForm3EburyAssociationAccountConfig, organisationId, parentOrganisationId, associationId, associationAccountId,
					accountNumberUpdate, accountNumberCodeUpdate, accountWithBicUpdate, accountWithBankIdUpdate, accountWithBankIdCodeUpdate,
					accountLabelsUpdate[0], currencyUpdate, countryUpdate, ibanUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEburyAssociationAccountExists("form3_ebury_association_account.association_account"),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "association_account_id", associationAccountId),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_number", accountNumberUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_number_code", accountNumberCodeUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_with_bic", accountWithBicUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_with_bank_id", accountWithBankIdUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_with_bank_id_code", accountWithBankIdCodeUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "account_labels.0", accountLabelsUpdate[0]),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "currency", currencyUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "country", countryUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association_account.association_account", "iban", ibanUpdate),
				),
			},
		},
	})
}

func testAccCheckEburyAssociationAccountDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_ebury_association_account" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetEburyAssociationIDAccountsID(associations.NewGetEburyAssociationIDAccountsIDParams().
			WithAssociationID(strfmt.UUID(rs.Primary.Attributes["association_id"])).
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("ebury association account record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckEburyAssociationAccountExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ebury association account Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetEburyAssociationIDAccountsID(associations.NewGetEburyAssociationIDAccountsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)).
			WithAssociationID(strfmt.UUID(rs.Primary.Attributes["association_id"])))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("ebury association account record not found, expected %s but found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3EburyAssociationAccountConfig = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-organisation"
}

resource "form3_ebury_association" "association" {
	organisation_id  		 = "${form3_organisation.organisation.organisation_id}"
	organisation_location 	 = "GB"
	association_id           = "%s"
	funding_currency 		 = "GBP"
	ebury_contact_id 		 = "CON100000"
	ebury_client_id  		 = "CLI100000"
	party_payment_fee		 = "2.00"
	organisation_payment_fee = "3.00"
	organisation_kyc_model   = "Reliance"
	party_name 				 = "Test party"
	party_address  			 = ["Test", "party", "address"]
	party_city 				 = "Test city"
	party_post_code 		 = "TE1 2ST"
}

resource "form3_ebury_association_account" "association_account" {
	organisation_id  		  = "${form3_organisation.organisation.organisation_id}"
	association_id            = "${form3_ebury_association.association.association_id}"
	association_account_id    = "%s"
	account_number            = "%s"
	account_number_code       = "%s"
	account_with_bic          = "%s"
	account_with_bank_id      = "%s"
	account_with_bank_id_code = "%s"
	account_labels            = ["%s"]
	currency                  = "%s"
	country                   = "%s"
	iban                      = "%s"
}`
