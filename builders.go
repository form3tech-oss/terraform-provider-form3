package form3

import (
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"time"
	"math/rand"
	"math"
	"fmt"
)

type PaymentBuilder struct {
	payment models.Payment
}

func (b *PaymentBuilder) WithDefaults() *PaymentBuilder {
	id, _ := uuid.NewV4()
	organisationId, _ := uuid.NewV4()
	b.payment = models.Payment{
		ID:             ConvertUUIDtoStrFmtUUID(id),
		OrganisationID: ConvertUUIDtoStrFmtUUID(organisationId),
		Attributes: &models.PaymentAttributes{
			Amount: "60.00",
			BeneficiaryParty: &models.PaymentAttributesBeneficiaryParty{
				AccountNumber:     "12345678",
				AccountNumberCode: "BBAN",
				AccountWith: &models.AccountHoldingEntity{
					BankID:     "888888",
					BankIDCode: "GBDSC",
				},
			},
			Currency: "GBP",
			DebtorParty: &models.PaymentAttributesDebtorParty{
				AccountNumber:     "87654321",
				AccountNumberCode: "BBAN",
				AccountWith: &models.AccountHoldingEntity{
					BankID:     "333333",
					BankIDCode: "GBDSC",
				},
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
	b.payment.OrganisationID = ConvertUUIDtoPointer(id)
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

func (b *PaymentBuilder) WithSchemeTransactionID(schemeTransactionID string) *PaymentBuilder {
	b.payment.Attributes.SchemeTransactionID = schemeTransactionID
	return b
}

func (b *PaymentBuilder) NewSchemeTransactionID() string {
	uniqueId := math.Mod(float64(rand.Int63n(math.MaxInt64)), 100000000000000000)
	return fmt.Sprintf("%17.0f", uniqueId)
}

func (b *PaymentBuilder) NewMessageID() string {
	timestamp := time.Now().UTC().Format("20060102150405")
	uniqueId := b.NewSchemeTransactionID()
	return timestamp + fmt.Sprintf("%17.0f", uniqueId)
}

func (b *PaymentBuilder) Build() *models.Payment {
	return &b.payment
}
