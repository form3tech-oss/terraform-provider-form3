package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/payments"
	"github.com/form3tech-oss/terraform-provider-form3/models"
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

	assertNoErrorOccurred(t, err)
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

	assertNoErrorOccurred(t, err)

	_, err = auth.TransactionClient.Payments.PostPaymentsIDSubmissions(
		payments.NewPostPaymentsIDSubmissionsParams().
			WithID(*paymentResponse.Payload.Data.ID).
			WithSubmissionCreationRequest(
				&models.PaymentSubmissionCreation{
					Data: &models.NewPaymentSubmission{
						ID:             NewUUID(),
						OrganisationID: &testOrganisationId,
						Type:           "payment_submissions",
					},
				},
			),
	)

	assertNoErrorOccurred(t, err)
}
