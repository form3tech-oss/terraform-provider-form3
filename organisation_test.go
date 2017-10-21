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

	auth.SecurityClient.Organisations.PostOrganisations(organisations.NewPostOrganisationsParams().
		WithOrganisationCreationRequest(&models.OrganisationCreation{
			Data: &models.Organisation{
				OrganisationID: organisationId,
				Type:           "organisations",
				ID:             strfmt.UUID("f6679900-10d2-44a1-9a46-2d972f4bf457"),
				Attributes: &models.OrganisationAttributes{
					Name: "TestOrganisation",
				},
			},
		}))

	/*
		assertNoErrorOccurred(err, t)

		_, err = auth.SecurityClient.Organisations.DeleteOrganisationsID(organisations.NewDeleteOrganisationsIDParams().
			WithID(createResponse.Payload.Data.ID),
		)

		assertNoErrorOccurred(err, t)

		_, err = auth.SecurityClient.Organisations.GetOrganisationsID(organisations.NewGetOrganisationsIDParams().
			WithID(createResponse.Payload.Data.ID))

		assertStatusCode(err, t, 404)
	*/
}
