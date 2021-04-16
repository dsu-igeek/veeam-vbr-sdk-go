// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ELoginGrantType Authorization grant type.
//
// Available values:
// - `password` — used to obtain an access token by providing a user name and password.
// - `refresh_token` — used to refresh an expired or lost access token by providing a refresh token.
// - `authorization_code` — used to obtain an access token by providing an authorization code.
//
//
// swagger:model ELoginGrantType
type ELoginGrantType string

func NewELoginGrantType(value ELoginGrantType) *ELoginGrantType {
	v := value
	return &v
}

const (

	// ELoginGrantTypePassword captures enum value "password"
	ELoginGrantTypePassword ELoginGrantType = "password"

	// ELoginGrantTypeRefreshToken captures enum value "refresh_token"
	ELoginGrantTypeRefreshToken ELoginGrantType = "refresh_token"

	// ELoginGrantTypeAuthorizationCode captures enum value "authorization_code"
	ELoginGrantTypeAuthorizationCode ELoginGrantType = "authorization_code"
)

// for schema
var eLoginGrantTypeEnum []interface{}

func init() {
	var res []ELoginGrantType
	if err := json.Unmarshal([]byte(`["password","refresh_token","authorization_code"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eLoginGrantTypeEnum = append(eLoginGrantTypeEnum, v)
	}
}

func (m ELoginGrantType) validateELoginGrantTypeEnum(path, location string, value ELoginGrantType) error {
	if err := validate.EnumCase(path, location, value, eLoginGrantTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e login grant type
func (m ELoginGrantType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateELoginGrantTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e login grant type based on context it is used
func (m ELoginGrantType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
