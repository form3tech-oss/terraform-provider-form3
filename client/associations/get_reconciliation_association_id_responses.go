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

// GetReconciliationAssociationIDReader is a Reader for the GetReconciliationAssociationID structure.
type GetReconciliationAssociationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReconciliationAssociationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReconciliationAssociationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReconciliationAssociationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReconciliationAssociationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReconciliationAssociationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReconciliationAssociationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetReconciliationAssociationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetReconciliationAssociationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReconciliationAssociationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetReconciliationAssociationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReconciliationAssociationIDOK creates a GetReconciliationAssociationIDOK with default headers values
func NewGetReconciliationAssociationIDOK() *GetReconciliationAssociationIDOK {
	return &GetReconciliationAssociationIDOK{}
}

/* GetReconciliationAssociationIDOK describes a response with status code 200, with default header values.

Associations details
*/
type GetReconciliationAssociationIDOK struct {
	Payload *models.ReconciliationAssociationDetailsResponse
}

func (o *GetReconciliationAssociationIDOK) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdOK  %+v", 200, o.Payload)
}
func (o *GetReconciliationAssociationIDOK) GetPayload() *models.ReconciliationAssociationDetailsResponse {
	return o.Payload
}

func (o *GetReconciliationAssociationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReconciliationAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationAssociationIDBadRequest creates a GetReconciliationAssociationIDBadRequest with default headers values
func NewGetReconciliationAssociationIDBadRequest() *GetReconciliationAssociationIDBadRequest {
	return &GetReconciliationAssociationIDBadRequest{}
}

/* GetReconciliationAssociationIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetReconciliationAssociationIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetReconciliationAssociationIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetReconciliationAssociationIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationAssociationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationAssociationIDUnauthorized creates a GetReconciliationAssociationIDUnauthorized with default headers values
func NewGetReconciliationAssociationIDUnauthorized() *GetReconciliationAssociationIDUnauthorized {
	return &GetReconciliationAssociationIDUnauthorized{}
}

/* GetReconciliationAssociationIDUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetReconciliationAssociationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetReconciliationAssociationIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetReconciliationAssociationIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationAssociationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationAssociationIDForbidden creates a GetReconciliationAssociationIDForbidden with default headers values
func NewGetReconciliationAssociationIDForbidden() *GetReconciliationAssociationIDForbidden {
	return &GetReconciliationAssociationIDForbidden{}
}

/* GetReconciliationAssociationIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReconciliationAssociationIDForbidden struct {
	Payload *models.APIError
}

func (o *GetReconciliationAssociationIDForbidden) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdForbidden  %+v", 403, o.Payload)
}
func (o *GetReconciliationAssociationIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationAssociationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationAssociationIDNotFound creates a GetReconciliationAssociationIDNotFound with default headers values
func NewGetReconciliationAssociationIDNotFound() *GetReconciliationAssociationIDNotFound {
	return &GetReconciliationAssociationIDNotFound{}
}

/* GetReconciliationAssociationIDNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetReconciliationAssociationIDNotFound struct {
	Payload *models.APIError
}

func (o *GetReconciliationAssociationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdNotFound  %+v", 404, o.Payload)
}
func (o *GetReconciliationAssociationIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationAssociationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationAssociationIDConflict creates a GetReconciliationAssociationIDConflict with default headers values
func NewGetReconciliationAssociationIDConflict() *GetReconciliationAssociationIDConflict {
	return &GetReconciliationAssociationIDConflict{}
}

/* GetReconciliationAssociationIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetReconciliationAssociationIDConflict struct {
	Payload *models.APIError
}

func (o *GetReconciliationAssociationIDConflict) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdConflict  %+v", 409, o.Payload)
}
func (o *GetReconciliationAssociationIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationAssociationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationAssociationIDTooManyRequests creates a GetReconciliationAssociationIDTooManyRequests with default headers values
func NewGetReconciliationAssociationIDTooManyRequests() *GetReconciliationAssociationIDTooManyRequests {
	return &GetReconciliationAssociationIDTooManyRequests{}
}

/* GetReconciliationAssociationIDTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetReconciliationAssociationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetReconciliationAssociationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetReconciliationAssociationIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationAssociationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationAssociationIDInternalServerError creates a GetReconciliationAssociationIDInternalServerError with default headers values
func NewGetReconciliationAssociationIDInternalServerError() *GetReconciliationAssociationIDInternalServerError {
	return &GetReconciliationAssociationIDInternalServerError{}
}

/* GetReconciliationAssociationIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetReconciliationAssociationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetReconciliationAssociationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReconciliationAssociationIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationAssociationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationAssociationIDServiceUnavailable creates a GetReconciliationAssociationIDServiceUnavailable with default headers values
func NewGetReconciliationAssociationIDServiceUnavailable() *GetReconciliationAssociationIDServiceUnavailable {
	return &GetReconciliationAssociationIDServiceUnavailable{}
}

/* GetReconciliationAssociationIDServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetReconciliationAssociationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetReconciliationAssociationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /reconciliation/{associationId}][%d] getReconciliationAssociationIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetReconciliationAssociationIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetReconciliationAssociationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
