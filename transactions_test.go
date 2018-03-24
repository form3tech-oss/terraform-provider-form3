package form3

import (
	"github.com/ewilde/go-form3/client/payments"
	"github.com/ewilde/go-form3/models"
	"testing"
)

func TestAccPostTransactionsPayment(t *testing.T) {
	payment := (&PaymentBuilder{}).
		WithDefaults().
		WithAmount("60.00").
		WithOrganisationID(testOrganisationId).
		Build()

	_, err := auth.TransactionClient.Payments.
		PostPayments(payments.NewPostPaymentsParams().
			WithPaymentCreationRequest(&models.PaymentCreation{
				Data: payment,
			}))

	assertNoErrorOccurred(err, t)
}

func TestAccPostTransactionsPaymentSubmission(t *testing.T) {
	payment := (&PaymentBuilder{}).
		WithDefaults().
		WithAmount("60.00").
		WithOrganisationID(testOrganisationId).
		Build()

	paymentResponse, err := auth.TransactionClient.Payments.
		PostPayments(payments.NewPostPaymentsParams().
			WithPaymentCreationRequest(&models.PaymentCreation{
				Data: payment,
			}))

	assertNoErrorOccurred(err, t)

	_, err = auth.TransactionClient.Payments.PostPaymentsIDSubmissions(
		payments.NewPostPaymentsIDSubmissionsParams().
			WithID(*paymentResponse.Payload.Data.ID).
			WithSubmissionCreationRequest(
				&models.PaymentSubmissionCreation{
					Data: &models.NewPaymentSubmission{
						ID:             NewStrFmtUUID(),
						OrganisationID: &testOrganisationId,
						Type:           "payment_submissions",
					},
				},
			),
	)

	assertNoErrorOccurred(err, t)
}
