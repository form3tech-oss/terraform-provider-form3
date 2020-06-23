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

// GetPayportIDReader is a Reader for the GetPayportID structure.
type GetPayportIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPayportIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPayportIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPayportIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPayportIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPayportIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPayportIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPayportIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPayportIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPayportIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPayportIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPayportIDOK creates a GetPayportIDOK with default headers values
func NewGetPayportIDOK() *GetPayportIDOK {
	return &GetPayportIDOK{}
}

/*GetPayportIDOK handles this case with default header values.

Associations details
*/
type GetPayportIDOK struct {
	Payload *models.PayportAssociationDetailsResponse
}

func (o *GetPayportIDOK) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdOK  %+v", 200, o.Payload)
}

func (o *GetPayportIDOK) GetPayload() *models.PayportAssociationDetailsResponse {
	return o.Payload
}

func (o *GetPayportIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayportAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportIDBadRequest creates a GetPayportIDBadRequest with default headers values
func NewGetPayportIDBadRequest() *GetPayportIDBadRequest {
	return &GetPayportIDBadRequest{}
}

/*GetPayportIDBadRequest handles this case with default header values.

Bad Request
*/
type GetPayportIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetPayportIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetPayportIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportIDUnauthorized creates a GetPayportIDUnauthorized with default headers values
func NewGetPayportIDUnauthorized() *GetPayportIDUnauthorized {
	return &GetPayportIDUnauthorized{}
}

/*GetPayportIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetPayportIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetPayportIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPayportIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportIDForbidden creates a GetPayportIDForbidden with default headers values
func NewGetPayportIDForbidden() *GetPayportIDForbidden {
	return &GetPayportIDForbidden{}
}

/*GetPayportIDForbidden handles this case with default header values.

Forbidden
*/
type GetPayportIDForbidden struct {
	Payload *models.APIError
}

func (o *GetPayportIDForbidden) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdForbidden  %+v", 403, o.Payload)
}

func (o *GetPayportIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportIDNotFound creates a GetPayportIDNotFound with default headers values
func NewGetPayportIDNotFound() *GetPayportIDNotFound {
	return &GetPayportIDNotFound{}
}

/*GetPayportIDNotFound handles this case with default header values.

Record not found
*/
type GetPayportIDNotFound struct {
	Payload *models.APIError
}

func (o *GetPayportIDNotFound) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdNotFound  %+v", 404, o.Payload)
}

func (o *GetPayportIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportIDConflict creates a GetPayportIDConflict with default headers values
func NewGetPayportIDConflict() *GetPayportIDConflict {
	return &GetPayportIDConflict{}
}

/*GetPayportIDConflict handles this case with default header values.

Conflict
*/
type GetPayportIDConflict struct {
	Payload *models.APIError
}

func (o *GetPayportIDConflict) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdConflict  %+v", 409, o.Payload)
}

func (o *GetPayportIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportIDTooManyRequests creates a GetPayportIDTooManyRequests with default headers values
func NewGetPayportIDTooManyRequests() *GetPayportIDTooManyRequests {
	return &GetPayportIDTooManyRequests{}
}

/*GetPayportIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetPayportIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetPayportIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPayportIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportIDInternalServerError creates a GetPayportIDInternalServerError with default headers values
func NewGetPayportIDInternalServerError() *GetPayportIDInternalServerError {
	return &GetPayportIDInternalServerError{}
}

/*GetPayportIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPayportIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetPayportIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPayportIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportIDServiceUnavailable creates a GetPayportIDServiceUnavailable with default headers values
func NewGetPayportIDServiceUnavailable() *GetPayportIDServiceUnavailable {
	return &GetPayportIDServiceUnavailable{}
}

/*GetPayportIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetPayportIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetPayportIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPayportIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
