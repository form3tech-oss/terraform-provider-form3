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

// PatchLhvAssociationIDReader is a Reader for the PatchLhvAssociationID structure.
type PatchLhvAssociationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLhvAssociationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLhvAssociationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLhvAssociationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchLhvAssociationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLhvAssociationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLhvAssociationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchLhvAssociationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchLhvAssociationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchLhvAssociationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchLhvAssociationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLhvAssociationIDOK creates a PatchLhvAssociationIDOK with default headers values
func NewPatchLhvAssociationIDOK() *PatchLhvAssociationIDOK {
	return &PatchLhvAssociationIDOK{}
}

/*PatchLhvAssociationIDOK handles this case with default header values.

Associations details
*/
type PatchLhvAssociationIDOK struct {
	Payload *models.LhvAssociationDetailsResponse
}

func (o *PatchLhvAssociationIDOK) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdOK  %+v", 200, o.Payload)
}

func (o *PatchLhvAssociationIDOK) GetPayload() *models.LhvAssociationDetailsResponse {
	return o.Payload
}

func (o *PatchLhvAssociationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LhvAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLhvAssociationIDBadRequest creates a PatchLhvAssociationIDBadRequest with default headers values
func NewPatchLhvAssociationIDBadRequest() *PatchLhvAssociationIDBadRequest {
	return &PatchLhvAssociationIDBadRequest{}
}

/*PatchLhvAssociationIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchLhvAssociationIDBadRequest struct {
	Payload *models.APIError
}

func (o *PatchLhvAssociationIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLhvAssociationIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLhvAssociationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLhvAssociationIDUnauthorized creates a PatchLhvAssociationIDUnauthorized with default headers values
func NewPatchLhvAssociationIDUnauthorized() *PatchLhvAssociationIDUnauthorized {
	return &PatchLhvAssociationIDUnauthorized{}
}

/*PatchLhvAssociationIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PatchLhvAssociationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *PatchLhvAssociationIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchLhvAssociationIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLhvAssociationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLhvAssociationIDForbidden creates a PatchLhvAssociationIDForbidden with default headers values
func NewPatchLhvAssociationIDForbidden() *PatchLhvAssociationIDForbidden {
	return &PatchLhvAssociationIDForbidden{}
}

/*PatchLhvAssociationIDForbidden handles this case with default header values.

Forbidden
*/
type PatchLhvAssociationIDForbidden struct {
	Payload *models.APIError
}

func (o *PatchLhvAssociationIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLhvAssociationIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLhvAssociationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLhvAssociationIDNotFound creates a PatchLhvAssociationIDNotFound with default headers values
func NewPatchLhvAssociationIDNotFound() *PatchLhvAssociationIDNotFound {
	return &PatchLhvAssociationIDNotFound{}
}

/*PatchLhvAssociationIDNotFound handles this case with default header values.

Record not found
*/
type PatchLhvAssociationIDNotFound struct {
	Payload *models.APIError
}

func (o *PatchLhvAssociationIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLhvAssociationIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLhvAssociationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLhvAssociationIDConflict creates a PatchLhvAssociationIDConflict with default headers values
func NewPatchLhvAssociationIDConflict() *PatchLhvAssociationIDConflict {
	return &PatchLhvAssociationIDConflict{}
}

/*PatchLhvAssociationIDConflict handles this case with default header values.

Conflict
*/
type PatchLhvAssociationIDConflict struct {
	Payload *models.APIError
}

func (o *PatchLhvAssociationIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdConflict  %+v", 409, o.Payload)
}

func (o *PatchLhvAssociationIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLhvAssociationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLhvAssociationIDTooManyRequests creates a PatchLhvAssociationIDTooManyRequests with default headers values
func NewPatchLhvAssociationIDTooManyRequests() *PatchLhvAssociationIDTooManyRequests {
	return &PatchLhvAssociationIDTooManyRequests{}
}

/*PatchLhvAssociationIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PatchLhvAssociationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *PatchLhvAssociationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchLhvAssociationIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLhvAssociationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLhvAssociationIDInternalServerError creates a PatchLhvAssociationIDInternalServerError with default headers values
func NewPatchLhvAssociationIDInternalServerError() *PatchLhvAssociationIDInternalServerError {
	return &PatchLhvAssociationIDInternalServerError{}
}

/*PatchLhvAssociationIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchLhvAssociationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *PatchLhvAssociationIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchLhvAssociationIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLhvAssociationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLhvAssociationIDServiceUnavailable creates a PatchLhvAssociationIDServiceUnavailable with default headers values
func NewPatchLhvAssociationIDServiceUnavailable() *PatchLhvAssociationIDServiceUnavailable {
	return &PatchLhvAssociationIDServiceUnavailable{}
}

/*PatchLhvAssociationIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PatchLhvAssociationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PatchLhvAssociationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /lhv/{associationId}][%d] patchLhvAssociationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchLhvAssociationIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLhvAssociationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}