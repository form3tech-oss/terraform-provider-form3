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

// DeletePayportIDReader is a Reader for the DeletePayportID structure.
type DeletePayportIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePayportIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePayportIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePayportIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePayportIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePayportIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePayportIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeletePayportIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePayportIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePayportIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeletePayportIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePayportIDNoContent creates a DeletePayportIDNoContent with default headers values
func NewDeletePayportIDNoContent() *DeletePayportIDNoContent {
	return &DeletePayportIDNoContent{}
}

/*DeletePayportIDNoContent handles this case with default header values.

Association deleted
*/
type DeletePayportIDNoContent struct {
}

func (o *DeletePayportIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdNoContent ", 204)
}

func (o *DeletePayportIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePayportIDBadRequest creates a DeletePayportIDBadRequest with default headers values
func NewDeletePayportIDBadRequest() *DeletePayportIDBadRequest {
	return &DeletePayportIDBadRequest{}
}

/*DeletePayportIDBadRequest handles this case with default header values.

Bad Request
*/
type DeletePayportIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeletePayportIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePayportIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePayportIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePayportIDUnauthorized creates a DeletePayportIDUnauthorized with default headers values
func NewDeletePayportIDUnauthorized() *DeletePayportIDUnauthorized {
	return &DeletePayportIDUnauthorized{}
}

/*DeletePayportIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeletePayportIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeletePayportIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePayportIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePayportIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePayportIDForbidden creates a DeletePayportIDForbidden with default headers values
func NewDeletePayportIDForbidden() *DeletePayportIDForbidden {
	return &DeletePayportIDForbidden{}
}

/*DeletePayportIDForbidden handles this case with default header values.

Forbidden
*/
type DeletePayportIDForbidden struct {
	Payload *models.APIError
}

func (o *DeletePayportIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdForbidden  %+v", 403, o.Payload)
}

func (o *DeletePayportIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePayportIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePayportIDNotFound creates a DeletePayportIDNotFound with default headers values
func NewDeletePayportIDNotFound() *DeletePayportIDNotFound {
	return &DeletePayportIDNotFound{}
}

/*DeletePayportIDNotFound handles this case with default header values.

Record not found
*/
type DeletePayportIDNotFound struct {
	Payload *models.APIError
}

func (o *DeletePayportIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdNotFound  %+v", 404, o.Payload)
}

func (o *DeletePayportIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePayportIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePayportIDConflict creates a DeletePayportIDConflict with default headers values
func NewDeletePayportIDConflict() *DeletePayportIDConflict {
	return &DeletePayportIDConflict{}
}

/*DeletePayportIDConflict handles this case with default header values.

Conflict
*/
type DeletePayportIDConflict struct {
	Payload *models.APIError
}

func (o *DeletePayportIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdConflict  %+v", 409, o.Payload)
}

func (o *DeletePayportIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePayportIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePayportIDTooManyRequests creates a DeletePayportIDTooManyRequests with default headers values
func NewDeletePayportIDTooManyRequests() *DeletePayportIDTooManyRequests {
	return &DeletePayportIDTooManyRequests{}
}

/*DeletePayportIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeletePayportIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeletePayportIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeletePayportIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePayportIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePayportIDInternalServerError creates a DeletePayportIDInternalServerError with default headers values
func NewDeletePayportIDInternalServerError() *DeletePayportIDInternalServerError {
	return &DeletePayportIDInternalServerError{}
}

/*DeletePayportIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeletePayportIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeletePayportIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePayportIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePayportIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePayportIDServiceUnavailable creates a DeletePayportIDServiceUnavailable with default headers values
func NewDeletePayportIDServiceUnavailable() *DeletePayportIDServiceUnavailable {
	return &DeletePayportIDServiceUnavailable{}
}

/*DeletePayportIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeletePayportIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeletePayportIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /payport/{id}][%d] deletePayportIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeletePayportIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePayportIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
