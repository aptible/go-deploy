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

// InlineResponse20041EmbeddedLinks inline response 200 41 embedded links
//
// swagger:model inline_response_200_41__embedded__links
type InlineResponse20041EmbeddedLinks struct {

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`

	// stack
	Stack *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"stack,omitempty"`
}

// Validate validates this inline response 200 41 embedded links
func (m *InlineResponse20041EmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20041EmbeddedLinks) validateSelf(formats strfmt.Registry) error {
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

func (m *InlineResponse20041EmbeddedLinks) validateStack(formats strfmt.Registry) error {
	if swag.IsZero(m.Stack) { // not required
		return nil
	}

	if m.Stack != nil {
		if err := m.Stack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stack")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this inline response 200 41 embedded links based on the context it is used
func (m *InlineResponse20041EmbeddedLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20041EmbeddedLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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

func (m *InlineResponse20041EmbeddedLinks) contextValidateStack(ctx context.Context, formats strfmt.Registry) error {

	if m.Stack != nil {
		if err := m.Stack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20041EmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20041EmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse20041EmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
