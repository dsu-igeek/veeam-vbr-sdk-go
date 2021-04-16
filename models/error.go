// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Error error
//
// swagger:model Error
type Error struct {

	// The error code is a string that uniquely identifies an error condition and should be understood by programs that detect and handle errors by type
	// Required: true
	// Enum: [AccessDenied ExpiredToken InvalidToken InvalidURI MethodNotAllowed NotFound NotImplemented ServiceUnavailable UnexpectedContent UnknownError]
	ErrorCode *string `json:"errorCode"`

	// The error message contains a generic description of the error condition in English. It is intended for a human audience
	// Required: true
	Message *string `json:"message"`

	// ID of the object that is involved in the error (or empty)
	ResourceID string `json:"resourceId,omitempty"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var errorTypeErrorCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AccessDenied","ExpiredToken","InvalidToken","InvalidURI","MethodNotAllowed","NotFound","NotImplemented","ServiceUnavailable","UnexpectedContent","UnknownError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorTypeErrorCodePropEnum = append(errorTypeErrorCodePropEnum, v)
	}
}

const (

	// ErrorErrorCodeAccessDenied captures enum value "AccessDenied"
	ErrorErrorCodeAccessDenied string = "AccessDenied"

	// ErrorErrorCodeExpiredToken captures enum value "ExpiredToken"
	ErrorErrorCodeExpiredToken string = "ExpiredToken"

	// ErrorErrorCodeInvalidToken captures enum value "InvalidToken"
	ErrorErrorCodeInvalidToken string = "InvalidToken"

	// ErrorErrorCodeInvalidURI captures enum value "InvalidURI"
	ErrorErrorCodeInvalidURI string = "InvalidURI"

	// ErrorErrorCodeMethodNotAllowed captures enum value "MethodNotAllowed"
	ErrorErrorCodeMethodNotAllowed string = "MethodNotAllowed"

	// ErrorErrorCodeNotFound captures enum value "NotFound"
	ErrorErrorCodeNotFound string = "NotFound"

	// ErrorErrorCodeNotImplemented captures enum value "NotImplemented"
	ErrorErrorCodeNotImplemented string = "NotImplemented"

	// ErrorErrorCodeServiceUnavailable captures enum value "ServiceUnavailable"
	ErrorErrorCodeServiceUnavailable string = "ServiceUnavailable"

	// ErrorErrorCodeUnexpectedContent captures enum value "UnexpectedContent"
	ErrorErrorCodeUnexpectedContent string = "UnexpectedContent"

	// ErrorErrorCodeUnknownError captures enum value "UnknownError"
	ErrorErrorCodeUnknownError string = "UnknownError"
)

// prop value enum
func (m *Error) validateErrorCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, errorTypeErrorCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Error) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("errorCode", "body", m.ErrorCode); err != nil {
		return err
	}

	// value enum
	if err := m.validateErrorCodeEnum("errorCode", "body", *m.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error based on context it is used
func (m *Error) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
