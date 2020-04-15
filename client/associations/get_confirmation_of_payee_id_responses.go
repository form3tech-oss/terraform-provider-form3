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

// GetConfirmationOfPayeeIDReader is a Reader for the GetConfirmationOfPayeeID structure.
type GetConfirmationOfPayeeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfirmationOfPayeeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConfirmationOfPayeeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetConfirmationOfPayeeIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetConfirmationOfPayeeIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetConfirmationOfPayeeIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetConfirmationOfPayeeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetConfirmationOfPayeeIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetConfirmationOfPayeeIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetConfirmationOfPayeeIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetConfirmationOfPayeeIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConfirmationOfPayeeIDOK creates a GetConfirmationOfPayeeIDOK with default headers values
func NewGetConfirmationOfPayeeIDOK() *GetConfirmationOfPayeeIDOK {
	return &GetConfirmationOfPayeeIDOK{}
}

/*GetConfirmationOfPayeeIDOK handles this case with default header values.

Associations details
*/
type GetConfirmationOfPayeeIDOK struct {
	Payload *models.CoPAssociationDetailsResponse
}

func (o *GetConfirmationOfPayeeIDOK) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdOK  %+v", 200, o.Payload)
}

func (o *GetConfirmationOfPayeeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoPAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfirmationOfPayeeIDBadRequest creates a GetConfirmationOfPayeeIDBadRequest with default headers values
func NewGetConfirmationOfPayeeIDBadRequest() *GetConfirmationOfPayeeIDBadRequest {
	return &GetConfirmationOfPayeeIDBadRequest{}
}

/*GetConfirmationOfPayeeIDBadRequest handles this case with default header values.

Bad Request
*/
type GetConfirmationOfPayeeIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetConfirmationOfPayeeIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfirmationOfPayeeIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfirmationOfPayeeIDUnauthorized creates a GetConfirmationOfPayeeIDUnauthorized with default headers values
func NewGetConfirmationOfPayeeIDUnauthorized() *GetConfirmationOfPayeeIDUnauthorized {
	return &GetConfirmationOfPayeeIDUnauthorized{}
}

/*GetConfirmationOfPayeeIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetConfirmationOfPayeeIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetConfirmationOfPayeeIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConfirmationOfPayeeIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfirmationOfPayeeIDForbidden creates a GetConfirmationOfPayeeIDForbidden with default headers values
func NewGetConfirmationOfPayeeIDForbidden() *GetConfirmationOfPayeeIDForbidden {
	return &GetConfirmationOfPayeeIDForbidden{}
}

/*GetConfirmationOfPayeeIDForbidden handles this case with default header values.

Forbidden
*/
type GetConfirmationOfPayeeIDForbidden struct {
	Payload *models.APIError
}

func (o *GetConfirmationOfPayeeIDForbidden) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdForbidden  %+v", 403, o.Payload)
}

func (o *GetConfirmationOfPayeeIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfirmationOfPayeeIDNotFound creates a GetConfirmationOfPayeeIDNotFound with default headers values
func NewGetConfirmationOfPayeeIDNotFound() *GetConfirmationOfPayeeIDNotFound {
	return &GetConfirmationOfPayeeIDNotFound{}
}

/*GetConfirmationOfPayeeIDNotFound handles this case with default header values.

Record not found
*/
type GetConfirmationOfPayeeIDNotFound struct {
	Payload *models.APIError
}

func (o *GetConfirmationOfPayeeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdNotFound  %+v", 404, o.Payload)
}

func (o *GetConfirmationOfPayeeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfirmationOfPayeeIDConflict creates a GetConfirmationOfPayeeIDConflict with default headers values
func NewGetConfirmationOfPayeeIDConflict() *GetConfirmationOfPayeeIDConflict {
	return &GetConfirmationOfPayeeIDConflict{}
}

/*GetConfirmationOfPayeeIDConflict handles this case with default header values.

Conflict
*/
type GetConfirmationOfPayeeIDConflict struct {
	Payload *models.APIError
}

func (o *GetConfirmationOfPayeeIDConflict) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdConflict  %+v", 409, o.Payload)
}

func (o *GetConfirmationOfPayeeIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfirmationOfPayeeIDTooManyRequests creates a GetConfirmationOfPayeeIDTooManyRequests with default headers values
func NewGetConfirmationOfPayeeIDTooManyRequests() *GetConfirmationOfPayeeIDTooManyRequests {
	return &GetConfirmationOfPayeeIDTooManyRequests{}
}

/*GetConfirmationOfPayeeIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetConfirmationOfPayeeIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetConfirmationOfPayeeIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfirmationOfPayeeIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfirmationOfPayeeIDInternalServerError creates a GetConfirmationOfPayeeIDInternalServerError with default headers values
func NewGetConfirmationOfPayeeIDInternalServerError() *GetConfirmationOfPayeeIDInternalServerError {
	return &GetConfirmationOfPayeeIDInternalServerError{}
}

/*GetConfirmationOfPayeeIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetConfirmationOfPayeeIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetConfirmationOfPayeeIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfirmationOfPayeeIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfirmationOfPayeeIDServiceUnavailable creates a GetConfirmationOfPayeeIDServiceUnavailable with default headers values
func NewGetConfirmationOfPayeeIDServiceUnavailable() *GetConfirmationOfPayeeIDServiceUnavailable {
	return &GetConfirmationOfPayeeIDServiceUnavailable{}
}

/*GetConfirmationOfPayeeIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetConfirmationOfPayeeIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetConfirmationOfPayeeIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee/{id}][%d] getConfirmationOfPayeeIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConfirmationOfPayeeIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
