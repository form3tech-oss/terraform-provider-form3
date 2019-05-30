package api

import (
	"github.com/form3tech-oss/go-form3/client/accounts"
	"github.com/form3tech-oss/go-form3/client/associations"
	"github.com/form3tech-oss/go-form3/client/organisations"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"testing"
)

func TestAccDeleteOrganisation(t *testing.T) {

	createResponse, err := auth.OrganisationClient.Organisations.PostUnits(organisations.NewPostUnitsParams().
		WithOrganisationCreationRequest(&models.OrganisationCreation{
			Data: &models.Organisation{
				OrganisationID: organisationId,
				Type:           "organisations",
				ID:             strfmt.UUID("58a78c22-efa6-4f67-b2ec-30c53fd9a438"),
				Attributes: &models.OrganisationAttributes{
					Name: "TestOrganisation",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.OrganisationClient.Organisations.DeleteUnitsID(organisations.NewDeleteUnitsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.OrganisationClient.Organisations.GetUnitsID(organisations.NewGetUnitsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func TestAccDeleteOrganisationAssociation(t *testing.T) {
	assocId, _ := uuid.NewV4()
	createResponse, err := auth.AssociationClient.Associations.PostStarling(associations.NewPostStarlingParams().
		WithCreationRequest(&models.AssociationCreation{
			Data: &models.NewAssociation{
				OrganisationID: testOrganisationId,
				Type:           "starling_associations",
				ID:             strfmt.UUID(assocId.String()),
				Attributes: &models.NewAssociationAttributes{
					StarlingAccountName: "TestStarlingAccountName",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.DeleteStarlingID(associations.NewDeleteStarlingIDParams().
		WithID(createResponse.Payload.Data.ID).
		WithVersion(*createResponse.Payload.Data.Version),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.GetStarlingID(associations.NewGetStarlingIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func TestAccDeleteBankids(t *testing.T) {

	createResponse, err := auth.AccountClient.Accounts.PostBankids(accounts.NewPostBankidsParams().
		WithBankIDCreationRequest(&models.BankIDCreation{
			Data: &models.BankID{
				OrganisationID: testOrganisationId,
				Type:           "bankids",
				ID:             strfmt.UUID("8a2f6b61-ac5a-4f8e-b578-e4da08a36dc6"),
				Attributes: &models.BankIDAttributes{
					BankID:     "400301",
					BankIDCode: "GBDSC",
					Country:    "GB",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.AccountClient.Accounts.DeleteBankidsID(accounts.NewDeleteBankidsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AccountClient.Accounts.GetBankidsID(accounts.NewGetBankidsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}

func TestAccDeleteBics(t *testing.T) {

	createResponse, err := auth.AccountClient.Accounts.PostBics(accounts.NewPostBicsParams().
		WithBicCreationRequest(&models.BicCreation{
			Data: &models.Bic{
				OrganisationID: testOrganisationId,
				Type:           "bics",
				ID:             strfmt.UUID("2f8f3856-a318-4d49-8162-d65a337a74fd"),
				Attributes: &models.BicAttributes{
					Bic: "NWBKGB41",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.AccountClient.Accounts.DeleteBicsID(accounts.NewDeleteBicsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AccountClient.Accounts.GetBicsID(accounts.NewGetBicsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}
