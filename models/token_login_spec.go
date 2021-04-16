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

// TokenLoginSpec token login spec
//
// swagger:model TokenLoginSpec
type TokenLoginSpec struct {

	// Authorization code. Required if the `grant_type` value is `authorization_code`.
	Code string `json:"code,omitempty"`

	// grant type
	// Required: true
	GrantType *ELoginGrantType `json:"grant_type"`

	// Password. Required if the `grant_type` value is `password`.
	// Format: password
	Password strfmt.Password `json:"password,omitempty"`

	// Refresh token. Required if the `grant_type` value is `refresh_token`.
	RefreshToken string `json:"refresh_token,omitempty"`

	// If *true*, a short-term refresh token is used. Lifetime of the short-term refresh token is the access token lifetime plus 15 minutes.
	UseShortTermRefresh bool `json:"use_short_term_refresh,omitempty"`

	// User name. Required if the `grant_type` value is `password`.
	Username string `json:"username,omitempty"`
}

// Validate validates this token login spec
func (m *TokenLoginSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrantType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenLoginSpec) validateGrantType(formats strfmt.Registry) error {

	if err := validate.Required("grant_type", "body", m.GrantType); err != nil {
		return err
	}

	if err := validate.Required("grant_type", "body", m.GrantType); err != nil {
		return err
	}

	if m.GrantType != nil {
		if err := m.GrantType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grant_type")
			}
			return err
		}
	}

	return nil
}

func (m *TokenLoginSpec) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this token login spec based on the context it is used
func (m *TokenLoginSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGrantType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenLoginSpec) contextValidateGrantType(ctx context.Context, formats strfmt.Registry) error {

	if m.GrantType != nil {
		if err := m.GrantType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grant_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenLoginSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenLoginSpec) UnmarshalBinary(b []byte) error {
	var res TokenLoginSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
