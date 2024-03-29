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

// AppRequest34 app request 34
// swagger:model app_request_34
type AppRequest34 struct {

	// acme
	Acme bool `json:"acme"`

	// certificate
	// Required: true
	Certificate *int64 `json:"certificate"`

	// container exposed ports
	ContainerExposedPorts []int64 `json:"container_exposed_ports"`

	// container port
	ContainerPort int64 `json:"container_port,omitempty"`

	// container ports
	ContainerPorts []int64 `json:"container_ports"`

	// default
	Default bool `json:"default"`

	// internal
	Internal bool `json:"internal"`

	// ip whitelist
	IPWhitelist []string `json:"ip_whitelist"`

	// platform
	Platform string `json:"platform,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`

	// user domain
	UserDomain string `json:"user_domain,omitempty"`
}

// Validate validates this app request 34
func (m *AppRequest34) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRequest34) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest34) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest34) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest34) UnmarshalBinary(b []byte) error {
	var res AppRequest34
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
