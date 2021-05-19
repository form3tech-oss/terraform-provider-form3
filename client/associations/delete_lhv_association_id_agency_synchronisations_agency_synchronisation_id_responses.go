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

// DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDReader is a Reader for the DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationID structure.
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent handles this case with default header values.

LHV Association agency synchronisation details deleted
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent struct {
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdNoContent ", 204)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound handles this case with default header values.

Record not found
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict handles this case with default header values.

Conflict
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable creates a DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable with default headers values
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable {
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable{}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}/agency_synchronisations/{agencySynchronisationId}][%d] deleteLhvAssociationIdAgencySynchronisationsAgencySynchronisationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
