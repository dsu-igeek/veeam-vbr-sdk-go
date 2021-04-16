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

// EMonth e month
//
// swagger:model EMonth
type EMonth string

func NewEMonth(value EMonth) *EMonth {
	v := value
	return &v
}

const (

	// EMonthJanuary captures enum value "January"
	EMonthJanuary EMonth = "January"

	// EMonthFebruary captures enum value "February"
	EMonthFebruary EMonth = "February"

	// EMonthMarch captures enum value "March"
	EMonthMarch EMonth = "March"

	// EMonthApril captures enum value "April"
	EMonthApril EMonth = "April"

	// EMonthMay captures enum value "May"
	EMonthMay EMonth = "May"

	// EMonthJune captures enum value "June"
	EMonthJune EMonth = "June"

	// EMonthJuly captures enum value "July"
	EMonthJuly EMonth = "July"

	// EMonthAugust captures enum value "August"
	EMonthAugust EMonth = "August"

	// EMonthSeptember captures enum value "September"
	EMonthSeptember EMonth = "September"

	// EMonthOctober captures enum value "October"
	EMonthOctober EMonth = "October"

	// EMonthNovember captures enum value "November"
	EMonthNovember EMonth = "November"

	// EMonthDecember captures enum value "December"
	EMonthDecember EMonth = "December"
)

// for schema
var eMonthEnum []interface{}

func init() {
	var res []EMonth
	if err := json.Unmarshal([]byte(`["January","February","March","April","May","June","July","August","September","October","November","December"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eMonthEnum = append(eMonthEnum, v)
	}
}

func (m EMonth) validateEMonthEnum(path, location string, value EMonth) error {
	if err := validate.EnumCase(path, location, value, eMonthEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e month
func (m EMonth) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEMonthEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e month based on context it is used
func (m EMonth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
