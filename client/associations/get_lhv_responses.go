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

// GetLhvReader is a Reader for the GetLhv structure.
type GetLhvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLhvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLhvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLhvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLhvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLhvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLhvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetLhvConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLhvTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLhvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLhvServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLhvOK creates a GetLhvOK with default headers values
func NewGetLhvOK() *GetLhvOK {
	return &GetLhvOK{}
}

/* GetLhvOK describes a response with status code 200, with default header values.

List of associations
*/
type GetLhvOK struct {
	Payload *models.LhvAssociationDetailsListResponse
}

func (o *GetLhvOK) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvOK  %+v", 200, o.Payload)
}
func (o *GetLhvOK) GetPayload() *models.LhvAssociationDetailsListResponse {
	return o.Payload
}

func (o *GetLhvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LhvAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvBadRequest creates a GetLhvBadRequest with default headers values
func NewGetLhvBadRequest() *GetLhvBadRequest {
	return &GetLhvBadRequest{}
}

/* GetLhvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLhvBadRequest struct {
	Payload *models.APIError
}

func (o *GetLhvBadRequest) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvBadRequest  %+v", 400, o.Payload)
}
func (o *GetLhvBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvUnauthorized creates a GetLhvUnauthorized with default headers values
func NewGetLhvUnauthorized() *GetLhvUnauthorized {
	return &GetLhvUnauthorized{}
}

/* GetLhvUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetLhvUnauthorized struct {
	Payload *models.APIError
}

func (o *GetLhvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvUnauthorized  %+v", 401, o.Payload)
}
func (o *GetLhvUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvForbidden creates a GetLhvForbidden with default headers values
func NewGetLhvForbidden() *GetLhvForbidden {
	return &GetLhvForbidden{}
}

/* GetLhvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLhvForbidden struct {
	Payload *models.APIError
}

func (o *GetLhvForbidden) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvForbidden  %+v", 403, o.Payload)
}
func (o *GetLhvForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvNotFound creates a GetLhvNotFound with default headers values
func NewGetLhvNotFound() *GetLhvNotFound {
	return &GetLhvNotFound{}
}

/* GetLhvNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetLhvNotFound struct {
	Payload *models.APIError
}

func (o *GetLhvNotFound) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvNotFound  %+v", 404, o.Payload)
}
func (o *GetLhvNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvConflict creates a GetLhvConflict with default headers values
func NewGetLhvConflict() *GetLhvConflict {
	return &GetLhvConflict{}
}

/* GetLhvConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetLhvConflict struct {
	Payload *models.APIError
}

func (o *GetLhvConflict) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvConflict  %+v", 409, o.Payload)
}
func (o *GetLhvConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvTooManyRequests creates a GetLhvTooManyRequests with default headers values
func NewGetLhvTooManyRequests() *GetLhvTooManyRequests {
	return &GetLhvTooManyRequests{}
}

/* GetLhvTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetLhvTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetLhvTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetLhvTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvInternalServerError creates a GetLhvInternalServerError with default headers values
func NewGetLhvInternalServerError() *GetLhvInternalServerError {
	return &GetLhvInternalServerError{}
}

/* GetLhvInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetLhvInternalServerError struct {
	Payload *models.APIError
}

func (o *GetLhvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLhvInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvServiceUnavailable creates a GetLhvServiceUnavailable with default headers values
func NewGetLhvServiceUnavailable() *GetLhvServiceUnavailable {
	return &GetLhvServiceUnavailable{}
}

/* GetLhvServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetLhvServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetLhvServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /lhv][%d] getLhvServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetLhvServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
