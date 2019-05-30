package api

import (
	"errors"
	"github.com/form3tech-oss/go-form3/client"
	"github.com/form3tech-oss/go-form3/client/organisations"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
)

var organisationId strfmt.UUID
var testOrganisationId strfmt.UUID

var auth *AuthenticatedClient
var authOnce = new(sync.Once)
var config = client.DefaultTransportConfig()

func TestMain(m *testing.M) {
	if err := testPreCheck(); err != nil {
		log.Fatalf("[FATAL] Error initializing test run: %+v", err)
		os.Exit(-1)
	}

	if v := os.Getenv("FORM3_HOST"); v != "" {
		config.WithHost(v)
	}

	createClient(config)
	log.Println("[INFO] Starting tests")
	if err := createOrganisation(); err != nil {
		log.Fatalf("[FATAL] Error creating test organisation: %+v", err)
		os.Exit(-1)
	}

	code := m.Run()

	if err := deleteOrganisation(); err != nil {
		log.Printf("[WARN] Error deleting test organisation: %+v\n", err)
		os.Exit(-1)
	}

	log.Println("[INFO] Stopping tests")

	os.Exit(code)
}

func createOrganisation() error {
	ensureAuthenticated()
	newId, _ := uuid.NewV4()
	testOrganisationId = strfmt.UUID(newId.String())
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
	log.Printf("[INFO] Deleting test organisation %v", testOrganisationId)

	if _, err := auth.OrganisationClient.Organisations.DeleteUnitsID(organisations.NewDeleteUnitsIDParams().
		WithID(testOrganisationId)); err != nil {
		return err
	}

	log.Printf("[INFO] Sucessfuly deleted test organisation %v", testOrganisationId)

	return nil
}

func createClient(config *client.TransportConfig) {
	authOnce.Do(func() {
		auth = NewAuthenticatedClient(config)
	})
}

func ensureAuthenticated() {
	if AccessToken == "" {
		Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
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

func testPreCheck() error {
	skip := len(os.Getenv("FORM3_ACC")) == 0
	if skip {
		return errors.New("form3 client_test.go tests require setting FORM3_ACC=1 environment variable")
	}

	if len(os.Getenv("FORM3_CLIENT_ID")) == 0 {
		return errors.New("FORM3_CLIENT_ID must be set for acceptance tests")
	}

	if len(os.Getenv("FORM3_ORGANISATION_ID")) == 0 {
		return errors.New("FORM3_ORGANISATION_ID must be set for acceptance tests")
	}
	organisationId = strfmt.UUID(os.Getenv("FORM3_ORGANISATION_ID"))

	if len(os.Getenv("FORM3_CLIENT_SECRET")) == 0 {
		return errors.New("FORM3_CLIENT_SECRET must be set for acceptance tests")
	}

	return nil
}
