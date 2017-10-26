package form3

import (
	"github.com/ewilde/go-form3/client/organisations"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/strfmt"
	"testing"
	"github.com/ewilde/go-form3/client/associations"
)

func TestAccDeleteOrganisation(t *testing.T) {
	testPreCheck(t)
	ensureAuthenticated()

	createResponse, err := auth.OrganisationClient.Organisations.PostUnits(organisations.NewPostUnitsParams().
		WithOrganisationCreationRequest(&models.OrganisationCreation{
			Data: &models.Organisation{
				OrganisationID: organisationId,
				Type:           "organisations",
				ID:             strfmt.UUID("58a78c22-efa6-4f67-b2ec-30c53fd9a437"),
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
	testPreCheck(t)
	ensureAuthenticated()

	createResponse, err := auth.AssociationClient.Associations.PostAssociations(associations.NewPostAssociationsParams().
		WithCreationRequest(&models.AssociationCreation{
			Data: &models.Association{
				OrganisationID: organisationId,
				Type:           "organisation_associations",
				ID:             strfmt.UUID("7c0763f6-4192-4e65-8f5d-fac1fc57eb21"),
				Attributes: &models.AssociationAttributes{
					StarlingAccountName: "TestStarlingAccountName",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.DeleteAssociationsID(associations.NewDeleteAssociationsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)

	_, err = auth.AssociationClient.Associations.GetAssociationsID(associations.NewGetAssociationsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}
