package api

import (
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/google/uuid"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
)

func TestAccPostMandateManagement(t *testing.T) {

	randomUid := uuid.New()

	createResponse, err := auth.AssociationClient.Associations.PostMandatemanagement(associations.NewPostMandatemanagementParams().
		WithMandateManagementCreationRequest(&models.MandateManagementAssociationCreation{
			Data: &models.MandateManagementAssociation{
				OrganisationID: testOrganisationId,
				ID:             strfmt.UUID(randomUid.String()),
				Attributes: &models.MandateManagementAssociationAttributes{
					PaymentScheme: models.PaymentSchemeBACS,
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.GetMandatemanagementID(associations.NewGetMandatemanagementIDParams().WithID(createResponse.Payload.Data.ID))
	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.DeleteMandatemanagementID(associations.NewDeleteMandatemanagementIDParams().
		WithID(createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(t, err)
}
