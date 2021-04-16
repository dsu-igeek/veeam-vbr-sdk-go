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

// ServerTimeModel server time model
//
// swagger:model ServerTimeModel
type ServerTimeModel struct {

	// Current date and time on the server.
	// Required: true
	// Format: date-time
	ServerTime *strfmt.DateTime `json:"serverTime"`

	// time zone
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this server time model
func (m *ServerTimeModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerTimeModel) validateServerTime(formats strfmt.Registry) error {

	if err := validate.Required("serverTime", "body", m.ServerTime); err != nil {
		return err
	}

	if err := validate.FormatOf("serverTime", "body", "date-time", m.ServerTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this server time model based on context it is used
func (m *ServerTimeModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerTimeModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerTimeModel) UnmarshalBinary(b []byte) error {
	var res ServerTimeModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}