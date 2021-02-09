package api

import (
	"strings"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/client/organisations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
)

func TestAccDeleteOrganisation(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	id := NewUUID()

	defer func() {
		if t.Failed() {
			_, _ = auth.OrganisationClient.Organisations.DeleteUnitsID(
				organisations.NewDeleteUnitsIDParams().
					WithID(*id).
					WithVersion(0))
		}
	}()

	_, err := auth.OrganisationClient.Organisations.PostUnits(organisations.NewPostUnitsParams().
		WithOrganisationCreationRequest(&models.OrganisationCreation{
			Data: &models.Organisation{
				OrganisationID: organisationId,
				Type:           "organisations",
				ID:             *id,
				Attributes: &models.OrganisationAttributes{
					Name: "TestOrganisation",
				},
			},
		}))

	if err != nil {
		if !strings.Contains(err.Error(), "409") {
			assertNoErrorOccurred(t, err)
		}
	}

	_, err = auth.OrganisationClient.Organisations.DeleteUnitsID(organisations.NewDeleteUnitsIDParams().
		WithID(*id),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.OrganisationClient.Organisations.GetUnitsID(organisations.NewGetUnitsIDParams().
		WithID(*id))

	assertStatusCode(t, err, 404)
}

func TestAccDeleteOrganisationAssociation(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	assocId := NewUUID()

	defer func() {
		if t.Failed() {
			_, _ = auth.AssociationClient.Associations.DeleteStarlingID(associations.NewDeleteStarlingIDParams().
				WithID(*assocId).
				WithVersion(0),
			)
		}
	}()

	createResponse, err := auth.AssociationClient.Associations.PostStarling(associations.NewPostStarlingParams().
		WithCreationRequest(&models.AssociationCreation{
			Data: &models.NewAssociation{
				OrganisationID: testOrganisationId,
				Type:           "starling_associations",
				ID:             *assocId,
				Attributes: &models.NewAssociationAttributes{
					StarlingAccountName: "TestStarlingAccountName",
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.DeleteStarlingID(associations.NewDeleteStarlingIDParams().
		WithID(*assocId).
		WithVersion(0),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.AssociationClient.Associations.GetStarlingID(associations.NewGetStarlingIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}

func TestAccDeleteBankids(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	bankIdUUID := NewUUID()

	defer func() {
		if t.Failed() {
			_, _ = auth.AccountClient.Accounts.DeleteBankidsID(accounts.NewDeleteBankidsIDParams().
				WithID(*bankIdUUID).WithVersion(0),
			)
		}
	}()

	createResponse, err := auth.AccountClient.Accounts.PostBankids(accounts.NewPostBankidsParams().
		WithBankIDCreationRequest(&models.BankIDCreation{
			Data: &models.BankID{
				OrganisationID: testOrganisationId,
				Type:           "bankids",
				ID:             *bankIdUUID,
				Attributes: &models.BankIDAttributes{
					BankID:     "400301",
					BankIDCode: "GBDSC",
					Country:    "GB",
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.AccountClient.Accounts.DeleteBankidsID(accounts.NewDeleteBankidsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.AccountClient.Accounts.GetBankidsID(accounts.NewGetBankidsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}

func TestAccDeleteBics(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	bicUUID := NewUUID()

	defer func() {
		_, _ = auth.AccountClient.Accounts.DeleteBicsID(
			accounts.NewDeleteBicsIDParams().
				WithID(*bicUUID).
				WithVersion(0),
		)
	}()

	createResponse, err := auth.AccountClient.Accounts.PostBics(accounts.NewPostBicsParams().
		WithBicCreationRequest(&models.BicCreation{
			Data: &models.Bic{
				OrganisationID: testOrganisationId,
				Type:           "bics",
				ID:             *bicUUID,
				Attributes: &models.BicAttributes{
					Bic: "NWBKGB41",
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.AccountClient.Accounts.DeleteBicsID(accounts.NewDeleteBicsIDParams().
		WithID(createResponse.Payload.Data.ID).WithVersion(0),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.AccountClient.Accounts.GetBicsID(accounts.NewGetBicsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}
