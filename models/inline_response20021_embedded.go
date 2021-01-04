// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20021Embedded inline response 200 21 embedded
// swagger:model inline_response_200_21__embedded
type InlineResponse20021Embedded struct {

	// ephemeral sessions
	EphemeralSessions []*InlineResponse20021EmbeddedEphemeralSessions `json:"ephemeral_sessions"`
}

// Validate validates this inline response 200 21 embedded
func (m *InlineResponse20021Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEphemeralSessions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20021Embedded) validateEphemeralSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.EphemeralSessions) { // not required
		return nil
	}

	for i := 0; i < len(m.EphemeralSessions); i++ {
		if swag.IsZero(m.EphemeralSessions[i]) { // not required
			continue
		}

		if m.EphemeralSessions[i] != nil {
			if err := m.EphemeralSessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ephemeral_sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20021Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20021Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20021Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
