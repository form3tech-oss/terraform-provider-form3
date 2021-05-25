// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetSwiftIDReader is a Reader for the GetSwiftID structure.
type GetSwiftIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSwiftIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSwiftIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSwiftIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSwiftIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSwiftIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSwiftIDOK creates a GetSwiftIDOK with default headers values
func NewGetSwiftIDOK() *GetSwiftIDOK {
	return &GetSwiftIDOK{}
}

/* GetSwiftIDOK describes a response with status code 200, with default header values.

Associations details
*/
type GetSwiftIDOK struct {
	Payload *models.SwiftAssociationDetailsResponse
}

func (o *GetSwiftIDOK) Error() string {
	return fmt.Sprintf("[GET /swift/{id}][%d] getSwiftIdOK  %+v", 200, o.Payload)
}
func (o *GetSwiftIDOK) GetPayload() *models.SwiftAssociationDetailsResponse {
	return o.Payload
}

func (o *GetSwiftIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwiftAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSwiftIDBadRequest creates a GetSwiftIDBadRequest with default headers values
func NewGetSwiftIDBadRequest() *GetSwiftIDBadRequest {
	return &GetSwiftIDBadRequest{}
}

/* GetSwiftIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSwiftIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetSwiftIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /swift/{id}][%d] getSwiftIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetSwiftIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSwiftIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSwiftIDForbidden creates a GetSwiftIDForbidden with default headers values
func NewGetSwiftIDForbidden() *GetSwiftIDForbidden {
	return &GetSwiftIDForbidden{}
}

/* GetSwiftIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSwiftIDForbidden struct {
	Payload *models.APIError
}

func (o *GetSwiftIDForbidden) Error() string {
	return fmt.Sprintf("[GET /swift/{id}][%d] getSwiftIdForbidden  %+v", 403, o.Payload)
}
func (o *GetSwiftIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSwiftIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSwiftIDNotFound creates a GetSwiftIDNotFound with default headers values
func NewGetSwiftIDNotFound() *GetSwiftIDNotFound {
	return &GetSwiftIDNotFound{}
}

/* GetSwiftIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSwiftIDNotFound struct {
	Payload *models.APIError
}

func (o *GetSwiftIDNotFound) Error() string {
	return fmt.Sprintf("[GET /swift/{id}][%d] getSwiftIdNotFound  %+v", 404, o.Payload)
}
func (o *GetSwiftIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSwiftIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
