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

// PatchSepainstantIDReader is a Reader for the PatchSepainstantID structure.
type PatchSepainstantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSepainstantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchSepainstantIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchSepainstantIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchSepainstantIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchSepainstantIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchSepainstantIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchSepainstantIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchSepainstantIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchSepainstantIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchSepainstantIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSepainstantIDOK creates a PatchSepainstantIDOK with default headers values
func NewPatchSepainstantIDOK() *PatchSepainstantIDOK {
	return &PatchSepainstantIDOK{}
}

/*PatchSepainstantIDOK handles this case with default header values.

Associations details
*/
type PatchSepainstantIDOK struct {
	Payload *models.SepaInstantAssociationDetailsResponse
}

func (o *PatchSepainstantIDOK) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdOK  %+v", 200, o.Payload)
}

func (o *PatchSepainstantIDOK) GetPayload() *models.SepaInstantAssociationDetailsResponse {
	return o.Payload
}

func (o *PatchSepainstantIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaInstantAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepainstantIDBadRequest creates a PatchSepainstantIDBadRequest with default headers values
func NewPatchSepainstantIDBadRequest() *PatchSepainstantIDBadRequest {
	return &PatchSepainstantIDBadRequest{}
}

/*PatchSepainstantIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchSepainstantIDBadRequest struct {
	Payload *models.APIError
}

func (o *PatchSepainstantIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSepainstantIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepainstantIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepainstantIDUnauthorized creates a PatchSepainstantIDUnauthorized with default headers values
func NewPatchSepainstantIDUnauthorized() *PatchSepainstantIDUnauthorized {
	return &PatchSepainstantIDUnauthorized{}
}

/*PatchSepainstantIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PatchSepainstantIDUnauthorized struct {
	Payload *models.APIError
}

func (o *PatchSepainstantIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchSepainstantIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepainstantIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepainstantIDForbidden creates a PatchSepainstantIDForbidden with default headers values
func NewPatchSepainstantIDForbidden() *PatchSepainstantIDForbidden {
	return &PatchSepainstantIDForbidden{}
}

/*PatchSepainstantIDForbidden handles this case with default header values.

Forbidden
*/
type PatchSepainstantIDForbidden struct {
	Payload *models.APIError
}

func (o *PatchSepainstantIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchSepainstantIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepainstantIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepainstantIDNotFound creates a PatchSepainstantIDNotFound with default headers values
func NewPatchSepainstantIDNotFound() *PatchSepainstantIDNotFound {
	return &PatchSepainstantIDNotFound{}
}

/*PatchSepainstantIDNotFound handles this case with default header values.

Record not found
*/
type PatchSepainstantIDNotFound struct {
	Payload *models.APIError
}

func (o *PatchSepainstantIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchSepainstantIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepainstantIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepainstantIDConflict creates a PatchSepainstantIDConflict with default headers values
func NewPatchSepainstantIDConflict() *PatchSepainstantIDConflict {
	return &PatchSepainstantIDConflict{}
}

/*PatchSepainstantIDConflict handles this case with default header values.

Conflict
*/
type PatchSepainstantIDConflict struct {
	Payload *models.APIError
}

func (o *PatchSepainstantIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdConflict  %+v", 409, o.Payload)
}

func (o *PatchSepainstantIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepainstantIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepainstantIDTooManyRequests creates a PatchSepainstantIDTooManyRequests with default headers values
func NewPatchSepainstantIDTooManyRequests() *PatchSepainstantIDTooManyRequests {
	return &PatchSepainstantIDTooManyRequests{}
}

/*PatchSepainstantIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PatchSepainstantIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *PatchSepainstantIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchSepainstantIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepainstantIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepainstantIDInternalServerError creates a PatchSepainstantIDInternalServerError with default headers values
func NewPatchSepainstantIDInternalServerError() *PatchSepainstantIDInternalServerError {
	return &PatchSepainstantIDInternalServerError{}
}

/*PatchSepainstantIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchSepainstantIDInternalServerError struct {
	Payload *models.APIError
}

func (o *PatchSepainstantIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchSepainstantIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepainstantIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSepainstantIDServiceUnavailable creates a PatchSepainstantIDServiceUnavailable with default headers values
func NewPatchSepainstantIDServiceUnavailable() *PatchSepainstantIDServiceUnavailable {
	return &PatchSepainstantIDServiceUnavailable{}
}

/*PatchSepainstantIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PatchSepainstantIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PatchSepainstantIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /sepainstant/{id}][%d] patchSepainstantIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchSepainstantIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchSepainstantIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
