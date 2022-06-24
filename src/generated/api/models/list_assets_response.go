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

// ListAssetsResponse list assets response
//
// swagger:model ListAssetsResponse
type ListAssetsResponse struct {

	// Generated cursor returned by previous query
	// Required: true
	Cursor *string `json:"cursor"`

	// Remaining results flag. 1: there are remaining results matching this query, 0: no remaining results
	// Required: true
	Remaining *int64 `json:"remaining"`

	// Assets matching query parameters
	// Required: true
	Result []*Asset `json:"result"`
}

// Validate validates this list assets response
func (m *ListAssetsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCursor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemaining(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListAssetsResponse) validateCursor(formats strfmt.Registry) error {

	if err := validate.Required("cursor", "body", m.Cursor); err != nil {
		return err
	}

	return nil
}

func (m *ListAssetsResponse) validateRemaining(formats strfmt.Registry) error {

	if err := validate.Required("remaining", "body", m.Remaining); err != nil {
		return err
	}

	return nil
}

func (m *ListAssetsResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for i := 0; i < len(m.Result); i++ {
		if swag.IsZero(m.Result[i]) { // not required
			continue
		}

		if m.Result[i] != nil {
			if err := m.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list assets response based on the context it is used
func (m *ListAssetsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListAssetsResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Result); i++ {

		if m.Result[i] != nil {
			if err := m.Result[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListAssetsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAssetsResponse) UnmarshalBinary(b []byte) error {
	var res ListAssetsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}