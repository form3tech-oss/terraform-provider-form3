// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostKeysKeyIDCertificatesReader is a Reader for the PostKeysKeyIDCertificates structure.
type PostKeysKeyIDCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKeysKeyIDCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostKeysKeyIDCertificatesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostKeysKeyIDCertificatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostKeysKeyIDCertificatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostKeysKeyIDCertificatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostKeysKeyIDCertificatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostKeysKeyIDCertificatesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostKeysKeyIDCertificatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostKeysKeyIDCertificatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPostKeysKeyIDCertificatesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostKeysKeyIDCertificatesCreated creates a PostKeysKeyIDCertificatesCreated with default headers values
func NewPostKeysKeyIDCertificatesCreated() *PostKeysKeyIDCertificatesCreated {
	return &PostKeysKeyIDCertificatesCreated{}
}

/*PostKeysKeyIDCertificatesCreated handles this case with default header values.

creation response
*/
type PostKeysKeyIDCertificatesCreated struct {
	Payload *models.CertificateCreationResponse
}

func (o *PostKeysKeyIDCertificatesCreated) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesCreated  %+v", 201, o.Payload)
}

func (o *PostKeysKeyIDCertificatesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysKeyIDCertificatesBadRequest creates a PostKeysKeyIDCertificatesBadRequest with default headers values
func NewPostKeysKeyIDCertificatesBadRequest() *PostKeysKeyIDCertificatesBadRequest {
	return &PostKeysKeyIDCertificatesBadRequest{}
}

/*PostKeysKeyIDCertificatesBadRequest handles this case with default header values.

Bad Request
*/
type PostKeysKeyIDCertificatesBadRequest struct {
	Payload *models.APIError
}

func (o *PostKeysKeyIDCertificatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesBadRequest  %+v", 400, o.Payload)
}

func (o *PostKeysKeyIDCertificatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysKeyIDCertificatesUnauthorized creates a PostKeysKeyIDCertificatesUnauthorized with default headers values
func NewPostKeysKeyIDCertificatesUnauthorized() *PostKeysKeyIDCertificatesUnauthorized {
	return &PostKeysKeyIDCertificatesUnauthorized{}
}

/*PostKeysKeyIDCertificatesUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostKeysKeyIDCertificatesUnauthorized struct {
	Payload *models.APIError
}

func (o *PostKeysKeyIDCertificatesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostKeysKeyIDCertificatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysKeyIDCertificatesForbidden creates a PostKeysKeyIDCertificatesForbidden with default headers values
func NewPostKeysKeyIDCertificatesForbidden() *PostKeysKeyIDCertificatesForbidden {
	return &PostKeysKeyIDCertificatesForbidden{}
}

/*PostKeysKeyIDCertificatesForbidden handles this case with default header values.

Forbidden
*/
type PostKeysKeyIDCertificatesForbidden struct {
	Payload *models.APIError
}

func (o *PostKeysKeyIDCertificatesForbidden) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesForbidden  %+v", 403, o.Payload)
}

func (o *PostKeysKeyIDCertificatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysKeyIDCertificatesNotFound creates a PostKeysKeyIDCertificatesNotFound with default headers values
func NewPostKeysKeyIDCertificatesNotFound() *PostKeysKeyIDCertificatesNotFound {
	return &PostKeysKeyIDCertificatesNotFound{}
}

/*PostKeysKeyIDCertificatesNotFound handles this case with default header values.

Record not found
*/
type PostKeysKeyIDCertificatesNotFound struct {
	Payload *models.APIError
}

func (o *PostKeysKeyIDCertificatesNotFound) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesNotFound  %+v", 404, o.Payload)
}

func (o *PostKeysKeyIDCertificatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysKeyIDCertificatesConflict creates a PostKeysKeyIDCertificatesConflict with default headers values
func NewPostKeysKeyIDCertificatesConflict() *PostKeysKeyIDCertificatesConflict {
	return &PostKeysKeyIDCertificatesConflict{}
}

/*PostKeysKeyIDCertificatesConflict handles this case with default header values.

Conflict
*/
type PostKeysKeyIDCertificatesConflict struct {
	Payload *models.APIError
}

func (o *PostKeysKeyIDCertificatesConflict) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesConflict  %+v", 409, o.Payload)
}

func (o *PostKeysKeyIDCertificatesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysKeyIDCertificatesTooManyRequests creates a PostKeysKeyIDCertificatesTooManyRequests with default headers values
func NewPostKeysKeyIDCertificatesTooManyRequests() *PostKeysKeyIDCertificatesTooManyRequests {
	return &PostKeysKeyIDCertificatesTooManyRequests{}
}

/*PostKeysKeyIDCertificatesTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostKeysKeyIDCertificatesTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostKeysKeyIDCertificatesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostKeysKeyIDCertificatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysKeyIDCertificatesInternalServerError creates a PostKeysKeyIDCertificatesInternalServerError with default headers values
func NewPostKeysKeyIDCertificatesInternalServerError() *PostKeysKeyIDCertificatesInternalServerError {
	return &PostKeysKeyIDCertificatesInternalServerError{}
}

/*PostKeysKeyIDCertificatesInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostKeysKeyIDCertificatesInternalServerError struct {
	Payload *models.APIError
}

func (o *PostKeysKeyIDCertificatesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostKeysKeyIDCertificatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysKeyIDCertificatesServiceUnavailable creates a PostKeysKeyIDCertificatesServiceUnavailable with default headers values
func NewPostKeysKeyIDCertificatesServiceUnavailable() *PostKeysKeyIDCertificatesServiceUnavailable {
	return &PostKeysKeyIDCertificatesServiceUnavailable{}
}

/*PostKeysKeyIDCertificatesServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostKeysKeyIDCertificatesServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostKeysKeyIDCertificatesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /keys/{key_id}/certificates][%d] postKeysKeyIdCertificatesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostKeysKeyIDCertificatesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
