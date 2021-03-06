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

// InlineResponse2007 inline response 200 7
// swagger:model inline_response_200_7
type InlineResponse2007 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// embedded
	// Required: true
	Embedded *InlineResponse2006EmbeddedEmbedded `json:"_embedded"`

	// links
	// Required: true
	Links *InlineResponse2006EmbeddedLinks `json:"_links"`

	// aws region
	// Required: true
	AwsRegion *string `json:"aws_region"`

	// aws snapshot id
	// Required: true
	AwsSnapshotID *string `json:"aws_snapshot_id"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// database handle
	// Required: true
	DatabaseHandle *string `json:"database_handle"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// kms arn
	// Required: true
	KmsArn *string `json:"kms_arn"`

	// manual
	// Required: true
	Manual *bool `json:"manual"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this inline response 200 7
func (m *InlineResponse2007) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKmsArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2007) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateEmbedded(formats strfmt.Registry) error {

	if err := validate.Required("_embedded", "body", m.Embedded); err != nil {
		return err
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2007) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
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

func (m *InlineResponse2007) validateAwsRegion(formats strfmt.Registry) error {

	if err := validate.Required("aws_region", "body", m.AwsRegion); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateAwsSnapshotID(formats strfmt.Registry) error {

	if err := validate.Required("aws_snapshot_id", "body", m.AwsSnapshotID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateDatabaseHandle(formats strfmt.Registry) error {

	if err := validate.Required("database_handle", "body", m.DatabaseHandle); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateKmsArn(formats strfmt.Registry) error {

	if err := validate.Required("kms_arn", "body", m.KmsArn); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateManual(formats strfmt.Registry) error {

	if err := validate.Required("manual", "body", m.Manual); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2007) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2007) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2007) UnmarshalBinary(b []byte) error {
	var res InlineResponse2007
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
