// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokenData token data
//
// swagger:model TokenData
type TokenData struct {

	// Number of decimals supported by this asset
	Decimals int64 `json:"decimals,omitempty"`

	// [DEPRECATED] Internal Immutable X Token ID
	ID string `json:"id,omitempty"`

	// Properties of this asset
	Properties *AssetProperties `json:"properties,omitempty"`

	// Quantity of this asset
	Quantity string `json:"quantity,omitempty"`

	// Quantity of this asset with the sum of all fees applied to the asset
	QuantityWithFees string `json:"quantity_with_fees,omitempty"`

	// Address of ERC721/ERC20 contract
	TokenAddress string `json:"token_address,omitempty"`

	// ERC721 Token ID
	TokenID string `json:"token_id,omitempty"`
}

// Validate validates this token data
func (m *TokenData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenData) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this token data based on the context it is used
func (m *TokenData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenData) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	if m.Properties != nil {
		if err := m.Properties.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenData) UnmarshalBinary(b []byte) error {
	var res TokenData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}