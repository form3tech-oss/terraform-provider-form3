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

// PostPaymentsIDSubmissionsReader is a Reader for the PostPaymentsIDSubmissions structure.
type PostPaymentsIDSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDSubmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostPaymentsIDSubmissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostPaymentsIDSubmissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostPaymentsIDSubmissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostPaymentsIDSubmissionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostPaymentsIDSubmissionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostPaymentsIDSubmissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPostPaymentsIDSubmissionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDSubmissionsCreated creates a PostPaymentsIDSubmissionsCreated with default headers values
func NewPostPaymentsIDSubmissionsCreated() *PostPaymentsIDSubmissionsCreated {
	return &PostPaymentsIDSubmissionsCreated{}
}

/*PostPaymentsIDSubmissionsCreated handles this case with default header values.

Submission creation response
*/
type PostPaymentsIDSubmissionsCreated struct {
	Payload *models.PaymentSubmissionCreationResponse
}

func (o *PostPaymentsIDSubmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDSubmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSubmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsBadRequest creates a PostPaymentsIDSubmissionsBadRequest with default headers values
func NewPostPaymentsIDSubmissionsBadRequest() *PostPaymentsIDSubmissionsBadRequest {
	return &PostPaymentsIDSubmissionsBadRequest{}
}

/*PostPaymentsIDSubmissionsBadRequest handles this case with default header values.

Bad Request
*/
type PostPaymentsIDSubmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsUnauthorized creates a PostPaymentsIDSubmissionsUnauthorized with default headers values
func NewPostPaymentsIDSubmissionsUnauthorized() *PostPaymentsIDSubmissionsUnauthorized {
	return &PostPaymentsIDSubmissionsUnauthorized{}
}

/*PostPaymentsIDSubmissionsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostPaymentsIDSubmissionsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPaymentsIDSubmissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsForbidden creates a PostPaymentsIDSubmissionsForbidden with default headers values
func NewPostPaymentsIDSubmissionsForbidden() *PostPaymentsIDSubmissionsForbidden {
	return &PostPaymentsIDSubmissionsForbidden{}
}

/*PostPaymentsIDSubmissionsForbidden handles this case with default header values.

Forbidden
*/
type PostPaymentsIDSubmissionsForbidden struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsForbidden) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsForbidden  %+v", 403, o.Payload)
}

func (o *PostPaymentsIDSubmissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsNotFound creates a PostPaymentsIDSubmissionsNotFound with default headers values
func NewPostPaymentsIDSubmissionsNotFound() *PostPaymentsIDSubmissionsNotFound {
	return &PostPaymentsIDSubmissionsNotFound{}
}

/*PostPaymentsIDSubmissionsNotFound handles this case with default header values.

Record not found
*/
type PostPaymentsIDSubmissionsNotFound struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsNotFound) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsNotFound  %+v", 404, o.Payload)
}

func (o *PostPaymentsIDSubmissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsConflict creates a PostPaymentsIDSubmissionsConflict with default headers values
func NewPostPaymentsIDSubmissionsConflict() *PostPaymentsIDSubmissionsConflict {
	return &PostPaymentsIDSubmissionsConflict{}
}

/*PostPaymentsIDSubmissionsConflict handles this case with default header values.

Conflict
*/
type PostPaymentsIDSubmissionsConflict struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsConflict) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsConflict  %+v", 409, o.Payload)
}

func (o *PostPaymentsIDSubmissionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsTooManyRequests creates a PostPaymentsIDSubmissionsTooManyRequests with default headers values
func NewPostPaymentsIDSubmissionsTooManyRequests() *PostPaymentsIDSubmissionsTooManyRequests {
	return &PostPaymentsIDSubmissionsTooManyRequests{}
}

/*PostPaymentsIDSubmissionsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostPaymentsIDSubmissionsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPaymentsIDSubmissionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsInternalServerError creates a PostPaymentsIDSubmissionsInternalServerError with default headers values
func NewPostPaymentsIDSubmissionsInternalServerError() *PostPaymentsIDSubmissionsInternalServerError {
	return &PostPaymentsIDSubmissionsInternalServerError{}
}

/*PostPaymentsIDSubmissionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPaymentsIDSubmissionsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPaymentsIDSubmissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsServiceUnavailable creates a PostPaymentsIDSubmissionsServiceUnavailable with default headers values
func NewPostPaymentsIDSubmissionsServiceUnavailable() *PostPaymentsIDSubmissionsServiceUnavailable {
	return &PostPaymentsIDSubmissionsServiceUnavailable{}
}

/*PostPaymentsIDSubmissionsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostPaymentsIDSubmissionsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPaymentsIDSubmissionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
