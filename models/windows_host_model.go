// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WindowsHostModel windows host model
//
// swagger:model WindowsHostModel
type WindowsHostModel struct {
	ManagedServerModel

	// network settings
	NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WindowsHostModel) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ManagedServerModel
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ManagedServerModel = aO0

	// AO1
	var dataAO1 struct {
		NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.NetworkSettings = dataAO1.NetworkSettings

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WindowsHostModel) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ManagedServerModel)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
	}

	dataAO1.NetworkSettings = m.NetworkSettings

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this windows host model
func (m *WindowsHostModel) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ManagedServerModel
	if err := m.ManagedServerModel.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WindowsHostModel) validateNetworkSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this windows host model based on the context it is used
func (m *WindowsHostModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ManagedServerModel
	if err := m.ManagedServerModel.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WindowsHostModel) contextValidateNetworkSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WindowsHostModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WindowsHostModel) UnmarshalBinary(b []byte) error {
	var res WindowsHostModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
