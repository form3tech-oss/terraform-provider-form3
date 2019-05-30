// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostRolesReader is a Reader for the PostRoles structure.
type PostRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostRolesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRolesCreated creates a PostRolesCreated with default headers values
func NewPostRolesCreated() *PostRolesCreated {
	return &PostRolesCreated{}
}

/*PostRolesCreated handles this case with default header values.

Role creation response
*/
type PostRolesCreated struct {
	Payload *models.RoleCreationResponse
}

func (o *PostRolesCreated) Error() string {
	return fmt.Sprintf("[POST /roles][%d] postRolesCreated  %+v", 201, o.Payload)
}

func (o *PostRolesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
