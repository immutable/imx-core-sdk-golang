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

// EncodeAssetResponse encode asset response
//
// swagger:model EncodeAssetResponse
type EncodeAssetResponse struct {

	// Stark encoded asset id
	// Required: true
	AssetID *string `json:"asset_id"`

	// Stark encoded asset type
	// Required: true
	AssetType *string `json:"asset_type"`
}

// Validate validates this encode asset response
func (m *EncodeAssetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssetType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncodeAssetResponse) validateAssetID(formats strfmt.Registry) error {

	if err := validate.Required("asset_id", "body", m.AssetID); err != nil {
		return err
	}

	return nil
}

func (m *EncodeAssetResponse) validateAssetType(formats strfmt.Registry) error {

	if err := validate.Required("asset_type", "body", m.AssetType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this encode asset response based on context it is used
func (m *EncodeAssetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncodeAssetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncodeAssetResponse) UnmarshalBinary(b []byte) error {
	var res EncodeAssetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}