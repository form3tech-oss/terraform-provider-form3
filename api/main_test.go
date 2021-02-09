package api

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

var (
	auth     *AuthenticatedClient
	authOnce = new(sync.Once)
	config   = client.DefaultTransportConfig()
)

const testOrgName string = "terraform-provider-form3-test-organisation"

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	defer log.Println("[INFO] Stopping tests")

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
	initOrgs, err := auth.OrganisationClient.Organisations.GetUnits(nil)
	if err != nil {
		log.Printf("Failed to fetch organisations to check")
		return 1
	}
	exitCode := 0
	defer func() {
		if errTestOrg := deleteOrganisation(); errTestOrg != nil {
			log.Printf("[Error] Error deleting test organisation: %+v\n", errTestOrg)
			exitCode = 1
		}
		if errLeakOrg := verifyTotalAmountOfTestOrgsIsSame(auth, initOrgs.Payload.Data); errLeakOrg != nil {
			log.Fatal(errLeakOrg)
		}
	}()

	exitCode = m.Run()
	return exitCode
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

func verifyTotalAmountOfTestOrgsIsSame(c *AuthenticatedClient, initialOrgs []*models.Organisation) error {
	orgsResp, _ := c.OrganisationClient.Organisations.GetUnits(nil)

	initTestOrgs := map[string]interface{}{}
	finalTestOrgs := map[string]interface{}{}

	for _, init := range initialOrgs {
		if init.Attributes.Name == testOrgName {
			initTestOrgs[init.ID.String()] = struct{}{}
		}
	}

	for _, v := range orgsResp.Payload.Data {
		if v.Attributes.Name == testOrgName {
			finalTestOrgs[v.ID.String()] = struct{}{}
		}
	}

	if len(finalTestOrgs) > len(initTestOrgs) {
		newTestOrgs := []string{}
		for k := range finalTestOrgs {
			_, ok := initTestOrgs[k]
			if !ok {
				newTestOrgs = append(newTestOrgs, k)
			}
		}
		log.Printf("Test Org is: %s and org is: %s", testOrganisationId.String(), organisationId.String())
		return fmt.Errorf("Organization leak: There are %d new orgs, %s", len(newTestOrgs), strings.Join(newTestOrgs, ","))
	}

	return nil
}
