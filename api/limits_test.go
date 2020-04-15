package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/limits"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
)

func TestAccPostLimit(t *testing.T) {

	createResponse, err := auth.LimitsClient.Limits.PostLimits(limits.NewPostLimitsParams().
		WithLimitCreationRequest(&models.LimitCreation{
			Data: &models.Limit{
				OrganisationID: organisationId,
				Type:           "limits",
				ID:             strfmt.UUID("7066746d-4c12-4c21-9b9f-ee1a6d401f44"),
				Attributes: &models.LimitAttributes{
					Amount:              "1000",
					Gateway:             "payport-interface",
					Scheme:              "FPS",
					SettlementCycleType: models.SettlementCycleTypePerScheme,
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.LimitsClient.Limits.GetLimitsID(limits.NewGetLimitsIDParams().WithID(strfmt.UUID(createResponse.Payload.Data.ID)))
	assertNoErrorOccurred(err, t)

	_, err = auth.LimitsClient.Limits.DeleteLimitsID(limits.NewDeleteLimitsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(err, t)

}

func TestAccGetLimits(t *testing.T) {

	createResponse, err := auth.LimitsClient.Limits.PostLimits(limits.NewPostLimitsParams().
		WithLimitCreationRequest(&models.LimitCreation{
			Data: &models.Limit{
				OrganisationID: organisationId,
				Type:           "limits",
				ID:             strfmt.UUID("7066746d-4c12-4c21-9b9f-ee1a6d401f44"),
				Attributes: &models.LimitAttributes{
					Amount:              "1000",
					Gateway:             "payport-interface",
					Scheme:              "FPS",
					SettlementCycleType: models.SettlementCycleTypePerScheme,
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	response, err := auth.LimitsClient.Limits.GetLimits(limits.NewGetLimitsParams())

	_, err = auth.LimitsClient.Limits.DeleteLimitsID(limits.NewDeleteLimitsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(err, t)

	if len(response.Payload.Data) == 0 {
		t.Error("Expected at least one limit")
	}

}

func TestAccDeleteLimit(t *testing.T) {

	createResponse, err := auth.LimitsClient.Limits.PostLimits(limits.NewPostLimitsParams().
		WithLimitCreationRequest(&models.LimitCreation{
			Data: &models.Limit{
				OrganisationID: organisationId,
				Type:           "limits",
				ID:             strfmt.UUID("7066746d-4c12-4c21-9b9f-ee1a6d401f44"),
				Attributes: &models.LimitAttributes{
					Amount:              "1000",
					Gateway:             "payport-interface",
					Scheme:              "FPS",
					SettlementCycleType: models.SettlementCycleTypePerScheme,
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.LimitsClient.Limits.DeleteLimitsID(limits.NewDeleteLimitsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	if err != nil {
		t.Error(err)
	}

	_, err = auth.LimitsClient.Limits.GetLimitsID(limits.NewGetLimitsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func TestAccGetLimitWithIdNotFound(t *testing.T) {

	_, err := auth.LimitsClient.Limits.GetLimitsID(limits.NewGetLimitsIDParams().WithID(strfmt.UUID("8ea57253-aea2-409b-ab59-e9f0a96adc12")))

	assertStatusCode(err, t, 404)

}
