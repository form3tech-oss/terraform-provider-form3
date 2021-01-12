// Code generated by go-swagger; DO NOT EDIT.

package limits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// PatchLimitsIDReader is a Reader for the PatchLimitsID structure.
type PatchLimitsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLimitsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLimitsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLimitsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchLimitsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLimitsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLimitsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchLimitsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchLimitsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchLimitsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchLimitsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLimitsIDOK creates a PatchLimitsIDOK with default headers values
func NewPatchLimitsIDOK() *PatchLimitsIDOK {
	return &PatchLimitsIDOK{}
}

/*PatchLimitsIDOK handles this case with default header values.

Limit updated
*/
type PatchLimitsIDOK struct {
	Payload *models.LimitDetailsResponse
}

func (o *PatchLimitsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdOK  %+v", 200, o.Payload)
}

func (o *PatchLimitsIDOK) GetPayload() *models.LimitDetailsResponse {
	return o.Payload
}

func (o *PatchLimitsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLimitsIDBadRequest creates a PatchLimitsIDBadRequest with default headers values
func NewPatchLimitsIDBadRequest() *PatchLimitsIDBadRequest {
	return &PatchLimitsIDBadRequest{}
}

/*PatchLimitsIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchLimitsIDBadRequest struct {
	Payload *models.APIError
}

func (o *PatchLimitsIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLimitsIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLimitsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLimitsIDUnauthorized creates a PatchLimitsIDUnauthorized with default headers values
func NewPatchLimitsIDUnauthorized() *PatchLimitsIDUnauthorized {
	return &PatchLimitsIDUnauthorized{}
}

/*PatchLimitsIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PatchLimitsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *PatchLimitsIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchLimitsIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLimitsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLimitsIDForbidden creates a PatchLimitsIDForbidden with default headers values
func NewPatchLimitsIDForbidden() *PatchLimitsIDForbidden {
	return &PatchLimitsIDForbidden{}
}

/*PatchLimitsIDForbidden handles this case with default header values.

Forbidden
*/
type PatchLimitsIDForbidden struct {
	Payload *models.APIError
}

func (o *PatchLimitsIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLimitsIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLimitsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLimitsIDNotFound creates a PatchLimitsIDNotFound with default headers values
func NewPatchLimitsIDNotFound() *PatchLimitsIDNotFound {
	return &PatchLimitsIDNotFound{}
}

/*PatchLimitsIDNotFound handles this case with default header values.

Record not found
*/
type PatchLimitsIDNotFound struct {
	Payload *models.APIError
}

func (o *PatchLimitsIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLimitsIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLimitsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLimitsIDConflict creates a PatchLimitsIDConflict with default headers values
func NewPatchLimitsIDConflict() *PatchLimitsIDConflict {
	return &PatchLimitsIDConflict{}
}

/*PatchLimitsIDConflict handles this case with default header values.

Conflict
*/
type PatchLimitsIDConflict struct {
	Payload *models.APIError
}

func (o *PatchLimitsIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdConflict  %+v", 409, o.Payload)
}

func (o *PatchLimitsIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLimitsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLimitsIDTooManyRequests creates a PatchLimitsIDTooManyRequests with default headers values
func NewPatchLimitsIDTooManyRequests() *PatchLimitsIDTooManyRequests {
	return &PatchLimitsIDTooManyRequests{}
}

/*PatchLimitsIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PatchLimitsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *PatchLimitsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchLimitsIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLimitsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLimitsIDInternalServerError creates a PatchLimitsIDInternalServerError with default headers values
func NewPatchLimitsIDInternalServerError() *PatchLimitsIDInternalServerError {
	return &PatchLimitsIDInternalServerError{}
}

/*PatchLimitsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchLimitsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *PatchLimitsIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchLimitsIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLimitsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLimitsIDServiceUnavailable creates a PatchLimitsIDServiceUnavailable with default headers values
func NewPatchLimitsIDServiceUnavailable() *PatchLimitsIDServiceUnavailable {
	return &PatchLimitsIDServiceUnavailable{}
}

/*PatchLimitsIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PatchLimitsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PatchLimitsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchLimitsIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchLimitsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
