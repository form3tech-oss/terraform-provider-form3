// Code generated by go-swagger; DO NOT EDIT.

package ace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetRolesRoleIDAcesReader is a Reader for the GetRolesRoleIDAces structure.
type GetRolesRoleIDAcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesRoleIDAcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRolesRoleIDAcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRolesRoleIDAcesOK creates a GetRolesRoleIDAcesOK with default headers values
func NewGetRolesRoleIDAcesOK() *GetRolesRoleIDAcesOK {
	return &GetRolesRoleIDAcesOK{}
}

/*GetRolesRoleIDAcesOK handles this case with default header values.

List of Ace details
*/
type GetRolesRoleIDAcesOK struct {
	Payload *models.AceDetailsListResponse
}

func (o *GetRolesRoleIDAcesOK) Error() string {
	return fmt.Sprintf("[GET /roles/{role_id}/aces][%d] getRolesRoleIdAcesOK  %+v", 200, o.Payload)
}

func (o *GetRolesRoleIDAcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AceDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
