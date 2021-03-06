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

// PostUsersUserIDCredentialsReader is a Reader for the PostUsersUserIDCredentials structure.
type PostUsersUserIDCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersUserIDCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostUsersUserIDCredentialsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsersUserIDCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUsersUserIDCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUsersUserIDCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUsersUserIDCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostUsersUserIDCredentialsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUsersUserIDCredentialsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUsersUserIDCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUsersUserIDCredentialsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUsersUserIDCredentialsCreated creates a PostUsersUserIDCredentialsCreated with default headers values
func NewPostUsersUserIDCredentialsCreated() *PostUsersUserIDCredentialsCreated {
	return &PostUsersUserIDCredentialsCreated{}
}

/*PostUsersUserIDCredentialsCreated handles this case with default header values.

Credential creation response
*/
type PostUsersUserIDCredentialsCreated struct {
	Payload *models.CredentialCreationResponse
}

func (o *PostUsersUserIDCredentialsCreated) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsCreated  %+v", 201, o.Payload)
}

func (o *PostUsersUserIDCredentialsCreated) GetPayload() *models.CredentialCreationResponse {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUserIDCredentialsBadRequest creates a PostUsersUserIDCredentialsBadRequest with default headers values
func NewPostUsersUserIDCredentialsBadRequest() *PostUsersUserIDCredentialsBadRequest {
	return &PostUsersUserIDCredentialsBadRequest{}
}

/*PostUsersUserIDCredentialsBadRequest handles this case with default header values.

Bad Request
*/
type PostUsersUserIDCredentialsBadRequest struct {
	Payload *models.APIError
}

func (o *PostUsersUserIDCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *PostUsersUserIDCredentialsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUserIDCredentialsUnauthorized creates a PostUsersUserIDCredentialsUnauthorized with default headers values
func NewPostUsersUserIDCredentialsUnauthorized() *PostUsersUserIDCredentialsUnauthorized {
	return &PostUsersUserIDCredentialsUnauthorized{}
}

/*PostUsersUserIDCredentialsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostUsersUserIDCredentialsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostUsersUserIDCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUsersUserIDCredentialsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUserIDCredentialsForbidden creates a PostUsersUserIDCredentialsForbidden with default headers values
func NewPostUsersUserIDCredentialsForbidden() *PostUsersUserIDCredentialsForbidden {
	return &PostUsersUserIDCredentialsForbidden{}
}

/*PostUsersUserIDCredentialsForbidden handles this case with default header values.

Forbidden
*/
type PostUsersUserIDCredentialsForbidden struct {
	Payload *models.APIError
}

func (o *PostUsersUserIDCredentialsForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *PostUsersUserIDCredentialsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUserIDCredentialsNotFound creates a PostUsersUserIDCredentialsNotFound with default headers values
func NewPostUsersUserIDCredentialsNotFound() *PostUsersUserIDCredentialsNotFound {
	return &PostUsersUserIDCredentialsNotFound{}
}

/*PostUsersUserIDCredentialsNotFound handles this case with default header values.

Record not found
*/
type PostUsersUserIDCredentialsNotFound struct {
	Payload *models.APIError
}

func (o *PostUsersUserIDCredentialsNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *PostUsersUserIDCredentialsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUserIDCredentialsConflict creates a PostUsersUserIDCredentialsConflict with default headers values
func NewPostUsersUserIDCredentialsConflict() *PostUsersUserIDCredentialsConflict {
	return &PostUsersUserIDCredentialsConflict{}
}

/*PostUsersUserIDCredentialsConflict handles this case with default header values.

Conflict
*/
type PostUsersUserIDCredentialsConflict struct {
	Payload *models.APIError
}

func (o *PostUsersUserIDCredentialsConflict) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsConflict  %+v", 409, o.Payload)
}

func (o *PostUsersUserIDCredentialsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUserIDCredentialsTooManyRequests creates a PostUsersUserIDCredentialsTooManyRequests with default headers values
func NewPostUsersUserIDCredentialsTooManyRequests() *PostUsersUserIDCredentialsTooManyRequests {
	return &PostUsersUserIDCredentialsTooManyRequests{}
}

/*PostUsersUserIDCredentialsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostUsersUserIDCredentialsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostUsersUserIDCredentialsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUsersUserIDCredentialsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUserIDCredentialsInternalServerError creates a PostUsersUserIDCredentialsInternalServerError with default headers values
func NewPostUsersUserIDCredentialsInternalServerError() *PostUsersUserIDCredentialsInternalServerError {
	return &PostUsersUserIDCredentialsInternalServerError{}
}

/*PostUsersUserIDCredentialsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostUsersUserIDCredentialsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostUsersUserIDCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUsersUserIDCredentialsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUserIDCredentialsServiceUnavailable creates a PostUsersUserIDCredentialsServiceUnavailable with default headers values
func NewPostUsersUserIDCredentialsServiceUnavailable() *PostUsersUserIDCredentialsServiceUnavailable {
	return &PostUsersUserIDCredentialsServiceUnavailable{}
}

/*PostUsersUserIDCredentialsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostUsersUserIDCredentialsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostUsersUserIDCredentialsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUsersUserIDCredentialsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUsersUserIDCredentialsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
