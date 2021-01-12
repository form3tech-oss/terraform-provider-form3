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

// GetSepaLiquidityReader is a Reader for the GetSepaLiquidity structure.
type GetSepaLiquidityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSepaLiquidityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSepaLiquidityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSepaLiquidityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSepaLiquidityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSepaLiquidityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSepaLiquidityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetSepaLiquidityConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSepaLiquidityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSepaLiquidityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSepaLiquidityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSepaLiquidityOK creates a GetSepaLiquidityOK with default headers values
func NewGetSepaLiquidityOK() *GetSepaLiquidityOK {
	return &GetSepaLiquidityOK{}
}

/*GetSepaLiquidityOK handles this case with default header values.

List of associations
*/
type GetSepaLiquidityOK struct {
	Payload *models.SepaLiquidityAssociationDetailsListResponse
}

func (o *GetSepaLiquidityOK) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityOK  %+v", 200, o.Payload)
}

func (o *GetSepaLiquidityOK) GetPayload() *models.SepaLiquidityAssociationDetailsListResponse {
	return o.Payload
}

func (o *GetSepaLiquidityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaLiquidityAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaLiquidityBadRequest creates a GetSepaLiquidityBadRequest with default headers values
func NewGetSepaLiquidityBadRequest() *GetSepaLiquidityBadRequest {
	return &GetSepaLiquidityBadRequest{}
}

/*GetSepaLiquidityBadRequest handles this case with default header values.

Bad Request
*/
type GetSepaLiquidityBadRequest struct {
	Payload *models.APIError
}

func (o *GetSepaLiquidityBadRequest) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityBadRequest  %+v", 400, o.Payload)
}

func (o *GetSepaLiquidityBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaLiquidityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaLiquidityUnauthorized creates a GetSepaLiquidityUnauthorized with default headers values
func NewGetSepaLiquidityUnauthorized() *GetSepaLiquidityUnauthorized {
	return &GetSepaLiquidityUnauthorized{}
}

/*GetSepaLiquidityUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetSepaLiquidityUnauthorized struct {
	Payload *models.APIError
}

func (o *GetSepaLiquidityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSepaLiquidityUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaLiquidityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaLiquidityForbidden creates a GetSepaLiquidityForbidden with default headers values
func NewGetSepaLiquidityForbidden() *GetSepaLiquidityForbidden {
	return &GetSepaLiquidityForbidden{}
}

/*GetSepaLiquidityForbidden handles this case with default header values.

Forbidden
*/
type GetSepaLiquidityForbidden struct {
	Payload *models.APIError
}

func (o *GetSepaLiquidityForbidden) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityForbidden  %+v", 403, o.Payload)
}

func (o *GetSepaLiquidityForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaLiquidityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaLiquidityNotFound creates a GetSepaLiquidityNotFound with default headers values
func NewGetSepaLiquidityNotFound() *GetSepaLiquidityNotFound {
	return &GetSepaLiquidityNotFound{}
}

/*GetSepaLiquidityNotFound handles this case with default header values.

Record not found
*/
type GetSepaLiquidityNotFound struct {
	Payload *models.APIError
}

func (o *GetSepaLiquidityNotFound) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityNotFound  %+v", 404, o.Payload)
}

func (o *GetSepaLiquidityNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaLiquidityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaLiquidityConflict creates a GetSepaLiquidityConflict with default headers values
func NewGetSepaLiquidityConflict() *GetSepaLiquidityConflict {
	return &GetSepaLiquidityConflict{}
}

/*GetSepaLiquidityConflict handles this case with default header values.

Conflict
*/
type GetSepaLiquidityConflict struct {
	Payload *models.APIError
}

func (o *GetSepaLiquidityConflict) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityConflict  %+v", 409, o.Payload)
}

func (o *GetSepaLiquidityConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaLiquidityConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaLiquidityTooManyRequests creates a GetSepaLiquidityTooManyRequests with default headers values
func NewGetSepaLiquidityTooManyRequests() *GetSepaLiquidityTooManyRequests {
	return &GetSepaLiquidityTooManyRequests{}
}

/*GetSepaLiquidityTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetSepaLiquidityTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetSepaLiquidityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSepaLiquidityTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaLiquidityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaLiquidityInternalServerError creates a GetSepaLiquidityInternalServerError with default headers values
func NewGetSepaLiquidityInternalServerError() *GetSepaLiquidityInternalServerError {
	return &GetSepaLiquidityInternalServerError{}
}

/*GetSepaLiquidityInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSepaLiquidityInternalServerError struct {
	Payload *models.APIError
}

func (o *GetSepaLiquidityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSepaLiquidityInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaLiquidityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaLiquidityServiceUnavailable creates a GetSepaLiquidityServiceUnavailable with default headers values
func NewGetSepaLiquidityServiceUnavailable() *GetSepaLiquidityServiceUnavailable {
	return &GetSepaLiquidityServiceUnavailable{}
}

/*GetSepaLiquidityServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetSepaLiquidityServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetSepaLiquidityServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /sepa-liquidity][%d] getSepaLiquidityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSepaLiquidityServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaLiquidityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
