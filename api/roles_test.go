package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/ace"
	"github.com/form3tech-oss/terraform-provider-form3/client/roles"
	"github.com/form3tech-oss/terraform-provider-form3/models"
)

func TestAccDeleteRole(t *testing.T) {
	id := NewUUID()

	defer func() {
		if t.Failed() {
			_, _ = auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
				WithRoleID(*id),
			)
		}
	}()

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

	assertNoErrorOccurred(t, err)

	_, err = auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
		WithRoleID(createResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.SecurityClient.Roles.GetRolesRoleID(roles.NewGetRolesRoleIDParams().
		WithRoleID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}

func TestAccDeleteRoleAce(t *testing.T) {
	roleId := NewUUID()
	aceId := NewUUID()

	defer func() {
		if t.Failed() {
			_, _ = auth.SecurityClient.Ace.DeleteRolesRoleIDAcesAceID(ace.NewDeleteRolesRoleIDAcesAceIDParams().
				WithRoleID(*roleId).
				WithAceID(*aceId),
			)

			_, _ = auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
				WithRoleID(*roleId),
			)
		}
	}()

	createRoleResponse, err := auth.SecurityClient.Roles.PostRoles(roles.NewPostRolesParams().
		WithRoleCreationRequest(&models.RoleCreation{
			Data: &models.Role{
				OrganisationID: organisationId,
				Type:           "roles",
				ID:             *roleId,
				Attributes: &models.RoleAttributes{
					Name: "TestRole",
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	createAceResponse, err := auth.SecurityClient.Ace.PostRolesRoleIDAces(ace.NewPostRolesRoleIDAcesParams().
		WithRoleID(createRoleResponse.Payload.Data.ID).
		WithAceCreationRequest(&models.AceCreation{
			Data: &models.Ace{
				OrganisationID: organisationId,
				Type:           "roles",
				ID:             *aceId,
				Attributes: &models.AceAttributes{
					RoleID:     createRoleResponse.Payload.Data.ID,
					Action:     "updated",
					RecordType: "User",
				},
			},
		}))
	assertNoErrorOccurred(t, err)

	_, err = auth.SecurityClient.Ace.DeleteRolesRoleIDAcesAceID(ace.NewDeleteRolesRoleIDAcesAceIDParams().
		WithRoleID(createAceResponse.Payload.Data.Attributes.RoleID).
		WithAceID(createAceResponse.Payload.Data.ID))

	assertNoErrorOccurred(t, err)

	_, err = auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
		WithRoleID(createRoleResponse.Payload.Data.ID),
	)

	assertNoErrorOccurred(t, err)
}
