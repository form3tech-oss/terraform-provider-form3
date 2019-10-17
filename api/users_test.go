package api

import (
	"github.com/form3tech-oss/terraform-provider-form3/client/roles"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestAccGetUsers(t *testing.T) {

	response, err := auth.SecurityClient.Users.GetUsers(users.NewGetUsersParams())
	assertNoErrorOccurred(err, t)

	if len(response.Payload.Data) == 0 {
		t.Error("Expected at least one user")
	}
}

func TestAccDeleteUser(t *testing.T) {
	roleID := uuid.NewV4().String()

	_, err := auth.SecurityClient.Roles.PostRoles(roles.NewPostRolesParams().
		WithRoleCreationRequest(&models.RoleCreation{
			Data: &models.Role{
				OrganisationID: organisationId,
				Type:           "roles",
				ID:             strfmt.UUID(roleID),
				Attributes: &models.RoleAttributes{
					Name: "terraform-test",
				},
			},
		}),
	)
	if err != nil {
		t.Errorf("failed to create role %s: %s", roleID, err)
	}

	createResponse, err := auth.SecurityClient.Users.PostUsers(users.NewPostUsersParams().
		WithUserCreationRequest(&models.UserCreation{
			Data: &models.User{
				OrganisationID: organisationId,
				Type:           "users",
				ID:             strfmt.UUID("8d1abeff-ec82-44b8-a9d4-5080756ebf4f"),
				Attributes: &models.UserAttributes{
					Email:    "go-form3@form3.tech",
					Username: "go-form3",
					RoleIds:  []strfmt.UUID{strfmt.UUID(roleID)},
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.SecurityClient.Users.DeleteUsersUserID(users.NewDeleteUsersUserIDParams().
		WithUserID(createResponse.Payload.Data.ID),
	)

	if err != nil {
		t.Error(err)
	}

	_, err = auth.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().
		WithUserID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)

	_, err = auth.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
		WithRoleID(strfmt.UUID(roleID)))
	if err != nil {
		t.Errorf("failed to delete role %s: %s", roleID, err)
	}
}

func TestAccGetUserWithIdNotFound(t *testing.T) {

	_, err := auth.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(strfmt.UUID("700e7327-3834-4fe1-95f6-7eea7773bf0f")))
	if err == nil {
		t.Error("Expected error to occur")
	}

	apiError := err.(*runtime.APIError)
	if apiError == nil {
		t.Errorf("Expected API Error not %+v", err)
	}

	if apiError.Code != 404 {
		t.Errorf("Expected 404 Not Found not %v", apiError.Code)
	}
}
