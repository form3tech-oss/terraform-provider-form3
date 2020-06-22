// Code generated by go-swagger; DO NOT EDIT.

package account_routings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// DeleteAccountRoutingsIDReader is a Reader for the DeleteAccountRoutingsID structure.
type DeleteAccountRoutingsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountRoutingsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAccountRoutingsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAccountRoutingsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteAccountRoutingsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteAccountRoutingsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteAccountRoutingsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteAccountRoutingsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewDeleteAccountRoutingsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteAccountRoutingsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteAccountRoutingsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAccountRoutingsIDNoContent creates a DeleteAccountRoutingsIDNoContent with default headers values
func NewDeleteAccountRoutingsIDNoContent() *DeleteAccountRoutingsIDNoContent {
	return &DeleteAccountRoutingsIDNoContent{}
}

/*DeleteAccountRoutingsIDNoContent handles this case with default header values.

Account Routing deleted
*/
type DeleteAccountRoutingsIDNoContent struct {
}

func (o *DeleteAccountRoutingsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdNoContent ", 204)
}

func (o *DeleteAccountRoutingsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAccountRoutingsIDBadRequest creates a DeleteAccountRoutingsIDBadRequest with default headers values
func NewDeleteAccountRoutingsIDBadRequest() *DeleteAccountRoutingsIDBadRequest {
	return &DeleteAccountRoutingsIDBadRequest{}
}

/*DeleteAccountRoutingsIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteAccountRoutingsIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteAccountRoutingsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAccountRoutingsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountRoutingsIDUnauthorized creates a DeleteAccountRoutingsIDUnauthorized with default headers values
func NewDeleteAccountRoutingsIDUnauthorized() *DeleteAccountRoutingsIDUnauthorized {
	return &DeleteAccountRoutingsIDUnauthorized{}
}

/*DeleteAccountRoutingsIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteAccountRoutingsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteAccountRoutingsIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAccountRoutingsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountRoutingsIDForbidden creates a DeleteAccountRoutingsIDForbidden with default headers values
func NewDeleteAccountRoutingsIDForbidden() *DeleteAccountRoutingsIDForbidden {
	return &DeleteAccountRoutingsIDForbidden{}
}

/*DeleteAccountRoutingsIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteAccountRoutingsIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteAccountRoutingsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAccountRoutingsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountRoutingsIDNotFound creates a DeleteAccountRoutingsIDNotFound with default headers values
func NewDeleteAccountRoutingsIDNotFound() *DeleteAccountRoutingsIDNotFound {
	return &DeleteAccountRoutingsIDNotFound{}
}

/*DeleteAccountRoutingsIDNotFound handles this case with default header values.

Record not found
*/
type DeleteAccountRoutingsIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteAccountRoutingsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAccountRoutingsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountRoutingsIDConflict creates a DeleteAccountRoutingsIDConflict with default headers values
func NewDeleteAccountRoutingsIDConflict() *DeleteAccountRoutingsIDConflict {
	return &DeleteAccountRoutingsIDConflict{}
}

/*DeleteAccountRoutingsIDConflict handles this case with default header values.

Conflict
*/
type DeleteAccountRoutingsIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteAccountRoutingsIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteAccountRoutingsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountRoutingsIDTooManyRequests creates a DeleteAccountRoutingsIDTooManyRequests with default headers values
func NewDeleteAccountRoutingsIDTooManyRequests() *DeleteAccountRoutingsIDTooManyRequests {
	return &DeleteAccountRoutingsIDTooManyRequests{}
}

/*DeleteAccountRoutingsIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteAccountRoutingsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteAccountRoutingsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAccountRoutingsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountRoutingsIDInternalServerError creates a DeleteAccountRoutingsIDInternalServerError with default headers values
func NewDeleteAccountRoutingsIDInternalServerError() *DeleteAccountRoutingsIDInternalServerError {
	return &DeleteAccountRoutingsIDInternalServerError{}
}

/*DeleteAccountRoutingsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteAccountRoutingsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteAccountRoutingsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAccountRoutingsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountRoutingsIDServiceUnavailable creates a DeleteAccountRoutingsIDServiceUnavailable with default headers values
func NewDeleteAccountRoutingsIDServiceUnavailable() *DeleteAccountRoutingsIDServiceUnavailable {
	return &DeleteAccountRoutingsIDServiceUnavailable{}
}

/*DeleteAccountRoutingsIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteAccountRoutingsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteAccountRoutingsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /account_routings/{id}][%d] deleteAccountRoutingsIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteAccountRoutingsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
