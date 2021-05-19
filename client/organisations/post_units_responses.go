// Code generated by go-swagger; DO NOT EDIT.

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostUnitsReader is a Reader for the PostUnits structure.
type PostUnitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUnitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostUnitsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUnitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUnitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUnitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUnitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostUnitsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUnitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUnitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUnitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUnitsCreated creates a PostUnitsCreated with default headers values
func NewPostUnitsCreated() *PostUnitsCreated {
	return &PostUnitsCreated{}
}

/*PostUnitsCreated handles this case with default header values.

Organisation creation response
*/
type PostUnitsCreated struct {
	Payload *models.OrganisationCreationResponse
}

func (o *PostUnitsCreated) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsCreated  %+v", 201, o.Payload)
}

func (o *PostUnitsCreated) GetPayload() *models.OrganisationCreationResponse {
	return o.Payload
}

func (o *PostUnitsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganisationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnitsBadRequest creates a PostUnitsBadRequest with default headers values
func NewPostUnitsBadRequest() *PostUnitsBadRequest {
	return &PostUnitsBadRequest{}
}

/*PostUnitsBadRequest handles this case with default header values.

Bad Request
*/
type PostUnitsBadRequest struct {
	Payload *models.APIError
}

func (o *PostUnitsBadRequest) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsBadRequest  %+v", 400, o.Payload)
}

func (o *PostUnitsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUnitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnitsUnauthorized creates a PostUnitsUnauthorized with default headers values
func NewPostUnitsUnauthorized() *PostUnitsUnauthorized {
	return &PostUnitsUnauthorized{}
}

/*PostUnitsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostUnitsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostUnitsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUnitsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUnitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnitsForbidden creates a PostUnitsForbidden with default headers values
func NewPostUnitsForbidden() *PostUnitsForbidden {
	return &PostUnitsForbidden{}
}

/*PostUnitsForbidden handles this case with default header values.

Forbidden
*/
type PostUnitsForbidden struct {
	Payload *models.APIError
}

func (o *PostUnitsForbidden) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsForbidden  %+v", 403, o.Payload)
}

func (o *PostUnitsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUnitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnitsNotFound creates a PostUnitsNotFound with default headers values
func NewPostUnitsNotFound() *PostUnitsNotFound {
	return &PostUnitsNotFound{}
}

/*PostUnitsNotFound handles this case with default header values.

Record not found
*/
type PostUnitsNotFound struct {
	Payload *models.APIError
}

func (o *PostUnitsNotFound) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsNotFound  %+v", 404, o.Payload)
}

func (o *PostUnitsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUnitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnitsConflict creates a PostUnitsConflict with default headers values
func NewPostUnitsConflict() *PostUnitsConflict {
	return &PostUnitsConflict{}
}

/*PostUnitsConflict handles this case with default header values.

Conflict
*/
type PostUnitsConflict struct {
	Payload *models.APIError
}

func (o *PostUnitsConflict) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsConflict  %+v", 409, o.Payload)
}

func (o *PostUnitsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUnitsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnitsTooManyRequests creates a PostUnitsTooManyRequests with default headers values
func NewPostUnitsTooManyRequests() *PostUnitsTooManyRequests {
	return &PostUnitsTooManyRequests{}
}

/*PostUnitsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostUnitsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostUnitsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUnitsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUnitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnitsInternalServerError creates a PostUnitsInternalServerError with default headers values
func NewPostUnitsInternalServerError() *PostUnitsInternalServerError {
	return &PostUnitsInternalServerError{}
}

/*PostUnitsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostUnitsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostUnitsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUnitsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUnitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnitsServiceUnavailable creates a PostUnitsServiceUnavailable with default headers values
func NewPostUnitsServiceUnavailable() *PostUnitsServiceUnavailable {
	return &PostUnitsServiceUnavailable{}
}

/*PostUnitsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostUnitsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostUnitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /units][%d] postUnitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUnitsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostUnitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
