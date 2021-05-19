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

// GetReconciliationReader is a Reader for the GetReconciliation structure.
type GetReconciliationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReconciliationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReconciliationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReconciliationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReconciliationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReconciliationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReconciliationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetReconciliationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetReconciliationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReconciliationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetReconciliationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReconciliationOK creates a GetReconciliationOK with default headers values
func NewGetReconciliationOK() *GetReconciliationOK {
	return &GetReconciliationOK{}
}

/*GetReconciliationOK handles this case with default header values.

List of associations
*/
type GetReconciliationOK struct {
	Payload *models.ReconciliationAssociationDetailsListResponse
}

func (o *GetReconciliationOK) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationOK  %+v", 200, o.Payload)
}

func (o *GetReconciliationOK) GetPayload() *models.ReconciliationAssociationDetailsListResponse {
	return o.Payload
}

func (o *GetReconciliationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReconciliationAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationBadRequest creates a GetReconciliationBadRequest with default headers values
func NewGetReconciliationBadRequest() *GetReconciliationBadRequest {
	return &GetReconciliationBadRequest{}
}

/*GetReconciliationBadRequest handles this case with default header values.

Bad Request
*/
type GetReconciliationBadRequest struct {
	Payload *models.APIError
}

func (o *GetReconciliationBadRequest) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationBadRequest  %+v", 400, o.Payload)
}

func (o *GetReconciliationBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationUnauthorized creates a GetReconciliationUnauthorized with default headers values
func NewGetReconciliationUnauthorized() *GetReconciliationUnauthorized {
	return &GetReconciliationUnauthorized{}
}

/*GetReconciliationUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetReconciliationUnauthorized struct {
	Payload *models.APIError
}

func (o *GetReconciliationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReconciliationUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationForbidden creates a GetReconciliationForbidden with default headers values
func NewGetReconciliationForbidden() *GetReconciliationForbidden {
	return &GetReconciliationForbidden{}
}

/*GetReconciliationForbidden handles this case with default header values.

Forbidden
*/
type GetReconciliationForbidden struct {
	Payload *models.APIError
}

func (o *GetReconciliationForbidden) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationForbidden  %+v", 403, o.Payload)
}

func (o *GetReconciliationForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationNotFound creates a GetReconciliationNotFound with default headers values
func NewGetReconciliationNotFound() *GetReconciliationNotFound {
	return &GetReconciliationNotFound{}
}

/*GetReconciliationNotFound handles this case with default header values.

Record not found
*/
type GetReconciliationNotFound struct {
	Payload *models.APIError
}

func (o *GetReconciliationNotFound) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationNotFound  %+v", 404, o.Payload)
}

func (o *GetReconciliationNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationConflict creates a GetReconciliationConflict with default headers values
func NewGetReconciliationConflict() *GetReconciliationConflict {
	return &GetReconciliationConflict{}
}

/*GetReconciliationConflict handles this case with default header values.

Conflict
*/
type GetReconciliationConflict struct {
	Payload *models.APIError
}

func (o *GetReconciliationConflict) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationConflict  %+v", 409, o.Payload)
}

func (o *GetReconciliationConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationTooManyRequests creates a GetReconciliationTooManyRequests with default headers values
func NewGetReconciliationTooManyRequests() *GetReconciliationTooManyRequests {
	return &GetReconciliationTooManyRequests{}
}

/*GetReconciliationTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetReconciliationTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetReconciliationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetReconciliationTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationInternalServerError creates a GetReconciliationInternalServerError with default headers values
func NewGetReconciliationInternalServerError() *GetReconciliationInternalServerError {
	return &GetReconciliationInternalServerError{}
}

/*GetReconciliationInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetReconciliationInternalServerError struct {
	Payload *models.APIError
}

func (o *GetReconciliationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReconciliationInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationServiceUnavailable creates a GetReconciliationServiceUnavailable with default headers values
func NewGetReconciliationServiceUnavailable() *GetReconciliationServiceUnavailable {
	return &GetReconciliationServiceUnavailable{}
}

/*GetReconciliationServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetReconciliationServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetReconciliationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /reconciliation][%d] getReconciliationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetReconciliationServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
