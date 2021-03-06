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

// InlineResponse20043EmbeddedSSHPortalConnections inline response 200 43 embedded ssh portal connections
// swagger:model inline_response_200_43__embedded_ssh_portal_connections
type InlineResponse20043EmbeddedSSHPortalConnections struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse20043EmbeddedLinks `json:"_links,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// ssh certificate body
	SSHCertificateBody string `json:"ssh_certificate_body,omitempty"`

	// ssh port forward socket
	SSHPortForwardSocket *int64 `json:"ssh_port_forward_socket,omitempty"`

	// ssh pty
	SSHPty bool `json:"ssh_pty,omitempty"`

	// ssh user
	SSHUser string `json:"ssh_user,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this inline response 200 43 embedded ssh portal connections
func (m *InlineResponse20043EmbeddedSSHPortalConnections) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20043EmbeddedSSHPortalConnections) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20043EmbeddedSSHPortalConnections) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20043EmbeddedSSHPortalConnections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20043EmbeddedSSHPortalConnections) UnmarshalBinary(b []byte) error {
	var res InlineResponse20043EmbeddedSSHPortalConnections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
