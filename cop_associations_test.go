package form3

import (
	"github.com/form3tech-oss/go-form3/client/associations"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteConfirmationOfPayeeAssociation(t *testing.T) {
	id, _ := uuid.NewV4()
	keyId, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostConfirmationOfPayee(associations.NewPostConfirmationOfPayeeParams().
		WithCreationRequest(&models.CoPAssociationCreation{
			Data: &models.CoPAssociation{
				ID:             *UUIDtoStrFmtUUID(id),
				OrganisationID: organisationId,
				Attributes: &models.CoPAssociationAttributes{
					OpenBankingOrganisationID: "1234",
					PublicKeyID:               "5678",
				},
				Relationships: &models.CoPAssociationRelationships{
					SigningCertificate: &models.SigningCertificate{
						Data: &models.SigningCertificateData{
							KeyID: *UUIDtoStrFmtUUID(keyId),
						},
					},
				},
			},
		}))

	assertNoErrorOccurred(err, t)
	_, err = auth.AssociationClient.Associations.DeleteConfirmationOfPayeeID(associations.NewDeleteConfirmationOfPayeeIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func TestGetConfirmationOfPayeeForNonExistingAssociation(t *testing.T) {
	randomId, _ := uuid.NewV4()
	randomUUID := strfmt.UUID(randomId.String())

	_, err := auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(randomUUID))

	assertStatusCode(err, t, 404)
}

func TestGetConfirmationOfPayeeAssociation(t *testing.T) {

	id, _ := uuid.NewV4()
	keyId, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostConfirmationOfPayee(associations.NewPostConfirmationOfPayeeParams().
		WithCreationRequest(&models.CoPAssociationCreation{
			Data: &models.CoPAssociation{
				ID:             *UUIDtoStrFmtUUID(id),
				OrganisationID: organisationId,
				Attributes: &models.CoPAssociationAttributes{
					OpenBankingOrganisationID: "1234",
					PublicKeyID:               "5678",
				},
				Relationships: &models.CoPAssociationRelationships{
					SigningCertificate: &models.SigningCertificate{
						Data: &models.SigningCertificateData{
							KeyID: *UUIDtoStrFmtUUID(keyId),
						},
					},
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.DeleteConfirmationOfPayeeID(associations.NewDeleteConfirmationOfPayeeIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)
}

func TestPostConfirmationOfPayeeAssociation(t *testing.T) {
	id, _ := uuid.NewV4()
	keyId, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostConfirmationOfPayee(associations.NewPostConfirmationOfPayeeParams().
		WithCreationRequest(&models.CoPAssociationCreation{
			Data: &models.CoPAssociation{
				ID:             *UUIDtoStrFmtUUID(id),
				OrganisationID: organisationId,
				Attributes: &models.CoPAssociationAttributes{
					OpenBankingOrganisationID: "1234",
					PublicKeyID:               "5678",
				},
				Relationships: &models.CoPAssociationRelationships{
					SigningCertificate: &models.SigningCertificate{
						Data: &models.SigningCertificateData{
							KeyID: *UUIDtoStrFmtUUID(keyId),
						},
					},
				},
			},
		}))

	assertNoErrorOccurred(err, t)
	actualOrganisationId := createResponse.Payload.Data.OrganisationID.String()
	if actualOrganisationId != organisationId.String() {
		t.Fatalf("Expected %s, got %s", organisationId.String(), actualOrganisationId)
	}

	assert.Equal(t, "1234", createResponse.Payload.Data.Attributes.OpenBankingOrganisationID)
	assert.Equal(t, "5678", createResponse.Payload.Data.Attributes.PublicKeyID)
	assert.Equal(t, *UUIDtoStrFmtUUID(keyId), createResponse.Payload.Data.Relationships.SigningCertificate.Data.KeyID)

	_, err = auth.AssociationClient.Associations.DeleteConfirmationOfPayeeID(associations.NewDeleteConfirmationOfPayeeIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

}

func TestGetConfirmationOfPayeeAssociationList(t *testing.T) {
	id, _ := uuid.NewV4()
	organisationIdUUID := strfmt.UUID(organisationId.String())
	keyId, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostConfirmationOfPayee(associations.NewPostConfirmationOfPayeeParams().
		WithCreationRequest(&models.CoPAssociationCreation{
			Data: &models.CoPAssociation{
				ID:             *UUIDtoStrFmtUUID(id),
				OrganisationID: organisationIdUUID,
				Attributes: &models.CoPAssociationAttributes{
					OpenBankingOrganisationID: "1234",
					PublicKeyID:               "5678",
				},
				Relationships: &models.CoPAssociationRelationships{
					SigningCertificate: &models.SigningCertificate{
						Data: &models.SigningCertificateData{
							KeyID: *UUIDtoStrFmtUUID(keyId),
						},
					},
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	getConfirmationOfPayeeResponse, err := auth.AssociationClient.Associations.GetConfirmationOfPayee(associations.NewGetConfirmationOfPayeeParams().
		WithFilterOrganisationID([]strfmt.UUID{organisationIdUUID}))

	assertNoErrorOccurred(err, t)
	assert.True(t, len(getConfirmationOfPayeeResponse.Payload.Data) > 0)

	found := false
	for _, assoc := range getConfirmationOfPayeeResponse.Payload.Data {
		if assoc.ID == *UUIDtoStrFmtUUID(id) {
			found = true
			assert.Equal(t, assoc.Attributes.OpenBankingOrganisationID, "1234")
			assert.Equal(t, assoc.Relationships.SigningCertificate.Data.KeyID, *UUIDtoStrFmtUUID(keyId))
		}
	}
	assert.True(t, found)

	_, err = auth.AssociationClient.Associations.DeleteConfirmationOfPayeeID(associations.NewDeleteConfirmationOfPayeeIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

}
