// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20033Links inline response 200 33 links
// swagger:model inline_response_200_33__links
type InlineResponse20033Links struct {

	// next
	Next *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"next,omitempty"`

	// prev
	Prev *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"prev,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`

	// service
	Service *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"service,omitempty"`
}

// Validate validates this inline response 200 33 links
func (m *InlineResponse20033Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20033Links) validateNext(formats strfmt.Registry) error {

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

func (m *InlineResponse20033Links) validatePrev(formats strfmt.Registry) error {

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

func (m *InlineResponse20033Links) validateSelf(formats strfmt.Registry) error {

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

func (m *InlineResponse20033Links) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20033Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20033Links) UnmarshalBinary(b []byte) error {
	var res InlineResponse20033Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
