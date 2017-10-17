package form3

import (
	"github.com/ewilde/go-form3/client"
	"github.com/ewilde/go-form3/client/users"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
	"testing"
)

var auth *AuthenticatedClient
var authOnce = new(sync.Once)
var organisationId strfmt.UUID

func TestAccLogin(t *testing.T) {
	testPreCheck(t)

	err := auth.Authenticate(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
	}

	if len(auth.AccessToken) < 32 {
		t.Errorf("expected access token to be set, found %s", auth.AccessToken)
	}
}

func TestAccGetUsers(t *testing.T) {
	testPreCheck(t)
	ensureAuthenticated()

	response, err := auth.ApiClients.Users.GetUsers(users.NewGetUsersParams())
	if err != nil {
		t.Error(err)
	}

	if len(response.Payload.Data) == 0 {
		t.Error("Expected at least one user")
	}
}

func TestAccDeleteUser(t *testing.T) {
	testPreCheck(t)
	ensureAuthenticated()

	createResponse, err := auth.ApiClients.Users.PostUsers(users.NewPostUsersParams().
		WithUserCreationRequest(&models.UserCreation{
			Data: &models.User{
				OrganisationID: organisationId,
				Type:           "users",
				ID:             strfmt.UUID("8d1abeff-ec82-44b8-a9d4-5080756ebf4f"),
				Attributes: &models.UserAttributes{
					Email:    "go-form3@form3.tech",
					Username: "go-form3",
					RoleIds:  []strfmt.UUID{strfmt.UUID("32881d6b-a000-4258-b779-56c59970590f")},
				},
			},
		}))

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

	_, err = auth.ApiClients.Users.DeleteUsersUserID(users.NewDeleteUsersUserIDParams().
		WithUserID(createResponse.Payload.Data.ID),
	)

	if err != nil {
		t.Error(err)
	}
}

func TestAccGetUserWithIdNotFound(t *testing.T) {
	testPreCheck(t)
	ensureAuthenticated()

	_, err := auth.ApiClients.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(strfmt.UUID("700e7327-3834-4fe1-95f6-7eea7773bf0f")))
	if err == nil {
		t.Error("Expected error to occur")
	}

	apiError := err.(*runtime.APIError)
	if apiError == nil {
		t.Errorf("Expected API Error not %+v", err)
	}

	if apiError.Code != 404 {
		t.Errorf("Expected 404 Not Found not %v", apiError.Code)
	}
}

func TestUUIDConversion(t *testing.T) {
	stringUUID := string("14cd29f6-114a-4325-ac5c-31808e7f77f6")
	uuid := strfmt.UUID(stringUUID)

	if uuid.String() != stringUUID {
		t.Errorf("Expected %s found %s", stringUUID, uuid.String())
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

	if len(os.Getenv("FORM3_ORGANISATION_ID")) == 0 {
		t.Fatal("FORM3_ORGANISATION_ID must be set for acceptance tests")
	}
	organisationId = strfmt.UUID(os.Getenv("FORM3_ORGANISATION_ID"))

	if len(os.Getenv("FORM3_CLIENT_SECRET")) == 0 {
		t.Fatal("FORM3_CLIENT_SECRET must be set for acceptance tests")
	}

	config := client.DefaultTransportConfig()
	if v := os.Getenv("FORM3_HOST"); v != "" {
		config.WithHost(v)
	}

	createClient(config)
	return config
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

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
