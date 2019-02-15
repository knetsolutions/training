// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// ListTransactionsSentReader is a Reader for the ListTransactionsSent structure.
type ListTransactionsSentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTransactionsSentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTransactionsSentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewListTransactionsSentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListTransactionsSentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListTransactionsSentOK creates a ListTransactionsSentOK with default headers values
func NewListTransactionsSentOK() *ListTransactionsSentOK {
	return &ListTransactionsSentOK{}
}

/*ListTransactionsSentOK handles this case with default header values.

Success!
*/
type ListTransactionsSentOK struct {
	Payload []*model.Transaction
}

func (o *ListTransactionsSentOK) Error() string {
	return fmt.Sprintf("[GET /account/{sender}/sent][%d] listTransactionsSentOK  %+v", 200, o.Payload)
}

func (o *ListTransactionsSentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransactionsSentNotFound creates a ListTransactionsSentNotFound with default headers values
func NewListTransactionsSentNotFound() *ListTransactionsSentNotFound {
	return &ListTransactionsSentNotFound{}
}

/*ListTransactionsSentNotFound handles this case with default header values.

No transactions found
*/
type ListTransactionsSentNotFound struct {
}

func (o *ListTransactionsSentNotFound) Error() string {
	return fmt.Sprintf("[GET /account/{sender}/sent][%d] listTransactionsSentNotFound ", 404)
}

func (o *ListTransactionsSentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListTransactionsSentInternalServerError creates a ListTransactionsSentInternalServerError with default headers values
func NewListTransactionsSentInternalServerError() *ListTransactionsSentInternalServerError {
	return &ListTransactionsSentInternalServerError{}
}

/*ListTransactionsSentInternalServerError handles this case with default header values.

Internal server error
*/
type ListTransactionsSentInternalServerError struct {
}

func (o *ListTransactionsSentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /account/{sender}/sent][%d] listTransactionsSentInternalServerError ", 500)
}

func (o *ListTransactionsSentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
