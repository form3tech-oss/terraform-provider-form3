package api

import (
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteConfirmationOfPayeeAssociation(t *testing.T) {
	createResponse := createAssociation(t)

	deleteAssociation(t, createResponse)

	_, err := auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(*createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func deleteAssociation(t *testing.T, createResponse *associations.PostConfirmationOfPayeeCreated) {
	_, err := auth.AssociationClient.Associations.DeleteConfirmationOfPayeeID(associations.NewDeleteConfirmationOfPayeeIDParams().
		WithID(*createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(err, t)
}

func createAssociation(t *testing.T) *associations.PostConfirmationOfPayeeCreated {
	id, _ := uuid.NewV4()
	keyId, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostConfirmationOfPayee(associations.NewPostConfirmationOfPayeeParams().
		WithCreationRequest(&models.CoPAssociationCreation{
			Data: &models.CoPAssociation{
				ID:             UUIDtoStrFmtUUID(id),
				OrganisationID: &testOrganisationId,
				Attributes: &models.CoPAssociationAttributes{
					OpenBankingOrganisationID: swag.String("1234"),
					PublicKeyID:               swag.String("5678"),
					MatchingCriteria: &models.MatchingCriteria{
						ExactMatchThreshold: swag.String("0.85"),
						CloseMatchThreshold: swag.String("0.65"),
					},
				},
				Relationships: &models.CoPAssociationRelationships{
					SigningCertificate: &models.SigningCertificate{
						Data: &models.SigningCertificateData{
							KeyID: UUIDtoStrFmtUUID(keyId),
							Dn:    swag.String("test"),
						},
					},
				},
			},
		}))
	assertNoErrorOccurred(err, t)
	return createResponse
}

func TestGetConfirmationOfPayeeForNonExistingAssociation(t *testing.T) {
	randomId, _ := uuid.NewV4()
	randomUUID := strfmt.UUID(randomId.String())

	_, err := auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(randomUUID))

	assertStatusCode(err, t, 404)
}

func TestGetConfirmationOfPayeeAssociation(t *testing.T) {
	createResponse := createAssociation(t)
	defer deleteAssociation(t, createResponse)

	_, err := auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(*createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)
}

func TestPostConfirmationOfPayeeAssociation(t *testing.T) {
	createResponse := createAssociation(t)
	defer deleteAssociation(t, createResponse)

	actualOrganisationId := createResponse.Payload.Data.OrganisationID.String()
	if actualOrganisationId != testOrganisationId.String() {
		t.Fatalf("Expected %s, got %s", testOrganisationId.String(), actualOrganisationId)
	}

	assert.Equal(t, "1234", *createResponse.Payload.Data.Attributes.OpenBankingOrganisationID)
	assert.Equal(t, "5678", *createResponse.Payload.Data.Attributes.PublicKeyID)
}

func TestGetConfirmationOfPayeeAssociationList(t *testing.T) {
	createResponse := createAssociation(t)
	defer deleteAssociation(t, createResponse)

	getConfirmationOfPayeeResponse, err := auth.AssociationClient.Associations.GetConfirmationOfPayee(associations.NewGetConfirmationOfPayeeParams().
		WithFilterOrganisationID([]strfmt.UUID{*createResponse.Payload.Data.OrganisationID}))

	assertNoErrorOccurred(err, t)
	assert.True(t, len(getConfirmationOfPayeeResponse.Payload.Data) > 0)

	found := false
	for _, assoc := range getConfirmationOfPayeeResponse.Payload.Data {
		if *assoc.ID == *createResponse.Payload.Data.ID {
			found = true
			assert.Equal(t, "1234", *assoc.Attributes.OpenBankingOrganisationID)
			assert.Equal(t, *createResponse.Payload.Data.Relationships.SigningCertificate.Data.KeyID, *assoc.Relationships.SigningCertificate.Data.KeyID)
		}
	}
	assert.True(t, found)
}
