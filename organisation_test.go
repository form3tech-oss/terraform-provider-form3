package form3

import (
	"github.com/ewilde/go-form3/client/organisations"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/strfmt"
	"testing"
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
