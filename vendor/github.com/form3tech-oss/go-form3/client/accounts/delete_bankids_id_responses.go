// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteBankidsIDReader is a Reader for the DeleteBankidsID structure.
type DeleteBankidsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBankidsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteBankidsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteBankidsIDNoContent creates a DeleteBankidsIDNoContent with default headers values
func NewDeleteBankidsIDNoContent() *DeleteBankidsIDNoContent {
	return &DeleteBankidsIDNoContent{}
}

/*DeleteBankidsIDNoContent handles this case with default header values.

BankId deleted
*/
type DeleteBankidsIDNoContent struct {
}

func (o *DeleteBankidsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /bankids/{id}][%d] deleteBankidsIdNoContent ", 204)
}

func (o *DeleteBankidsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
