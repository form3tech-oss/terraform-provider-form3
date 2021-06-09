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

// GetPaymentdefaultsIDReader is a Reader for the GetPaymentdefaultsID structure.
type GetPaymentdefaultsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentdefaultsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPaymentdefaultsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPaymentdefaultsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPaymentdefaultsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPaymentdefaultsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPaymentdefaultsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPaymentdefaultsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPaymentdefaultsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPaymentdefaultsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPaymentdefaultsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentdefaultsIDOK creates a GetPaymentdefaultsIDOK with default headers values
func NewGetPaymentdefaultsIDOK() *GetPaymentdefaultsIDOK {
	return &GetPaymentdefaultsIDOK{}
}

/*GetPaymentdefaultsIDOK handles this case with default header values.

Payment default details
*/
type GetPaymentdefaultsIDOK struct {
	Payload *models.PaymentDefaultsResponse
}

func (o *GetPaymentdefaultsIDOK) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentdefaultsIDOK) GetPayload() *models.PaymentDefaultsResponse {
	return o.Payload
}

func (o *GetPaymentdefaultsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentDefaultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentdefaultsIDBadRequest creates a GetPaymentdefaultsIDBadRequest with default headers values
func NewGetPaymentdefaultsIDBadRequest() *GetPaymentdefaultsIDBadRequest {
	return &GetPaymentdefaultsIDBadRequest{}
}

/*GetPaymentdefaultsIDBadRequest handles this case with default header values.

Bad Request
*/
type GetPaymentdefaultsIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetPaymentdefaultsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetPaymentdefaultsIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentdefaultsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentdefaultsIDUnauthorized creates a GetPaymentdefaultsIDUnauthorized with default headers values
func NewGetPaymentdefaultsIDUnauthorized() *GetPaymentdefaultsIDUnauthorized {
	return &GetPaymentdefaultsIDUnauthorized{}
}

/*GetPaymentdefaultsIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetPaymentdefaultsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetPaymentdefaultsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPaymentdefaultsIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentdefaultsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentdefaultsIDForbidden creates a GetPaymentdefaultsIDForbidden with default headers values
func NewGetPaymentdefaultsIDForbidden() *GetPaymentdefaultsIDForbidden {
	return &GetPaymentdefaultsIDForbidden{}
}

/*GetPaymentdefaultsIDForbidden handles this case with default header values.

Forbidden
*/
type GetPaymentdefaultsIDForbidden struct {
	Payload *models.APIError
}

func (o *GetPaymentdefaultsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetPaymentdefaultsIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentdefaultsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentdefaultsIDNotFound creates a GetPaymentdefaultsIDNotFound with default headers values
func NewGetPaymentdefaultsIDNotFound() *GetPaymentdefaultsIDNotFound {
	return &GetPaymentdefaultsIDNotFound{}
}

/*GetPaymentdefaultsIDNotFound handles this case with default header values.

Record not found
*/
type GetPaymentdefaultsIDNotFound struct {
	Payload *models.APIError
}

func (o *GetPaymentdefaultsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetPaymentdefaultsIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentdefaultsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentdefaultsIDConflict creates a GetPaymentdefaultsIDConflict with default headers values
func NewGetPaymentdefaultsIDConflict() *GetPaymentdefaultsIDConflict {
	return &GetPaymentdefaultsIDConflict{}
}

/*GetPaymentdefaultsIDConflict handles this case with default header values.

Conflict
*/
type GetPaymentdefaultsIDConflict struct {
	Payload *models.APIError
}

func (o *GetPaymentdefaultsIDConflict) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdConflict  %+v", 409, o.Payload)
}

func (o *GetPaymentdefaultsIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentdefaultsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentdefaultsIDTooManyRequests creates a GetPaymentdefaultsIDTooManyRequests with default headers values
func NewGetPaymentdefaultsIDTooManyRequests() *GetPaymentdefaultsIDTooManyRequests {
	return &GetPaymentdefaultsIDTooManyRequests{}
}

/*GetPaymentdefaultsIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetPaymentdefaultsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetPaymentdefaultsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPaymentdefaultsIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentdefaultsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentdefaultsIDInternalServerError creates a GetPaymentdefaultsIDInternalServerError with default headers values
func NewGetPaymentdefaultsIDInternalServerError() *GetPaymentdefaultsIDInternalServerError {
	return &GetPaymentdefaultsIDInternalServerError{}
}

/*GetPaymentdefaultsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPaymentdefaultsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetPaymentdefaultsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPaymentdefaultsIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentdefaultsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentdefaultsIDServiceUnavailable creates a GetPaymentdefaultsIDServiceUnavailable with default headers values
func NewGetPaymentdefaultsIDServiceUnavailable() *GetPaymentdefaultsIDServiceUnavailable {
	return &GetPaymentdefaultsIDServiceUnavailable{}
}

/*GetPaymentdefaultsIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetPaymentdefaultsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetPaymentdefaultsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPaymentdefaultsIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentdefaultsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
