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

// GetKeysKeyIDCertificatesCertificateIDReader is a Reader for the GetKeysKeyIDCertificatesCertificateID structure.
type GetKeysKeyIDCertificatesCertificateIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeysKeyIDCertificatesCertificateIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKeysKeyIDCertificatesCertificateIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKeysKeyIDCertificatesCertificateIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKeysKeyIDCertificatesCertificateIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKeysKeyIDCertificatesCertificateIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKeysKeyIDCertificatesCertificateIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetKeysKeyIDCertificatesCertificateIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKeysKeyIDCertificatesCertificateIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKeysKeyIDCertificatesCertificateIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKeysKeyIDCertificatesCertificateIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKeysKeyIDCertificatesCertificateIDOK creates a GetKeysKeyIDCertificatesCertificateIDOK with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDOK() *GetKeysKeyIDCertificatesCertificateIDOK {
	return &GetKeysKeyIDCertificatesCertificateIDOK{}
}

/* GetKeysKeyIDCertificatesCertificateIDOK describes a response with status code 200, with default header values.

Certificate details
*/
type GetKeysKeyIDCertificatesCertificateIDOK struct {
	Payload *models.CertificateDetailsResponse
}

func (o *GetKeysKeyIDCertificatesCertificateIDOK) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdOK  %+v", 200, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDOK) GetPayload() *models.CertificateDetailsResponse {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysKeyIDCertificatesCertificateIDBadRequest creates a GetKeysKeyIDCertificatesCertificateIDBadRequest with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDBadRequest() *GetKeysKeyIDCertificatesCertificateIDBadRequest {
	return &GetKeysKeyIDCertificatesCertificateIDBadRequest{}
}

/* GetKeysKeyIDCertificatesCertificateIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetKeysKeyIDCertificatesCertificateIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetKeysKeyIDCertificatesCertificateIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysKeyIDCertificatesCertificateIDUnauthorized creates a GetKeysKeyIDCertificatesCertificateIDUnauthorized with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDUnauthorized() *GetKeysKeyIDCertificatesCertificateIDUnauthorized {
	return &GetKeysKeyIDCertificatesCertificateIDUnauthorized{}
}

/* GetKeysKeyIDCertificatesCertificateIDUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetKeysKeyIDCertificatesCertificateIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetKeysKeyIDCertificatesCertificateIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysKeyIDCertificatesCertificateIDForbidden creates a GetKeysKeyIDCertificatesCertificateIDForbidden with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDForbidden() *GetKeysKeyIDCertificatesCertificateIDForbidden {
	return &GetKeysKeyIDCertificatesCertificateIDForbidden{}
}

/* GetKeysKeyIDCertificatesCertificateIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetKeysKeyIDCertificatesCertificateIDForbidden struct {
	Payload *models.APIError
}

func (o *GetKeysKeyIDCertificatesCertificateIDForbidden) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdForbidden  %+v", 403, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysKeyIDCertificatesCertificateIDNotFound creates a GetKeysKeyIDCertificatesCertificateIDNotFound with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDNotFound() *GetKeysKeyIDCertificatesCertificateIDNotFound {
	return &GetKeysKeyIDCertificatesCertificateIDNotFound{}
}

/* GetKeysKeyIDCertificatesCertificateIDNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetKeysKeyIDCertificatesCertificateIDNotFound struct {
	Payload *models.APIError
}

func (o *GetKeysKeyIDCertificatesCertificateIDNotFound) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdNotFound  %+v", 404, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysKeyIDCertificatesCertificateIDConflict creates a GetKeysKeyIDCertificatesCertificateIDConflict with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDConflict() *GetKeysKeyIDCertificatesCertificateIDConflict {
	return &GetKeysKeyIDCertificatesCertificateIDConflict{}
}

/* GetKeysKeyIDCertificatesCertificateIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetKeysKeyIDCertificatesCertificateIDConflict struct {
	Payload *models.APIError
}

func (o *GetKeysKeyIDCertificatesCertificateIDConflict) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdConflict  %+v", 409, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysKeyIDCertificatesCertificateIDTooManyRequests creates a GetKeysKeyIDCertificatesCertificateIDTooManyRequests with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDTooManyRequests() *GetKeysKeyIDCertificatesCertificateIDTooManyRequests {
	return &GetKeysKeyIDCertificatesCertificateIDTooManyRequests{}
}

/* GetKeysKeyIDCertificatesCertificateIDTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetKeysKeyIDCertificatesCertificateIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetKeysKeyIDCertificatesCertificateIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysKeyIDCertificatesCertificateIDInternalServerError creates a GetKeysKeyIDCertificatesCertificateIDInternalServerError with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDInternalServerError() *GetKeysKeyIDCertificatesCertificateIDInternalServerError {
	return &GetKeysKeyIDCertificatesCertificateIDInternalServerError{}
}

/* GetKeysKeyIDCertificatesCertificateIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetKeysKeyIDCertificatesCertificateIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetKeysKeyIDCertificatesCertificateIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysKeyIDCertificatesCertificateIDServiceUnavailable creates a GetKeysKeyIDCertificatesCertificateIDServiceUnavailable with default headers values
func NewGetKeysKeyIDCertificatesCertificateIDServiceUnavailable() *GetKeysKeyIDCertificatesCertificateIDServiceUnavailable {
	return &GetKeysKeyIDCertificatesCertificateIDServiceUnavailable{}
}

/* GetKeysKeyIDCertificatesCertificateIDServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetKeysKeyIDCertificatesCertificateIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetKeysKeyIDCertificatesCertificateIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates/{certificate_id}][%d] getKeysKeyIdCertificatesCertificateIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetKeysKeyIDCertificatesCertificateIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesCertificateIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
