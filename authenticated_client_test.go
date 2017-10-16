package form3

import (
	"fmt"
	"github.com/ewilde/go-form3/client"
	"github.com/ewilde/go-form3/client/users"
	"os"
	"testing"
)

func TestAccLogin(t *testing.T) {
	config := testPreCheck(t)

	auth := NewAuthenticatedClient(config)
	err := auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
	}

	if len(auth.AccessToken) < 32 {
		fmt.Errorf("expected access token to be set, found %s", auth.AccessToken)
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
