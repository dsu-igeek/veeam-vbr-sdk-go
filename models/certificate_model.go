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

// CertificateModel certificate model
//
// swagger:model CertificateModel
type CertificateModel struct {

	// Issuer of the certificate.
	// Required: true
	IssuedBy *string `json:"issuedBy"`

	// Acquirer of the certificate.
	// Required: true
	IssuedTo *string `json:"issuedTo"`

	// Key algorithm of the certificate.
	// Required: true
	KeyAlgorithm *string `json:"keyAlgorithm"`

	// Key size of the certificate.
	// Required: true
	KeySize *string `json:"keySize"`

	// Serial number of the certificate.
	// Required: true
	SerialNumber *string `json:"serialNumber"`

	// Subject of the certificate.
	// Required: true
	Subject *string `json:"subject"`

	// Thumbprint of the certificate.
	// Required: true
	Thumbprint *string `json:"thumbprint"`

	// Expiration date and time of the certificate.
	// Required: true
	// Format: date-time
	ValidBy *strfmt.DateTime `json:"validBy"`

	// Date and time the certificate is valid from.
	// Required: true
	// Format: date-time
	ValidFrom *strfmt.DateTime `json:"validFrom"`
}

// Validate validates this certificate model
func (m *CertificateModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssuedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeySize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateModel) validateIssuedBy(formats strfmt.Registry) error {

	if err := validate.Required("issuedBy", "body", m.IssuedBy); err != nil {
		return err
	}

	return nil
}

func (m *CertificateModel) validateIssuedTo(formats strfmt.Registry) error {

	if err := validate.Required("issuedTo", "body", m.IssuedTo); err != nil {
		return err
	}

	return nil
}

func (m *CertificateModel) validateKeyAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("keyAlgorithm", "body", m.KeyAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *CertificateModel) validateKeySize(formats strfmt.Registry) error {

	if err := validate.Required("keySize", "body", m.KeySize); err != nil {
		return err
	}

	return nil
}

func (m *CertificateModel) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.Required("serialNumber", "body", m.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (m *CertificateModel) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

func (m *CertificateModel) validateThumbprint(formats strfmt.Registry) error {

	if err := validate.Required("thumbprint", "body", m.Thumbprint); err != nil {
		return err
	}

	return nil
}

func (m *CertificateModel) validateValidBy(formats strfmt.Registry) error {

	if err := validate.Required("validBy", "body", m.ValidBy); err != nil {
		return err
	}

	if err := validate.FormatOf("validBy", "body", "date-time", m.ValidBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificateModel) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", m.ValidFrom); err != nil {
		return err
	}

	if err := validate.FormatOf("validFrom", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate model based on context it is used
func (m *CertificateModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateModel) UnmarshalBinary(b []byte) error {
	var res CertificateModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}