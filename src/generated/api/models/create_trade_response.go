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

// CreateTradeResponse create trade response
//
// swagger:model CreateTradeResponse
type CreateTradeResponse struct {

	// Request ID that returns when a trade initiated through risk-manager
	RequestID string `json:"request_id,omitempty"`

	// Current status of trade
	// Required: true
	Status *string `json:"status"`

	// ID of trade within Immutable X
	// Required: true
	TradeID *int64 `json:"trade_id"`
}

// Validate validates this create trade response
func (m *CreateTradeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTradeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTradeResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CreateTradeResponse) validateTradeID(formats strfmt.Registry) error {

	if err := validate.Required("trade_id", "body", m.TradeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create trade response based on context it is used
func (m *CreateTradeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateTradeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTradeResponse) UnmarshalBinary(b []byte) error {
	var res CreateTradeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}