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

// AddMetadataSchemaToCollectionRequest add metadata schema to collection request
//
// swagger:model AddMetadataSchemaToCollectionRequest
type AddMetadataSchemaToCollectionRequest struct {

	// Not required from API user
	ContractAddress string `json:"contract_address,omitempty"`

	// The metadata container
	// Required: true
	// Min Items: 1
	Metadata []*MetadataSchemaRequest `json:"metadata"`
}

// Validate validates this add metadata schema to collection request
func (m *AddMetadataSchemaToCollectionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddMetadataSchemaToCollectionRequest) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	iMetadataSize := int64(len(m.Metadata))

	if err := validate.MinItems("metadata", "body", iMetadataSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add metadata schema to collection request based on the context it is used
func (m *AddMetadataSchemaToCollectionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddMetadataSchemaToCollectionRequest) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metadata); i++ {

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddMetadataSchemaToCollectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddMetadataSchemaToCollectionRequest) UnmarshalBinary(b []byte) error {
	var res AddMetadataSchemaToCollectionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}