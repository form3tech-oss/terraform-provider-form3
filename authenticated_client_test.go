package form3

import (
	"github.com/ewilde/go-form3/client"
	"github.com/go-openapi/strfmt"
	"os"
	"sync"
	"testing"
)

var auth *AuthenticatedClient
var authOnce = new(sync.Once)
var organisationId strfmt.UUID

func TestAccLogin(t *testing.T) {
	testPreCheck(t)

	err := auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
	}

	if len(auth.AccessToken) < 32 {
		t.Errorf("expected access token to be set, found %s", auth.AccessToken)
	}
}

func TestUUIDConversion(t *testing.T) {
	stringUUID := string("14cd29f6-114a-4325-ac5c-31808e7f77f6")
	uuid := strfmt.UUID(stringUUID)

	if uuid.String() != stringUUID {
		t.Errorf("Expected %s found %s", stringUUID, uuid.String())
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

	if len(os.Getenv("FORM3_ORGANISATION_ID")) == 0 {
		t.Fatal("FORM3_ORGANISATION_ID must be set for acceptance tests")
	}
	organisationId = strfmt.UUID(os.Getenv("FORM3_ORGANISATION_ID"))

	if len(os.Getenv("FORM3_CLIENT_SECRET")) == 0 {
		t.Fatal("FORM3_CLIENT_SECRET must be set for acceptance tests")
	}

	config := client.DefaultTransportConfig()
	if v := os.Getenv("FORM3_HOST"); v != "" {
		config.WithHost(v)
	}

	createClient(config)
	return config
}

func createClient(config *client.TransportConfig) {
	authOnce.Do(func() {
		auth = NewAuthenticatedClient(config)
	})
}

func ensureAuthenticated() {
	if auth.AccessToken == "" {
		auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	}
}
