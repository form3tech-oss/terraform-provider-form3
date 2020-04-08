package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/ace"
	"github.com/form3tech-oss/terraform-provider-form3/client/roles"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	uuid "github.com/nu7hatch/gouuid"
)

func TestAccDeleteRole(t *testing.T) {

	id := NewUUID()
	createResponse, err := auth.SecurityClient.Roles.PostRoles(roles.NewPostRolesParams().
		WithRoleCreationRequest(&models.RoleCreation{
			Data: &models.Role{
				OrganisationID: organisationId,
				Type:           "roles",
				ID:             *id,
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

	assertNoErrorOccurred(err, t)

	_, err = auth.SecurityClient.Ace.DeleteRolesRoleIDAcesAceID(ace.NewDeleteRolesRoleIDAcesAceIDParams().
		WithRoleID(createAceResponse.Payload.Data.Attributes.RoleID).
		WithAceID(createAceResponse.Payload.Data.ID))

	assertNoErrorOccurred(err, t)

	_, err = auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
		WithRoleID(createRoleResponse.Payload.Data.ID).
		WithVersion(0),
	)

	assertNoErrorOccurred(err, t)
}
