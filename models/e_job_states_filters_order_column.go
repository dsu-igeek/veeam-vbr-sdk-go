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

// EJobStatesFiltersOrderColumn Orders job states by the specified column.
//
// swagger:model EJobStatesFiltersOrderColumn
type EJobStatesFiltersOrderColumn string

func NewEJobStatesFiltersOrderColumn(value EJobStatesFiltersOrderColumn) *EJobStatesFiltersOrderColumn {
	v := value
	return &v
}

const (

	// EJobStatesFiltersOrderColumnName captures enum value "Name"
	EJobStatesFiltersOrderColumnName EJobStatesFiltersOrderColumn = "Name"

	// EJobStatesFiltersOrderColumnType captures enum value "Type"
	EJobStatesFiltersOrderColumnType EJobStatesFiltersOrderColumn = "Type"

	// EJobStatesFiltersOrderColumnStatus captures enum value "Status"
	EJobStatesFiltersOrderColumnStatus EJobStatesFiltersOrderColumn = "Status"

	// EJobStatesFiltersOrderColumnLastRun captures enum value "LastRun"
	EJobStatesFiltersOrderColumnLastRun EJobStatesFiltersOrderColumn = "LastRun"

	// EJobStatesFiltersOrderColumnLastResult captures enum value "LastResult"
	EJobStatesFiltersOrderColumnLastResult EJobStatesFiltersOrderColumn = "LastResult"

	// EJobStatesFiltersOrderColumnNextRun captures enum value "NextRun"
	EJobStatesFiltersOrderColumnNextRun EJobStatesFiltersOrderColumn = "NextRun"

	// EJobStatesFiltersOrderColumnDescription captures enum value "Description"
	EJobStatesFiltersOrderColumnDescription EJobStatesFiltersOrderColumn = "Description"

	// EJobStatesFiltersOrderColumnRepositoryID captures enum value "RepositoryId"
	EJobStatesFiltersOrderColumnRepositoryID EJobStatesFiltersOrderColumn = "RepositoryId"

	// EJobStatesFiltersOrderColumnObjectsCount captures enum value "ObjectsCount"
	EJobStatesFiltersOrderColumnObjectsCount EJobStatesFiltersOrderColumn = "ObjectsCount"
)

// for schema
var eJobStatesFiltersOrderColumnEnum []interface{}

func init() {
	var res []EJobStatesFiltersOrderColumn
	if err := json.Unmarshal([]byte(`["Name","Type","Status","LastRun","LastResult","NextRun","Description","RepositoryId","ObjectsCount"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eJobStatesFiltersOrderColumnEnum = append(eJobStatesFiltersOrderColumnEnum, v)
	}
}

func (m EJobStatesFiltersOrderColumn) validateEJobStatesFiltersOrderColumnEnum(path, location string, value EJobStatesFiltersOrderColumn) error {
	if err := validate.EnumCase(path, location, value, eJobStatesFiltersOrderColumnEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e job states filters order column
func (m EJobStatesFiltersOrderColumn) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEJobStatesFiltersOrderColumnEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e job states filters order column based on context it is used
func (m EJobStatesFiltersOrderColumn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
