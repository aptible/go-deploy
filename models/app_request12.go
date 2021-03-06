// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppRequest12 app request 12
// swagger:model app_request_12
type AppRequest12 struct {

	// Mapping of environment variables
	// Required: true
	Env interface{} `json:"env"`
}

// Validate validates this app request 12
func (m *AppRequest12) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRequest12) validateEnv(formats strfmt.Registry) error {

	if err := validate.Required("env", "body", m.Env); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest12) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest12) UnmarshalBinary(b []byte) error {
	var res AppRequest12
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
