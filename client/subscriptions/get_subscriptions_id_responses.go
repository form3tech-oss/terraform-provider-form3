// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetSubscriptionsIDReader is a Reader for the GetSubscriptionsID structure.
type GetSubscriptionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSubscriptionsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSubscriptionsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSubscriptionsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSubscriptionsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetSubscriptionsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetSubscriptionsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSubscriptionsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetSubscriptionsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSubscriptionsIDOK creates a GetSubscriptionsIDOK with default headers values
func NewGetSubscriptionsIDOK() *GetSubscriptionsIDOK {
	return &GetSubscriptionsIDOK{}
}

/*GetSubscriptionsIDOK handles this case with default header values.

Subscription details
*/
type GetSubscriptionsIDOK struct {
	Payload *models.SubscriptionDetailsResponse
}

func (o *GetSubscriptionsIDOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsIDBadRequest creates a GetSubscriptionsIDBadRequest with default headers values
func NewGetSubscriptionsIDBadRequest() *GetSubscriptionsIDBadRequest {
	return &GetSubscriptionsIDBadRequest{}
}

/*GetSubscriptionsIDBadRequest handles this case with default header values.

Bad Request
*/
type GetSubscriptionsIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetSubscriptionsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetSubscriptionsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsIDUnauthorized creates a GetSubscriptionsIDUnauthorized with default headers values
func NewGetSubscriptionsIDUnauthorized() *GetSubscriptionsIDUnauthorized {
	return &GetSubscriptionsIDUnauthorized{}
}

/*GetSubscriptionsIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetSubscriptionsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetSubscriptionsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSubscriptionsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsIDForbidden creates a GetSubscriptionsIDForbidden with default headers values
func NewGetSubscriptionsIDForbidden() *GetSubscriptionsIDForbidden {
	return &GetSubscriptionsIDForbidden{}
}

/*GetSubscriptionsIDForbidden handles this case with default header values.

Forbidden
*/
type GetSubscriptionsIDForbidden struct {
	Payload *models.APIError
}

func (o *GetSubscriptionsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetSubscriptionsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsIDNotFound creates a GetSubscriptionsIDNotFound with default headers values
func NewGetSubscriptionsIDNotFound() *GetSubscriptionsIDNotFound {
	return &GetSubscriptionsIDNotFound{}
}

/*GetSubscriptionsIDNotFound handles this case with default header values.

Record not found
*/
type GetSubscriptionsIDNotFound struct {
	Payload *models.APIError
}

func (o *GetSubscriptionsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSubscriptionsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsIDConflict creates a GetSubscriptionsIDConflict with default headers values
func NewGetSubscriptionsIDConflict() *GetSubscriptionsIDConflict {
	return &GetSubscriptionsIDConflict{}
}

/*GetSubscriptionsIDConflict handles this case with default header values.

Conflict
*/
type GetSubscriptionsIDConflict struct {
	Payload *models.APIError
}

func (o *GetSubscriptionsIDConflict) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdConflict  %+v", 409, o.Payload)
}

func (o *GetSubscriptionsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsIDTooManyRequests creates a GetSubscriptionsIDTooManyRequests with default headers values
func NewGetSubscriptionsIDTooManyRequests() *GetSubscriptionsIDTooManyRequests {
	return &GetSubscriptionsIDTooManyRequests{}
}

/*GetSubscriptionsIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetSubscriptionsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetSubscriptionsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSubscriptionsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsIDInternalServerError creates a GetSubscriptionsIDInternalServerError with default headers values
func NewGetSubscriptionsIDInternalServerError() *GetSubscriptionsIDInternalServerError {
	return &GetSubscriptionsIDInternalServerError{}
}

/*GetSubscriptionsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSubscriptionsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetSubscriptionsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubscriptionsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsIDServiceUnavailable creates a GetSubscriptionsIDServiceUnavailable with default headers values
func NewGetSubscriptionsIDServiceUnavailable() *GetSubscriptionsIDServiceUnavailable {
	return &GetSubscriptionsIDServiceUnavailable{}
}

/*GetSubscriptionsIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetSubscriptionsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetSubscriptionsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSubscriptionsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
