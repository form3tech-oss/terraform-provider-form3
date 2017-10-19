package form3

import (
	"github.com/ewilde/go-form3/client/roles"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/strfmt"
	"testing"
)

func TestAccDeleteRole(t *testing.T) {
	testPreCheck(t)
	ensureAuthenticated()

	createResponse, err := auth.SecurityClient.Roles.PostRoles(roles.NewPostRolesParams().
		WithRoleCreationRequest(&models.RoleCreation{
			Data: &models.Role{
				OrganisationID: organisationId,
				Type:           "roles",
				ID:             strfmt.UUID("f6679900-10d2-44a1-9a46-2d972f4bf457"),
				Attributes: &models.RoleAttributes{
					Name: "TestRole",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
		WithRoleID(createResponse.Payload.Data.ID),
	)

	if err != nil {
		t.Error(err)
	}

	_, err = auth.SecurityClient.Roles.GetRolesRoleID(roles.NewGetRolesRoleIDParams().
		WithRoleID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}
