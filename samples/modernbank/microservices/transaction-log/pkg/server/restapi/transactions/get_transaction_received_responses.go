// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// GetTransactionReceivedOKCode is the HTTP code returned for type GetTransactionReceivedOK
const GetTransactionReceivedOKCode int = 200

/*GetTransactionReceivedOK Success!

swagger:response getTransactionReceivedOK
*/
type GetTransactionReceivedOK struct {

	/*
	  In: Body
	*/
	Payload *model.Transaction `json:"body,omitempty"`
}

// NewGetTransactionReceivedOK creates GetTransactionReceivedOK with default headers values
func NewGetTransactionReceivedOK() *GetTransactionReceivedOK {

	return &GetTransactionReceivedOK{}
}

// WithPayload adds the payload to the get transaction received o k response
func (o *GetTransactionReceivedOK) WithPayload(payload *model.Transaction) *GetTransactionReceivedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction received o k response
func (o *GetTransactionReceivedOK) SetPayload(payload *model.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionReceivedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTransactionReceivedNotFoundCode is the HTTP code returned for type GetTransactionReceivedNotFound
const GetTransactionReceivedNotFoundCode int = 404

/*GetTransactionReceivedNotFound Transaction not found

swagger:response getTransactionReceivedNotFound
*/
type GetTransactionReceivedNotFound struct {
}

// NewGetTransactionReceivedNotFound creates GetTransactionReceivedNotFound with default headers values
func NewGetTransactionReceivedNotFound() *GetTransactionReceivedNotFound {

	return &GetTransactionReceivedNotFound{}
}

// WriteResponse to the client
func (o *GetTransactionReceivedNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTransactionReceivedInternalServerErrorCode is the HTTP code returned for type GetTransactionReceivedInternalServerError
const GetTransactionReceivedInternalServerErrorCode int = 500

/*GetTransactionReceivedInternalServerError Internal server error

swagger:response getTransactionReceivedInternalServerError
*/
type GetTransactionReceivedInternalServerError struct {
}

// NewGetTransactionReceivedInternalServerError creates GetTransactionReceivedInternalServerError with default headers values
func NewGetTransactionReceivedInternalServerError() *GetTransactionReceivedInternalServerError {

	return &GetTransactionReceivedInternalServerError{}
}

// WriteResponse to the client
func (o *GetTransactionReceivedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}