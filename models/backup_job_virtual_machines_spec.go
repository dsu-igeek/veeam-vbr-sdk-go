// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupJobVirtualMachinesSpec backup job virtual machines spec
//
// swagger:model BackupJobVirtualMachinesSpec
type BackupJobVirtualMachinesSpec struct {

	// excludes
	Excludes *BackupJobExclusionsSpec `json:"excludes,omitempty"`

	// includes
	// Required: true
	Includes []*VmwareObjectModel `json:"includes"`
}

// Validate validates this backup job virtual machines spec
func (m *BackupJobVirtualMachinesSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExcludes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupJobVirtualMachinesSpec) validateExcludes(formats strfmt.Registry) error {
	if swag.IsZero(m.Excludes) { // not required
		return nil
	}

	if m.Excludes != nil {
		if err := m.Excludes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excludes")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobVirtualMachinesSpec) validateIncludes(formats strfmt.Registry) error {

	if err := validate.Required("includes", "body", m.Includes); err != nil {
		return err
	}

	for i := 0; i < len(m.Includes); i++ {
		if swag.IsZero(m.Includes[i]) { // not required
			continue
		}

		if m.Includes[i] != nil {
			if err := m.Includes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("includes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this backup job virtual machines spec based on the context it is used
func (m *BackupJobVirtualMachinesSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExcludes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncludes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupJobVirtualMachinesSpec) contextValidateExcludes(ctx context.Context, formats strfmt.Registry) error {

	if m.Excludes != nil {
		if err := m.Excludes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excludes")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobVirtualMachinesSpec) contextValidateIncludes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Includes); i++ {

		if m.Includes[i] != nil {
			if err := m.Includes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("includes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupJobVirtualMachinesSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupJobVirtualMachinesSpec) UnmarshalBinary(b []byte) error {
	var res BackupJobVirtualMachinesSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
