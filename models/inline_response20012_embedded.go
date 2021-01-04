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

// InlineResponse20012Embedded inline response 200 12 embedded
// swagger:model inline_response_200_12__embedded
type InlineResponse20012Embedded struct {

	// database credentials
	DatabaseCredentials []*InlineResponse20012EmbeddedDatabaseCredentials `json:"database_credentials"`
}

// Validate validates this inline response 200 12 embedded
func (m *InlineResponse20012Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20012Embedded) validateDatabaseCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.DatabaseCredentials) { // not required
		return nil
	}

	for i := 0; i < len(m.DatabaseCredentials); i++ {
		if swag.IsZero(m.DatabaseCredentials[i]) { // not required
			continue
		}

		if m.DatabaseCredentials[i] != nil {
			if err := m.DatabaseCredentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("database_credentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20012Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20012Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20012Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
