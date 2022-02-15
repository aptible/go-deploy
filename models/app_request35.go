// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppRequest35 app request 35
//
// swagger:model app_request_35
type AppRequest35 struct {

	// certificate
	Certificate int64 `json:"certificate,omitempty"`

	// container port
	ContainerPort int64 `json:"container_port,omitempty"`

	// container ports
	ContainerPorts []int64 `json:"container_ports"`

	// ip whitelist
	IPWhitelist []string `json:"ip_whitelist"`

	// platform
	Platform string `json:"platform,omitempty"`

	// user domain
	UserDomain string `json:"user_domain,omitempty"`
}

// Validate validates this app request 35
func (m *AppRequest35) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app request 35 based on context it is used
func (m *AppRequest35) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest35) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest35) UnmarshalBinary(b []byte) error {
	var res AppRequest35
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
