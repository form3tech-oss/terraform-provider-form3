// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// DeleteSepaLiquidityIDReader is a Reader for the DeleteSepaLiquidityID structure.
type DeleteSepaLiquidityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSepaLiquidityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSepaLiquidityIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSepaLiquidityIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSepaLiquidityIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSepaLiquidityIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSepaLiquidityIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteSepaLiquidityIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteSepaLiquidityIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSepaLiquidityIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteSepaLiquidityIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSepaLiquidityIDNoContent creates a DeleteSepaLiquidityIDNoContent with default headers values
func NewDeleteSepaLiquidityIDNoContent() *DeleteSepaLiquidityIDNoContent {
	return &DeleteSepaLiquidityIDNoContent{}
}

/* DeleteSepaLiquidityIDNoContent describes a response with status code 204, with default header values.

Association deleted
*/
type DeleteSepaLiquidityIDNoContent struct {
}

func (o *DeleteSepaLiquidityIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdNoContent ", 204)
}

func (o *DeleteSepaLiquidityIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSepaLiquidityIDBadRequest creates a DeleteSepaLiquidityIDBadRequest with default headers values
func NewDeleteSepaLiquidityIDBadRequest() *DeleteSepaLiquidityIDBadRequest {
	return &DeleteSepaLiquidityIDBadRequest{}
}

/* DeleteSepaLiquidityIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteSepaLiquidityIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteSepaLiquidityIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteSepaLiquidityIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaLiquidityIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaLiquidityIDUnauthorized creates a DeleteSepaLiquidityIDUnauthorized with default headers values
func NewDeleteSepaLiquidityIDUnauthorized() *DeleteSepaLiquidityIDUnauthorized {
	return &DeleteSepaLiquidityIDUnauthorized{}
}

/* DeleteSepaLiquidityIDUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteSepaLiquidityIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteSepaLiquidityIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteSepaLiquidityIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaLiquidityIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaLiquidityIDForbidden creates a DeleteSepaLiquidityIDForbidden with default headers values
func NewDeleteSepaLiquidityIDForbidden() *DeleteSepaLiquidityIDForbidden {
	return &DeleteSepaLiquidityIDForbidden{}
}

/* DeleteSepaLiquidityIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteSepaLiquidityIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteSepaLiquidityIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdForbidden  %+v", 403, o.Payload)
}
func (o *DeleteSepaLiquidityIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaLiquidityIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaLiquidityIDNotFound creates a DeleteSepaLiquidityIDNotFound with default headers values
func NewDeleteSepaLiquidityIDNotFound() *DeleteSepaLiquidityIDNotFound {
	return &DeleteSepaLiquidityIDNotFound{}
}

/* DeleteSepaLiquidityIDNotFound describes a response with status code 404, with default header values.

Record not found
*/
type DeleteSepaLiquidityIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteSepaLiquidityIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteSepaLiquidityIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaLiquidityIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaLiquidityIDConflict creates a DeleteSepaLiquidityIDConflict with default headers values
func NewDeleteSepaLiquidityIDConflict() *DeleteSepaLiquidityIDConflict {
	return &DeleteSepaLiquidityIDConflict{}
}

/* DeleteSepaLiquidityIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteSepaLiquidityIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteSepaLiquidityIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdConflict  %+v", 409, o.Payload)
}
func (o *DeleteSepaLiquidityIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaLiquidityIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaLiquidityIDTooManyRequests creates a DeleteSepaLiquidityIDTooManyRequests with default headers values
func NewDeleteSepaLiquidityIDTooManyRequests() *DeleteSepaLiquidityIDTooManyRequests {
	return &DeleteSepaLiquidityIDTooManyRequests{}
}

/* DeleteSepaLiquidityIDTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteSepaLiquidityIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteSepaLiquidityIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteSepaLiquidityIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaLiquidityIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaLiquidityIDInternalServerError creates a DeleteSepaLiquidityIDInternalServerError with default headers values
func NewDeleteSepaLiquidityIDInternalServerError() *DeleteSepaLiquidityIDInternalServerError {
	return &DeleteSepaLiquidityIDInternalServerError{}
}

/* DeleteSepaLiquidityIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteSepaLiquidityIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteSepaLiquidityIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteSepaLiquidityIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaLiquidityIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaLiquidityIDServiceUnavailable creates a DeleteSepaLiquidityIDServiceUnavailable with default headers values
func NewDeleteSepaLiquidityIDServiceUnavailable() *DeleteSepaLiquidityIDServiceUnavailable {
	return &DeleteSepaLiquidityIDServiceUnavailable{}
}

/* DeleteSepaLiquidityIDServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteSepaLiquidityIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteSepaLiquidityIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /sepa-liquidity/{id}][%d] deleteSepaLiquidityIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DeleteSepaLiquidityIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaLiquidityIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
