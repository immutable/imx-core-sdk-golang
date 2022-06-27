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

// GetSignableTransferResponse get signable transfer response
//
// swagger:model GetSignableTransferResponse
type GetSignableTransferResponse struct {

	// Sender of the transfer
	// Required: true
	SenderStarkKey *string `json:"sender_stark_key"`

	// Message to sign with L1 wallet to confirm transfer request
	// Required: true
	SignableMessage *string `json:"signable_message"`

	// List of transfer responses without the sender stark key
	// Required: true
	SignableResponses []*SignableTransferResponseDetails `json:"signable_responses"`
}

// Validate validates this get signable transfer response
func (m *GetSignableTransferResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSenderStarkKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignableMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignableResponses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSignableTransferResponse) validateSenderStarkKey(formats strfmt.Registry) error {

	if err := validate.Required("sender_stark_key", "body", m.SenderStarkKey); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableTransferResponse) validateSignableMessage(formats strfmt.Registry) error {

	if err := validate.Required("signable_message", "body", m.SignableMessage); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableTransferResponse) validateSignableResponses(formats strfmt.Registry) error {

	if err := validate.Required("signable_responses", "body", m.SignableResponses); err != nil {
		return err
	}

	for i := 0; i < len(m.SignableResponses); i++ {
		if swag.IsZero(m.SignableResponses[i]) { // not required
			continue
		}

		if m.SignableResponses[i] != nil {
			if err := m.SignableResponses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signable_responses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signable_responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get signable transfer response based on the context it is used
func (m *GetSignableTransferResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSignableResponses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSignableTransferResponse) contextValidateSignableResponses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SignableResponses); i++ {

		if m.SignableResponses[i] != nil {
			if err := m.SignableResponses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signable_responses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signable_responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSignableTransferResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSignableTransferResponse) UnmarshalBinary(b []byte) error {
	var res GetSignableTransferResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
