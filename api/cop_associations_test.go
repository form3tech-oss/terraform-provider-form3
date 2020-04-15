package api

import (
	"testing"

	"github.com/google/uuid"

	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
)

func TestDeleteConfirmationOfPayeeAssociation(t *testing.T) {
	createResponse := createAssociation(t)

	deleteAssociation(t, createResponse)

	_, err := auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(*createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}

func deleteAssociation(t *testing.T, createResponse *associations.PostConfirmationOfPayeeCreated) {
	_, err := auth.AssociationClient.Associations.DeleteConfirmationOfPayeeID(associations.NewDeleteConfirmationOfPayeeIDParams().
		WithID(*createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(t, err)
}

func createAssociation(t *testing.T) *associations.PostConfirmationOfPayeeCreated {
	id := uuid.New()
	keyId := uuid.New()
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
	assertNoErrorOccurred(t, err)
	return createResponse
}

func TestGetConfirmationOfPayeeForNonExistingAssociation(t *testing.T) {
	randomId := uuid.New()
	randomUUID := strfmt.UUID(randomId.String())

	_, err := auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(randomUUID))

	assertStatusCode(t, err, 404)
}

func TestGetConfirmationOfPayeeAssociation(t *testing.T) {
	createResponse := createAssociation(t)
	defer deleteAssociation(t, createResponse)

	_, err := auth.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(*createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)
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

	assertNoErrorOccurred(t, err)
	assert.True(t, len(getConfirmationOfPayeeResponse.Payload.Data) > 0)

	found := false
	for _, assoc := range getConfirmationOfPayeeResponse.Payload.Data {
		if *assoc.ID == *createResponse.Payload.Data.ID {
			found = true
			assert.Equal(t, *createResponse.Payload.Data.Relationships.SigningCertificate.Data.KeyID, *assoc.Relationships.SigningCertificate.Data.KeyID)
		}
	}
	assert.True(t, found)
}
