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

// GetSignableRegistrationOffchainResponse get signable registration offchain response
//
// swagger:model GetSignableRegistrationOffchainResponse
type GetSignableRegistrationOffchainResponse struct {

	// Hash of the payload to be signed for user registration offchain
	// Required: true
	PayloadHash *string `json:"payload_hash"`

	// Message to sign with L1 wallet to register user offchain
	// Required: true
	SignableMessage *string `json:"signable_message"`
}

// Validate validates this get signable registration offchain response
func (m *GetSignableRegistrationOffchainResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayloadHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignableMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSignableRegistrationOffchainResponse) validatePayloadHash(formats strfmt.Registry) error {

	if err := validate.Required("payload_hash", "body", m.PayloadHash); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableRegistrationOffchainResponse) validateSignableMessage(formats strfmt.Registry) error {

	if err := validate.Required("signable_message", "body", m.SignableMessage); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get signable registration offchain response based on context it is used
func (m *GetSignableRegistrationOffchainResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSignableRegistrationOffchainResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSignableRegistrationOffchainResponse) UnmarshalBinary(b []byte) error {
	var res GetSignableRegistrationOffchainResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}