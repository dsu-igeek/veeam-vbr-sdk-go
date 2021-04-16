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

// ScheduleMonthlyModel Monthly scheduling options.
//
// swagger:model ScheduleMonthlyModel
type ScheduleMonthlyModel struct {

	// day number in month
	DayNumberInMonth EDayNumberInMonth `json:"dayNumberInMonth,omitempty"`

	// Day of the month when the job must start.
	DayOfMonth int64 `json:"dayOfMonth,omitempty"`

	// day of week
	DayOfWeek EDayOfWeek `json:"dayOfWeek,omitempty"`

	// If *true*, monthly schedule is enabled.
	// Required: true
	IsEnabled bool `json:"isEnabled"`

	// Local time when the job must start.
	LocalTime string `json:"localTime,omitempty"`

	// Months when the job must start.
	// Unique: true
	Months []EMonth `json:"months"`
}

// Validate validates this schedule monthly model
func (m *ScheduleMonthlyModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDayNumberInMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDayOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonths(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleMonthlyModel) validateDayNumberInMonth(formats strfmt.Registry) error {
	if swag.IsZero(m.DayNumberInMonth) { // not required
		return nil
	}

	if err := m.DayNumberInMonth.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dayNumberInMonth")
		}
		return err
	}

	return nil
}

func (m *ScheduleMonthlyModel) validateDayOfWeek(formats strfmt.Registry) error {
	if swag.IsZero(m.DayOfWeek) { // not required
		return nil
	}

	if err := m.DayOfWeek.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dayOfWeek")
		}
		return err
	}

	return nil
}

func (m *ScheduleMonthlyModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", bool(m.IsEnabled)); err != nil {
		return err
	}

	return nil
}

func (m *ScheduleMonthlyModel) validateMonths(formats strfmt.Registry) error {
	if swag.IsZero(m.Months) { // not required
		return nil
	}

	if err := validate.UniqueItems("months", "body", m.Months); err != nil {
		return err
	}

	for i := 0; i < len(m.Months); i++ {

		if err := m.Months[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("months" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this schedule monthly model based on the context it is used
func (m *ScheduleMonthlyModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDayNumberInMonth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDayOfWeek(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleMonthlyModel) contextValidateDayNumberInMonth(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DayNumberInMonth.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dayNumberInMonth")
		}
		return err
	}

	return nil
}

func (m *ScheduleMonthlyModel) contextValidateDayOfWeek(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DayOfWeek.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dayOfWeek")
		}
		return err
	}

	return nil
}

func (m *ScheduleMonthlyModel) contextValidateMonths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Months); i++ {

		if err := m.Months[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("months" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleMonthlyModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleMonthlyModel) UnmarshalBinary(b []byte) error {
	var res ScheduleMonthlyModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
