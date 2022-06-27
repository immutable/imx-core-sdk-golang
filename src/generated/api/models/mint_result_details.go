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

// MintResultDetails mint result details
//
// swagger:model MintResultDetails
type MintResultDetails struct {

	// Contract address of this token
	// Required: true
	ContractAddress *string `json:"contract_address"`

	// IMX ID of this token
	// Required: true
	TokenID *string `json:"token_id"`

	// Mint Transaction ID
	// Required: true
	TxID *int64 `json:"tx_id"`
}

// Validate validates this mint result details
func (m *MintResultDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContractAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MintResultDetails) validateContractAddress(formats strfmt.Registry) error {

	if err := validate.Required("contract_address", "body", m.ContractAddress); err != nil {
		return err
	}

	return nil
}

func (m *MintResultDetails) validateTokenID(formats strfmt.Registry) error {

	if err := validate.Required("token_id", "body", m.TokenID); err != nil {
		return err
	}

	return nil
}

func (m *MintResultDetails) validateTxID(formats strfmt.Registry) error {

	if err := validate.Required("tx_id", "body", m.TxID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mint result details based on context it is used
func (m *MintResultDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MintResultDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MintResultDetails) UnmarshalBinary(b []byte) error {
	var res MintResultDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
