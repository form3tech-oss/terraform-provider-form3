// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRolesRoleIDReader is a Reader for the DeleteRolesRoleID structure.
type DeleteRolesRoleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRolesRoleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteRolesRoleIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRolesRoleIDNoContent creates a DeleteRolesRoleIDNoContent with default headers values
func NewDeleteRolesRoleIDNoContent() *DeleteRolesRoleIDNoContent {
	return &DeleteRolesRoleIDNoContent{}
}

/*DeleteRolesRoleIDNoContent handles this case with default header values.

Role deleted
*/
type DeleteRolesRoleIDNoContent struct {
}

func (o *DeleteRolesRoleIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /roles/{role_id}][%d] deleteRolesRoleIdNoContent ", 204)
}

func (o *DeleteRolesRoleIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
