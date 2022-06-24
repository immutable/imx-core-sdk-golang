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

// GetSignableDepositResponse get signable deposit response
//
// swagger:model GetSignableDepositResponse
type GetSignableDepositResponse struct {

	// Amount this user is depositing
	// Required: true
	Amount *string `json:"amount"`

	// ID of the asset this user is depositing
	// Required: true
	AssetID *string `json:"asset_id"`

	// Nonce of the deposit
	// Required: true
	Nonce *int64 `json:"nonce"`

	// Public stark key of the depositing user
	// Required: true
	StarkKey *string `json:"stark_key"`

	// ID of the vault this user is depositing to
	// Required: true
	VaultID *int64 `json:"vault_id"`
}

// Validate validates this get signable deposit response
func (m *GetSignableDepositResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarkKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSignableDepositResponse) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableDepositResponse) validateAssetID(formats strfmt.Registry) error {

	if err := validate.Required("asset_id", "body", m.AssetID); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableDepositResponse) validateNonce(formats strfmt.Registry) error {

	if err := validate.Required("nonce", "body", m.Nonce); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableDepositResponse) validateStarkKey(formats strfmt.Registry) error {

	if err := validate.Required("stark_key", "body", m.StarkKey); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableDepositResponse) validateVaultID(formats strfmt.Registry) error {

	if err := validate.Required("vault_id", "body", m.VaultID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get signable deposit response based on context it is used
func (m *GetSignableDepositResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSignableDepositResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSignableDepositResponse) UnmarshalBinary(b []byte) error {
	var res GetSignableDepositResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}