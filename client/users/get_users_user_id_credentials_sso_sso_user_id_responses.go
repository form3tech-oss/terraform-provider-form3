// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetUsersUserIDCredentialsSsoSsoUserIDReader is a Reader for the GetUsersUserIDCredentialsSsoSsoUserID structure.
type GetUsersUserIDCredentialsSsoSsoUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDCredentialsSsoSsoUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDOK creates a GetUsersUserIDCredentialsSsoSsoUserIDOK with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDOK() *GetUsersUserIDCredentialsSsoSsoUserIDOK {
	return &GetUsersUserIDCredentialsSsoSsoUserIDOK{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDOK describes a response with status code 200, with default header values.

Sso user data
*/
type GetUsersUserIDCredentialsSsoSsoUserIDOK struct {
	Payload *models.SsoUserDetailsResponse
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdOK  %+v", 200, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDOK) GetPayload() *models.SsoUserDetailsResponse {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SsoUserDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDBadRequest creates a GetUsersUserIDCredentialsSsoSsoUserIDBadRequest with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDBadRequest() *GetUsersUserIDCredentialsSsoSsoUserIDBadRequest {
	return &GetUsersUserIDCredentialsSsoSsoUserIDBadRequest{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetUsersUserIDCredentialsSsoSsoUserIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDUnauthorized creates a GetUsersUserIDCredentialsSsoSsoUserIDUnauthorized with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDUnauthorized() *GetUsersUserIDCredentialsSsoSsoUserIDUnauthorized {
	return &GetUsersUserIDCredentialsSsoSsoUserIDUnauthorized{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetUsersUserIDCredentialsSsoSsoUserIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDForbidden creates a GetUsersUserIDCredentialsSsoSsoUserIDForbidden with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDForbidden() *GetUsersUserIDCredentialsSsoSsoUserIDForbidden {
	return &GetUsersUserIDCredentialsSsoSsoUserIDForbidden{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsersUserIDCredentialsSsoSsoUserIDForbidden struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdForbidden  %+v", 403, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDNotFound creates a GetUsersUserIDCredentialsSsoSsoUserIDNotFound with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDNotFound() *GetUsersUserIDCredentialsSsoSsoUserIDNotFound {
	return &GetUsersUserIDCredentialsSsoSsoUserIDNotFound{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetUsersUserIDCredentialsSsoSsoUserIDNotFound struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdNotFound  %+v", 404, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDConflict creates a GetUsersUserIDCredentialsSsoSsoUserIDConflict with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDConflict() *GetUsersUserIDCredentialsSsoSsoUserIDConflict {
	return &GetUsersUserIDCredentialsSsoSsoUserIDConflict{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetUsersUserIDCredentialsSsoSsoUserIDConflict struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDConflict) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdConflict  %+v", 409, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests creates a GetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests() *GetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests {
	return &GetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDInternalServerError creates a GetUsersUserIDCredentialsSsoSsoUserIDInternalServerError with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDInternalServerError() *GetUsersUserIDCredentialsSsoSsoUserIDInternalServerError {
	return &GetUsersUserIDCredentialsSsoSsoUserIDInternalServerError{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetUsersUserIDCredentialsSsoSsoUserIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable creates a GetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable with default headers values
func NewGetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable() *GetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable {
	return &GetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable{}
}

/* GetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/sso/{sso_user_id}][%d] getUsersUserIdCredentialsSsoSsoUserIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsSsoSsoUserIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
