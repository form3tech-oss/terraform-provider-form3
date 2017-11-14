package form3

import (
	"github.com/ewilde/go-form3/client"
	"github.com/ewilde/go-form3/client/organisations"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
)

var organisationId strfmt.UUID
var testOrganisationId = strfmt.UUID("b85bf973-93f1-4857-9351-f0a74babd687")

var auth *AuthenticatedClient
var authOnce = new(sync.Once)
var config = client.DefaultTransportConfig()

type testContext struct {
}

func TestMain(m *testing.M) {
	if v := os.Getenv("FORM3_HOST"); v != "" {
		config.WithHost(v)
	}

	createClient(config)
	log.Println("[INFO] Starting tests")
	if err := createOrganisation(); err != nil {
		log.Fatalf("[FATAL] Error creating test organisation: %+v", err)
		os.Exit(-1)
	}

	defer func() {
		if err := deleteOrganisation(); err != nil {
			log.Printf("[WARN] Error deleting test organisation: %+v\n", err)
			os.Exit(-1)
		}
	}()
	code := m.Run()

	log.Println("[INFO] Stopping tests")

	os.Exit(code)
}

func createOrganisation() error {
	ensureAuthenticated()
	_, err := auth.OrganisationClient.Organisations.PostUnits(organisations.NewPostUnitsParams().
		WithOrganisationCreationRequest(&models.OrganisationCreation{
		Data: &models.Organisation{
			OrganisationID: organisationId,
			Type:           "organisations",
			ID:             testOrganisationId,
			Attributes: &models.OrganisationAttributes{
				Name: "TestOrganisation",
			},
		},
	}))

	return err
}

func deleteOrganisation() error {
	_, err := auth.OrganisationClient.Organisations.DeleteUnitsID(organisations.NewDeleteUnitsIDParams().
		WithID(testOrganisationId),
	)

	return err
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

func assertNoErrorOccurred(err error, t *testing.T) {
	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok {
			response, ok := apiError.Response.(runtime.ClientResponse)
			if ok {
				bodyBytes, _ := ioutil.ReadAll(response.Body())
				body := string(bodyBytes)
				t.Fatalf("%v %v %v", response.Message(), response.Code(), body)
			}
			t.Fatalf("%s", getType(apiError.Response))
		}

		t.Fatal(err)
	}
}

func assertStatusCode(err error, t *testing.T, code int) {
	if err == nil {
		t.Fatal("No error, expected an api error")
	}

	apiError, ok := err.(*runtime.APIError)
	if !ok {
		t.Fatalf("Expected api error, got %+v", err)
	}

	if apiError.Code != code {
		t.Fatalf("Expected %d got %d", code, apiError.Code)
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
