// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostPaymentsReader is a Reader for the PostPayments structure.
type PostPaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostPaymentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPaymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPaymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPaymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPaymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostPaymentsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPaymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPaymentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPaymentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsCreated creates a PostPaymentsCreated with default headers values
func NewPostPaymentsCreated() *PostPaymentsCreated {
	return &PostPaymentsCreated{}
}

/*PostPaymentsCreated handles this case with default header values.

Payment creation response
*/
type PostPaymentsCreated struct {
	Payload *models.PaymentCreationResponse
}

func (o *PostPaymentsCreated) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsCreated) GetPayload() *models.PaymentCreationResponse {
	return o.Payload
}

func (o *PostPaymentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsBadRequest creates a PostPaymentsBadRequest with default headers values
func NewPostPaymentsBadRequest() *PostPaymentsBadRequest {
	return &PostPaymentsBadRequest{}
}

/*PostPaymentsBadRequest handles this case with default header values.

Bad Request
*/
type PostPaymentsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsUnauthorized creates a PostPaymentsUnauthorized with default headers values
func NewPostPaymentsUnauthorized() *PostPaymentsUnauthorized {
	return &PostPaymentsUnauthorized{}
}

/*PostPaymentsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostPaymentsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostPaymentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPaymentsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsForbidden creates a PostPaymentsForbidden with default headers values
func NewPostPaymentsForbidden() *PostPaymentsForbidden {
	return &PostPaymentsForbidden{}
}

/*PostPaymentsForbidden handles this case with default header values.

Forbidden
*/
type PostPaymentsForbidden struct {
	Payload *models.APIError
}

func (o *PostPaymentsForbidden) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsForbidden  %+v", 403, o.Payload)
}

func (o *PostPaymentsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsNotFound creates a PostPaymentsNotFound with default headers values
func NewPostPaymentsNotFound() *PostPaymentsNotFound {
	return &PostPaymentsNotFound{}
}

/*PostPaymentsNotFound handles this case with default header values.

Record not found
*/
type PostPaymentsNotFound struct {
	Payload *models.APIError
}

func (o *PostPaymentsNotFound) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsNotFound  %+v", 404, o.Payload)
}

func (o *PostPaymentsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsConflict creates a PostPaymentsConflict with default headers values
func NewPostPaymentsConflict() *PostPaymentsConflict {
	return &PostPaymentsConflict{}
}

/*PostPaymentsConflict handles this case with default header values.

Conflict
*/
type PostPaymentsConflict struct {
	Payload *models.APIError
}

func (o *PostPaymentsConflict) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsConflict  %+v", 409, o.Payload)
}

func (o *PostPaymentsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsTooManyRequests creates a PostPaymentsTooManyRequests with default headers values
func NewPostPaymentsTooManyRequests() *PostPaymentsTooManyRequests {
	return &PostPaymentsTooManyRequests{}
}

/*PostPaymentsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostPaymentsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostPaymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPaymentsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsInternalServerError creates a PostPaymentsInternalServerError with default headers values
func NewPostPaymentsInternalServerError() *PostPaymentsInternalServerError {
	return &PostPaymentsInternalServerError{}
}

/*PostPaymentsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPaymentsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostPaymentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPaymentsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsServiceUnavailable creates a PostPaymentsServiceUnavailable with default headers values
func NewPostPaymentsServiceUnavailable() *PostPaymentsServiceUnavailable {
	return &PostPaymentsServiceUnavailable{}
}

/*PostPaymentsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostPaymentsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostPaymentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /payments][%d] postPaymentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPaymentsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
