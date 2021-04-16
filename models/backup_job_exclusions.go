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
)

// BackupJobExclusions backup job exclusions
//
// swagger:model BackupJobExclusions
type BackupJobExclusions struct {

	// Array of VM disks excluded from the backup.
	Disks []*VmwareObjectDiskModel `json:"disks"`

	// Array of VM templates excluded from the backup. OUTSIDE
	Templates *BackupJobExclusionsTemplates `json:"templates,omitempty"`

	// Array of VMs excluded from the backup.
	Vms []*VmwareObjectSizeModel `json:"vms"`
}

// Validate validates this backup job exclusions
func (m *BackupJobExclusions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupJobExclusions) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupJobExclusions) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	if m.Templates != nil {
		if err := m.Templates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templates")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobExclusions) validateVms(formats strfmt.Registry) error {
	if swag.IsZero(m.Vms) { // not required
		return nil
	}

	for i := 0; i < len(m.Vms); i++ {
		if swag.IsZero(m.Vms[i]) { // not required
			continue
		}

		if m.Vms[i] != nil {
			if err := m.Vms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this backup job exclusions based on the context it is used
func (m *BackupJobExclusions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupJobExclusions) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupJobExclusions) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	if m.Templates != nil {
		if err := m.Templates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templates")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobExclusions) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vms); i++ {

		if m.Vms[i] != nil {
			if err := m.Vms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupJobExclusions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupJobExclusions) UnmarshalBinary(b []byte) error {
	var res BackupJobExclusions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
