// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSepainstantIDReader is a Reader for the DeleteSepainstantID structure.
type DeleteSepainstantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSepainstantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSepainstantIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSepainstantIDNoContent creates a DeleteSepainstantIDNoContent with default headers values
func NewDeleteSepainstantIDNoContent() *DeleteSepainstantIDNoContent {
	return &DeleteSepainstantIDNoContent{}
}

/*DeleteSepainstantIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteSepainstantIDNoContent struct {
}

func (o *DeleteSepainstantIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sepainstant/{id}][%d] deleteSepainstantIdNoContent ", 204)
}

func (o *DeleteSepainstantIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
