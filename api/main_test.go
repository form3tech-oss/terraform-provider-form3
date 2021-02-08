package api

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client"
	"github.com/form3tech-oss/terraform-provider-form3/client/organisations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

var (
	organisationId     strfmt.UUID
	testOrganisationId strfmt.UUID
)

const testOrgName string = "terraform-provider-form3-test-organisation"

var (
	auth     *AuthenticatedClient
	authOnce = new(sync.Once)
	config   = client.DefaultTransportConfig()
)

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	defer log.Println("[INFO] Stopping tests")
	orgCount := getOrgAmount(testOrgName)
	defer verifyNoTestOrganizationLeak(orgCount)

	flag.Parse()

	if !testing.Verbose() {
		log.SetOutput(ioutil.Discard)
	}

	skip := len(os.Getenv("FORM3_ACC")) == 0
	if skip {
		log.Println("Client tests skipped as FORM3_ACC environment variable not set")
		return 0
	}

	if err := testPreCheck(); err != nil {
		log.Fatalf("[FATAL] Error initializing test run: %+v", err)
	}

	if v := os.Getenv("FORM3_HOST"); v != "" {
		config.WithHost(v)
	}

	createClient(config)
	log.Println("[INFO] Starting tests")
	if err := createOrganisation(); err != nil {
		log.Fatalf("[FATAL] Error creating test organisation: %s", JsonErrorPrettyPrint(err))
	}

	defer func() {
		if errTestOrg := deleteOrganisation(); errTestOrg != nil {
			log.Fatalf("[Error] Error deleting test organisation: %+v\n", errTestOrg)
		}
	}()

	return m.Run()
}

func createOrganisation() error {
	if err := ensureAuthenticated(); err != nil {
		return err
	}

	newId := uuid.New()
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

func verifyNoTestOrganizationLeak(initCount int) error {
	log.Printf("[INFO] Verifying there are no `terraform-provider-form3-test-organisation` leftover.")
	count := getOrgAmount(testOrgName)
	if count > initCount {
		log.Fatalf("[Error] Organization leak: had %d organizations with name: %s before, and now %d \n", initCount, testOrgName, count)
	}
	return nil
}

func getOrgAmount(name string) int {
	count := 0
	orgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	for _, v := range orgs.Payload.Data {
		if v.Attributes.Name == name {
			count++
		}
	}
	return count
}

func deleteOrganisation() error {
	log.Printf("[INFO] Deleting test organisation %v", testOrganisationId)

	if _, err := auth.OrganisationClient.Organisations.DeleteUnitsID(organisations.NewDeleteUnitsIDParams().
		WithID(testOrganisationId).WithVersion(0)); err != nil {
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

func ensureAuthenticated() error {
	if auth.AccessToken == "" {
		return auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	}
	return nil
}

func assertNoErrorOccurred(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(JsonErrorPrettyPrint(err))
	}
}

func assertStatusCode(t *testing.T, err error, code int) {
	if err == nil {
		t.Fatal("No error, expected an api error")
	}

	if !IsJsonErrorStatusCode(err, code) {
		t.Fatalf("Expected api error, got %+v", JsonErrorPrettyPrint(err))
	}
}

func testPreCheck() error {
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
