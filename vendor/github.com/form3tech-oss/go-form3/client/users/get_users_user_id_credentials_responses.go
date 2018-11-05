// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/models"
)

// GetUsersUserIDCredentialsReader is a Reader for the GetUsersUserIDCredentials structure.
type GetUsersUserIDCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersUserIDCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersUserIDCredentialsOK creates a GetUsersUserIDCredentialsOK with default headers values
func NewGetUsersUserIDCredentialsOK() *GetUsersUserIDCredentialsOK {
	return &GetUsersUserIDCredentialsOK{}
}

/*GetUsersUserIDCredentialsOK handles this case with default header values.

List of credentials for user
*/
type GetUsersUserIDCredentialsOK struct {
	Payload *models.UserCredentialListResponse
}

func (o *GetUsersUserIDCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetUsersUserIDCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserCredentialListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
