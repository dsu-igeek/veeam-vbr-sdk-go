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

// VmwareFcdInstantRecoverySpec vmware fcd instant recovery spec
//
// swagger:model VmwareFcdInstantRecoverySpec
type VmwareFcdInstantRecoverySpec struct {

	// Destination cluster.
	// Required: true
	DestinationCluster struct {
		VmwareObjectModel
	} `json:"destinationCluster"`

	// Array of disks that will be restored.
	// Required: true
	DisksMapping []*VmwareFcdInstantRecoveryDiskSpec `json:"disksMapping"`

	// ID of the restore point.
	// Required: true
	// Format: uuid
	ObjectRestorePointID *strfmt.UUID `json:"objectRestorePointId"`

	// write cache
	WriteCache *VmwareFcdWriteCacheSpec `json:"writeCache,omitempty"`
}

// Validate validates this vmware fcd instant recovery spec
func (m *VmwareFcdInstantRecoverySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisksMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectRestorePointID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWriteCache(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareFcdInstantRecoverySpec) validateDestinationCluster(formats strfmt.Registry) error {

	return nil
}

func (m *VmwareFcdInstantRecoverySpec) validateDisksMapping(formats strfmt.Registry) error {

	if err := validate.Required("disksMapping", "body", m.DisksMapping); err != nil {
		return err
	}

	for i := 0; i < len(m.DisksMapping); i++ {
		if swag.IsZero(m.DisksMapping[i]) { // not required
			continue
		}

		if m.DisksMapping[i] != nil {
			if err := m.DisksMapping[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disksMapping" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VmwareFcdInstantRecoverySpec) validateObjectRestorePointID(formats strfmt.Registry) error {

	if err := validate.Required("objectRestorePointId", "body", m.ObjectRestorePointID); err != nil {
		return err
	}

	if err := validate.FormatOf("objectRestorePointId", "body", "uuid", m.ObjectRestorePointID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VmwareFcdInstantRecoverySpec) validateWriteCache(formats strfmt.Registry) error {
	if swag.IsZero(m.WriteCache) { // not required
		return nil
	}

	if m.WriteCache != nil {
		if err := m.WriteCache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writeCache")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vmware fcd instant recovery spec based on the context it is used
func (m *VmwareFcdInstantRecoverySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisksMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWriteCache(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareFcdInstantRecoverySpec) contextValidateDestinationCluster(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *VmwareFcdInstantRecoverySpec) contextValidateDisksMapping(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisksMapping); i++ {

		if m.DisksMapping[i] != nil {
			if err := m.DisksMapping[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disksMapping" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VmwareFcdInstantRecoverySpec) contextValidateWriteCache(ctx context.Context, formats strfmt.Registry) error {

	if m.WriteCache != nil {
		if err := m.WriteCache.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writeCache")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareFcdInstantRecoverySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareFcdInstantRecoverySpec) UnmarshalBinary(b []byte) error {
	var res VmwareFcdInstantRecoverySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
