// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20014EmbeddedEmbeddedDiskLinks inline response 200 14 embedded embedded disk links
// swagger:model inline_response_200_14__embedded__embedded_disk__links
type InlineResponse20014EmbeddedEmbeddedDiskLinks struct {

	// account
	Account *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"account,omitempty"`

	// database
	Database *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"database,omitempty"`

	// operations
	Operations *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"operations,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"self,omitempty"`
}

// Validate validates this inline response 200 14 embedded embedded disk links
func (m *InlineResponse20014EmbeddedEmbeddedDiskLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperations(formats); err != nil {
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

func (m *InlineResponse20014EmbeddedEmbeddedDiskLinks) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20014EmbeddedEmbeddedDiskLinks) validateDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.Database) { // not required
		return nil
	}

	if m.Database != nil {
		if err := m.Database.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20014EmbeddedEmbeddedDiskLinks) validateOperations(formats strfmt.Registry) error {

	if swag.IsZero(m.Operations) { // not required
		return nil
	}

	if m.Operations != nil {
		if err := m.Operations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operations")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20014EmbeddedEmbeddedDiskLinks) validateSelf(formats strfmt.Registry) error {

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
func (m *InlineResponse20014EmbeddedEmbeddedDiskLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20014EmbeddedEmbeddedDiskLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse20014EmbeddedEmbeddedDiskLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
