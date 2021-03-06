// Code generated by go-swagger; DO NOT EDIT.

package platformsecurityapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetPlatformSecuritySigningKeysSigningkeyIDReader is a Reader for the GetPlatformSecuritySigningKeysSigningkeyID structure.
type GetPlatformSecuritySigningKeysSigningkeyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformSecuritySigningKeysSigningkeyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDOK creates a GetPlatformSecuritySigningKeysSigningkeyIDOK with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDOK() *GetPlatformSecuritySigningKeysSigningkeyIDOK {
	return &GetPlatformSecuritySigningKeysSigningkeyIDOK{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDOK handles this case with default header values.

signing key response
*/
type GetPlatformSecuritySigningKeysSigningkeyIDOK struct {
	Payload *models.SigningKeysResponse
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDOK) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdOK  %+v", 200, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDOK) GetPayload() *models.SigningKeysResponse {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SigningKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDBadRequest creates a GetPlatformSecuritySigningKeysSigningkeyIDBadRequest with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDBadRequest() *GetPlatformSecuritySigningKeysSigningkeyIDBadRequest {
	return &GetPlatformSecuritySigningKeysSigningkeyIDBadRequest{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDBadRequest handles this case with default header values.

Bad Request
*/
type GetPlatformSecuritySigningKeysSigningkeyIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDUnauthorized creates a GetPlatformSecuritySigningKeysSigningkeyIDUnauthorized with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDUnauthorized() *GetPlatformSecuritySigningKeysSigningkeyIDUnauthorized {
	return &GetPlatformSecuritySigningKeysSigningkeyIDUnauthorized{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetPlatformSecuritySigningKeysSigningkeyIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDForbidden creates a GetPlatformSecuritySigningKeysSigningkeyIDForbidden with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDForbidden() *GetPlatformSecuritySigningKeysSigningkeyIDForbidden {
	return &GetPlatformSecuritySigningKeysSigningkeyIDForbidden{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDForbidden handles this case with default header values.

Forbidden
*/
type GetPlatformSecuritySigningKeysSigningkeyIDForbidden struct {
	Payload *models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDForbidden) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdForbidden  %+v", 403, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDNotFound creates a GetPlatformSecuritySigningKeysSigningkeyIDNotFound with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDNotFound() *GetPlatformSecuritySigningKeysSigningkeyIDNotFound {
	return &GetPlatformSecuritySigningKeysSigningkeyIDNotFound{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDNotFound handles this case with default header values.

Record not found
*/
type GetPlatformSecuritySigningKeysSigningkeyIDNotFound struct {
	Payload *models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdNotFound  %+v", 404, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDConflict creates a GetPlatformSecuritySigningKeysSigningkeyIDConflict with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDConflict() *GetPlatformSecuritySigningKeysSigningkeyIDConflict {
	return &GetPlatformSecuritySigningKeysSigningkeyIDConflict{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDConflict handles this case with default header values.

Conflict
*/
type GetPlatformSecuritySigningKeysSigningkeyIDConflict struct {
	Payload *models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDConflict) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdConflict  %+v", 409, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests creates a GetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests() *GetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests {
	return &GetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDInternalServerError creates a GetPlatformSecuritySigningKeysSigningkeyIDInternalServerError with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDInternalServerError() *GetPlatformSecuritySigningKeysSigningkeyIDInternalServerError {
	return &GetPlatformSecuritySigningKeysSigningkeyIDInternalServerError{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPlatformSecuritySigningKeysSigningkeyIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable creates a GetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable() *GetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable {
	return &GetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
