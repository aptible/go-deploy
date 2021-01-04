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

// InlineResponse20043Embedded inline response 200 43 embedded
// swagger:model inline_response_200_43__embedded
type InlineResponse20043Embedded struct {

	// vpn tunnels
	VpnTunnels []*InlineResponse20043EmbeddedVpnTunnels `json:"vpn_tunnels"`
}

// Validate validates this inline response 200 43 embedded
func (m *InlineResponse20043Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVpnTunnels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20043Embedded) validateVpnTunnels(formats strfmt.Registry) error {

	if swag.IsZero(m.VpnTunnels) { // not required
		return nil
	}

	for i := 0; i < len(m.VpnTunnels); i++ {
		if swag.IsZero(m.VpnTunnels[i]) { // not required
			continue
		}

		if m.VpnTunnels[i] != nil {
			if err := m.VpnTunnels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vpn_tunnels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20043Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20043Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20043Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
