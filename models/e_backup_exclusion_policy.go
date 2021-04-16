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

// EBackupExclusionPolicy Exclusion policy.
//
// swagger:model EBackupExclusionPolicy
type EBackupExclusionPolicy string

func NewEBackupExclusionPolicy(value EBackupExclusionPolicy) *EBackupExclusionPolicy {
	v := value
	return &v
}

const (

	// EBackupExclusionPolicyDisabled captures enum value "disabled"
	EBackupExclusionPolicyDisabled EBackupExclusionPolicy = "disabled"

	// EBackupExclusionPolicyExcludeOnly captures enum value "excludeOnly"
	EBackupExclusionPolicyExcludeOnly EBackupExclusionPolicy = "excludeOnly"

	// EBackupExclusionPolicyIncludeOnly captures enum value "includeOnly"
	EBackupExclusionPolicyIncludeOnly EBackupExclusionPolicy = "includeOnly"
)

// for schema
var eBackupExclusionPolicyEnum []interface{}

func init() {
	var res []EBackupExclusionPolicy
	if err := json.Unmarshal([]byte(`["disabled","excludeOnly","includeOnly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eBackupExclusionPolicyEnum = append(eBackupExclusionPolicyEnum, v)
	}
}

func (m EBackupExclusionPolicy) validateEBackupExclusionPolicyEnum(path, location string, value EBackupExclusionPolicy) error {
	if err := validate.EnumCase(path, location, value, eBackupExclusionPolicyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e backup exclusion policy
func (m EBackupExclusionPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEBackupExclusionPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e backup exclusion policy based on context it is used
func (m EBackupExclusionPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
