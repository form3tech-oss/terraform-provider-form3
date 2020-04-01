package api

import (
	"log"

	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/form3tech-oss/terraform-provider-form3/client/organisations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
)

func TestAccDeleteAccount(t *testing.T) {

	bankId := "100202"
	accountNumber := "12345678"

	existingAccounts, err := auth.AccountClient.Accounts.GetAccounts(accounts.NewGetAccountsParams().
		WithFilterAccountNumber([]string{accountNumber}))

	assertNoErrorOccurred(err, t)

	log.Printf("[DEBUG]  accounts: %v", existingAccounts)

	createdBankID, err := auth.AccountClient.Accounts.PostBankids(accounts.NewPostBankidsParams().
		WithBankIDCreationRequest(&models.BankIDCreation{
			Data: &models.BankID{
				Type:           "bankids",
				ID:             NewUUIDValue(),
				OrganisationID: organisationId,
				Attributes: &models.BankIDAttributes{
					Country:    "GB",
					BankIDCode: "GBDSC",
					BankID:     bankId,
				},
			},
		}))

	defer func() {
		if createdBankID != nil {
			if _, err := auth.AccountClient.Accounts.DeleteBankidsID(accounts.NewDeleteBankidsIDParams().WithID(createdBankID.Payload.Data.ID)); err != nil {
				log.Printf("[WARN] Did not delete bank id, error %s\n", err.Error())
			} else {
				log.Printf("[INFO] Successfully deleted bank id\n")
			}
		}
	}()

	createdBicID, _ := auth.AccountClient.Accounts.PostBics(accounts.NewPostBicsParams().
		WithBicCreationRequest(&models.BicCreation{
			Data: &models.Bic{
				Type:           "bankids",
				ID:             NewUUIDValue(),
				OrganisationID: organisationId,
				Attributes: &models.BicAttributes{
					Bic: "NWABCD12",
				},
			},
		}))

	defer func() {
		if createdBicID != nil {

			if _, err := auth.AccountClient.Accounts.DeleteBicsID(accounts.NewDeleteBicsIDParams().WithID(createdBicID.Payload.Data.ID)); err != nil {
				log.Printf("[WARN] Did not delete bic id, error %s\n", err.Error())
			} else {
				log.Printf("[INFO] Successfully deleted bic\n")
			}
		}
	}()

	// if no  account is found, then we create it
	// and we obtain its uuid and version. Otherwise, we use the values
	// from the  account

	var accountID *strfmt.UUID
	var accountVersion int64

	if len(existingAccounts.Payload.Data) > 0 {

		account := existingAccounts.Payload.Data[0]
		accountID = account.ID
		accountVersion = *account.Version

		log.Printf("[DEBUG] Existing account found with id=%v, version=%v", accountID, accountVersion)
	} else {

		log.Printf("[DEBUG] Creating new account ")
		createResponse, err := auth.AccountClient.Accounts.PostAccounts(accounts.NewPostAccountsParams().
			WithAccountCreationRequest(&models.AccountCreation{
				Data: &models.Account{
					OrganisationID: UUID(organisationId),
					Type:           "accounts",
					ID:             NewUUID(),
					Attributes: &models.AccountAttributes{
						AccountNumber: accountNumber,
						BankID:        bankId,
						Bic:           "NWABCD12",
						BankIDCode:    "GBDSC",
						Country:       String("GB"),
					},
				},
			}))

		assertNoErrorOccurred(err, t)

		accountID = createResponse.Payload.Data.ID
		accountVersion = 1
	}

	_, err = auth.AccountClient.Accounts.DeleteAccountsID(accounts.NewDeleteAccountsIDParams().
		WithID(UUIDValue(accountID)).WithVersion(accountVersion),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AccountClient.Organisations.GetUnitsID(organisations.NewGetUnitsIDParams().
		WithID(UUIDValue(accountID)))

	assertStatusCode(err, t, 404)
}
