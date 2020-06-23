// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// DeleteKeysKeyIDCertificatesCertificateIDReader is a Reader for the DeleteKeysKeyIDCertificatesCertificateID structure.
type DeleteKeysKeyIDCertificatesCertificateIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKeysKeyIDCertificatesCertificateIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteKeysKeyIDCertificatesCertificateIDNoContent creates a DeleteKeysKeyIDCertificatesCertificateIDNoContent with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDNoContent() *DeleteKeysKeyIDCertificatesCertificateIDNoContent {
	return &DeleteKeysKeyIDCertificatesCertificateIDNoContent{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDNoContent handles this case with default header values.

Certificate deleted
*/
type DeleteKeysKeyIDCertificatesCertificateIDNoContent struct {
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdNoContent ", 204)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKeysKeyIDCertificatesCertificateIDBadRequest creates a DeleteKeysKeyIDCertificatesCertificateIDBadRequest with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDBadRequest() *DeleteKeysKeyIDCertificatesCertificateIDBadRequest {
	return &DeleteKeysKeyIDCertificatesCertificateIDBadRequest{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteKeysKeyIDCertificatesCertificateIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDCertificatesCertificateIDUnauthorized creates a DeleteKeysKeyIDCertificatesCertificateIDUnauthorized with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDUnauthorized() *DeleteKeysKeyIDCertificatesCertificateIDUnauthorized {
	return &DeleteKeysKeyIDCertificatesCertificateIDUnauthorized{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteKeysKeyIDCertificatesCertificateIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDCertificatesCertificateIDForbidden creates a DeleteKeysKeyIDCertificatesCertificateIDForbidden with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDForbidden() *DeleteKeysKeyIDCertificatesCertificateIDForbidden {
	return &DeleteKeysKeyIDCertificatesCertificateIDForbidden{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteKeysKeyIDCertificatesCertificateIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDCertificatesCertificateIDNotFound creates a DeleteKeysKeyIDCertificatesCertificateIDNotFound with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDNotFound() *DeleteKeysKeyIDCertificatesCertificateIDNotFound {
	return &DeleteKeysKeyIDCertificatesCertificateIDNotFound{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDNotFound handles this case with default header values.

Record not found
*/
type DeleteKeysKeyIDCertificatesCertificateIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDCertificatesCertificateIDConflict creates a DeleteKeysKeyIDCertificatesCertificateIDConflict with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDConflict() *DeleteKeysKeyIDCertificatesCertificateIDConflict {
	return &DeleteKeysKeyIDCertificatesCertificateIDConflict{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDConflict handles this case with default header values.

Conflict
*/
type DeleteKeysKeyIDCertificatesCertificateIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDCertificatesCertificateIDTooManyRequests creates a DeleteKeysKeyIDCertificatesCertificateIDTooManyRequests with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDTooManyRequests() *DeleteKeysKeyIDCertificatesCertificateIDTooManyRequests {
	return &DeleteKeysKeyIDCertificatesCertificateIDTooManyRequests{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteKeysKeyIDCertificatesCertificateIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDCertificatesCertificateIDInternalServerError creates a DeleteKeysKeyIDCertificatesCertificateIDInternalServerError with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDInternalServerError() *DeleteKeysKeyIDCertificatesCertificateIDInternalServerError {
	return &DeleteKeysKeyIDCertificatesCertificateIDInternalServerError{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteKeysKeyIDCertificatesCertificateIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable creates a DeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable with default headers values
func NewDeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable() *DeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable {
	return &DeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable{}
}

/*DeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}/certificates/{certificate_id}][%d] deleteKeysKeyIdCertificatesCertificateIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteKeysKeyIDCertificatesCertificateIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
