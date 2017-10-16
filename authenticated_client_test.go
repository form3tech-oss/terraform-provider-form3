package form3

import (
	"github.com/ewilde/go-form3/client"
	"github.com/ewilde/go-form3/client/users"
	"os"
	"testing"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/runtime"
)

func TestAccLogin(t *testing.T) {
	config := testPreCheck(t)

	auth := NewAuthenticatedClient(config)
	err := auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
	}

	if len(auth.AccessToken) < 32 {
		t.Errorf("expected access token to be set, found %s", auth.AccessToken)
	}
}

func TestAccGetUsers(t *testing.T) {
	config := testPreCheck(t)

	auth := NewAuthenticatedClient(config)
	err := auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
	}

	response, err := auth.ApiClients.Users.GetUsers(users.NewGetUsersParams())
	if err != nil {
		t.Error(err)
	}

	if len(response.Payload.Data) == 0 {
		t.Error("Expected at least one user")
	}
}

func TestAccGetUserWithIdNotFound(t *testing.T) {
	config := testPreCheck(t)
	auth := NewAuthenticatedClient(config)
	err := auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
	}

	_, err = auth.ApiClients.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(strfmt.UUID("700e7327-3834-4fe1-95f6-7eea7773bf0f")))
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

func testPreCheck(t *testing.T) *client.TransportConfig {
	skip := len(os.Getenv("FORM3_ACC")) == 0
	if skip {
		t.Log("form3 client_test.go tests require setting FORM3_ACC=1 environment variable")
		t.Skip()
	}

	if len(os.Getenv("FORM3_CLIENT_ID")) == 0 {
		t.Fatal("FORM3_CLIENT_ID must be set for acceptance tests")
	}

	if len(os.Getenv("FORM3_CLIENT_SECRET")) == 0 {
		t.Fatal("FORM3_CLIENT_SECRET must be set for acceptance tests")
	}

	config := client.DefaultTransportConfig()
	if v := os.Getenv("FORM3_HOST"); v != "" {
		config.WithHost(v)
	}

	return config
}
