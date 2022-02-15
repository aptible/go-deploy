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

// InlineResponse20019EmbeddedLinks inline response 200 19 embedded links
//
// swagger:model inline_response_200_19__embedded__links
type InlineResponse20019EmbeddedLinks struct {

	// ephemeral session
	EphemeralSession *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"ephemeral_session,omitempty"`

	// log drain
	LogDrain *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"log_drain,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`
}

// Validate validates this inline response 200 19 embedded links
func (m *InlineResponse20019EmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEphemeralSession(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogDrain(formats); err != nil {
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

func (m *InlineResponse20019EmbeddedLinks) validateEphemeralSession(formats strfmt.Registry) error {
	if swag.IsZero(m.EphemeralSession) { // not required
		return nil
	}

	if m.EphemeralSession != nil {
		if err := m.EphemeralSession.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeral_session")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ephemeral_session")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20019EmbeddedLinks) validateLogDrain(formats strfmt.Registry) error {
	if swag.IsZero(m.LogDrain) { // not required
		return nil
	}

	if m.LogDrain != nil {
		if err := m.LogDrain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_drain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_drain")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20019EmbeddedLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this inline response 200 19 embedded links based on the context it is used
func (m *InlineResponse20019EmbeddedLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEphemeralSession(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogDrain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20019EmbeddedLinks) contextValidateEphemeralSession(ctx context.Context, formats strfmt.Registry) error {

	if m.EphemeralSession != nil {
		if err := m.EphemeralSession.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeral_session")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ephemeral_session")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20019EmbeddedLinks) contextValidateLogDrain(ctx context.Context, formats strfmt.Registry) error {

	if m.LogDrain != nil {
		if err := m.LogDrain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_drain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_drain")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20019EmbeddedLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *InlineResponse20019EmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20019EmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse20019EmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
