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

// Fee fee
//
// swagger:model Fee
type Fee struct {

	// Wallet address
	// Required: true
	Address *string `json:"address"`

	// The percentage of fee
	// Required: true
	// Maximum: 100
	Percentage *float64 `json:"percentage"`

	// Type of fee. Examples: `royalty`, `maker`, `taker` or `protocol`
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this fee
func (m *Fee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Fee) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *Fee) validatePercentage(formats strfmt.Registry) error {

	if err := validate.Required("percentage", "body", m.Percentage); err != nil {
		return err
	}

	if err := validate.Maximum("percentage", "body", *m.Percentage, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Fee) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fee based on context it is used
func (m *Fee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Fee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Fee) UnmarshalBinary(b []byte) error {
	var res Fee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
