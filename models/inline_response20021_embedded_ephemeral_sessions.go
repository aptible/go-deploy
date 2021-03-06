// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20021EmbeddedEphemeralSessions inline response 200 21 embedded ephemeral sessions
// swagger:model inline_response_200_21__embedded_ephemeral_sessions
type InlineResponse20021EmbeddedEphemeralSessions struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse20021EmbeddedLinks `json:"_links,omitempty"`

	// aws instance id
	AwsInstanceID string `json:"aws_instance_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this inline response 200 21 embedded ephemeral sessions
func (m *InlineResponse20021EmbeddedEphemeralSessions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20021EmbeddedEphemeralSessions) validateLinks(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *InlineResponse20021EmbeddedEphemeralSessions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20021EmbeddedEphemeralSessions) UnmarshalBinary(b []byte) error {
	var res InlineResponse20021EmbeddedEphemeralSessions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
