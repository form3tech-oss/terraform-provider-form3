package form3

import (
	"github.com/form3tech-oss/go-form3/client/ace"
	"github.com/form3tech-oss/go-form3/client/roles"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/strfmt"
	"testing"
	"github.com/nu7hatch/gouuid"
)

func TestAccDeleteRole(t *testing.T) {

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

func TestAccDeleteRoleAce(t *testing.T) {

	roleId := Must(uuid.NewV4()).String()

	createRoleResponse, err := auth.SecurityClient.Roles.PostRoles(roles.NewPostRolesParams().
		WithRoleCreationRequest(&models.RoleCreation{
			Data: &models.Role{
				OrganisationID: organisationId,
				Type:           "roles",
				ID:             strfmt.UUID(roleId),
				Attributes: &models.RoleAttributes{
					Name: "TestRole",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	aceId := Must(uuid.NewV4()).String()

	createAceResponse, err := auth.SecurityClient.Ace.PostRolesRoleIDAces(ace.NewPostRolesRoleIDAcesParams().
		WithRoleID(createRoleResponse.Payload.Data.ID).
		WithAceCreationRequest(&models.AceCreation{
			Data: &models.Ace{
				OrganisationID: organisationId,
				Type:           "roles",
				ID:             strfmt.UUID(aceId),
				Attributes: &models.AceAttributes{
					RoleID:     createRoleResponse.Payload.Data.ID,
					Action:     "updated",
					RecordType: "User",
				},
			},
		}))

	_, err = auth.SecurityClient.Ace.DeleteRolesRoleIDAcesAceID(ace.NewDeleteRolesRoleIDAcesAceIDParams().
		WithRoleID(createAceResponse.Payload.Data.Attributes.RoleID).
		WithAceID(createAceResponse.Payload.Data.ID))

	assertNoErrorOccurred(err, t)

	_, err = auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
		WithRoleID(createRoleResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(err, t)
}
