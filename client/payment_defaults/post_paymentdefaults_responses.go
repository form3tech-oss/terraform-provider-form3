// Code generated by go-swagger; DO NOT EDIT.

package payment_defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostPaymentdefaultsReader is a Reader for the PostPaymentdefaults structure.
type PostPaymentdefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentdefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostPaymentdefaultsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPaymentdefaultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPaymentdefaultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPaymentdefaultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPaymentdefaultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostPaymentdefaultsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPaymentdefaultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPaymentdefaultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPaymentdefaultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostPaymentdefaultsCreated creates a PostPaymentdefaultsCreated with default headers values
func NewPostPaymentdefaultsCreated() *PostPaymentdefaultsCreated {
	return &PostPaymentdefaultsCreated{}
}

/*PostPaymentdefaultsCreated handles this case with default header values.

Default settings created
*/
type PostPaymentdefaultsCreated struct {
	Payload *models.PaymentDefaultsCreateResponse
}

func (o *PostPaymentdefaultsCreated) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentdefaultsCreated) GetPayload() *models.PaymentDefaultsCreateResponse {
	return o.Payload
}

func (o *PostPaymentdefaultsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentDefaultsCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentdefaultsBadRequest creates a PostPaymentdefaultsBadRequest with default headers values
func NewPostPaymentdefaultsBadRequest() *PostPaymentdefaultsBadRequest {
	return &PostPaymentdefaultsBadRequest{}
}

/*PostPaymentdefaultsBadRequest handles this case with default header values.

Bad Request
*/
type PostPaymentdefaultsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentdefaultsBadRequest) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentdefaultsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentdefaultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentdefaultsUnauthorized creates a PostPaymentdefaultsUnauthorized with default headers values
func NewPostPaymentdefaultsUnauthorized() *PostPaymentdefaultsUnauthorized {
	return &PostPaymentdefaultsUnauthorized{}
}

/*PostPaymentdefaultsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostPaymentdefaultsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostPaymentdefaultsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPaymentdefaultsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentdefaultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentdefaultsForbidden creates a PostPaymentdefaultsForbidden with default headers values
func NewPostPaymentdefaultsForbidden() *PostPaymentdefaultsForbidden {
	return &PostPaymentdefaultsForbidden{}
}

/*PostPaymentdefaultsForbidden handles this case with default header values.

Forbidden
*/
type PostPaymentdefaultsForbidden struct {
	Payload *models.APIError
}

func (o *PostPaymentdefaultsForbidden) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsForbidden  %+v", 403, o.Payload)
}

func (o *PostPaymentdefaultsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentdefaultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentdefaultsNotFound creates a PostPaymentdefaultsNotFound with default headers values
func NewPostPaymentdefaultsNotFound() *PostPaymentdefaultsNotFound {
	return &PostPaymentdefaultsNotFound{}
}

/*PostPaymentdefaultsNotFound handles this case with default header values.

Record not found
*/
type PostPaymentdefaultsNotFound struct {
	Payload *models.APIError
}

func (o *PostPaymentdefaultsNotFound) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsNotFound  %+v", 404, o.Payload)
}

func (o *PostPaymentdefaultsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentdefaultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentdefaultsConflict creates a PostPaymentdefaultsConflict with default headers values
func NewPostPaymentdefaultsConflict() *PostPaymentdefaultsConflict {
	return &PostPaymentdefaultsConflict{}
}

/*PostPaymentdefaultsConflict handles this case with default header values.

Conflict
*/
type PostPaymentdefaultsConflict struct {
	Payload *models.APIError
}

func (o *PostPaymentdefaultsConflict) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsConflict  %+v", 409, o.Payload)
}

func (o *PostPaymentdefaultsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentdefaultsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentdefaultsTooManyRequests creates a PostPaymentdefaultsTooManyRequests with default headers values
func NewPostPaymentdefaultsTooManyRequests() *PostPaymentdefaultsTooManyRequests {
	return &PostPaymentdefaultsTooManyRequests{}
}

/*PostPaymentdefaultsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostPaymentdefaultsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostPaymentdefaultsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPaymentdefaultsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentdefaultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentdefaultsInternalServerError creates a PostPaymentdefaultsInternalServerError with default headers values
func NewPostPaymentdefaultsInternalServerError() *PostPaymentdefaultsInternalServerError {
	return &PostPaymentdefaultsInternalServerError{}
}

/*PostPaymentdefaultsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPaymentdefaultsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostPaymentdefaultsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPaymentdefaultsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentdefaultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentdefaultsServiceUnavailable creates a PostPaymentdefaultsServiceUnavailable with default headers values
func NewPostPaymentdefaultsServiceUnavailable() *PostPaymentdefaultsServiceUnavailable {
	return &PostPaymentdefaultsServiceUnavailable{}
}

/*PostPaymentdefaultsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostPaymentdefaultsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostPaymentdefaultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /paymentdefaults][%d] postPaymentdefaultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPaymentdefaultsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentdefaultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
