// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InlineResponse20029Embedded inline response 200 29 embedded
//
// swagger:model inline_response_200_29__embedded
type InlineResponse20029Embedded struct {

	// metric drains
	MetricDrains []*InlineResponse20029EmbeddedMetricDrains `json:"metric_drains"`
}

// Validate validates this inline response 200 29 embedded
func (m *InlineResponse20029Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetricDrains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20029Embedded) validateMetricDrains(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricDrains) { // not required
		return nil
	}

	for i := 0; i < len(m.MetricDrains); i++ {
		if swag.IsZero(m.MetricDrains[i]) { // not required
			continue
		}

		if m.MetricDrains[i] != nil {
			if err := m.MetricDrains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metric_drains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metric_drains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this inline response 200 29 embedded based on the context it is used
func (m *InlineResponse20029Embedded) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetricDrains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20029Embedded) contextValidateMetricDrains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetricDrains); i++ {

		if m.MetricDrains[i] != nil {
			if err := m.MetricDrains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metric_drains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metric_drains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20029Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20029Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20029Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
