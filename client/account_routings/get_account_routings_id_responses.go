// Code generated by go-swagger; DO NOT EDIT.

package account_routings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetAccountRoutingsIDReader is a Reader for the GetAccountRoutingsID structure.
type GetAccountRoutingsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountRoutingsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountRoutingsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountRoutingsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountRoutingsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountRoutingsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountRoutingsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAccountRoutingsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountRoutingsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountRoutingsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAccountRoutingsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountRoutingsIDOK creates a GetAccountRoutingsIDOK with default headers values
func NewGetAccountRoutingsIDOK() *GetAccountRoutingsIDOK {
	return &GetAccountRoutingsIDOK{}
}

/* GetAccountRoutingsIDOK describes a response with status code 200, with default header values.

Account Routing details
*/
type GetAccountRoutingsIDOK struct {
	Payload *models.AccountRoutingDetailsResponse
}

func (o *GetAccountRoutingsIDOK) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdOK  %+v", 200, o.Payload)
}
func (o *GetAccountRoutingsIDOK) GetPayload() *models.AccountRoutingDetailsResponse {
	return o.Payload
}

func (o *GetAccountRoutingsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountRoutingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRoutingsIDBadRequest creates a GetAccountRoutingsIDBadRequest with default headers values
func NewGetAccountRoutingsIDBadRequest() *GetAccountRoutingsIDBadRequest {
	return &GetAccountRoutingsIDBadRequest{}
}

/* GetAccountRoutingsIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAccountRoutingsIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetAccountRoutingsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccountRoutingsIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAccountRoutingsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRoutingsIDUnauthorized creates a GetAccountRoutingsIDUnauthorized with default headers values
func NewGetAccountRoutingsIDUnauthorized() *GetAccountRoutingsIDUnauthorized {
	return &GetAccountRoutingsIDUnauthorized{}
}

/* GetAccountRoutingsIDUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetAccountRoutingsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetAccountRoutingsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAccountRoutingsIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAccountRoutingsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRoutingsIDForbidden creates a GetAccountRoutingsIDForbidden with default headers values
func NewGetAccountRoutingsIDForbidden() *GetAccountRoutingsIDForbidden {
	return &GetAccountRoutingsIDForbidden{}
}

/* GetAccountRoutingsIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAccountRoutingsIDForbidden struct {
	Payload *models.APIError
}

func (o *GetAccountRoutingsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdForbidden  %+v", 403, o.Payload)
}
func (o *GetAccountRoutingsIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAccountRoutingsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRoutingsIDNotFound creates a GetAccountRoutingsIDNotFound with default headers values
func NewGetAccountRoutingsIDNotFound() *GetAccountRoutingsIDNotFound {
	return &GetAccountRoutingsIDNotFound{}
}

/* GetAccountRoutingsIDNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetAccountRoutingsIDNotFound struct {
	Payload *models.APIError
}

func (o *GetAccountRoutingsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdNotFound  %+v", 404, o.Payload)
}
func (o *GetAccountRoutingsIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAccountRoutingsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRoutingsIDConflict creates a GetAccountRoutingsIDConflict with default headers values
func NewGetAccountRoutingsIDConflict() *GetAccountRoutingsIDConflict {
	return &GetAccountRoutingsIDConflict{}
}

/* GetAccountRoutingsIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAccountRoutingsIDConflict struct {
	Payload *models.APIError
}

func (o *GetAccountRoutingsIDConflict) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdConflict  %+v", 409, o.Payload)
}
func (o *GetAccountRoutingsIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAccountRoutingsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRoutingsIDTooManyRequests creates a GetAccountRoutingsIDTooManyRequests with default headers values
func NewGetAccountRoutingsIDTooManyRequests() *GetAccountRoutingsIDTooManyRequests {
	return &GetAccountRoutingsIDTooManyRequests{}
}

/* GetAccountRoutingsIDTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetAccountRoutingsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetAccountRoutingsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetAccountRoutingsIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAccountRoutingsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRoutingsIDInternalServerError creates a GetAccountRoutingsIDInternalServerError with default headers values
func NewGetAccountRoutingsIDInternalServerError() *GetAccountRoutingsIDInternalServerError {
	return &GetAccountRoutingsIDInternalServerError{}
}

/* GetAccountRoutingsIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccountRoutingsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetAccountRoutingsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountRoutingsIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAccountRoutingsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRoutingsIDServiceUnavailable creates a GetAccountRoutingsIDServiceUnavailable with default headers values
func NewGetAccountRoutingsIDServiceUnavailable() *GetAccountRoutingsIDServiceUnavailable {
	return &GetAccountRoutingsIDServiceUnavailable{}
}

/* GetAccountRoutingsIDServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetAccountRoutingsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetAccountRoutingsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /account_routings/{id}][%d] getAccountRoutingsIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetAccountRoutingsIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAccountRoutingsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
