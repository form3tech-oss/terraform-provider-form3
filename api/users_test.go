package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/roles"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
)

func TestAccGetUsers(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	response, err := auth.SecurityClient.Users.GetUsers(users.NewGetUsersParams())
	assertNoErrorOccurred(t, err)

	if len(response.Payload.Data) == 0 {
		t.Error("Expected at least one user")
	}
}

func TestAccDeleteUser(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	roleID := NewUUID()

	defer func() {
		_, err := auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().WithRoleID(*roleID))
		assertNoErrorOccurred(t, err)
	}()

	_, err := auth.SecurityClient.Roles.PostRoles(roles.NewPostRolesParams().
		WithRoleCreationRequest(&models.RoleCreation{
			Data: &models.Role{
				OrganisationID: organisationId,
				Type:           "roles",
				ID:             *roleID,
				Attributes: &models.RoleAttributes{
					Name: "terraform-test",
				},
			},
		}),
	)
	assertNoErrorOccurred(t, err)

	userUUID := NewUUID()

	defer func() {
		if t.Failed() {
			_, _ = auth.SecurityClient.Users.DeleteUsersUserID(users.NewDeleteUsersUserIDParams().WithUserID(*userUUID))
		}
	}()

	createResponse, err := auth.SecurityClient.Users.PostUsers(users.NewPostUsersParams().
		WithUserCreationRequest(&models.UserCreation{
			Data: &models.User{
				OrganisationID: organisationId,
				Type:           "users",
				ID:             *userUUID,
				Attributes: &models.UserAttributes{
					Email:    "go-form3@form3.tech",
					Username: "go-form3",
					RoleIds:  []strfmt.UUID{*roleID},
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.SecurityClient.Users.DeleteUsersUserID(users.NewDeleteUsersUserIDParams().
		WithUserID(createResponse.Payload.Data.ID),
	)
	assertNoErrorOccurred(t, err)

	_, err = auth.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().
		WithUserID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}

func TestAccGetUserWithIdNotFound(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	_, err := auth.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID("700e7327-3834-4fe1-95f6-7eea7773bf0f"))
	assertStatusCode(t, err, 404)
}
