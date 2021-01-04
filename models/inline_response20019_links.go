// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20019Links inline response 200 19 links
// swagger:model inline_response_200_19__links
type InlineResponse20019Links struct {

	// ephemeral session
	EphemeralSession *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"ephemeral_session,omitempty"`

	// log drain
	LogDrain *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"log_drain,omitempty"`

	// next
	Next *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"next,omitempty"`

	// prev
	Prev *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"prev,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`
}

// Validate validates this inline response 200 19 links
func (m *InlineResponse20019Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEphemeralSession(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogDrain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20019Links) validateEphemeralSession(formats strfmt.Registry) error {

	if swag.IsZero(m.EphemeralSession) { // not required
		return nil
	}

	if m.EphemeralSession != nil {
		if err := m.EphemeralSession.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeral_session")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20019Links) validateLogDrain(formats strfmt.Registry) error {

	if swag.IsZero(m.LogDrain) { // not required
		return nil
	}

	if m.LogDrain != nil {
		if err := m.LogDrain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_drain")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20019Links) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20019Links) validatePrev(formats strfmt.Registry) error {

	if swag.IsZero(m.Prev) { // not required
		return nil
	}

	if m.Prev != nil {
		if err := m.Prev.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20019Links) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20019Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20019Links) UnmarshalBinary(b []byte) error {
	var res InlineResponse20019Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
