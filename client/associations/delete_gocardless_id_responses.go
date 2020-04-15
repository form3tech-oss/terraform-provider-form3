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

// DeleteGocardlessIDReader is a Reader for the DeleteGocardlessID structure.
type DeleteGocardlessIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGocardlessIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteGocardlessIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteGocardlessIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteGocardlessIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteGocardlessIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteGocardlessIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteGocardlessIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewDeleteGocardlessIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteGocardlessIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteGocardlessIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteGocardlessIDNoContent creates a DeleteGocardlessIDNoContent with default headers values
func NewDeleteGocardlessIDNoContent() *DeleteGocardlessIDNoContent {
	return &DeleteGocardlessIDNoContent{}
}

/*DeleteGocardlessIDNoContent handles this case with default header values.

association deleted successfully
*/
type DeleteGocardlessIDNoContent struct {
}

func (o *DeleteGocardlessIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdNoContent ", 204)
}

func (o *DeleteGocardlessIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGocardlessIDBadRequest creates a DeleteGocardlessIDBadRequest with default headers values
func NewDeleteGocardlessIDBadRequest() *DeleteGocardlessIDBadRequest {
	return &DeleteGocardlessIDBadRequest{}
}

/*DeleteGocardlessIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteGocardlessIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteGocardlessIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGocardlessIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGocardlessIDUnauthorized creates a DeleteGocardlessIDUnauthorized with default headers values
func NewDeleteGocardlessIDUnauthorized() *DeleteGocardlessIDUnauthorized {
	return &DeleteGocardlessIDUnauthorized{}
}

/*DeleteGocardlessIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteGocardlessIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteGocardlessIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteGocardlessIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGocardlessIDForbidden creates a DeleteGocardlessIDForbidden with default headers values
func NewDeleteGocardlessIDForbidden() *DeleteGocardlessIDForbidden {
	return &DeleteGocardlessIDForbidden{}
}

/*DeleteGocardlessIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteGocardlessIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteGocardlessIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteGocardlessIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGocardlessIDNotFound creates a DeleteGocardlessIDNotFound with default headers values
func NewDeleteGocardlessIDNotFound() *DeleteGocardlessIDNotFound {
	return &DeleteGocardlessIDNotFound{}
}

/*DeleteGocardlessIDNotFound handles this case with default header values.

Record not found
*/
type DeleteGocardlessIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteGocardlessIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGocardlessIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGocardlessIDConflict creates a DeleteGocardlessIDConflict with default headers values
func NewDeleteGocardlessIDConflict() *DeleteGocardlessIDConflict {
	return &DeleteGocardlessIDConflict{}
}

/*DeleteGocardlessIDConflict handles this case with default header values.

Conflict
*/
type DeleteGocardlessIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteGocardlessIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteGocardlessIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGocardlessIDTooManyRequests creates a DeleteGocardlessIDTooManyRequests with default headers values
func NewDeleteGocardlessIDTooManyRequests() *DeleteGocardlessIDTooManyRequests {
	return &DeleteGocardlessIDTooManyRequests{}
}

/*DeleteGocardlessIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteGocardlessIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteGocardlessIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteGocardlessIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGocardlessIDInternalServerError creates a DeleteGocardlessIDInternalServerError with default headers values
func NewDeleteGocardlessIDInternalServerError() *DeleteGocardlessIDInternalServerError {
	return &DeleteGocardlessIDInternalServerError{}
}

/*DeleteGocardlessIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteGocardlessIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteGocardlessIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGocardlessIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGocardlessIDServiceUnavailable creates a DeleteGocardlessIDServiceUnavailable with default headers values
func NewDeleteGocardlessIDServiceUnavailable() *DeleteGocardlessIDServiceUnavailable {
	return &DeleteGocardlessIDServiceUnavailable{}
}

/*DeleteGocardlessIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteGocardlessIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteGocardlessIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /gocardless/{id}][%d] deleteGocardlessIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteGocardlessIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
