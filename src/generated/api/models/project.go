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

// Project project
//
// swagger:model Project
type Project struct {

	// The current period expiry date for collection limit
	// Required: true
	CollectionLimitExpiresAt *string `json:"collection_limit_expires_at"`

	// The total monthly collection limit
	// Required: true
	CollectionMonthlyLimit *int64 `json:"collection_monthly_limit"`

	// The number of collection remaining in the current period
	// Required: true
	CollectionRemaining *int64 `json:"collection_remaining"`

	// The company name
	// Required: true
	CompanyName *string `json:"company_name"`

	// The project contact email
	// Required: true
	ContactEmail *string `json:"contact_email"`

	// The project ID
	// Required: true
	ID *int64 `json:"id"`

	// The current period expiry date for mint operation limit
	// Required: true
	MintLimitExpiresAt *string `json:"mint_limit_expires_at"`

	// The total monthly mint operation limit
	// Required: true
	MintMonthlyLimit *int64 `json:"mint_monthly_limit"`

	// The number of mint operation remaining in the current period
	// Required: true
	MintRemaining *int64 `json:"mint_remaining"`

	// The project name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectionLimitExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectionMonthlyLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectionRemaining(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompanyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMintLimitExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMintMonthlyLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMintRemaining(formats); err != nil {
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

func (m *Project) validateCollectionLimitExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("collection_limit_expires_at", "body", m.CollectionLimitExpiresAt); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateCollectionMonthlyLimit(formats strfmt.Registry) error {

	if err := validate.Required("collection_monthly_limit", "body", m.CollectionMonthlyLimit); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateCollectionRemaining(formats strfmt.Registry) error {

	if err := validate.Required("collection_remaining", "body", m.CollectionRemaining); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateCompanyName(formats strfmt.Registry) error {

	if err := validate.Required("company_name", "body", m.CompanyName); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateContactEmail(formats strfmt.Registry) error {

	if err := validate.Required("contact_email", "body", m.ContactEmail); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateMintLimitExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("mint_limit_expires_at", "body", m.MintLimitExpiresAt); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateMintMonthlyLimit(formats strfmt.Registry) error {

	if err := validate.Required("mint_monthly_limit", "body", m.MintMonthlyLimit); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateMintRemaining(formats strfmt.Registry) error {

	if err := validate.Required("mint_remaining", "body", m.MintRemaining); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project based on context it is used
func (m *Project) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}