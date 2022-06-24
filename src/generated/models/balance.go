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

// Balance balance
//
// swagger:model Balance
type Balance struct {

	// Amount which is currently inside the exchange
	// Required: true
	Balance *string `json:"balance"`

	// Amount which is currently preparing withdrawal from the exchange
	// Required: true
	PreparingWithdrawal *string `json:"preparing_withdrawal"`

	// Symbol of the token (e.g. ETH, IMX)
	// Required: true
	Symbol *string `json:"symbol"`

	// Amount which is currently withdrawable from the exchange
	// Required: true
	Withdrawable *string `json:"withdrawable"`
}

// Validate validates this balance
func (m *Balance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreparingWithdrawal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymbol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithdrawable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Balance) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validatePreparingWithdrawal(formats strfmt.Registry) error {

	if err := validate.Required("preparing_withdrawal", "body", m.PreparingWithdrawal); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateSymbol(formats strfmt.Registry) error {

	if err := validate.Required("symbol", "body", m.Symbol); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateWithdrawable(formats strfmt.Registry) error {

	if err := validate.Required("withdrawable", "body", m.Withdrawable); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this balance based on context it is used
func (m *Balance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Balance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Balance) UnmarshalBinary(b []byte) error {
	var res Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}