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

// GetSignableCancelOrderResponse get signable cancel order response
//
// swagger:model GetSignableCancelOrderResponse
type GetSignableCancelOrderResponse struct {

	// ID of the order to be cancelled
	// Required: true
	OrderID *int64 `json:"order_id"`

	// Hash of the payload to be signed for cancel order
	// Required: true
	PayloadHash *string `json:"payload_hash"`

	// Message to sign from wallet to confirm cancel order
	// Required: true
	SignableMessage *string `json:"signable_message"`
}

// Validate validates this get signable cancel order response
func (m *GetSignableCancelOrderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

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

func (m *GetSignableCancelOrderResponse) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableCancelOrderResponse) validatePayloadHash(formats strfmt.Registry) error {

	if err := validate.Required("payload_hash", "body", m.PayloadHash); err != nil {
		return err
	}

	return nil
}

func (m *GetSignableCancelOrderResponse) validateSignableMessage(formats strfmt.Registry) error {

	if err := validate.Required("signable_message", "body", m.SignableMessage); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get signable cancel order response based on context it is used
func (m *GetSignableCancelOrderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSignableCancelOrderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSignableCancelOrderResponse) UnmarshalBinary(b []byte) error {
	var res GetSignableCancelOrderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}