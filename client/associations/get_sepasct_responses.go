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

// GetSepasctReader is a Reader for the GetSepasct structure.
type GetSepasctReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSepasctReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSepasctOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSepasctBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSepasctUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSepasctForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSepasctNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetSepasctConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetSepasctTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSepasctInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetSepasctServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSepasctOK creates a GetSepasctOK with default headers values
func NewGetSepasctOK() *GetSepasctOK {
	return &GetSepasctOK{}
}

/*GetSepasctOK handles this case with default header values.

List of associations
*/
type GetSepasctOK struct {
	Payload *models.SepaSctAssociationDetailsListResponse
}

func (o *GetSepasctOK) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctOK  %+v", 200, o.Payload)
}

func (o *GetSepasctOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaSctAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepasctBadRequest creates a GetSepasctBadRequest with default headers values
func NewGetSepasctBadRequest() *GetSepasctBadRequest {
	return &GetSepasctBadRequest{}
}

/*GetSepasctBadRequest handles this case with default header values.

Bad Request
*/
type GetSepasctBadRequest struct {
	Payload *models.APIError
}

func (o *GetSepasctBadRequest) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctBadRequest  %+v", 400, o.Payload)
}

func (o *GetSepasctBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepasctUnauthorized creates a GetSepasctUnauthorized with default headers values
func NewGetSepasctUnauthorized() *GetSepasctUnauthorized {
	return &GetSepasctUnauthorized{}
}

/*GetSepasctUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetSepasctUnauthorized struct {
	Payload *models.APIError
}

func (o *GetSepasctUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSepasctUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepasctForbidden creates a GetSepasctForbidden with default headers values
func NewGetSepasctForbidden() *GetSepasctForbidden {
	return &GetSepasctForbidden{}
}

/*GetSepasctForbidden handles this case with default header values.

Forbidden
*/
type GetSepasctForbidden struct {
	Payload *models.APIError
}

func (o *GetSepasctForbidden) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctForbidden  %+v", 403, o.Payload)
}

func (o *GetSepasctForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepasctNotFound creates a GetSepasctNotFound with default headers values
func NewGetSepasctNotFound() *GetSepasctNotFound {
	return &GetSepasctNotFound{}
}

/*GetSepasctNotFound handles this case with default header values.

Record not found
*/
type GetSepasctNotFound struct {
	Payload *models.APIError
}

func (o *GetSepasctNotFound) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctNotFound  %+v", 404, o.Payload)
}

func (o *GetSepasctNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepasctConflict creates a GetSepasctConflict with default headers values
func NewGetSepasctConflict() *GetSepasctConflict {
	return &GetSepasctConflict{}
}

/*GetSepasctConflict handles this case with default header values.

Conflict
*/
type GetSepasctConflict struct {
	Payload *models.APIError
}

func (o *GetSepasctConflict) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctConflict  %+v", 409, o.Payload)
}

func (o *GetSepasctConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepasctTooManyRequests creates a GetSepasctTooManyRequests with default headers values
func NewGetSepasctTooManyRequests() *GetSepasctTooManyRequests {
	return &GetSepasctTooManyRequests{}
}

/*GetSepasctTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetSepasctTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetSepasctTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSepasctTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepasctInternalServerError creates a GetSepasctInternalServerError with default headers values
func NewGetSepasctInternalServerError() *GetSepasctInternalServerError {
	return &GetSepasctInternalServerError{}
}

/*GetSepasctInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSepasctInternalServerError struct {
	Payload *models.APIError
}

func (o *GetSepasctInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSepasctInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepasctServiceUnavailable creates a GetSepasctServiceUnavailable with default headers values
func NewGetSepasctServiceUnavailable() *GetSepasctServiceUnavailable {
	return &GetSepasctServiceUnavailable{}
}

/*GetSepasctServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetSepasctServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetSepasctServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSepasctServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
