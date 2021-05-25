package unittests_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

func TestFailedBacsAssociationWithDebugEnabled(t *testing.T) {
	skip := len(os.Getenv("FORM3_ACC")) != 0
	if skip {
		t.Skip("skipped as FORM3_ACC environment variable is set")
	}

	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("panic happen: %#v", err)
		}
	}()

	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stderr)

	defer os.Setenv("TF_LOG", os.Getenv("TF_LOG"))
	os.Setenv("TF_LOG", "DEBUG")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/oauth2/token" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"AccessToken":"token","TokenType":"Bearer"}`)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprintf(w, `{"data": {"type": "value"}}`)
	}))
	defer ts.Close()

	tsURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	authClient := api.NewAuthenticatedClient(&client.TransportConfig{
		Host:     tsURL.Host,
		BasePath: "",
		Schemes:  []string{tsURL.Scheme},
	})

	err = authClient.Authenticate("client", "secret")
	if err != nil {
		t.Error(err)
	}

	randomUUID := strfmt.UUID(uuid.New().String())

	_, err = authClient.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().WithID(randomUUID))
	if err == nil {
		t.Error("expected error")
	}
}
