package api

import (
	"fmt"
	"github.com/form3tech-oss/go-form3/client/associations"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"testing"
)

func TestDeleteVocalinkreportAssociation(t *testing.T) {
	id, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostVocalinkreport(associations.NewPostVocalinkreportParams().
		WithCreationRequest(&models.VocalinkReportAssociationCreation{
			Data: &models.NewVocalinkReportAssociation{
				ID:             UUIDtoStrFmtUUID(id),
				OrganisationID: &organisationId,
				Relationships:  &models.VocalinkReportAssociationRelationships{},
			},
		}))

	assertNoErrorOccurred(err, t)
	_, err = auth.AssociationClient.Associations.DeleteVocalinkreportID(associations.NewDeleteVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func TestGetVocalinkreportForNonExistingAssociation(t *testing.T) {
	randomId, _ := uuid.NewV4()
	randomUUID := strfmt.UUID(randomId.String())

	_, err := auth.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
		WithID(randomUUID))

	assertStatusCode(err, t, 404)
}

func TestGetVocalinkreportAssociation(t *testing.T) {

	id, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostVocalinkreport(associations.NewPostVocalinkreportParams().
		WithCreationRequest(&models.VocalinkReportAssociationCreation{
			Data: &models.NewVocalinkReportAssociation{
				ID:             UUIDtoStrFmtUUID(id),
				OrganisationID: &organisationId,
				Relationships:  &models.VocalinkReportAssociationRelationships{},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.DeleteVocalinkreportID(associations.NewDeleteVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)
}

func TestPostVocalinkreportAssociation(t *testing.T) {

	id, _ := uuid.NewV4()
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

	assertNoErrorOccurred(err, t)
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

	assertNoErrorOccurred(err, t)

}

func TestGetVocalinkreportAssociationList(t *testing.T) {
	id, _ := uuid.NewV4()
	organisationIdUUID := strfmt.UUID(organisationId.String())

	createResponse, err := auth.AssociationClient.Associations.PostVocalinkreport(associations.NewPostVocalinkreportParams().
		WithCreationRequest(&models.VocalinkReportAssociationCreation{
			Data: &models.NewVocalinkReportAssociation{
				ID:             UUIDtoStrFmtUUID(id),
				OrganisationID: &organisationId,
				Relationships:  &models.VocalinkReportAssociationRelationships{},
			},
		}))

	assertNoErrorOccurred(err, t)

	getVocalinkreportResponse, err := auth.AssociationClient.Associations.GetVocalinkreport(associations.NewGetVocalinkreportParams().
		WithFilterOrganisationID([]strfmt.UUID{organisationIdUUID}))

	fmt.Println(getVocalinkreportResponse)
	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.DeleteVocalinkreportID(associations.NewDeleteVocalinkreportIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

}
