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

// AssetProperties asset properties
//
// swagger:model AssetProperties
type AssetProperties struct {

	// Details of this asset's collection
	Collection *CollectionDetails `json:"collection,omitempty"`

	// Image URL of this asset
	ImageURL string `json:"image_url,omitempty"`

	// Name of this asset
	Name string `json:"name,omitempty"`
}

// Validate validates this asset properties
func (m *AssetProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetProperties) validateCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.Collection) { // not required
		return nil
	}

	if m.Collection != nil {
		if err := m.Collection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this asset properties based on the context it is used
func (m *AssetProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetProperties) contextValidateCollection(ctx context.Context, formats strfmt.Registry) error {

	if m.Collection != nil {
		if err := m.Collection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetProperties) UnmarshalBinary(b []byte) error {
	var res AssetProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}