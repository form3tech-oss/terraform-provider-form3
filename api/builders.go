package api

import (
	"fmt"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"math"
	"math/rand"
	"time"
)

type PaymentBuilder struct {
	payment models.Payment
}

func (b *PaymentBuilder) WithDefaults() *PaymentBuilder {
	id, _ := uuid.NewV4()
	organisationId, _ := uuid.NewV4()
	b.payment = models.Payment{
		ID:             UUIDtoStrFmtUUID(id),
		OrganisationID: UUIDtoStrFmtUUID(organisationId),
		Attributes: &models.PaymentAttributes{
			Amount: "60.00",
			BeneficiaryParty: &models.PaymentAttributesBeneficiaryParty{
				AccountNumber:     "12345678",
				AccountNumberCode: "BBAN",
				AccountWith: &models.BeneficiaryDebtorAccountHoldingEntity{
					BankID:     "888888",
					BankIDCode: "GBDSC",
				},
				Country: "GB",
			},
			Currency: "GBP",
			DebtorParty: &models.PaymentAttributesDebtorParty{
				AccountNumber:     "87654321",
				AccountNumberCode: "BBAN",
				AccountWith: &models.BeneficiaryDebtorAccountHoldingEntity{
					BankID:     "333333",
					BankIDCode: "GBDSC",
				},
				Country: "GB",
			},
			EndToEndReference:   "00151519632ZCBBBJQ",
			SchemeTransactionID: b.NewSchemeTransactionID(),
			ProcessingDate:      strfmt.Date(time.Now()),
			PaymentScheme:       "Bacs",
			SchemePaymentType:   "TelephoneBanking",
			PaymentType:         "Dividend",
		},
	}

	return b
}

func (b *PaymentBuilder) WithAmount(amount string) *PaymentBuilder {
	b.payment.Attributes.Amount = amount
	return b
}

func (b *PaymentBuilder) WithOrganisationID(id strfmt.UUID) *PaymentBuilder {
	b.payment.OrganisationID = UUID(id)
	return b
}

func (b *PaymentBuilder) WithPaymentScheme(scheme string) *PaymentBuilder {
	b.payment.Attributes.PaymentScheme = scheme
	return b
}

func (b *PaymentBuilder) WithSchemePaymentType(schemePaymentType string) *PaymentBuilder {
	b.payment.Attributes.SchemePaymentType = schemePaymentType
	return b
}

func (b *PaymentBuilder) WithDebtorPartyAccountWithBankID(bankID string) *PaymentBuilder {
	b.payment.Attributes.DebtorParty.AccountWith.BankID = bankID
	return b
}

func (b *PaymentBuilder) WithDebtorPartyAccountNumber(accountNumber string) *PaymentBuilder {
	b.payment.Attributes.DebtorParty.AccountNumber = accountNumber
	return b
}

func (b *PaymentBuilder) WithBeneficiaryPartyAccountWithBankID(bankID string) *PaymentBuilder {
	b.payment.Attributes.BeneficiaryParty.AccountWith.BankID = bankID
	return b
}

func (b *PaymentBuilder) WithBeneficiaryPartyAccountNumber(accountNumber string) *PaymentBuilder {
	b.payment.Attributes.BeneficiaryParty.AccountNumber = accountNumber
	return b
}

func (b *PaymentBuilder) WithSchemeTransactionID(schemeTransactionID string) *PaymentBuilder {
	b.payment.Attributes.SchemeTransactionID = schemeTransactionID
	return b
}

func (b *PaymentBuilder) WithReference(reference string) *PaymentBuilder {
	b.payment.Attributes.Reference = reference
	return b
}

func (b *PaymentBuilder) NewSchemeTransactionID() string {
	uniqueId := float64(rand.Int63n(math.MaxInt64 / 100))
	return fmt.Sprintf("%17.0f", uniqueId)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (b *PaymentBuilder) NewMessageID() string {
	timestamp := time.Now().UTC().Format("20060102150405")
	uniqueId := b.NewSchemeTransactionID()
	return timestamp + uniqueId
}

func (b *PaymentBuilder) Build() *models.Payment {
	return &b.payment
}
