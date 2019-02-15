// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/tetrateio/training/samples/modernbank/microservices/user/pkg/model"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/*CreateUserCreated handles this case with default header values.

Created
*/
type CreateUserCreated struct {
	Payload *model.User
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserConflict creates a CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {
	return &CreateUserConflict{}
}

/*CreateUserConflict handles this case with default header values.

User alreadys exists
*/
type CreateUserConflict struct {
}

func (o *CreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserConflict ", 409)
}

func (o *CreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserInternalServerError creates a CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {
	return &CreateUserInternalServerError{}
}

/*CreateUserInternalServerError handles this case with default header values.

Internal server error
*/
type CreateUserInternalServerError struct {
}

func (o *CreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserInternalServerError ", 500)
}

func (o *CreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}