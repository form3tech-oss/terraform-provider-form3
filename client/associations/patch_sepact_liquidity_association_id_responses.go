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

// PatchSepactLiquidityAssociationIDReader is a Reader for the PatchSepactLiquidityAssociationID structure.
type PatchSepactLiquidityAssociationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSepactLiquidityAssociationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchSepactLiquidityAssociationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchSepactLiquidityAssociationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchSepactLiquidityAssociationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchSepactLiquidityAssociationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchSepactLiquidityAssociationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchSepactLiquidityAssociationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchSepactLiquidityAssociationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchSepactLiquidityAssociationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchSepactLiquidityAssociationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchSepactLiquidityAssociationIDOK creates a PatchSepactLiquidityAssociationIDOK with default headers values
func NewPatchSepactLiquidityAssociationIDOK() *PatchSepactLiquidityAssociationIDOK {
	return &PatchSepactLiquidityAssociationIDOK{}
}

/*PatchSepactLiquidityAssociationIDOK handles this case with default header values.

Associations details
*/
type PatchSepactLiquidityAssociationIDOK struct {
	Payload *models.SepactLiquidityAssociationDetailsResponse
}

func (o *PatchSepactLiquidityAssociationIDOK) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdOK  %+v", 200, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDOK) GetPayload() *models.SepactLiquidityAssociationDetailsResponse {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepactLiquidityAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepactLiquidityAssociationIDBadRequest creates a PatchSepactLiquidityAssociationIDBadRequest with default headers values
func NewPatchSepactLiquidityAssociationIDBadRequest() *PatchSepactLiquidityAssociationIDBadRequest {
	return &PatchSepactLiquidityAssociationIDBadRequest{}
}

/*PatchSepactLiquidityAssociationIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchSepactLiquidityAssociationIDBadRequest struct {
	Payload *models.APIError
}

func (o *PatchSepactLiquidityAssociationIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepactLiquidityAssociationIDUnauthorized creates a PatchSepactLiquidityAssociationIDUnauthorized with default headers values
func NewPatchSepactLiquidityAssociationIDUnauthorized() *PatchSepactLiquidityAssociationIDUnauthorized {
	return &PatchSepactLiquidityAssociationIDUnauthorized{}
}

/*PatchSepactLiquidityAssociationIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PatchSepactLiquidityAssociationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *PatchSepactLiquidityAssociationIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepactLiquidityAssociationIDForbidden creates a PatchSepactLiquidityAssociationIDForbidden with default headers values
func NewPatchSepactLiquidityAssociationIDForbidden() *PatchSepactLiquidityAssociationIDForbidden {
	return &PatchSepactLiquidityAssociationIDForbidden{}
}

/*PatchSepactLiquidityAssociationIDForbidden handles this case with default header values.

Forbidden
*/
type PatchSepactLiquidityAssociationIDForbidden struct {
	Payload *models.APIError
}

func (o *PatchSepactLiquidityAssociationIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepactLiquidityAssociationIDNotFound creates a PatchSepactLiquidityAssociationIDNotFound with default headers values
func NewPatchSepactLiquidityAssociationIDNotFound() *PatchSepactLiquidityAssociationIDNotFound {
	return &PatchSepactLiquidityAssociationIDNotFound{}
}

/*PatchSepactLiquidityAssociationIDNotFound handles this case with default header values.

Record not found
*/
type PatchSepactLiquidityAssociationIDNotFound struct {
	Payload *models.APIError
}

func (o *PatchSepactLiquidityAssociationIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepactLiquidityAssociationIDConflict creates a PatchSepactLiquidityAssociationIDConflict with default headers values
func NewPatchSepactLiquidityAssociationIDConflict() *PatchSepactLiquidityAssociationIDConflict {
	return &PatchSepactLiquidityAssociationIDConflict{}
}

/*PatchSepactLiquidityAssociationIDConflict handles this case with default header values.

Conflict
*/
type PatchSepactLiquidityAssociationIDConflict struct {
	Payload *models.APIError
}

func (o *PatchSepactLiquidityAssociationIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdConflict  %+v", 409, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepactLiquidityAssociationIDTooManyRequests creates a PatchSepactLiquidityAssociationIDTooManyRequests with default headers values
func NewPatchSepactLiquidityAssociationIDTooManyRequests() *PatchSepactLiquidityAssociationIDTooManyRequests {
	return &PatchSepactLiquidityAssociationIDTooManyRequests{}
}

/*PatchSepactLiquidityAssociationIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PatchSepactLiquidityAssociationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *PatchSepactLiquidityAssociationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepactLiquidityAssociationIDInternalServerError creates a PatchSepactLiquidityAssociationIDInternalServerError with default headers values
func NewPatchSepactLiquidityAssociationIDInternalServerError() *PatchSepactLiquidityAssociationIDInternalServerError {
	return &PatchSepactLiquidityAssociationIDInternalServerError{}
}

/*PatchSepactLiquidityAssociationIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchSepactLiquidityAssociationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *PatchSepactLiquidityAssociationIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepactLiquidityAssociationIDServiceUnavailable creates a PatchSepactLiquidityAssociationIDServiceUnavailable with default headers values
func NewPatchSepactLiquidityAssociationIDServiceUnavailable() *PatchSepactLiquidityAssociationIDServiceUnavailable {
	return &PatchSepactLiquidityAssociationIDServiceUnavailable{}
}

/*PatchSepactLiquidityAssociationIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PatchSepactLiquidityAssociationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PatchSepactLiquidityAssociationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /sepact-liquidity/{associationId}][%d] patchSepactLiquidityAssociationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchSepactLiquidityAssociationIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepactLiquidityAssociationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
