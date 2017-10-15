package form3

import (
	"testing"
	"os"
	"github.com/ewilde/go-form3/client"
)
func TestLogin(t *testing.T) {
	config := testPreCheck(t)

	auth := new(AuthenticatedClient)
	err := auth.Login(client.NewHTTPClientWithConfig(nil, config), os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
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
