package api

import (
	"log"

	"github.com/form3tech-oss/go-form3/client/accounts"
	"github.com/form3tech-oss/go-form3/client/organisations"
	"github.com/form3tech-oss/go-form3/models"
	"testing"
)

func TestAccDeleteAccount(t *testing.T) {

	createdBankID, err := auth.AccountClient.Accounts.PostBankids(accounts.NewPostBankidsParams().
		WithBankIDCreationRequest(&models.BankIDCreation{
			Data: &models.BankID{
				Type:           "bankids",
				ID:             NewUUIDValue(),
				OrganisationID: organisationId,
				Attributes: &models.BankIDAttributes{
					Country:    "GB",
					BankIDCode: "GBDSC",
					BankID:     "100202",
				},
			},
		}))

	defer func() {
		if _, err := auth.AccountClient.Accounts.DeleteBankidsID(accounts.NewDeleteBankidsIDParams().WithID(createdBankID.Payload.Data.ID)); err != nil {
			log.Printf("[WARN] Did not delete bank id, error %s\n", err.Error())
		} else {
			log.Printf("[INFO] Successfully deleted bank id\n")
		}
	}()

	createdBicID, err := auth.AccountClient.Accounts.PostBics(accounts.NewPostBicsParams().
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
		if _, err := auth.AccountClient.Accounts.DeleteBicsID(accounts.NewDeleteBicsIDParams().WithID(createdBicID.Payload.Data.ID)); err != nil {
			log.Printf("[WARN] Did not delete bic id, error %s\n", err.Error())
		} else {
			log.Printf("[INFO] Successfully deleted bic\n")
		}
	}()

	createResponse, err := auth.AccountClient.Accounts.PostAccounts(accounts.NewPostAccountsParams().
		WithAccountCreationRequest(&models.AccountCreation{
			Data: &models.Account{
				OrganisationID: UUID(organisationId),
				Type:           "accounts",
				ID:             NewUUID(),
				Attributes: &models.AccountAttributes{
					AccountNumber: "12345678",
					BankID:        "100202",
					Bic:           "NWABCD12",
					BankIDCode:    "GBDSC",
					Country:       String("GB"),
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.AccountClient.Accounts.DeleteAccountsID(accounts.NewDeleteAccountsIDParams().
		WithID(UUIDValue(createResponse.Payload.Data.ID)),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AccountClient.Organisations.GetUnitsID(organisations.NewGetUnitsIDParams().
		WithID(UUIDValue(createResponse.Payload.Data.ID)))

	assertStatusCode(err, t, 404)
}
