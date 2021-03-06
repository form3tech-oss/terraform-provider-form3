// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostSubscriptionsReader is a Reader for the PostSubscriptions structure.
type PostSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSubscriptionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSubscriptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSubscriptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSubscriptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSubscriptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSubscriptionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostSubscriptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSubscriptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostSubscriptionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSubscriptionsCreated creates a PostSubscriptionsCreated with default headers values
func NewPostSubscriptionsCreated() *PostSubscriptionsCreated {
	return &PostSubscriptionsCreated{}
}

/*PostSubscriptionsCreated handles this case with default header values.

Subscription creation response
*/
type PostSubscriptionsCreated struct {
	Payload *models.SubscriptionCreationResponse
}

func (o *PostSubscriptionsCreated) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsCreated  %+v", 201, o.Payload)
}

func (o *PostSubscriptionsCreated) GetPayload() *models.SubscriptionCreationResponse {
	return o.Payload
}

func (o *PostSubscriptionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubscriptionsBadRequest creates a PostSubscriptionsBadRequest with default headers values
func NewPostSubscriptionsBadRequest() *PostSubscriptionsBadRequest {
	return &PostSubscriptionsBadRequest{}
}

/*PostSubscriptionsBadRequest handles this case with default header values.

Bad Request
*/
type PostSubscriptionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostSubscriptionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostSubscriptionsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSubscriptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubscriptionsUnauthorized creates a PostSubscriptionsUnauthorized with default headers values
func NewPostSubscriptionsUnauthorized() *PostSubscriptionsUnauthorized {
	return &PostSubscriptionsUnauthorized{}
}

/*PostSubscriptionsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostSubscriptionsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostSubscriptionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSubscriptionsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSubscriptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubscriptionsForbidden creates a PostSubscriptionsForbidden with default headers values
func NewPostSubscriptionsForbidden() *PostSubscriptionsForbidden {
	return &PostSubscriptionsForbidden{}
}

/*PostSubscriptionsForbidden handles this case with default header values.

Forbidden
*/
type PostSubscriptionsForbidden struct {
	Payload *models.APIError
}

func (o *PostSubscriptionsForbidden) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsForbidden  %+v", 403, o.Payload)
}

func (o *PostSubscriptionsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSubscriptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubscriptionsNotFound creates a PostSubscriptionsNotFound with default headers values
func NewPostSubscriptionsNotFound() *PostSubscriptionsNotFound {
	return &PostSubscriptionsNotFound{}
}

/*PostSubscriptionsNotFound handles this case with default header values.

Record not found
*/
type PostSubscriptionsNotFound struct {
	Payload *models.APIError
}

func (o *PostSubscriptionsNotFound) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsNotFound  %+v", 404, o.Payload)
}

func (o *PostSubscriptionsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSubscriptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubscriptionsConflict creates a PostSubscriptionsConflict with default headers values
func NewPostSubscriptionsConflict() *PostSubscriptionsConflict {
	return &PostSubscriptionsConflict{}
}

/*PostSubscriptionsConflict handles this case with default header values.

Conflict
*/
type PostSubscriptionsConflict struct {
	Payload *models.APIError
}

func (o *PostSubscriptionsConflict) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsConflict  %+v", 409, o.Payload)
}

func (o *PostSubscriptionsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSubscriptionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubscriptionsTooManyRequests creates a PostSubscriptionsTooManyRequests with default headers values
func NewPostSubscriptionsTooManyRequests() *PostSubscriptionsTooManyRequests {
	return &PostSubscriptionsTooManyRequests{}
}

/*PostSubscriptionsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostSubscriptionsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostSubscriptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSubscriptionsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSubscriptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubscriptionsInternalServerError creates a PostSubscriptionsInternalServerError with default headers values
func NewPostSubscriptionsInternalServerError() *PostSubscriptionsInternalServerError {
	return &PostSubscriptionsInternalServerError{}
}

/*PostSubscriptionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostSubscriptionsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostSubscriptionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSubscriptionsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSubscriptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubscriptionsServiceUnavailable creates a PostSubscriptionsServiceUnavailable with default headers values
func NewPostSubscriptionsServiceUnavailable() *PostSubscriptionsServiceUnavailable {
	return &PostSubscriptionsServiceUnavailable{}
}

/*PostSubscriptionsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostSubscriptionsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostSubscriptionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] postSubscriptionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostSubscriptionsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostSubscriptionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
