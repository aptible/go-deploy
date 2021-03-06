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

// InlineResponse20013Embedded inline response 200 13 embedded
// swagger:model inline_response_200_13__embedded
type InlineResponse20013Embedded struct {

	// database images
	DatabaseImages []*InlineResponse20013EmbeddedDatabaseImages `json:"database_images"`
}

// Validate validates this inline response 200 13 embedded
func (m *InlineResponse20013Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20013Embedded) validateDatabaseImages(formats strfmt.Registry) error {

	if swag.IsZero(m.DatabaseImages) { // not required
		return nil
	}

	for i := 0; i < len(m.DatabaseImages); i++ {
		if swag.IsZero(m.DatabaseImages[i]) { // not required
			continue
		}

		if m.DatabaseImages[i] != nil {
			if err := m.DatabaseImages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("database_images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20013Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20013Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20013Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
