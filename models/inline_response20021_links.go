// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20021Links inline response 200 21 links
// swagger:model inline_response_200_21__links
type InlineResponse20021Links struct {

	// app
	App *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"app,omitempty"`

	// next
	Next *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"next,omitempty"`

	// operation
	Operation *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"operation,omitempty"`

	// prev
	Prev *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"prev,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`
}

// Validate validates this inline response 200 21 links
func (m *InlineResponse20021Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
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

func (m *InlineResponse20021Links) validateApp(formats strfmt.Registry) error {

	if swag.IsZero(m.App) { // not required
		return nil
	}

	if m.App != nil {
		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20021Links) validateNext(formats strfmt.Registry) error {

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

func (m *InlineResponse20021Links) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20021Links) validatePrev(formats strfmt.Registry) error {

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

func (m *InlineResponse20021Links) validateSelf(formats strfmt.Registry) error {

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
func (m *InlineResponse20021Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20021Links) UnmarshalBinary(b []byte) error {
	var res InlineResponse20021Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
