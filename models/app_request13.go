// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppRequest13 app request 13
// swagger:model app_request_13
type AppRequest13 struct {

	// current kms arn
	CurrentKmsArn int64 `json:"current_kms_arn,omitempty"`

	// Alternate name for `database_image_id`
	DatabaseImage int64 `json:"database_image,omitempty"`

	// database image id
	// Required: true
	DatabaseImageID *int64 `json:"database_image_id"`

	// handle
	// Required: true
	Handle *string `json:"handle"`

	// initial container size
	InitialContainerSize int64 `json:"initial_container_size,omitempty"`

	// initial disk size
	InitialDiskSize int64 `json:"initial_disk_size,omitempty"`

	// initialize from
	InitializeFrom string `json:"initialize_from,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this app request 13
func (m *AppRequest13) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRequest13) validateDatabaseImageID(formats strfmt.Registry) error {

	if err := validate.Required("database_image_id", "body", m.DatabaseImageID); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest13) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest13) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest13) UnmarshalBinary(b []byte) error {
	var res AppRequest13
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
