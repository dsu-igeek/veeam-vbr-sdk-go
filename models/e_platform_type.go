// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EPlatformType e platform type
//
// swagger:model EPlatformType
type EPlatformType string

func NewEPlatformType(value EPlatformType) *EPlatformType {
	v := value
	return &v
}

const (

	// EPlatformTypeVMWare captures enum value "VmWare"
	EPlatformTypeVMWare EPlatformType = "VmWare"

	// EPlatformTypeHyperV captures enum value "HyperV"
	EPlatformTypeHyperV EPlatformType = "HyperV"

	// EPlatformTypeVcd captures enum value "Vcd"
	EPlatformTypeVcd EPlatformType = "Vcd"

	// EPlatformTypeWindowsPhysical captures enum value "WindowsPhysical"
	EPlatformTypeWindowsPhysical EPlatformType = "WindowsPhysical"

	// EPlatformTypeLinuxPhysical captures enum value "LinuxPhysical"
	EPlatformTypeLinuxPhysical EPlatformType = "LinuxPhysical"

	// EPlatformTypeTape captures enum value "Tape"
	EPlatformTypeTape EPlatformType = "Tape"

	// EPlatformTypeNasBackup captures enum value "NasBackup"
	EPlatformTypeNasBackup EPlatformType = "NasBackup"

	// EPlatformTypeCustomPlatform captures enum value "CustomPlatform"
	EPlatformTypeCustomPlatform EPlatformType = "CustomPlatform"
)

// for schema
var ePlatformTypeEnum []interface{}

func init() {
	var res []EPlatformType
	if err := json.Unmarshal([]byte(`["VmWare","HyperV","Vcd","WindowsPhysical","LinuxPhysical","Tape","NasBackup","CustomPlatform"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ePlatformTypeEnum = append(ePlatformTypeEnum, v)
	}
}

func (m EPlatformType) validateEPlatformTypeEnum(path, location string, value EPlatformType) error {
	if err := validate.EnumCase(path, location, value, ePlatformTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e platform type
func (m EPlatformType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEPlatformTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e platform type based on context it is used
func (m EPlatformType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
