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

// GetUsersUserIDCredentialsReader is a Reader for the GetUsersUserIDCredentials structure.
type GetUsersUserIDCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersUserIDCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersUserIDCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersUserIDCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersUserIDCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersUserIDCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetUsersUserIDCredentialsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUsersUserIDCredentialsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersUserIDCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUsersUserIDCredentialsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersUserIDCredentialsOK creates a GetUsersUserIDCredentialsOK with default headers values
func NewGetUsersUserIDCredentialsOK() *GetUsersUserIDCredentialsOK {
	return &GetUsersUserIDCredentialsOK{}
}

/*GetUsersUserIDCredentialsOK handles this case with default header values.

List of credentials for user
*/
type GetUsersUserIDCredentialsOK struct {
	Payload *models.UserCredentialListResponse
}

func (o *GetUsersUserIDCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetUsersUserIDCredentialsOK) GetPayload() *models.UserCredentialListResponse {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserCredentialListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsBadRequest creates a GetUsersUserIDCredentialsBadRequest with default headers values
func NewGetUsersUserIDCredentialsBadRequest() *GetUsersUserIDCredentialsBadRequest {
	return &GetUsersUserIDCredentialsBadRequest{}
}

/*GetUsersUserIDCredentialsBadRequest handles this case with default header values.

Bad Request
*/
type GetUsersUserIDCredentialsBadRequest struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersUserIDCredentialsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsUnauthorized creates a GetUsersUserIDCredentialsUnauthorized with default headers values
func NewGetUsersUserIDCredentialsUnauthorized() *GetUsersUserIDCredentialsUnauthorized {
	return &GetUsersUserIDCredentialsUnauthorized{}
}

/*GetUsersUserIDCredentialsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetUsersUserIDCredentialsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsersUserIDCredentialsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsForbidden creates a GetUsersUserIDCredentialsForbidden with default headers values
func NewGetUsersUserIDCredentialsForbidden() *GetUsersUserIDCredentialsForbidden {
	return &GetUsersUserIDCredentialsForbidden{}
}

/*GetUsersUserIDCredentialsForbidden handles this case with default header values.

Forbidden
*/
type GetUsersUserIDCredentialsForbidden struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersUserIDCredentialsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsNotFound creates a GetUsersUserIDCredentialsNotFound with default headers values
func NewGetUsersUserIDCredentialsNotFound() *GetUsersUserIDCredentialsNotFound {
	return &GetUsersUserIDCredentialsNotFound{}
}

/*GetUsersUserIDCredentialsNotFound handles this case with default header values.

Record not found
*/
type GetUsersUserIDCredentialsNotFound struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *GetUsersUserIDCredentialsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsConflict creates a GetUsersUserIDCredentialsConflict with default headers values
func NewGetUsersUserIDCredentialsConflict() *GetUsersUserIDCredentialsConflict {
	return &GetUsersUserIDCredentialsConflict{}
}

/*GetUsersUserIDCredentialsConflict handles this case with default header values.

Conflict
*/
type GetUsersUserIDCredentialsConflict struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsConflict) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsConflict  %+v", 409, o.Payload)
}

func (o *GetUsersUserIDCredentialsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsTooManyRequests creates a GetUsersUserIDCredentialsTooManyRequests with default headers values
func NewGetUsersUserIDCredentialsTooManyRequests() *GetUsersUserIDCredentialsTooManyRequests {
	return &GetUsersUserIDCredentialsTooManyRequests{}
}

/*GetUsersUserIDCredentialsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetUsersUserIDCredentialsTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUsersUserIDCredentialsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsInternalServerError creates a GetUsersUserIDCredentialsInternalServerError with default headers values
func NewGetUsersUserIDCredentialsInternalServerError() *GetUsersUserIDCredentialsInternalServerError {
	return &GetUsersUserIDCredentialsInternalServerError{}
}

/*GetUsersUserIDCredentialsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetUsersUserIDCredentialsInternalServerError struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsersUserIDCredentialsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDCredentialsServiceUnavailable creates a GetUsersUserIDCredentialsServiceUnavailable with default headers values
func NewGetUsersUserIDCredentialsServiceUnavailable() *GetUsersUserIDCredentialsServiceUnavailable {
	return &GetUsersUserIDCredentialsServiceUnavailable{}
}

/*GetUsersUserIDCredentialsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetUsersUserIDCredentialsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDCredentialsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials][%d] getUsersUserIdCredentialsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUsersUserIDCredentialsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDCredentialsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
