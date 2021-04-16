// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// CreateTokenReader is a Reader for the CreateToken structure.
type CreateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTokenOK()
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
	case 401:
		result := NewCreateTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTokenOK creates a CreateTokenOK with default headers values
func NewCreateTokenOK() *CreateTokenOK {
	return &CreateTokenOK{}
}

/* CreateTokenOK describes a response with status code 200, with default header values.

OK
*/
type CreateTokenOK struct {
	Payload *models.TokenModel
}

func (o *CreateTokenOK) Error() string {
	return fmt.Sprintf("[POST /api/oauth2/token][%d] createTokenOK  %+v", 200, o.Payload)
}
func (o *CreateTokenOK) GetPayload() *models.TokenModel {
	return o.Payload
}

func (o *CreateTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenModel)

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

/* CreateTokenBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type CreateTokenBadRequest struct {
	Payload *models.Error
}

func (o *CreateTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/oauth2/token][%d] createTokenBadRequest  %+v", 400, o.Payload)
}
func (o *CreateTokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenUnauthorized creates a CreateTokenUnauthorized with default headers values
func NewCreateTokenUnauthorized() *CreateTokenUnauthorized {
	return &CreateTokenUnauthorized{}
}

/* CreateTokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type CreateTokenUnauthorized struct {
	Payload *models.Error
}

func (o *CreateTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/oauth2/token][%d] createTokenUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateTokenUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenForbidden creates a CreateTokenForbidden with default headers values
func NewCreateTokenForbidden() *CreateTokenForbidden {
	return &CreateTokenForbidden{}
}

/* CreateTokenForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type CreateTokenForbidden struct {
	Payload *models.Error
}

func (o *CreateTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /api/oauth2/token][%d] createTokenForbidden  %+v", 403, o.Payload)
}
func (o *CreateTokenForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenInternalServerError creates a CreateTokenInternalServerError with default headers values
func NewCreateTokenInternalServerError() *CreateTokenInternalServerError {
	return &CreateTokenInternalServerError{}
}

/* CreateTokenInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type CreateTokenInternalServerError struct {
	Payload *models.Error
}

func (o *CreateTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/oauth2/token][%d] createTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateTokenInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
