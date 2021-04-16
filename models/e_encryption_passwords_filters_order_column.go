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

// EEncryptionPasswordsFiltersOrderColumn e encryption passwords filters order column
//
// swagger:model EEncryptionPasswordsFiltersOrderColumn
type EEncryptionPasswordsFiltersOrderColumn string

func NewEEncryptionPasswordsFiltersOrderColumn(value EEncryptionPasswordsFiltersOrderColumn) *EEncryptionPasswordsFiltersOrderColumn {
	v := value
	return &v
}

const (

	// EEncryptionPasswordsFiltersOrderColumnHint captures enum value "Hint"
	EEncryptionPasswordsFiltersOrderColumnHint EEncryptionPasswordsFiltersOrderColumn = "Hint"

	// EEncryptionPasswordsFiltersOrderColumnModificationTime captures enum value "ModificationTime"
	EEncryptionPasswordsFiltersOrderColumnModificationTime EEncryptionPasswordsFiltersOrderColumn = "ModificationTime"
)

// for schema
var eEncryptionPasswordsFiltersOrderColumnEnum []interface{}

func init() {
	var res []EEncryptionPasswordsFiltersOrderColumn
	if err := json.Unmarshal([]byte(`["Hint","ModificationTime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eEncryptionPasswordsFiltersOrderColumnEnum = append(eEncryptionPasswordsFiltersOrderColumnEnum, v)
	}
}

func (m EEncryptionPasswordsFiltersOrderColumn) validateEEncryptionPasswordsFiltersOrderColumnEnum(path, location string, value EEncryptionPasswordsFiltersOrderColumn) error {
	if err := validate.EnumCase(path, location, value, eEncryptionPasswordsFiltersOrderColumnEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e encryption passwords filters order column
func (m EEncryptionPasswordsFiltersOrderColumn) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEEncryptionPasswordsFiltersOrderColumnEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e encryption passwords filters order column based on context it is used
func (m EEncryptionPasswordsFiltersOrderColumn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
