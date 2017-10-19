package form3

import (
	"github.com/go-openapi/runtime"
	"io/ioutil"
	"reflect"
	"testing"
)

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
