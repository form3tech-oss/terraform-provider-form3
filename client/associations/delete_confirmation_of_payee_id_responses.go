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

// DeleteConfirmationOfPayeeIDReader is a Reader for the DeleteConfirmationOfPayeeID structure.
type DeleteConfirmationOfPayeeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConfirmationOfPayeeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConfirmationOfPayeeIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConfirmationOfPayeeIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConfirmationOfPayeeIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConfirmationOfPayeeIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConfirmationOfPayeeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteConfirmationOfPayeeIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConfirmationOfPayeeIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConfirmationOfPayeeIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConfirmationOfPayeeIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConfirmationOfPayeeIDNoContent creates a DeleteConfirmationOfPayeeIDNoContent with default headers values
func NewDeleteConfirmationOfPayeeIDNoContent() *DeleteConfirmationOfPayeeIDNoContent {
	return &DeleteConfirmationOfPayeeIDNoContent{}
}

/*DeleteConfirmationOfPayeeIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteConfirmationOfPayeeIDNoContent struct {
}

func (o *DeleteConfirmationOfPayeeIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdNoContent ", 204)
}

func (o *DeleteConfirmationOfPayeeIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConfirmationOfPayeeIDBadRequest creates a DeleteConfirmationOfPayeeIDBadRequest with default headers values
func NewDeleteConfirmationOfPayeeIDBadRequest() *DeleteConfirmationOfPayeeIDBadRequest {
	return &DeleteConfirmationOfPayeeIDBadRequest{}
}

/*DeleteConfirmationOfPayeeIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteConfirmationOfPayeeIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteConfirmationOfPayeeIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConfirmationOfPayeeIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConfirmationOfPayeeIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfirmationOfPayeeIDUnauthorized creates a DeleteConfirmationOfPayeeIDUnauthorized with default headers values
func NewDeleteConfirmationOfPayeeIDUnauthorized() *DeleteConfirmationOfPayeeIDUnauthorized {
	return &DeleteConfirmationOfPayeeIDUnauthorized{}
}

/*DeleteConfirmationOfPayeeIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteConfirmationOfPayeeIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteConfirmationOfPayeeIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConfirmationOfPayeeIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConfirmationOfPayeeIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfirmationOfPayeeIDForbidden creates a DeleteConfirmationOfPayeeIDForbidden with default headers values
func NewDeleteConfirmationOfPayeeIDForbidden() *DeleteConfirmationOfPayeeIDForbidden {
	return &DeleteConfirmationOfPayeeIDForbidden{}
}

/*DeleteConfirmationOfPayeeIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteConfirmationOfPayeeIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteConfirmationOfPayeeIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConfirmationOfPayeeIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConfirmationOfPayeeIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfirmationOfPayeeIDNotFound creates a DeleteConfirmationOfPayeeIDNotFound with default headers values
func NewDeleteConfirmationOfPayeeIDNotFound() *DeleteConfirmationOfPayeeIDNotFound {
	return &DeleteConfirmationOfPayeeIDNotFound{}
}

/*DeleteConfirmationOfPayeeIDNotFound handles this case with default header values.

Record not found
*/
type DeleteConfirmationOfPayeeIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteConfirmationOfPayeeIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConfirmationOfPayeeIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConfirmationOfPayeeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfirmationOfPayeeIDConflict creates a DeleteConfirmationOfPayeeIDConflict with default headers values
func NewDeleteConfirmationOfPayeeIDConflict() *DeleteConfirmationOfPayeeIDConflict {
	return &DeleteConfirmationOfPayeeIDConflict{}
}

/*DeleteConfirmationOfPayeeIDConflict handles this case with default header values.

Conflict
*/
type DeleteConfirmationOfPayeeIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteConfirmationOfPayeeIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteConfirmationOfPayeeIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConfirmationOfPayeeIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfirmationOfPayeeIDTooManyRequests creates a DeleteConfirmationOfPayeeIDTooManyRequests with default headers values
func NewDeleteConfirmationOfPayeeIDTooManyRequests() *DeleteConfirmationOfPayeeIDTooManyRequests {
	return &DeleteConfirmationOfPayeeIDTooManyRequests{}
}

/*DeleteConfirmationOfPayeeIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteConfirmationOfPayeeIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteConfirmationOfPayeeIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConfirmationOfPayeeIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConfirmationOfPayeeIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfirmationOfPayeeIDInternalServerError creates a DeleteConfirmationOfPayeeIDInternalServerError with default headers values
func NewDeleteConfirmationOfPayeeIDInternalServerError() *DeleteConfirmationOfPayeeIDInternalServerError {
	return &DeleteConfirmationOfPayeeIDInternalServerError{}
}

/*DeleteConfirmationOfPayeeIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteConfirmationOfPayeeIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteConfirmationOfPayeeIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConfirmationOfPayeeIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConfirmationOfPayeeIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfirmationOfPayeeIDServiceUnavailable creates a DeleteConfirmationOfPayeeIDServiceUnavailable with default headers values
func NewDeleteConfirmationOfPayeeIDServiceUnavailable() *DeleteConfirmationOfPayeeIDServiceUnavailable {
	return &DeleteConfirmationOfPayeeIDServiceUnavailable{}
}

/*DeleteConfirmationOfPayeeIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteConfirmationOfPayeeIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteConfirmationOfPayeeIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /confirmation-of-payee/{id}][%d] deleteConfirmationOfPayeeIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConfirmationOfPayeeIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteConfirmationOfPayeeIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
