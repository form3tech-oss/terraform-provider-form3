package api

import (
	"github.com/google/uuid"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/mandates"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
)

func TestAccPostMandateManagement(t *testing.T) {

	randomUid := uuid.New()

	createResponse, err := auth.AssociationClient.Mandates.PostMandatemanagement(mandates.NewPostMandatemanagementParams().
		WithMandateManagementCreationRequest(&models.MandateManagementCreation{
			Data: &models.MandateManagement{
				OrganisationID: testOrganisationId,
				ID: strfmt.UUID(randomUid.String()),
				Attributes: &models.MandateManagementAttributes{
					PaymentScheme: models.PaymentSchemeBACS,
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Mandates.GetMandatemanagementID(mandates.NewGetMandatemanagementIDParams().WithID(createResponse.Payload.Data.ID))
	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Mandates.DeleteMandatemanagementID(mandates.NewDeleteMandatemanagementIDParams().
		WithID(createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(t, err)
}
