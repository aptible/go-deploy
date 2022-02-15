// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InlineResponse2008EmbeddedLinks inline response 200 8 embedded links
//
// swagger:model inline_response_200_8__embedded__links
type InlineResponse2008EmbeddedLinks struct {

	// account
	Account *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"account,omitempty"`

	// apps
	Apps *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"apps,omitempty"`

	// operations
	Operations *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"operations,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`

	// vhosts
	Vhosts *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"vhosts,omitempty"`
}

// Validate validates this inline response 200 8 embedded links
func (m *InlineResponse2008EmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperations(formats); err != nil {
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

func (m *InlineResponse2008EmbeddedLinks) validateAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008EmbeddedLinks) validateApps(formats strfmt.Registry) error {
	if swag.IsZero(m.Apps) { // not required
		return nil
	}

	if m.Apps != nil {
		if err := m.Apps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apps")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008EmbeddedLinks) validateOperations(formats strfmt.Registry) error {
	if swag.IsZero(m.Operations) { // not required
		return nil
	}

	if m.Operations != nil {
		if err := m.Operations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operations")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008EmbeddedLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008EmbeddedLinks) validateVhosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Vhosts) { // not required
		return nil
	}

	if m.Vhosts != nil {
		if err := m.Vhosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vhosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vhosts")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this inline response 200 8 embedded links based on the context it is used
func (m *InlineResponse2008EmbeddedLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVhosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2008EmbeddedLinks) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.Account != nil {
		if err := m.Account.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008EmbeddedLinks) contextValidateApps(ctx context.Context, formats strfmt.Registry) error {

	if m.Apps != nil {
		if err := m.Apps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apps")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008EmbeddedLinks) contextValidateOperations(ctx context.Context, formats strfmt.Registry) error {

	if m.Operations != nil {
		if err := m.Operations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operations")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008EmbeddedLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008EmbeddedLinks) contextValidateVhosts(ctx context.Context, formats strfmt.Registry) error {

	if m.Vhosts != nil {
		if err := m.Vhosts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vhosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vhosts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2008EmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2008EmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse2008EmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
