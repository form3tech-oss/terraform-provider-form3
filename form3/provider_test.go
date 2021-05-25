package form3

import (
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		// provider is called terraform-provider-form3 ie form3
		"form3": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", form3.JsonErrorPrettyPrint(err))
	}

}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("FORM3_CLIENT_ID"); v == "" {
		t.Fatal("FORM3_CLIENT_ID must be set for acceptance tests")
	}

	if v := os.Getenv("FORM3_CLIENT_SECRET"); v == "" {
		t.Fatal("FORM3_CLIENT_SECRET must be set for acceptance tests")
	}

	if v := os.Getenv("FORM3_ORGANISATION_ID"); v == "" {
		t.Fatal("FORM3_ORGANISATION_ID must be set for acceptance tests")
	}
}
