// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostConfirmationOfPayeeReader is a Reader for the PostConfirmationOfPayee structure.
type PostConfirmationOfPayeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConfirmationOfPayeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostConfirmationOfPayeeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostConfirmationOfPayeeCreated creates a PostConfirmationOfPayeeCreated with default headers values
func NewPostConfirmationOfPayeeCreated() *PostConfirmationOfPayeeCreated {
	return &PostConfirmationOfPayeeCreated{}
}

/*PostConfirmationOfPayeeCreated handles this case with default header values.

creation response
*/
type PostConfirmationOfPayeeCreated struct {
	Payload *models.CoPAssociationCreationResponse
}

func (o *PostConfirmationOfPayeeCreated) Error() string {
	return fmt.Sprintf("[POST /confirmation-of-payee][%d] postConfirmationOfPayeeCreated  %+v", 201, o.Payload)
}

func (o *PostConfirmationOfPayeeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoPAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
