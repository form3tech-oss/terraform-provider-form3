package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

func TestDeleteBacsAssociation(t *testing.T) {

	serviceUserNumber := "123458"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(1)

	id := uuid.New()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(testOrganisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(t, err)
	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}

func TestGetBacsForNonExistingAssociation(t *testing.T) {
	randomId := uuid.New()
	randomUUID := strfmt.UUID(randomId.String())

	_, err := auth.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(randomUUID))

	assertStatusCode(t, err, 404)
}

func TestGetBacsAssociation(t *testing.T) {
	serviceUserNumber := "987892"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(1)

	id := uuid.New()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(testOrganisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	bacsAssociation, err := auth.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

	actualServiceUserNumber := bacsAssociation.Payload.Data.Attributes.ServiceUserNumber
	actualOrganisationId := bacsAssociation.Payload.Data.OrganisationID
	if actualServiceUserNumber != serviceUserNumber {
		t.Fatalf("Expected %s, got %s", serviceUserNumber, actualServiceUserNumber)
	}

	if actualOrganisationId != testOrganisationId {
		t.Fatalf("Expected %s, got %s", testOrganisationId, actualOrganisationId)
	}

	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)
}

func TestPostBacsAssociation(t *testing.T) {
	serviceUserNumber := "987897"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(1)

	id := uuid.New()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(testOrganisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(t, err)
	actualOrganisationId := createResponse.Payload.Data.OrganisationID.String()
	actualServiceUserNumber := createResponse.Payload.Data.Attributes.ServiceUserNumber
	if actualOrganisationId != testOrganisationId.String() {
		t.Fatalf("Expected %s, got %s", testOrganisationId.String(), actualOrganisationId)
	}

	if actualServiceUserNumber != serviceUserNumber {
		t.Fatalf("Expected %s, got %s", serviceUserNumber, actualServiceUserNumber)
	}
	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

}

func TestPostBacsAssociationIncludingOptionalServiceCentre(t *testing.T) {
	serviceUserNumber := "987897"
	accountNumber := "12345699"
	sortingCode := "123456"
	accountType := int64(1)

	id := uuid.New()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(testOrganisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
					BankCode:          "1234",
					CentreNumber:      "99",
				},
			},
		}))

	assertNoErrorOccurred(t, err)
	actualBankCode := createResponse.Payload.Data.Attributes.BankCode
	actualCentreNumber := createResponse.Payload.Data.Attributes.CentreNumber
	if actualBankCode != "1234" {
		t.Fatalf("Expected %s, got %s", "1234", actualBankCode)
	}

	if actualCentreNumber != "99" {
		t.Fatalf("Expected %s, got %s", "99", actualCentreNumber)
	}
	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

}

func TestPostBacsAssociation_DoNotIgnoreAccountTypeWhenValueIsZero(t *testing.T) {
	serviceUserNumber := "987897"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(0)

	id := uuid.New()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(testOrganisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(t, err)
	actualAccountType := createResponse.Payload.Data.Attributes.AccountType
	if *actualAccountType != accountType {
		t.Fatalf("Expected %v, got %v", accountType, *actualAccountType)
	}

	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)
}

func TestGetBacsAssociationList(t *testing.T) {
	id := uuid.New()
	organisationIdUUID := strfmt.UUID(testOrganisationId.String())
	serviceUserNumber := "123456"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(1)

	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(testOrganisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.GetBacs(associations.NewGetBacsParams().
		WithFilterOrganisationID(&organisationIdUUID))

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)
}
