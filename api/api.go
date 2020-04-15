package api

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"

	"github.com/go-openapi/runtime"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

var httpCodeRegex = regexp.MustCompile(`\[(?P<Code>\d{3})]`)

// JsonError represents an error in json format
// it is a ApiError wrapper
type JsonError interface {
	GetPayload() *models.APIError
	Error() string
}

func IsJsonErrorStatusCode(err error, statusCode int) bool {
	if err == nil {
		return false
	}

	apiError, ok := err.(JsonError)
	if !ok {
		return false
	}

	foundCode := httpCodeRegex.FindStringSubmatch(apiError.Error())

	return len(foundCode) == 2 && foundCode[1] == strconv.Itoa(statusCode)
}

func JsonErrorPrettyPrint(err error) string {
	if err == nil {
		return ""
	}

	switch v := err.(type) {
	case JsonError:
		payload := v.GetPayload()
		return fmt.Sprintf("ErrorCode: %s Message: %s", payload.ErrorCode.String(), payload.ErrorMessage)
	case *runtime.APIError:
		response, ok := v.Response.(runtime.ClientResponse)
		if !ok {
			return err.Error()
		}
		bodyBytes, _ := ioutil.ReadAll(response.Body())
		return fmt.Sprintf("Code: [%d] Message: %s Body: %v", response.Code(), response.Message(), string(bodyBytes))
	default:
		return err.Error()
	}
}
