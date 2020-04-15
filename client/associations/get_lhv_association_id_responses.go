// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetLhvAssociationIDReader is a Reader for the GetLhvAssociationID structure.
type GetLhvAssociationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLhvAssociationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLhvAssociationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetLhvAssociationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetLhvAssociationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetLhvAssociationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLhvAssociationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetLhvAssociationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetLhvAssociationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetLhvAssociationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetLhvAssociationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLhvAssociationIDOK creates a GetLhvAssociationIDOK with default headers values
func NewGetLhvAssociationIDOK() *GetLhvAssociationIDOK {
	return &GetLhvAssociationIDOK{}
}

/*GetLhvAssociationIDOK handles this case with default header values.

Associations details
*/
type GetLhvAssociationIDOK struct {
	Payload *models.LhvAssociationDetailsResponse
}

func (o *GetLhvAssociationIDOK) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdOK  %+v", 200, o.Payload)
}

func (o *GetLhvAssociationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LhvAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDBadRequest creates a GetLhvAssociationIDBadRequest with default headers values
func NewGetLhvAssociationIDBadRequest() *GetLhvAssociationIDBadRequest {
	return &GetLhvAssociationIDBadRequest{}
}

/*GetLhvAssociationIDBadRequest handles this case with default header values.

Bad Request
*/
type GetLhvAssociationIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLhvAssociationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDUnauthorized creates a GetLhvAssociationIDUnauthorized with default headers values
func NewGetLhvAssociationIDUnauthorized() *GetLhvAssociationIDUnauthorized {
	return &GetLhvAssociationIDUnauthorized{}
}

/*GetLhvAssociationIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetLhvAssociationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLhvAssociationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDForbidden creates a GetLhvAssociationIDForbidden with default headers values
func NewGetLhvAssociationIDForbidden() *GetLhvAssociationIDForbidden {
	return &GetLhvAssociationIDForbidden{}
}

/*GetLhvAssociationIDForbidden handles this case with default header values.

Forbidden
*/
type GetLhvAssociationIDForbidden struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDForbidden) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLhvAssociationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDNotFound creates a GetLhvAssociationIDNotFound with default headers values
func NewGetLhvAssociationIDNotFound() *GetLhvAssociationIDNotFound {
	return &GetLhvAssociationIDNotFound{}
}

/*GetLhvAssociationIDNotFound handles this case with default header values.

Record not found
*/
type GetLhvAssociationIDNotFound struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLhvAssociationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDConflict creates a GetLhvAssociationIDConflict with default headers values
func NewGetLhvAssociationIDConflict() *GetLhvAssociationIDConflict {
	return &GetLhvAssociationIDConflict{}
}

/*GetLhvAssociationIDConflict handles this case with default header values.

Conflict
*/
type GetLhvAssociationIDConflict struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDConflict) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdConflict  %+v", 409, o.Payload)
}

func (o *GetLhvAssociationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDTooManyRequests creates a GetLhvAssociationIDTooManyRequests with default headers values
func NewGetLhvAssociationIDTooManyRequests() *GetLhvAssociationIDTooManyRequests {
	return &GetLhvAssociationIDTooManyRequests{}
}

/*GetLhvAssociationIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetLhvAssociationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLhvAssociationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDInternalServerError creates a GetLhvAssociationIDInternalServerError with default headers values
func NewGetLhvAssociationIDInternalServerError() *GetLhvAssociationIDInternalServerError {
	return &GetLhvAssociationIDInternalServerError{}
}

/*GetLhvAssociationIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetLhvAssociationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLhvAssociationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLhvAssociationIDServiceUnavailable creates a GetLhvAssociationIDServiceUnavailable with default headers values
func NewGetLhvAssociationIDServiceUnavailable() *GetLhvAssociationIDServiceUnavailable {
	return &GetLhvAssociationIDServiceUnavailable{}
}

/*GetLhvAssociationIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetLhvAssociationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetLhvAssociationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /lhv/{associationId}][%d] getLhvAssociationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLhvAssociationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
