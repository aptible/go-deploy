// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2009EmbeddedLinks inline response 200 9 embedded links
// swagger:model inline_response_200_9__embedded__links
type InlineResponse2009EmbeddedLinks struct {

	// resource
	Resource *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"resource,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`
}

// Validate validates this inline response 200 9 embedded links
func (m *InlineResponse2009EmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
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

func (m *InlineResponse2009EmbeddedLinks) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2009EmbeddedLinks) validateSelf(formats strfmt.Registry) error {

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
func (m *InlineResponse2009EmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2009EmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse2009EmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
