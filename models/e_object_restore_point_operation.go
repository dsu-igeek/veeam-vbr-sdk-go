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

// EObjectRestorePointOperation e object restore point operation
//
// swagger:model EObjectRestorePointOperation
type EObjectRestorePointOperation string

func NewEObjectRestorePointOperation(value EObjectRestorePointOperation) *EObjectRestorePointOperation {
	v := value
	return &v
}

const (

	// EObjectRestorePointOperationVmwareInstantRecoveryFcd captures enum value "VmwareInstantRecoveryFcd"
	EObjectRestorePointOperationVmwareInstantRecoveryFcd EObjectRestorePointOperation = "VmwareInstantRecoveryFcd"
)

// for schema
var eObjectRestorePointOperationEnum []interface{}

func init() {
	var res []EObjectRestorePointOperation
	if err := json.Unmarshal([]byte(`["VmwareInstantRecoveryFcd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eObjectRestorePointOperationEnum = append(eObjectRestorePointOperationEnum, v)
	}
}

func (m EObjectRestorePointOperation) validateEObjectRestorePointOperationEnum(path, location string, value EObjectRestorePointOperation) error {
	if err := validate.EnumCase(path, location, value, eObjectRestorePointOperationEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e object restore point operation
func (m EObjectRestorePointOperation) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEObjectRestorePointOperationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e object restore point operation based on context it is used
func (m EObjectRestorePointOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
