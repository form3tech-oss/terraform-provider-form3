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

// GetLhvAssociationIDAgencySynchronisationsReader is a Reader for the GetLhvAssociationIDAgencySynchronisations structure.
type GetLhvAssociationIDAgencySynchronisationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLhvAssociationIDAgencySynchronisationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLhvAssociationIDAgencySynchronisationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLhvAssociationIDAgencySynchronisationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLhvAssociationIDAgencySynchronisationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLhvAssociationIDAgencySynchronisationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLhvAssociationIDAgencySynchronisationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetLhvAssociationIDAgencySynchronisationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLhvAssociationIDAgencySynchronisationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLhvAssociationIDAgencySynchronisationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLhvAssociationIDAgencySynchronisationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLhvAssociationIDAgencySynchronisationsOK creates a GetLhvAssociationIDAgencySynchronisationsOK with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsOK() *GetLhvAssociationIDAgencySynchronisationsOK {
	return &GetLhvAssociationIDAgencySynchronisationsOK{}
}

/* GetLhvAssociationIDAgencySynchronisationsOK describes a response with status code 200, with default header values.

List of LHV agency synchronisation details
*/
type GetLhvAssociationIDAgencySynchronisationsOK struct {
	Payload *models.LhvAgencySynchronisationListResponse
}

func (o *GetLhvAssociationIDAgencySynchronisationsOK) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsOK  %+v", 200, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsOK) GetPayload() *models.LhvAgencySynchronisationListResponse {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LhvAgencySynchronisationListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDAgencySynchronisationsBadRequest creates a GetLhvAssociationIDAgencySynchronisationsBadRequest with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsBadRequest() *GetLhvAssociationIDAgencySynchronisationsBadRequest {
	return &GetLhvAssociationIDAgencySynchronisationsBadRequest{}
}

/* GetLhvAssociationIDAgencySynchronisationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLhvAssociationIDAgencySynchronisationsBadRequest struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDAgencySynchronisationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsBadRequest  %+v", 400, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDAgencySynchronisationsUnauthorized creates a GetLhvAssociationIDAgencySynchronisationsUnauthorized with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsUnauthorized() *GetLhvAssociationIDAgencySynchronisationsUnauthorized {
	return &GetLhvAssociationIDAgencySynchronisationsUnauthorized{}
}

/* GetLhvAssociationIDAgencySynchronisationsUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetLhvAssociationIDAgencySynchronisationsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDAgencySynchronisationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDAgencySynchronisationsForbidden creates a GetLhvAssociationIDAgencySynchronisationsForbidden with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsForbidden() *GetLhvAssociationIDAgencySynchronisationsForbidden {
	return &GetLhvAssociationIDAgencySynchronisationsForbidden{}
}

/* GetLhvAssociationIDAgencySynchronisationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLhvAssociationIDAgencySynchronisationsForbidden struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDAgencySynchronisationsForbidden) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsForbidden  %+v", 403, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDAgencySynchronisationsNotFound creates a GetLhvAssociationIDAgencySynchronisationsNotFound with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsNotFound() *GetLhvAssociationIDAgencySynchronisationsNotFound {
	return &GetLhvAssociationIDAgencySynchronisationsNotFound{}
}

/* GetLhvAssociationIDAgencySynchronisationsNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetLhvAssociationIDAgencySynchronisationsNotFound struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDAgencySynchronisationsNotFound) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsNotFound  %+v", 404, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDAgencySynchronisationsConflict creates a GetLhvAssociationIDAgencySynchronisationsConflict with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsConflict() *GetLhvAssociationIDAgencySynchronisationsConflict {
	return &GetLhvAssociationIDAgencySynchronisationsConflict{}
}

/* GetLhvAssociationIDAgencySynchronisationsConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetLhvAssociationIDAgencySynchronisationsConflict struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDAgencySynchronisationsConflict) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsConflict  %+v", 409, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDAgencySynchronisationsTooManyRequests creates a GetLhvAssociationIDAgencySynchronisationsTooManyRequests with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsTooManyRequests() *GetLhvAssociationIDAgencySynchronisationsTooManyRequests {
	return &GetLhvAssociationIDAgencySynchronisationsTooManyRequests{}
}

/* GetLhvAssociationIDAgencySynchronisationsTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetLhvAssociationIDAgencySynchronisationsTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDAgencySynchronisationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDAgencySynchronisationsInternalServerError creates a GetLhvAssociationIDAgencySynchronisationsInternalServerError with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsInternalServerError() *GetLhvAssociationIDAgencySynchronisationsInternalServerError {
	return &GetLhvAssociationIDAgencySynchronisationsInternalServerError{}
}

/* GetLhvAssociationIDAgencySynchronisationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetLhvAssociationIDAgencySynchronisationsInternalServerError struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDAgencySynchronisationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDAgencySynchronisationsServiceUnavailable creates a GetLhvAssociationIDAgencySynchronisationsServiceUnavailable with default headers values
func NewGetLhvAssociationIDAgencySynchronisationsServiceUnavailable() *GetLhvAssociationIDAgencySynchronisationsServiceUnavailable {
	return &GetLhvAssociationIDAgencySynchronisationsServiceUnavailable{}
}

/* GetLhvAssociationIDAgencySynchronisationsServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetLhvAssociationIDAgencySynchronisationsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDAgencySynchronisationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}/agency_synchronisations][%d] getLhvAssociationIdAgencySynchronisationsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetLhvAssociationIDAgencySynchronisationsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetLhvAssociationIDAgencySynchronisationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
