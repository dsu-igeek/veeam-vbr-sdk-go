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

// GuestOsCredentialsPerMachineModel guest os credentials per machine model
//
// swagger:model GuestOsCredentialsPerMachineModel
type GuestOsCredentialsPerMachineModel struct {

	// Credentials ID for a Linux VM.
	// Format: uuid
	LinuxCredsID strfmt.UUID `json:"linuxCredsId,omitempty"`

	// vm object
	// Required: true
	VMObject *VmwareObjectModel `json:"vmObject"`

	// Credentials ID for a Microsoft Windows VM.
	// Format: uuid
	WindowsCredsID strfmt.UUID `json:"windowsCredsId,omitempty"`
}

// Validate validates this guest os credentials per machine model
func (m *GuestOsCredentialsPerMachineModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinuxCredsID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWindowsCredsID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GuestOsCredentialsPerMachineModel) validateLinuxCredsID(formats strfmt.Registry) error {
	if swag.IsZero(m.LinuxCredsID) { // not required
		return nil
	}

	if err := validate.FormatOf("linuxCredsId", "body", "uuid", m.LinuxCredsID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GuestOsCredentialsPerMachineModel) validateVMObject(formats strfmt.Registry) error {

	if err := validate.Required("vmObject", "body", m.VMObject); err != nil {
		return err
	}

	if m.VMObject != nil {
		if err := m.VMObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmObject")
			}
			return err
		}
	}

	return nil
}

func (m *GuestOsCredentialsPerMachineModel) validateWindowsCredsID(formats strfmt.Registry) error {
	if swag.IsZero(m.WindowsCredsID) { // not required
		return nil
	}

	if err := validate.FormatOf("windowsCredsId", "body", "uuid", m.WindowsCredsID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this guest os credentials per machine model based on the context it is used
func (m *GuestOsCredentialsPerMachineModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVMObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GuestOsCredentialsPerMachineModel) contextValidateVMObject(ctx context.Context, formats strfmt.Registry) error {

	if m.VMObject != nil {
		if err := m.VMObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmObject")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GuestOsCredentialsPerMachineModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GuestOsCredentialsPerMachineModel) UnmarshalBinary(b []byte) error {
	var res GuestOsCredentialsPerMachineModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
