package form3

import (
	"github.com/go-openapi/runtime"
	rc "github.com/go-openapi/runtime/client"
	"github.com/ewilde/go-form3/client"
	"github.com/go-openapi/strfmt"
	"fmt"
	"io"
	"time"
	"net/http"
	"golang.org/x/net/context"
	"github.com/go-openapi/errors"
)

type AuthenticatedClient struct {

}

func (r *AuthenticatedClient) Login(a *client.Form3CorelibDataStructures, clientId string, clientSecret string) error {
	result, err := a.Transport.Submit(&runtime.ClientOperation{
		AuthInfo:			rc.BasicAuth(clientId, clientSecret),
		ID:                 "Authenticate",
		Method:             "POST",
		PathPattern:        "/oauth2/token",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Reader:             &LoginReader{

			formats:        strfmt.Default,
		},
		Params:				NewLoginParams(rc.BasicAuth(clientId, clientSecret)),
	})

	if err != nil {
		return err
	}

	println(result)
	return nil
}


type LoginReader struct {
	formats        strfmt.Registry
}

func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRolesRoleIDAcesAceIDOK creates a DeleteRolesRoleIDAcesAceIDOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

type LoginOK struct {
	Payload *LoginResponse
}

type LoginResponse struct {
	TokenType string `json:"token_type,omitempty"`
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
}

func (o *LoginOK) Error() string {
	return fmt.Sprintf("[GET /oauth2/token][%d] loginOK  %+v", 200, o.Payload)
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LoginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

type LoginParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
	clientAuth runtime.ClientAuthInfoWriter
}

func NewLoginParams(clientAuth runtime.ClientAuthInfoWriter) *LoginParams {
	var ()
	return &LoginParams{
		clientAuth: clientAuth,
		timeout: 	rc.DefaultTimeout,
	}
}

func NewLoginParamsWithTimeout(timeout time.Duration) *LoginParams {
	var ()
	return &LoginParams{

		timeout: timeout,
	}
}

func NewLoginParamsWithContext(ctx context.Context) *LoginParams {
	var ()
	return &LoginParams{

		Context: ctx,
	}
}

// NewGetUsersUserIDParamsWithHTTPClient creates a new GetUsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoginParamsWithHTTPClient(client *http.Client) *LoginParams {
	var ()
	return &LoginParams{
		HTTPClient: client,
	}
}

func (o *LoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	//o.clientAuth.AuthenticateRequest(r, reg)
	r.SetFormParam("grant_type", "client_credentials")
	//r.SetHeaderParam("Content-Type", "application/x-www-form-urlencoded")
	//r.SetHeaderParam("goo", "bboo")
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error


	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
