// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2002EmbeddedActivityReports inline response 200 2 embedded activity reports
// swagger:model inline_response_200_2__embedded_activity_reports
type InlineResponse2002EmbeddedActivityReports struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse2002EmbeddedLinks `json:"_links,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// ends at
	EndsAt string `json:"ends_at,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// starts at
	StartsAt string `json:"starts_at,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this inline response 200 2 embedded activity reports
func (m *InlineResponse2002EmbeddedActivityReports) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2002EmbeddedActivityReports) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2002EmbeddedActivityReports) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2002EmbeddedActivityReports) UnmarshalBinary(b []byte) error {
	var res InlineResponse2002EmbeddedActivityReports
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
