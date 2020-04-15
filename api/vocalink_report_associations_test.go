package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

func TestDeleteVocalinkreportAssociation(t *testing.T) {
	id := uuid.New()
	createResponse, err := auth.AssociationClient.Associations.PostVocalinkreport(associations.NewPostVocalinkreportParams().
		WithCreationRequest(&models.VocalinkReportAssociationCreation{
			Data: &models.NewVocalinkReportAssociation{
				ID:             UUIDtoStrFmtUUID(id),
				OrganisationID: &organisationId,
				Relationships:  &models.VocalinkReportAssociationRelationships{},
			},
		}))

	assertNoErrorOccurred(t, err)
	_, err = auth.AssociationClient.Associations.DeleteVocalinkreportID(associations.NewDeleteVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}

func TestGetVocalinkreportForNonExistingAssociation(t *testing.T) {
	randomId := uuid.New()
	randomUUID := strfmt.UUID(randomId.String())

	_, err := auth.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
		WithID(randomUUID))

	assertStatusCode(t, err, 404)
}

func TestGetVocalinkreportAssociation(t *testing.T) {

	id := uuid.New()
	createResponse, err := auth.AssociationClient.Associations.PostVocalinkreport(associations.NewPostVocalinkreportParams().
		WithCreationRequest(&models.VocalinkReportAssociationCreation{
			Data: &models.NewVocalinkReportAssociation{
				ID:             UUIDtoStrFmtUUID(id),
				OrganisationID: &organisationId,
				Relationships:  &models.VocalinkReportAssociationRelationships{},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.DeleteVocalinkreportID(associations.NewDeleteVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)
}

func TestPostVocalinkreportAssociation(t *testing.T) {

	id := uuid.New()
	createResponse, err := auth.AssociationClient.Associations.PostVocalinkreport(associations.NewPostVocalinkreportParams().
		WithCreationRequest(&models.VocalinkReportAssociationCreation{
			Data: &models.NewVocalinkReportAssociation{
				ID:             UUIDtoStrFmtUUID(id),
				OrganisationID: &organisationId,
				Relationships:  &models.VocalinkReportAssociationRelationships{},
				Attributes: &models.VocalinkReportAssociationAttributes{
					BacsServiceUserNumber: "123456",
				},
			},
		}))

	assertNoErrorOccurred(t, err)
	actualOrganisationId := createResponse.Payload.Data.OrganisationID.String()
	if actualOrganisationId != organisationId.String() {
		t.Fatalf("Expected %s, got %s", organisationId.String(), actualOrganisationId)
	}

	actualSun := createResponse.Payload.Data.Attributes.BacsServiceUserNumber
	if actualSun != "123456" {
		t.Errorf("Expected SUN %s, got %s", "123456", actualSun)
	}

	_, err = auth.AssociationClient.Associations.DeleteVocalinkreportID(associations.NewDeleteVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

}

func TestGetVocalinkreportAssociationList(t *testing.T) {
	id := NewUUID()
	organisationIdUUID := strfmt.UUID(organisationId.String())

	defer func() {
		_, err := auth.AssociationClient.Associations.DeleteVocalinkreportID(associations.NewDeleteVocalinkreportIDParams().
			WithID(*id).WithVersion(0),
		)
		assertNoErrorOccurred(t, err)
	}()

	_, err := auth.AssociationClient.Associations.PostVocalinkreport(associations.NewPostVocalinkreportParams().
		WithCreationRequest(&models.VocalinkReportAssociationCreation{
			Data: &models.NewVocalinkReportAssociation{
				ID:             id,
				OrganisationID: &organisationId,
				Relationships:  &models.VocalinkReportAssociationRelationships{},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.GetVocalinkreport(associations.NewGetVocalinkreportParams().
		WithFilterOrganisationID([]strfmt.UUID{organisationIdUUID}))

	assertNoErrorOccurred(t, err)
}
