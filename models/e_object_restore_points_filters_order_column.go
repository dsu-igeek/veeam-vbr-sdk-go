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

// EObjectRestorePointsFiltersOrderColumn e object restore points filters order column
//
// swagger:model EObjectRestorePointsFiltersOrderColumn
type EObjectRestorePointsFiltersOrderColumn string

func NewEObjectRestorePointsFiltersOrderColumn(value EObjectRestorePointsFiltersOrderColumn) *EObjectRestorePointsFiltersOrderColumn {
	v := value
	return &v
}

const (

	// EObjectRestorePointsFiltersOrderColumnCreationTime captures enum value "CreationTime"
	EObjectRestorePointsFiltersOrderColumnCreationTime EObjectRestorePointsFiltersOrderColumn = "CreationTime"

	// EObjectRestorePointsFiltersOrderColumnPlatformID captures enum value "PlatformId"
	EObjectRestorePointsFiltersOrderColumnPlatformID EObjectRestorePointsFiltersOrderColumn = "PlatformId"

	// EObjectRestorePointsFiltersOrderColumnBackupID captures enum value "BackupId"
	EObjectRestorePointsFiltersOrderColumnBackupID EObjectRestorePointsFiltersOrderColumn = "BackupId"
)

// for schema
var eObjectRestorePointsFiltersOrderColumnEnum []interface{}

func init() {
	var res []EObjectRestorePointsFiltersOrderColumn
	if err := json.Unmarshal([]byte(`["CreationTime","PlatformId","BackupId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eObjectRestorePointsFiltersOrderColumnEnum = append(eObjectRestorePointsFiltersOrderColumnEnum, v)
	}
}

func (m EObjectRestorePointsFiltersOrderColumn) validateEObjectRestorePointsFiltersOrderColumnEnum(path, location string, value EObjectRestorePointsFiltersOrderColumn) error {
	if err := validate.EnumCase(path, location, value, eObjectRestorePointsFiltersOrderColumnEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e object restore points filters order column
func (m EObjectRestorePointsFiltersOrderColumn) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEObjectRestorePointsFiltersOrderColumnEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e object restore points filters order column based on context it is used
func (m EObjectRestorePointsFiltersOrderColumn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
