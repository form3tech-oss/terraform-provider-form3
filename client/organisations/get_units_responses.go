// Code generated by go-swagger; DO NOT EDIT.

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetUnitsReader is a Reader for the GetUnits structure.
type GetUnitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUnitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUnitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUnitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUnitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUnitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUnitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetUnitsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetUnitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUnitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetUnitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUnitsOK creates a GetUnitsOK with default headers values
func NewGetUnitsOK() *GetUnitsOK {
	return &GetUnitsOK{}
}

/*GetUnitsOK handles this case with default header values.

List of organisation details
*/
type GetUnitsOK struct {
	Payload *models.OrganisationDetailsListResponse
}

func (o *GetUnitsOK) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsOK  %+v", 200, o.Payload)
}

func (o *GetUnitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganisationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitsBadRequest creates a GetUnitsBadRequest with default headers values
func NewGetUnitsBadRequest() *GetUnitsBadRequest {
	return &GetUnitsBadRequest{}
}

/*GetUnitsBadRequest handles this case with default header values.

Bad Request
*/
type GetUnitsBadRequest struct {
	Payload *models.APIError
}

func (o *GetUnitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUnitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitsUnauthorized creates a GetUnitsUnauthorized with default headers values
func NewGetUnitsUnauthorized() *GetUnitsUnauthorized {
	return &GetUnitsUnauthorized{}
}

/*GetUnitsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetUnitsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetUnitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUnitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitsForbidden creates a GetUnitsForbidden with default headers values
func NewGetUnitsForbidden() *GetUnitsForbidden {
	return &GetUnitsForbidden{}
}

/*GetUnitsForbidden handles this case with default header values.

Forbidden
*/
type GetUnitsForbidden struct {
	Payload *models.APIError
}

func (o *GetUnitsForbidden) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsForbidden  %+v", 403, o.Payload)
}

func (o *GetUnitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitsNotFound creates a GetUnitsNotFound with default headers values
func NewGetUnitsNotFound() *GetUnitsNotFound {
	return &GetUnitsNotFound{}
}

/*GetUnitsNotFound handles this case with default header values.

Record not found
*/
type GetUnitsNotFound struct {
	Payload *models.APIError
}

func (o *GetUnitsNotFound) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsNotFound  %+v", 404, o.Payload)
}

func (o *GetUnitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitsConflict creates a GetUnitsConflict with default headers values
func NewGetUnitsConflict() *GetUnitsConflict {
	return &GetUnitsConflict{}
}

/*GetUnitsConflict handles this case with default header values.

Conflict
*/
type GetUnitsConflict struct {
	Payload *models.APIError
}

func (o *GetUnitsConflict) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsConflict  %+v", 409, o.Payload)
}

func (o *GetUnitsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitsTooManyRequests creates a GetUnitsTooManyRequests with default headers values
func NewGetUnitsTooManyRequests() *GetUnitsTooManyRequests {
	return &GetUnitsTooManyRequests{}
}

/*GetUnitsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetUnitsTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetUnitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUnitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitsInternalServerError creates a GetUnitsInternalServerError with default headers values
func NewGetUnitsInternalServerError() *GetUnitsInternalServerError {
	return &GetUnitsInternalServerError{}
}

/*GetUnitsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetUnitsInternalServerError struct {
	Payload *models.APIError
}

func (o *GetUnitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUnitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitsServiceUnavailable creates a GetUnitsServiceUnavailable with default headers values
func NewGetUnitsServiceUnavailable() *GetUnitsServiceUnavailable {
	return &GetUnitsServiceUnavailable{}
}

/*GetUnitsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetUnitsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetUnitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUnitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
