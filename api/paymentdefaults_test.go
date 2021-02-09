package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/payment_defaults"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

var version = int64(0)

func TestAccPostPaymentDefaults(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	paymentId := uuid.New().String()

	createResponse, err := auth.PaymentdefaultsClient.PaymentDefaults.PostPaymentdefaults(payment_defaults.NewPostPaymentdefaultsParams().
		WithDefaultConfiguration(&models.PaymentDefaultsCreate{
			Data: &models.PaymentDefaults{
				OrganisationID: organisationId,
				Version:        &version,
				ID:             strfmt.UUID(paymentId),
				Attributes: &models.PaymentDefaultsAttributes{
					DefaultPaymentScheme: "FPS",
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
		WithID(strfmt.UUID(createResponse.Payload.Data.ID)))
	assertNoErrorOccurred(t, err)

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.DeletePaymentdefaultsID(payment_defaults.NewDeletePaymentdefaultsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(t, err)
}

func TestAccGetPaymentDefaultsList(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	paymentId := uuid.New().String()

	createResponse, err := auth.PaymentdefaultsClient.PaymentDefaults.PostPaymentdefaults(payment_defaults.NewPostPaymentdefaultsParams().
		WithDefaultConfiguration(&models.PaymentDefaultsCreate{
			Data: &models.PaymentDefaults{
				OrganisationID: organisationId,
				Version:        &version,
				ID:             strfmt.UUID(paymentId),
				Attributes: &models.PaymentDefaultsAttributes{
					DefaultPaymentScheme: "FPS",
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	response, err := auth.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaults(payment_defaults.NewGetPaymentdefaultsParams())
	assertNoErrorOccurred(t, err)

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.DeletePaymentdefaultsID(payment_defaults.NewDeletePaymentdefaultsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

	if len(response.Payload.Data) == 0 {
		t.Error("expected at least one payment")
	}
}

func TestAccDeletePaymentDefaults(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	paymentId := uuid.New().String()

	createResponse, err := auth.PaymentdefaultsClient.PaymentDefaults.PostPaymentdefaults(payment_defaults.NewPostPaymentdefaultsParams().
		WithDefaultConfiguration(&models.PaymentDefaultsCreate{
			Data: &models.PaymentDefaults{
				OrganisationID: organisationId,
				Version:        &version,
				ID:             strfmt.UUID(paymentId),
				Attributes: &models.PaymentDefaultsAttributes{
					DefaultPaymentScheme: "FPS",
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.DeletePaymentdefaultsID(payment_defaults.NewDeletePaymentdefaultsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	if err != nil {
		t.Error(err)
	}

	_, err = auth.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}

func TestAccGetPaymentDefaultsWithIdNotFound(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	_, err := auth.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
		WithID("8ea57253-aea2-409b-ab59-e9f0a96adc12"))

	assertStatusCode(t, err, 404)
}
