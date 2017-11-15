package form3

import (
	"github.com/ewilde/go-form3/client/users"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
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

	createResponse, err := auth.SecurityClient.Users.PostUsers(users.NewPostUsersParams().
		WithUserCreationRequest(&models.UserCreation{
			Data: &models.User{
				OrganisationID: organisationId,
				Type:           "users",
				ID:             strfmt.UUID("8d1abeff-ec82-44b8-a9d4-5080756ebf4f"),
				Attributes: &models.UserAttributes{
					Email:    "go-form3@form3.tech",
					Username: "go-form3",
					RoleIds:  []strfmt.UUID{strfmt.UUID("32881d6b-a000-4258-b779-56c59970590f")},
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
