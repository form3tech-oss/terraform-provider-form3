package form3

import (
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"time"
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
			SchemeTransactionID: "12345",
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

func (b *PaymentBuilder) WithPaymentType(paymentType string) *PaymentBuilder {
	b.payment.Attributes.PaymentType = paymentType
	return b
}

func (b *PaymentBuilder) Build() *models.Payment {
	return &b.payment
}
