// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/models"
)

// GetAccountsIDReader is a Reader for the GetAccountsID structure.
type GetAccountsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountsIDOK creates a GetAccountsIDOK with default headers values
func NewGetAccountsIDOK() *GetAccountsIDOK {
	return &GetAccountsIDOK{}
}

/*GetAccountsIDOK handles this case with default header values.

Account details
*/
type GetAccountsIDOK struct {
	Payload *models.AccountDetailsResponse
}

func (o *GetAccountsIDOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{id}][%d] getAccountsIdOK  %+v", 200, o.Payload)
}

func (o *GetAccountsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
