// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2003EmbeddedEmbeddedLinks inline response 200 3 embedded embedded links
// swagger:model inline_response_200_3__embedded__embedded__links
type InlineResponse2003EmbeddedEmbeddedLinks struct {

	// account
	Account *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"account,omitempty"`

	// app
	App *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"app,omitempty"`

	// current release
	CurrentRelease *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"current_release,omitempty"`

	// database
	Database *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"database,omitempty"`

	// operations
	Operations *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"operations,omitempty"`

	// releases
	Releases *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"releases,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`

	// vhosts
	Vhosts *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"vhosts,omitempty"`
}

// Validate validates this inline response 200 3 embedded embedded links
func (m *InlineResponse2003EmbeddedEmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVhosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2003EmbeddedEmbeddedLinks) validateAccount(formats strfmt.Registry) error {

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

func (m *InlineResponse2003EmbeddedEmbeddedLinks) validateApp(formats strfmt.Registry) error {

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

func (m *InlineResponse2003EmbeddedEmbeddedLinks) validateCurrentRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentRelease) { // not required
		return nil
	}

	if m.CurrentRelease != nil {
		if err := m.CurrentRelease.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_release")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedEmbeddedLinks) validateDatabase(formats strfmt.Registry) error {

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

func (m *InlineResponse2003EmbeddedEmbeddedLinks) validateOperations(formats strfmt.Registry) error {

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

func (m *InlineResponse2003EmbeddedEmbeddedLinks) validateReleases(formats strfmt.Registry) error {

	if swag.IsZero(m.Releases) { // not required
		return nil
	}

	if m.Releases != nil {
		if err := m.Releases.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releases")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedEmbeddedLinks) validateSelf(formats strfmt.Registry) error {

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

func (m *InlineResponse2003EmbeddedEmbeddedLinks) validateVhosts(formats strfmt.Registry) error {

	if swag.IsZero(m.Vhosts) { // not required
		return nil
	}

	if m.Vhosts != nil {
		if err := m.Vhosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vhosts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedEmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedEmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse2003EmbeddedEmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
