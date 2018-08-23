package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3/client/associations"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"testing"
)

func TestDeleteBacsAssociation(t *testing.T) {

	serviceUserNumber := "123458"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(1)

	id, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(organisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(err, t)
	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func TestGetBacsForNonExistingAssociation(t *testing.T) {
	randomId, _ := uuid.NewV4()
	randomUUID := strfmt.UUID(randomId.String())

	_, err := auth.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(randomUUID))

	assertStatusCode(err, t, 404)
}

func TestGetBacsAssociation(t *testing.T) {
	serviceUserNumber := "987892"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(1)

	id, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(organisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	bacsAssociation, err := auth.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	actualServiceUserNumber := bacsAssociation.Payload.Data.Attributes.ServiceUserNumber
	actualOrganisationId := bacsAssociation.Payload.Data.OrganisationID
	if actualServiceUserNumber != serviceUserNumber {
		t.Fatalf("Expected %s, got %s", serviceUserNumber, actualServiceUserNumber)
	}

	if actualOrganisationId != organisationId {
		t.Fatalf("Expected %s, got %s", organisationId, actualOrganisationId)
	}

	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)
}

func TestPostBacsAssociation(t *testing.T) {
	serviceUserNumber := "987897"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(1)

	id, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(organisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(err, t)
	actualOrganisationId := createResponse.Payload.Data.OrganisationID.String()
	actualServiceUserNumber := createResponse.Payload.Data.Attributes.ServiceUserNumber
	if actualOrganisationId != organisationId.String() {
		t.Fatalf("Expected %s, got %s", organisationId.String(), actualOrganisationId)
	}

	if actualServiceUserNumber != serviceUserNumber {
		t.Fatalf("Expected %s, got %s", serviceUserNumber, actualServiceUserNumber)
	}
	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

}

func TestPostBacsAssociation_DoNotIgnoreAccountTypeWhenValueIsZero(t *testing.T) {
	serviceUserNumber := "987897"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(0)

	id, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(organisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(err, t)
	actualAccountType := createResponse.Payload.Data.Attributes.AccountType
	if *actualAccountType != accountType {
		t.Fatalf("Expected %v, got %v", accountType, *actualAccountType)
	}

	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)
}

func TestGetBacsAssociationList(t *testing.T) {
	id, _ := uuid.NewV4()
	organisationIdUUID := strfmt.UUID(organisationId.String())
	serviceUserNumber := "123456"
	accountNumber := "12345678"
	sortingCode := "123456"
	accountType := int64(1)

	createResponse, err := auth.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: &models.BacsNewAssociation{
				ID:             strfmt.UUID(id.String()),
				OrganisationID: strfmt.UUID(organisationId.String()),
				Attributes: &models.BacsAssociationAttributes{
					ServiceUserNumber: serviceUserNumber,
					AccountNumber:     accountNumber,
					SortingCode:       sortingCode,
					AccountType:       &accountType,
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	getBacsResponse, err := auth.AssociationClient.Associations.GetBacs(associations.NewGetBacsParams().
		WithFilterOrganisationID(&organisationIdUUID))

	fmt.Println(getBacsResponse)
	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

}
