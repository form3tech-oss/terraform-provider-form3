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

// GetStarlingIDReader is a Reader for the GetStarlingID structure.
type GetStarlingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStarlingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStarlingIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStarlingIDOK creates a GetStarlingIDOK with default headers values
func NewGetStarlingIDOK() *GetStarlingIDOK {
	return &GetStarlingIDOK{}
}

/*GetStarlingIDOK handles this case with default header values.

Associations details
*/
type GetStarlingIDOK struct {
	Payload *models.AssociationDetailsResponse
}

func (o *GetStarlingIDOK) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdOK  %+v", 200, o.Payload)
}

func (o *GetStarlingIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
