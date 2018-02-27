// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/models"
)

// PostBacsReader is a Reader for the PostBacs structure.
type PostBacsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBacsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostBacsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostBacsCreated creates a PostBacsCreated with default headers values
func NewPostBacsCreated() *PostBacsCreated {
	return &PostBacsCreated{}
}

/*PostBacsCreated handles this case with default header values.

creation response
*/
type PostBacsCreated struct {
	Payload *models.BacsAssociationCreationResponse
}

func (o *PostBacsCreated) Error() string {
	return fmt.Sprintf("[POST /bacs][%d] postBacsCreated  %+v", 201, o.Payload)
}

func (o *PostBacsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BacsAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
