// Code generated by go-swagger; DO NOT EDIT.

package limits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetLimitsReader is a Reader for the GetLimits structure.
type GetLimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLimitsOK creates a GetLimitsOK with default headers values
func NewGetLimitsOK() *GetLimitsOK {
	return &GetLimitsOK{}
}

/*GetLimitsOK handles this case with default header values.

List of limit details
*/
type GetLimitsOK struct {
	Payload *models.LimitDetailsListResponse
}

func (o *GetLimitsOK) Error() string {
	return fmt.Sprintf("[GET /limits][%d] getLimitsOK  %+v", 200, o.Payload)
}

func (o *GetLimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
