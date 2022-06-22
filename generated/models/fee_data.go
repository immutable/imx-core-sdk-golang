// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FeeData fee data
//
// swagger:model FeeData
type FeeData struct {

	// Address of ERC721/ERC20 contract
	ContractAddress string `json:"contract_address,omitempty"`

	// Number of decimals supported by this asset
	Decimals int64 `json:"decimals,omitempty"`
}

// Validate validates this fee data
func (m *FeeData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fee data based on context it is used
func (m *FeeData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeeData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeeData) UnmarshalBinary(b []byte) error {
	var res FeeData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
