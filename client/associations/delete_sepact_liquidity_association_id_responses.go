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

// DeleteSepactLiquidityAssociationIDReader is a Reader for the DeleteSepactLiquidityAssociationID structure.
type DeleteSepactLiquidityAssociationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSepactLiquidityAssociationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSepactLiquidityAssociationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSepactLiquidityAssociationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSepactLiquidityAssociationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSepactLiquidityAssociationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSepactLiquidityAssociationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteSepactLiquidityAssociationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteSepactLiquidityAssociationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSepactLiquidityAssociationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteSepactLiquidityAssociationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSepactLiquidityAssociationIDNoContent creates a DeleteSepactLiquidityAssociationIDNoContent with default headers values
func NewDeleteSepactLiquidityAssociationIDNoContent() *DeleteSepactLiquidityAssociationIDNoContent {
	return &DeleteSepactLiquidityAssociationIDNoContent{}
}

/*DeleteSepactLiquidityAssociationIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteSepactLiquidityAssociationIDNoContent struct {
}

func (o *DeleteSepactLiquidityAssociationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdNoContent ", 204)
}

func (o *DeleteSepactLiquidityAssociationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSepactLiquidityAssociationIDBadRequest creates a DeleteSepactLiquidityAssociationIDBadRequest with default headers values
func NewDeleteSepactLiquidityAssociationIDBadRequest() *DeleteSepactLiquidityAssociationIDBadRequest {
	return &DeleteSepactLiquidityAssociationIDBadRequest{}
}

/*DeleteSepactLiquidityAssociationIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteSepactLiquidityAssociationIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteSepactLiquidityAssociationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSepactLiquidityAssociationIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepactLiquidityAssociationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepactLiquidityAssociationIDUnauthorized creates a DeleteSepactLiquidityAssociationIDUnauthorized with default headers values
func NewDeleteSepactLiquidityAssociationIDUnauthorized() *DeleteSepactLiquidityAssociationIDUnauthorized {
	return &DeleteSepactLiquidityAssociationIDUnauthorized{}
}

/*DeleteSepactLiquidityAssociationIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteSepactLiquidityAssociationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteSepactLiquidityAssociationIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSepactLiquidityAssociationIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepactLiquidityAssociationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepactLiquidityAssociationIDForbidden creates a DeleteSepactLiquidityAssociationIDForbidden with default headers values
func NewDeleteSepactLiquidityAssociationIDForbidden() *DeleteSepactLiquidityAssociationIDForbidden {
	return &DeleteSepactLiquidityAssociationIDForbidden{}
}

/*DeleteSepactLiquidityAssociationIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteSepactLiquidityAssociationIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteSepactLiquidityAssociationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSepactLiquidityAssociationIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepactLiquidityAssociationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepactLiquidityAssociationIDNotFound creates a DeleteSepactLiquidityAssociationIDNotFound with default headers values
func NewDeleteSepactLiquidityAssociationIDNotFound() *DeleteSepactLiquidityAssociationIDNotFound {
	return &DeleteSepactLiquidityAssociationIDNotFound{}
}

/*DeleteSepactLiquidityAssociationIDNotFound handles this case with default header values.

Record not found
*/
type DeleteSepactLiquidityAssociationIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteSepactLiquidityAssociationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSepactLiquidityAssociationIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepactLiquidityAssociationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepactLiquidityAssociationIDConflict creates a DeleteSepactLiquidityAssociationIDConflict with default headers values
func NewDeleteSepactLiquidityAssociationIDConflict() *DeleteSepactLiquidityAssociationIDConflict {
	return &DeleteSepactLiquidityAssociationIDConflict{}
}

/*DeleteSepactLiquidityAssociationIDConflict handles this case with default header values.

Conflict
*/
type DeleteSepactLiquidityAssociationIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteSepactLiquidityAssociationIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteSepactLiquidityAssociationIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepactLiquidityAssociationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepactLiquidityAssociationIDTooManyRequests creates a DeleteSepactLiquidityAssociationIDTooManyRequests with default headers values
func NewDeleteSepactLiquidityAssociationIDTooManyRequests() *DeleteSepactLiquidityAssociationIDTooManyRequests {
	return &DeleteSepactLiquidityAssociationIDTooManyRequests{}
}

/*DeleteSepactLiquidityAssociationIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteSepactLiquidityAssociationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteSepactLiquidityAssociationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteSepactLiquidityAssociationIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepactLiquidityAssociationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepactLiquidityAssociationIDInternalServerError creates a DeleteSepactLiquidityAssociationIDInternalServerError with default headers values
func NewDeleteSepactLiquidityAssociationIDInternalServerError() *DeleteSepactLiquidityAssociationIDInternalServerError {
	return &DeleteSepactLiquidityAssociationIDInternalServerError{}
}

/*DeleteSepactLiquidityAssociationIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteSepactLiquidityAssociationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteSepactLiquidityAssociationIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSepactLiquidityAssociationIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepactLiquidityAssociationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepactLiquidityAssociationIDServiceUnavailable creates a DeleteSepactLiquidityAssociationIDServiceUnavailable with default headers values
func NewDeleteSepactLiquidityAssociationIDServiceUnavailable() *DeleteSepactLiquidityAssociationIDServiceUnavailable {
	return &DeleteSepactLiquidityAssociationIDServiceUnavailable{}
}

/*DeleteSepactLiquidityAssociationIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteSepactLiquidityAssociationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteSepactLiquidityAssociationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /sepact-liquidity/{associationId}][%d] deleteSepactLiquidityAssociationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteSepactLiquidityAssociationIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepactLiquidityAssociationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
