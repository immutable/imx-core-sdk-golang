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

// CollectionDetails collection details
//
// swagger:model CollectionDetails
type CollectionDetails struct {

	// URL of the icon of the collection
	// Required: true
	IconURL *string `json:"icon_url"`

	// Name of the collection
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this collection details
func (m *CollectionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIconURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectionDetails) validateIconURL(formats strfmt.Registry) error {

	if err := validate.Required("icon_url", "body", m.IconURL); err != nil {
		return err
	}

	return nil
}

func (m *CollectionDetails) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this collection details based on context it is used
func (m *CollectionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectionDetails) UnmarshalBinary(b []byte) error {
	var res CollectionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
