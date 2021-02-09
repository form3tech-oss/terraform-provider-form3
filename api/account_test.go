package api

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/form3tech-oss/terraform-provider-form3/client/organisations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
)

func bicGenerator(r *rand.Rand) string {
	var (
		chars  = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		digits = []rune("0123456789")
		b      strings.Builder
	)

	for i := 0; i < 6; i++ {
		b.WriteRune(chars[r.Intn(len(chars))])
	}

	for i := 0; i < 5; i++ {
		b.WriteRune(digits[r.Intn(len(digits))])
	}

	return b.String()
}

func bankIdGenerator(r *rand.Rand) string {
	var (
		b      strings.Builder
		digits = []rune("0123456789")
	)

	for i := 0; i < 6; i++ {
		b.WriteRune(digits[r.Intn(len(digits))])
	}

	return b.String()
}

func TestAccDeleteAccount(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	bankIdUUID := NewUUIDValue()
	bic := bicGenerator(r)
	bankID := bankIdGenerator(r)

	defer func() {
		if t.Failed() {
			if _, err := auth.AccountClient.Accounts.DeleteBankidsID(accounts.NewDeleteBankidsIDParams().WithID(bankIdUUID)); err != nil {
				log.Printf("[CLEANUP] Did not delete bank id, error %s\n", JsonErrorPrettyPrint(err))
			} else {
				log.Printf("[CLEANUP] Successfully deleted bank id\n")
			}
		}
	}()

	_, err := auth.AccountClient.Accounts.PostBankids(accounts.NewPostBankidsParams().
		WithBankIDCreationRequest(&models.BankIDCreation{
			Data: &models.BankID{
				Type:           "bankids",
				ID:             bankIdUUID,
				OrganisationID: organisationId,
				Attributes: &models.BankIDAttributes{
					Country:    "GB",
					BankIDCode: "GBDSC",
					BankID:     bankID,
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	bicIdUUID := NewUUIDValue()

	defer func() {
		if t.Failed() {
			if _, err := auth.AccountClient.Accounts.DeleteBicsID(accounts.NewDeleteBicsIDParams().
				WithID(bicIdUUID).
				WithVersion(0),
			); err != nil {
				log.Printf("[CLEANUP] Did not delete bic id, error %s\n", JsonErrorPrettyPrint(err))
			} else {
				log.Printf("[CLEANUP] Successfully deleted bic\n")
			}
		}
	}()

	_, err = auth.AccountClient.Accounts.PostBics(accounts.NewPostBicsParams().
		WithBicCreationRequest(&models.BicCreation{
			Data: &models.Bic{
				Type:           "bankids",
				ID:             bicIdUUID,
				OrganisationID: organisationId,
				Attributes: &models.BicAttributes{
					Bic: bic,
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	accountIdUUID := NewUUID()

	defer func() {
		if t.Failed() {
			if _, err := auth.AccountClient.Accounts.DeleteAccountsID(
				accounts.NewDeleteAccountsIDParams().
					WithID(*accountIdUUID).
					WithVersion(1),
			); err != nil {
				log.Printf("[CLEANUP] Did not delete account id, error %s\n", JsonErrorPrettyPrint(err))
			} else {
				log.Printf("[CLEANUP] Successfully deleted account\n")
			}
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
					BankID:        bankID,
					Bic:           bic,
					BankIDCode:    "GBDSC",
					Country:       String("GB"),
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	fetchResponse, err := auth.AccountClient.Accounts.GetAccountsID(accounts.NewGetAccountsIDParams().WithID(UUIDValue(createResponse.Payload.Data.ID)))
	assertNoErrorOccurred(t, err)

	_, err = auth.AccountClient.Accounts.DeleteAccountsID(accounts.NewDeleteAccountsIDParams().
		WithID(UUIDValue(createResponse.Payload.Data.ID)).WithVersion(*fetchResponse.Payload.Data.Version),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.AccountClient.Organisations.GetUnitsID(organisations.NewGetUnitsIDParams().
		WithID(UUIDValue(createResponse.Payload.Data.ID)))

	assertStatusCode(t, err, 404)
}
