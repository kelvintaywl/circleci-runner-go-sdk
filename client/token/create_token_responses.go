// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kelvintaywl/circleci-runner-go-sdk/models"
)

// CreateTokenReader is a Reader for the CreateToken structure.
type CreateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTokenCreated creates a CreateTokenCreated with default headers values
func NewCreateTokenCreated() *CreateTokenCreated {
	return &CreateTokenCreated{}
}

/*
CreateTokenCreated describes a response with status code 201, with default header values.

A Runner token object
*/
type CreateTokenCreated struct {
	Payload *models.TokenCreated
}

// IsSuccess returns true when this create token created response has a 2xx status code
func (o *CreateTokenCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create token created response has a 3xx status code
func (o *CreateTokenCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token created response has a 4xx status code
func (o *CreateTokenCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create token created response has a 5xx status code
func (o *CreateTokenCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create token created response a status code equal to that given
func (o *CreateTokenCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateTokenCreated) Error() string {
	return fmt.Sprintf("[POST /v2/runner/token][%d] createTokenCreated  %+v", 201, o.Payload)
}

func (o *CreateTokenCreated) String() string {
	return fmt.Sprintf("[POST /v2/runner/token][%d] createTokenCreated  %+v", 201, o.Payload)
}

func (o *CreateTokenCreated) GetPayload() *models.TokenCreated {
	return o.Payload
}

func (o *CreateTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenCreated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenBadRequest creates a CreateTokenBadRequest with default headers values
func NewCreateTokenBadRequest() *CreateTokenBadRequest {
	return &CreateTokenBadRequest{}
}

/*
CreateTokenBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type CreateTokenBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this create token bad request response has a 2xx status code
func (o *CreateTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token bad request response has a 3xx status code
func (o *CreateTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token bad request response has a 4xx status code
func (o *CreateTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create token bad request response has a 5xx status code
func (o *CreateTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create token bad request response a status code equal to that given
func (o *CreateTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/runner/token][%d] createTokenBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTokenBadRequest) String() string {
	return fmt.Sprintf("[POST /v2/runner/token][%d] createTokenBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTokenBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *CreateTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenNotFound creates a CreateTokenNotFound with default headers values
func NewCreateTokenNotFound() *CreateTokenNotFound {
	return &CreateTokenNotFound{}
}

/*
CreateTokenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateTokenNotFound struct {
	Payload *models.Errored
}

// IsSuccess returns true when this create token not found response has a 2xx status code
func (o *CreateTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token not found response has a 3xx status code
func (o *CreateTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token not found response has a 4xx status code
func (o *CreateTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create token not found response has a 5xx status code
func (o *CreateTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create token not found response a status code equal to that given
func (o *CreateTokenNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/runner/token][%d] createTokenNotFound  %+v", 404, o.Payload)
}

func (o *CreateTokenNotFound) String() string {
	return fmt.Sprintf("[POST /v2/runner/token][%d] createTokenNotFound  %+v", 404, o.Payload)
}

func (o *CreateTokenNotFound) GetPayload() *models.Errored {
	return o.Payload
}

func (o *CreateTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
