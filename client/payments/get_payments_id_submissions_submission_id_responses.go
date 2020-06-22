// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetPaymentsIDSubmissionsSubmissionIDReader is a Reader for the GetPaymentsIDSubmissionsSubmissionID structure.
type GetPaymentsIDSubmissionsSubmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDSubmissionsSubmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDSubmissionsSubmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPaymentsIDSubmissionsSubmissionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetPaymentsIDSubmissionsSubmissionIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPaymentsIDSubmissionsSubmissionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPaymentsIDSubmissionsSubmissionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetPaymentsIDSubmissionsSubmissionIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetPaymentsIDSubmissionsSubmissionIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetPaymentsIDSubmissionsSubmissionIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetPaymentsIDSubmissionsSubmissionIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDSubmissionsSubmissionIDOK creates a GetPaymentsIDSubmissionsSubmissionIDOK with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDOK() *GetPaymentsIDSubmissionsSubmissionIDOK {
	return &GetPaymentsIDSubmissionsSubmissionIDOK{}
}

/*GetPaymentsIDSubmissionsSubmissionIDOK handles this case with default header values.

Submission details
*/
type GetPaymentsIDSubmissionsSubmissionIDOK struct {
	Payload *models.PaymentSubmissionDetailsResponse
}

func (o *GetPaymentsIDSubmissionsSubmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSubmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDSubmissionsSubmissionIDBadRequest creates a GetPaymentsIDSubmissionsSubmissionIDBadRequest with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDBadRequest() *GetPaymentsIDSubmissionsSubmissionIDBadRequest {
	return &GetPaymentsIDSubmissionsSubmissionIDBadRequest{}
}

/*GetPaymentsIDSubmissionsSubmissionIDBadRequest handles this case with default header values.

Bad Request
*/
type GetPaymentsIDSubmissionsSubmissionIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDSubmissionsSubmissionIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDSubmissionsSubmissionIDUnauthorized creates a GetPaymentsIDSubmissionsSubmissionIDUnauthorized with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDUnauthorized() *GetPaymentsIDSubmissionsSubmissionIDUnauthorized {
	return &GetPaymentsIDSubmissionsSubmissionIDUnauthorized{}
}

/*GetPaymentsIDSubmissionsSubmissionIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetPaymentsIDSubmissionsSubmissionIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDSubmissionsSubmissionIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDSubmissionsSubmissionIDForbidden creates a GetPaymentsIDSubmissionsSubmissionIDForbidden with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDForbidden() *GetPaymentsIDSubmissionsSubmissionIDForbidden {
	return &GetPaymentsIDSubmissionsSubmissionIDForbidden{}
}

/*GetPaymentsIDSubmissionsSubmissionIDForbidden handles this case with default header values.

Forbidden
*/
type GetPaymentsIDSubmissionsSubmissionIDForbidden struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDSubmissionsSubmissionIDForbidden) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdForbidden  %+v", 403, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDSubmissionsSubmissionIDNotFound creates a GetPaymentsIDSubmissionsSubmissionIDNotFound with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDNotFound() *GetPaymentsIDSubmissionsSubmissionIDNotFound {
	return &GetPaymentsIDSubmissionsSubmissionIDNotFound{}
}

/*GetPaymentsIDSubmissionsSubmissionIDNotFound handles this case with default header values.

Record not found
*/
type GetPaymentsIDSubmissionsSubmissionIDNotFound struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDSubmissionsSubmissionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdNotFound  %+v", 404, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDSubmissionsSubmissionIDConflict creates a GetPaymentsIDSubmissionsSubmissionIDConflict with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDConflict() *GetPaymentsIDSubmissionsSubmissionIDConflict {
	return &GetPaymentsIDSubmissionsSubmissionIDConflict{}
}

/*GetPaymentsIDSubmissionsSubmissionIDConflict handles this case with default header values.

Conflict
*/
type GetPaymentsIDSubmissionsSubmissionIDConflict struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDSubmissionsSubmissionIDConflict) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdConflict  %+v", 409, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDSubmissionsSubmissionIDTooManyRequests creates a GetPaymentsIDSubmissionsSubmissionIDTooManyRequests with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDTooManyRequests() *GetPaymentsIDSubmissionsSubmissionIDTooManyRequests {
	return &GetPaymentsIDSubmissionsSubmissionIDTooManyRequests{}
}

/*GetPaymentsIDSubmissionsSubmissionIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetPaymentsIDSubmissionsSubmissionIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDSubmissionsSubmissionIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDSubmissionsSubmissionIDInternalServerError creates a GetPaymentsIDSubmissionsSubmissionIDInternalServerError with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDInternalServerError() *GetPaymentsIDSubmissionsSubmissionIDInternalServerError {
	return &GetPaymentsIDSubmissionsSubmissionIDInternalServerError{}
}

/*GetPaymentsIDSubmissionsSubmissionIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPaymentsIDSubmissionsSubmissionIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDSubmissionsSubmissionIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDSubmissionsSubmissionIDServiceUnavailable creates a GetPaymentsIDSubmissionsSubmissionIDServiceUnavailable with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDServiceUnavailable() *GetPaymentsIDSubmissionsSubmissionIDServiceUnavailable {
	return &GetPaymentsIDSubmissionsSubmissionIDServiceUnavailable{}
}

/*GetPaymentsIDSubmissionsSubmissionIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetPaymentsIDSubmissionsSubmissionIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDSubmissionsSubmissionIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}][%d] getPaymentsIdSubmissionsSubmissionIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
