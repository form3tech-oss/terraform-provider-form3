// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostPaymentsIDReturnsReader is a Reader for the PostPaymentsIDReturns structure.
type PostPaymentsIDReturnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReturnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDReturnsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDReturnsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostPaymentsIDReturnsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostPaymentsIDReturnsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostPaymentsIDReturnsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostPaymentsIDReturnsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostPaymentsIDReturnsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostPaymentsIDReturnsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPostPaymentsIDReturnsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReturnsCreated creates a PostPaymentsIDReturnsCreated with default headers values
func NewPostPaymentsIDReturnsCreated() *PostPaymentsIDReturnsCreated {
	return &PostPaymentsIDReturnsCreated{}
}

/*PostPaymentsIDReturnsCreated handles this case with default header values.

Return creation response
*/
type PostPaymentsIDReturnsCreated struct {
	Payload *models.ReturnCreationResponse
}

func (o *PostPaymentsIDReturnsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReturnsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsBadRequest creates a PostPaymentsIDReturnsBadRequest with default headers values
func NewPostPaymentsIDReturnsBadRequest() *PostPaymentsIDReturnsBadRequest {
	return &PostPaymentsIDReturnsBadRequest{}
}

/*PostPaymentsIDReturnsBadRequest handles this case with default header values.

Bad Request
*/
type PostPaymentsIDReturnsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReturnsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsUnauthorized creates a PostPaymentsIDReturnsUnauthorized with default headers values
func NewPostPaymentsIDReturnsUnauthorized() *PostPaymentsIDReturnsUnauthorized {
	return &PostPaymentsIDReturnsUnauthorized{}
}

/*PostPaymentsIDReturnsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostPaymentsIDReturnsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPaymentsIDReturnsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsForbidden creates a PostPaymentsIDReturnsForbidden with default headers values
func NewPostPaymentsIDReturnsForbidden() *PostPaymentsIDReturnsForbidden {
	return &PostPaymentsIDReturnsForbidden{}
}

/*PostPaymentsIDReturnsForbidden handles this case with default header values.

Forbidden
*/
type PostPaymentsIDReturnsForbidden struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsForbidden) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsForbidden  %+v", 403, o.Payload)
}

func (o *PostPaymentsIDReturnsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsNotFound creates a PostPaymentsIDReturnsNotFound with default headers values
func NewPostPaymentsIDReturnsNotFound() *PostPaymentsIDReturnsNotFound {
	return &PostPaymentsIDReturnsNotFound{}
}

/*PostPaymentsIDReturnsNotFound handles this case with default header values.

Record not found
*/
type PostPaymentsIDReturnsNotFound struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsNotFound) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsNotFound  %+v", 404, o.Payload)
}

func (o *PostPaymentsIDReturnsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsConflict creates a PostPaymentsIDReturnsConflict with default headers values
func NewPostPaymentsIDReturnsConflict() *PostPaymentsIDReturnsConflict {
	return &PostPaymentsIDReturnsConflict{}
}

/*PostPaymentsIDReturnsConflict handles this case with default header values.

Conflict
*/
type PostPaymentsIDReturnsConflict struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsConflict) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsConflict  %+v", 409, o.Payload)
}

func (o *PostPaymentsIDReturnsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsTooManyRequests creates a PostPaymentsIDReturnsTooManyRequests with default headers values
func NewPostPaymentsIDReturnsTooManyRequests() *PostPaymentsIDReturnsTooManyRequests {
	return &PostPaymentsIDReturnsTooManyRequests{}
}

/*PostPaymentsIDReturnsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostPaymentsIDReturnsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPaymentsIDReturnsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsInternalServerError creates a PostPaymentsIDReturnsInternalServerError with default headers values
func NewPostPaymentsIDReturnsInternalServerError() *PostPaymentsIDReturnsInternalServerError {
	return &PostPaymentsIDReturnsInternalServerError{}
}

/*PostPaymentsIDReturnsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPaymentsIDReturnsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPaymentsIDReturnsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsServiceUnavailable creates a PostPaymentsIDReturnsServiceUnavailable with default headers values
func NewPostPaymentsIDReturnsServiceUnavailable() *PostPaymentsIDReturnsServiceUnavailable {
	return &PostPaymentsIDReturnsServiceUnavailable{}
}

/*PostPaymentsIDReturnsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostPaymentsIDReturnsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns][%d] postPaymentsIdReturnsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPaymentsIDReturnsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
