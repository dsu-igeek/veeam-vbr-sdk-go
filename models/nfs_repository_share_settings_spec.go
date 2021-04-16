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

// NfsRepositoryShareSettingsSpec nfs repository share settings spec
//
// swagger:model NfsRepositoryShareSettingsSpec
type NfsRepositoryShareSettingsSpec struct {

	// gateway server
	GatewayServer *RepositoryShareGatewayImportSpec `json:"gatewayServer,omitempty"`

	// Path to the shared folder that that is used as a backup repository.
	// Required: true
	SharePath *string `json:"sharePath"`
}

// Validate validates this nfs repository share settings spec
func (m *NfsRepositoryShareSettingsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGatewayServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsRepositoryShareSettingsSpec) validateGatewayServer(formats strfmt.Registry) error {
	if swag.IsZero(m.GatewayServer) { // not required
		return nil
	}

	if m.GatewayServer != nil {
		if err := m.GatewayServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gatewayServer")
			}
			return err
		}
	}

	return nil
}

func (m *NfsRepositoryShareSettingsSpec) validateSharePath(formats strfmt.Registry) error {

	if err := validate.Required("sharePath", "body", m.SharePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nfs repository share settings spec based on the context it is used
func (m *NfsRepositoryShareSettingsSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGatewayServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsRepositoryShareSettingsSpec) contextValidateGatewayServer(ctx context.Context, formats strfmt.Registry) error {

	if m.GatewayServer != nil {
		if err := m.GatewayServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gatewayServer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NfsRepositoryShareSettingsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NfsRepositoryShareSettingsSpec) UnmarshalBinary(b []byte) error {
	var res NfsRepositoryShareSettingsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}