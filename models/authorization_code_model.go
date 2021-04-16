// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthorizationCodeModel authorization code model
//
// swagger:model AuthorizationCodeModel
type AuthorizationCodeModel struct {

	// String that is used to obtain an access token. Lifetime of the authorization code is 60 seconds.
	// Required: true
	Code *string `json:"code"`
}

// Validate validates this authorization code model
func (m *AuthorizationCodeModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationCodeModel) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this authorization code model based on context it is used
func (m *AuthorizationCodeModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizationCodeModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizationCodeModel) UnmarshalBinary(b []byte) error {
	var res AuthorizationCodeModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}