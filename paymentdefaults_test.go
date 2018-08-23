package form3

import (
	"testing"

	"github.com/form3tech-oss/go-form3/client/payment_defaults"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/strfmt"
)

var version = int64(0)

func TestAccPostPaymentDefaults(t *testing.T) {
	createResponse, err := auth.PaymentdefaultsClient.PaymentDefaults.PostPaymentdefaults(payment_defaults.NewPostPaymentdefaultsParams().
		WithDefaultConfiguration(&models.PaymentDefaultsCreate{
			Data: &models.PaymentDefaults{
				OrganisationID: organisationId,
				Version:        &version,
				ID:             strfmt.UUID("7066746d-4c12-4c21-9b9f-ee1a6d401f44"),
				Attributes: &models.PaymentDefaultsAttributes{
					DefaultPaymentScheme: "FPS",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
		WithID(strfmt.UUID(createResponse.Payload.Data.ID)))
	assertNoErrorOccurred(err, t)

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.DeletePaymentdefaultsID(payment_defaults.NewDeletePaymentdefaultsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(err, t)
}

func TestAccGetPaymentDefaultsList(t *testing.T) {
	createResponse, err := auth.PaymentdefaultsClient.PaymentDefaults.PostPaymentdefaults(payment_defaults.NewPostPaymentdefaultsParams().
		WithDefaultConfiguration(&models.PaymentDefaultsCreate{
			Data: &models.PaymentDefaults{
				OrganisationID: organisationId,
				Version:        &version,
				ID:             strfmt.UUID("7066746d-4c12-4c21-9b9f-ee1a6d401f44"),
				Attributes: &models.PaymentDefaultsAttributes{
					DefaultPaymentScheme: "FPS",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	response, err := auth.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaults(payment_defaults.NewGetPaymentdefaultsParams())

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.DeletePaymentdefaultsID(payment_defaults.NewDeletePaymentdefaultsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	if len(response.Payload.Data) == 0 {
		t.Error("expected at least one payment")
	}
}

func TestAccDeletePaymentDefaults(t *testing.T) {
	createResponse, err := auth.PaymentdefaultsClient.PaymentDefaults.PostPaymentdefaults(payment_defaults.NewPostPaymentdefaultsParams().
		WithDefaultConfiguration(&models.PaymentDefaultsCreate{
			Data: &models.PaymentDefaults{
				OrganisationID: organisationId,
				Version:        &version,
				ID:             strfmt.UUID("7066746d-4c12-4c21-9b9f-ee1a6d401f44"),
				Attributes: &models.PaymentDefaultsAttributes{
					DefaultPaymentScheme: "FPS",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.DeletePaymentdefaultsID(payment_defaults.NewDeletePaymentdefaultsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	if err != nil {
		t.Error(err)
	}

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func TestAccGetPaymentDefaultsWithIdNotFound(t *testing.T) {
	_, err := auth.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
		WithID(strfmt.UUID("8ea57253-aea2-409b-ab59-e9f0a96adc12")))

	assertStatusCode(err, t, 404)
}
