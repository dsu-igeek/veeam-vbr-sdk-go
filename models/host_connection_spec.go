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

// HostConnectionSpec host connection spec
//
// swagger:model HostConnectionSpec
type HostConnectionSpec struct {

	// ID of a credentials record used to connect to the server.
	// Required: true
	// Format: uuid
	CredentialsID *strfmt.UUID `json:"credentialsId"`

	// Port used to communicate with the server.
	Port int64 `json:"port,omitempty"`

	// Full DNS name or IP address of the server.
	// Required: true
	ServerName *string `json:"serverName"`

	// type
	// Required: true
	Type *EManagedServerType `json:"type"`
}

// Validate validates this host connection spec
func (m *HostConnectionSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostConnectionSpec) validateCredentialsID(formats strfmt.Registry) error {

	if err := validate.Required("credentialsId", "body", m.CredentialsID); err != nil {
		return err
	}

	if err := validate.FormatOf("credentialsId", "body", "uuid", m.CredentialsID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HostConnectionSpec) validateServerName(formats strfmt.Registry) error {

	if err := validate.Required("serverName", "body", m.ServerName); err != nil {
		return err
	}

	return nil
}

func (m *HostConnectionSpec) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this host connection spec based on the context it is used
func (m *HostConnectionSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostConnectionSpec) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostConnectionSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostConnectionSpec) UnmarshalBinary(b []byte) error {
	var res HostConnectionSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
