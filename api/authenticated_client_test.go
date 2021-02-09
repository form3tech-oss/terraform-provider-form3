package api

import (
	"os"
	"testing"

	"github.com/go-openapi/strfmt"
)

func TestAccLogin(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	err := auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
	}

	if len(auth.AccessToken) < 32 {
		t.Errorf("expected access token to be set, found %s", auth.AccessToken)
	}
}

func TestUUIDConversion(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	stringUUID := string("14cd29f6-114a-4325-ac5c-31808e7f77f6")
	uuid := strfmt.UUID(stringUUID)

	if uuid.String() != stringUUID {
		t.Errorf("Expected %s found %s", stringUUID, uuid.String())
	}
}
