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

// JobStateModel job state model
//
// swagger:model JobStateModel
type JobStateModel struct {

	// Description of the job.
	// Required: true
	Description *string `json:"description"`

	// ID of the job.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// last result
	// Required: true
	LastResult *ESessionResult `json:"lastResult"`

	// Last run of the job.
	// Format: date-time
	LastRun strfmt.DateTime `json:"lastRun,omitempty"`

	// Name of the job.
	// Required: true
	Name *string `json:"name"`

	// Next run of the job.
	// Format: date-time
	NextRun strfmt.DateTime `json:"nextRun,omitempty"`

	// Number of objects processed by the job.
	// Required: true
	ObjectsCount *int64 `json:"objectsCount"`

	// ID of the backup repository.
	// Format: uuid
	RepositoryID strfmt.UUID `json:"repositoryId,omitempty"`

	// Name of the backup repository.
	RepositoryName string `json:"repositoryName,omitempty"`

	// ID of the last job session.
	// Format: uuid
	SessionID strfmt.UUID `json:"sessionId,omitempty"`

	// status
	// Required: true
	Status *EJobStatus `json:"status"`

	// type
	// Required: true
	Type *EJobType `json:"type"`

	// workload
	// Required: true
	Workload *EJobWorkload `json:"workload"`
}

// Validate validates this job state model
func (m *JobStateModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectsCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobStateModel) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *JobStateModel) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobStateModel) validateLastResult(formats strfmt.Registry) error {

	if err := validate.Required("lastResult", "body", m.LastResult); err != nil {
		return err
	}

	if err := validate.Required("lastResult", "body", m.LastResult); err != nil {
		return err
	}

	if m.LastResult != nil {
		if err := m.LastResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastResult")
			}
			return err
		}
	}

	return nil
}

func (m *JobStateModel) validateLastRun(formats strfmt.Registry) error {
	if swag.IsZero(m.LastRun) { // not required
		return nil
	}

	if err := validate.FormatOf("lastRun", "body", "date-time", m.LastRun.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobStateModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *JobStateModel) validateNextRun(formats strfmt.Registry) error {
	if swag.IsZero(m.NextRun) { // not required
		return nil
	}

	if err := validate.FormatOf("nextRun", "body", "date-time", m.NextRun.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobStateModel) validateObjectsCount(formats strfmt.Registry) error {

	if err := validate.Required("objectsCount", "body", m.ObjectsCount); err != nil {
		return err
	}

	return nil
}

func (m *JobStateModel) validateRepositoryID(formats strfmt.Registry) error {
	if swag.IsZero(m.RepositoryID) { // not required
		return nil
	}

	if err := validate.FormatOf("repositoryId", "body", "uuid", m.RepositoryID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobStateModel) validateSessionID(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionID) { // not required
		return nil
	}

	if err := validate.FormatOf("sessionId", "body", "uuid", m.SessionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobStateModel) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *JobStateModel) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *JobStateModel) validateWorkload(formats strfmt.Registry) error {

	if err := validate.Required("workload", "body", m.Workload); err != nil {
		return err
	}

	if err := validate.Required("workload", "body", m.Workload); err != nil {
		return err
	}

	if m.Workload != nil {
		if err := m.Workload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this job state model based on the context it is used
func (m *JobStateModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobStateModel) contextValidateLastResult(ctx context.Context, formats strfmt.Registry) error {

	if m.LastResult != nil {
		if err := m.LastResult.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastResult")
			}
			return err
		}
	}

	return nil
}

func (m *JobStateModel) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *JobStateModel) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *JobStateModel) contextValidateWorkload(ctx context.Context, formats strfmt.Registry) error {

	if m.Workload != nil {
		if err := m.Workload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobStateModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobStateModel) UnmarshalBinary(b []byte) error {
	var res JobStateModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}