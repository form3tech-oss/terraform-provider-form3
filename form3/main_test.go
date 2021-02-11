package form3

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/google/uuid"
)

var cl *api.AuthenticatedClient

var testOrgNamePrefix string = "terraform-provider-form3-test-organisation"
var testOrgName string = ""

func TestMain(m *testing.M) {
	testOrgSuffix := uuid.New().String()
	testOrgName = fmt.Sprintf("%s-%s", testOrgNamePrefix, testOrgSuffix)

	flag.Parse()
	if !testing.Verbose() {
		log.SetOutput(ioutil.Discard)
	}
	var err error
	cl, err = createClient()
	if err != nil {
		log.Fatalf("cannot setup authentication client, %+v", err)
	}
	orgResp, _ := cl.OrganisationClient.Organisations.GetUnits(nil)
	exitCode := 0
	defer func() {
		err := verifyTotalAmountOfTestOrgsIsSame(cl, orgResp.Payload.Data)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(exitCode)
	}()
	exitCode = m.Run()
}

func verifyOrgDoesNotExist(t *testing.T, ID string) error {
	org, err := cl.OrganisationClient.Organisations.GetUnits(nil)
	if err != nil {
		t.Error("failed to get organisations")
	}
	for _, v := range org.Payload.Data {
		if v.ID.String() == ID {
			t.Error("organisations leaked.")
		}
	}
	return nil
}

func createClient() (*api.AuthenticatedClient, error) {
	config := client.DefaultTransportConfig()
	if v := os.Getenv("FORM3_HOST"); v != "" {
		config.WithHost(v)
	}
	cl := api.NewAuthenticatedClient(config)
	if cl.AccessToken == "" {
		err := cl.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
		if err != nil {
			return nil, err
		}
	}
	return cl, nil
}

func verifyTotalAmountOfTestOrgsIsSame(c *api.AuthenticatedClient, initialOrgs []*models.Organisation) error {
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
		return fmt.Errorf("Organization leak: there are %d new orgs, %s", len(newTestOrgs), strings.Join(newTestOrgs, ","))
	}

	return nil
}
