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

// PostKeysReader is a Reader for the PostKeys structure.
type PostKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostKeysCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostKeysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostKeysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostKeysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostKeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostKeysConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostKeysTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostKeysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostKeysServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostKeysCreated creates a PostKeysCreated with default headers values
func NewPostKeysCreated() *PostKeysCreated {
	return &PostKeysCreated{}
}

/* PostKeysCreated describes a response with status code 201, with default header values.

creation response
*/
type PostKeysCreated struct {
	Payload *models.KeyCreationResponse
}

func (o *PostKeysCreated) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysCreated  %+v", 201, o.Payload)
}
func (o *PostKeysCreated) GetPayload() *models.KeyCreationResponse {
	return o.Payload
}

func (o *PostKeysCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysBadRequest creates a PostKeysBadRequest with default headers values
func NewPostKeysBadRequest() *PostKeysBadRequest {
	return &PostKeysBadRequest{}
}

/* PostKeysBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostKeysBadRequest struct {
	Payload *models.APIError
}

func (o *PostKeysBadRequest) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysBadRequest  %+v", 400, o.Payload)
}
func (o *PostKeysBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostKeysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysUnauthorized creates a PostKeysUnauthorized with default headers values
func NewPostKeysUnauthorized() *PostKeysUnauthorized {
	return &PostKeysUnauthorized{}
}

/* PostKeysUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type PostKeysUnauthorized struct {
	Payload *models.APIError
}

func (o *PostKeysUnauthorized) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysUnauthorized  %+v", 401, o.Payload)
}
func (o *PostKeysUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostKeysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysForbidden creates a PostKeysForbidden with default headers values
func NewPostKeysForbidden() *PostKeysForbidden {
	return &PostKeysForbidden{}
}

/* PostKeysForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostKeysForbidden struct {
	Payload *models.APIError
}

func (o *PostKeysForbidden) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysForbidden  %+v", 403, o.Payload)
}
func (o *PostKeysForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostKeysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysNotFound creates a PostKeysNotFound with default headers values
func NewPostKeysNotFound() *PostKeysNotFound {
	return &PostKeysNotFound{}
}

/* PostKeysNotFound describes a response with status code 404, with default header values.

Record not found
*/
type PostKeysNotFound struct {
	Payload *models.APIError
}

func (o *PostKeysNotFound) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysNotFound  %+v", 404, o.Payload)
}
func (o *PostKeysNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostKeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysConflict creates a PostKeysConflict with default headers values
func NewPostKeysConflict() *PostKeysConflict {
	return &PostKeysConflict{}
}

/* PostKeysConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostKeysConflict struct {
	Payload *models.APIError
}

func (o *PostKeysConflict) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysConflict  %+v", 409, o.Payload)
}
func (o *PostKeysConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostKeysConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysTooManyRequests creates a PostKeysTooManyRequests with default headers values
func NewPostKeysTooManyRequests() *PostKeysTooManyRequests {
	return &PostKeysTooManyRequests{}
}

/* PostKeysTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostKeysTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostKeysTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysTooManyRequests  %+v", 429, o.Payload)
}
func (o *PostKeysTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostKeysTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysInternalServerError creates a PostKeysInternalServerError with default headers values
func NewPostKeysInternalServerError() *PostKeysInternalServerError {
	return &PostKeysInternalServerError{}
}

/* PostKeysInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostKeysInternalServerError struct {
	Payload *models.APIError
}

func (o *PostKeysInternalServerError) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysInternalServerError  %+v", 500, o.Payload)
}
func (o *PostKeysInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostKeysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKeysServiceUnavailable creates a PostKeysServiceUnavailable with default headers values
func NewPostKeysServiceUnavailable() *PostKeysServiceUnavailable {
	return &PostKeysServiceUnavailable{}
}

/* PostKeysServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostKeysServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostKeysServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostKeysServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostKeysServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
