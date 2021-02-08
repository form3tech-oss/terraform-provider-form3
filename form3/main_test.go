package form3

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/api"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client"
)

const testOrgName string = "terraform-provider-form3-test-organisation"

var config = client.DefaultTransportConfig()

func TestMain(m *testing.M) {
	flag.Parse()

	if !testing.Verbose() {
		log.SetOutput(ioutil.Discard)
	}
	cl, err := createClient(config)
	if err != nil {
		log.Fatalf("failed to setup client %+v", err)
	}
	orgCount, err := getOrgAmount(testOrgName, cl)
	if err != nil {
		log.Fatalf("Failed to retrieve test organizations amount %+v", err)
	}
	defer verifyNoTestOrganizationLeak(orgCount, cl)

	os.Exit(m.Run())
}

func verifyNoTestOrganizationLeak(initCount int, client *form3.AuthenticatedClient) error {
	log.Printf("[INFO] Verifying there are no %s leftover.", testOrgName)
	count, err := getOrgAmount(testOrgName, client)
	if err != nil {
		return err
	}
	if count > initCount {
		log.Fatalf("[Error] Organization leak: had %d organizations with name: %s before, and now %d \n", initCount, testOrgName, count)
	}
	return nil
}

func getOrgAmount(name string, client *form3.AuthenticatedClient) (int, error) {
	count := 0
	orgs, _ := client.OrganisationClient.Organisations.GetUnits(nil)
	for _, v := range orgs.Payload.Data {
		if v.Attributes.Name == name {
			count++
		}
	}
	return count, nil
}

func createClient(config *client.TransportConfig) (*form3.AuthenticatedClient, error) {
	cl := api.NewAuthenticatedClient(config)
	if cl.AccessToken == "" {
		err := cl.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
		if err != nil {
			return nil, err
		}
	}
	return cl, nil
}
