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

// VmwareFcdInstantRecoveryDiskSpec vmware fcd instant recovery disk spec
//
// swagger:model VmwareFcdInstantRecoveryDiskSpec
type VmwareFcdInstantRecoveryDiskSpec struct {

	// Name of the VMDK file that will be stored in the datastore.
	// Required: true
	MountedDiskName *string `json:"mountedDiskName"`

	// Disk name displayed in the backup.
	// Required: true
	NameInBackup *string `json:"nameInBackup"`

	// Name under which the disk will be registered as an FCD in the vCenter Managed Object Browser.
	// Required: true
	RegisteredFcdName *string `json:"registeredFcdName"`
}

// Validate validates this vmware fcd instant recovery disk spec
func (m *VmwareFcdInstantRecoveryDiskSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMountedDiskName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameInBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredFcdName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareFcdInstantRecoveryDiskSpec) validateMountedDiskName(formats strfmt.Registry) error {

	if err := validate.Required("mountedDiskName", "body", m.MountedDiskName); err != nil {
		return err
	}

	return nil
}

func (m *VmwareFcdInstantRecoveryDiskSpec) validateNameInBackup(formats strfmt.Registry) error {

	if err := validate.Required("nameInBackup", "body", m.NameInBackup); err != nil {
		return err
	}

	return nil
}

func (m *VmwareFcdInstantRecoveryDiskSpec) validateRegisteredFcdName(formats strfmt.Registry) error {

	if err := validate.Required("registeredFcdName", "body", m.RegisteredFcdName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vmware fcd instant recovery disk spec based on context it is used
func (m *VmwareFcdInstantRecoveryDiskSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VmwareFcdInstantRecoveryDiskSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareFcdInstantRecoveryDiskSpec) UnmarshalBinary(b []byte) error {
	var res VmwareFcdInstantRecoveryDiskSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
