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

// MintRequest mint request
//
// swagger:model MintRequest
type MintRequest struct {

	// Signature from authorised minter
	// Required: true
	AuthSignature *string `json:"auth_signature"`

	// minting contract
	// Required: true
	ContractAddress *string `json:"contract_address"`

	// Global contract-level royalty fees
	Royalties []*MintFee `json:"royalties"`

	// Users to mint to
	// Required: true
	Users []*MintUser `json:"users"`
}

// Validate validates this mint request
func (m *MintRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthSignature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoyalties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MintRequest) validateAuthSignature(formats strfmt.Registry) error {

	if err := validate.Required("auth_signature", "body", m.AuthSignature); err != nil {
		return err
	}

	return nil
}

func (m *MintRequest) validateContractAddress(formats strfmt.Registry) error {

	if err := validate.Required("contract_address", "body", m.ContractAddress); err != nil {
		return err
	}

	return nil
}

func (m *MintRequest) validateRoyalties(formats strfmt.Registry) error {
	if swag.IsZero(m.Royalties) { // not required
		return nil
	}

	for i := 0; i < len(m.Royalties); i++ {
		if swag.IsZero(m.Royalties[i]) { // not required
			continue
		}

		if m.Royalties[i] != nil {
			if err := m.Royalties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("royalties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("royalties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MintRequest) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("users", "body", m.Users); err != nil {
		return err
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this mint request based on the context it is used
func (m *MintRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoyalties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MintRequest) contextValidateRoyalties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Royalties); i++ {

		if m.Royalties[i] != nil {
			if err := m.Royalties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("royalties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("royalties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MintRequest) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {
			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MintRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MintRequest) UnmarshalBinary(b []byte) error {
	var res MintRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}