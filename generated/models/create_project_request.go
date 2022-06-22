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

// CreateProjectRequest create project request
//
// swagger:model CreateProjectRequest
type CreateProjectRequest struct {

	// The company name
	// Required: true
	CompanyName *string `json:"company_name"`

	// The project contact email
	// Required: true
	ContactEmail *string `json:"contact_email"`

	// The project name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this create project request
func (m *CreateProjectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompanyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactEmail(formats); err != nil {
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

func (m *CreateProjectRequest) validateCompanyName(formats strfmt.Registry) error {

	if err := validate.Required("company_name", "body", m.CompanyName); err != nil {
		return err
	}

	return nil
}

func (m *CreateProjectRequest) validateContactEmail(formats strfmt.Registry) error {

	if err := validate.Required("contact_email", "body", m.ContactEmail); err != nil {
		return err
	}

	return nil
}

func (m *CreateProjectRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create project request based on context it is used
func (m *CreateProjectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateProjectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateProjectRequest) UnmarshalBinary(b []byte) error {
	var res CreateProjectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
