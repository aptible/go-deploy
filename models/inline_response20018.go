// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InlineResponse20018 inline response 200 18
//
// swagger:model inline_response_200_18
type InlineResponse20018 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// links
	// Required: true
	Links *InlineResponse20016EmbeddedEmbeddedDiskLinks `json:"_links"`

	// attached
	// Required: true
	Attached *bool `json:"attached"`

	// availability zone
	// Required: true
	AvailabilityZone *string `json:"availability_zone"`

	// baseline iops
	// Required: true
	BaselineIops *int64 `json:"baseline_iops"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// current kms arn
	// Required: true
	CurrentKmsArn *string `json:"current_kms_arn"`

	// device
	// Required: true
	Device *string `json:"device"`

	// ebs volume id
	// Required: true
	EbsVolumeID *string `json:"ebs_volume_id"`

	// ebs volume type
	// Required: true
	EbsVolumeType *string `json:"ebs_volume_type"`

	// ec2 instance id
	// Required: true
	Ec2InstanceID *string `json:"ec2_instance_id"`

	// filesystem
	// Required: true
	Filesystem *string `json:"filesystem"`

	// handle
	// Required: true
	Handle *string `json:"handle"`

	// host
	// Required: true
	Host *string `json:"host"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// key bytes
	// Required: true
	KeyBytes *int64 `json:"key_bytes"`

	// passphrase
	// Required: true
	Passphrase *string `json:"passphrase"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this inline response 200 18
func (m *InlineResponse20018) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttached(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailabilityZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaselineIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentKmsArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEbsVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEbsVolumeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEc2InstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassphrase(formats); err != nil {
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

func (m *InlineResponse20018) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20018) validateAttached(formats strfmt.Registry) error {

	if err := validate.Required("attached", "body", m.Attached); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateAvailabilityZone(formats strfmt.Registry) error {

	if err := validate.Required("availability_zone", "body", m.AvailabilityZone); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateBaselineIops(formats strfmt.Registry) error {

	if err := validate.Required("baseline_iops", "body", m.BaselineIops); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateCurrentKmsArn(formats strfmt.Registry) error {

	if err := validate.Required("current_kms_arn", "body", m.CurrentKmsArn); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateEbsVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("ebs_volume_id", "body", m.EbsVolumeID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateEbsVolumeType(formats strfmt.Registry) error {

	if err := validate.Required("ebs_volume_type", "body", m.EbsVolumeType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateEc2InstanceID(formats strfmt.Registry) error {

	if err := validate.Required("ec2_instance_id", "body", m.Ec2InstanceID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateFilesystem(formats strfmt.Registry) error {

	if err := validate.Required("filesystem", "body", m.Filesystem); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateKeyBytes(formats strfmt.Registry) error {

	if err := validate.Required("key_bytes", "body", m.KeyBytes); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validatePassphrase(formats strfmt.Registry) error {

	if err := validate.Required("passphrase", "body", m.Passphrase); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this inline response 200 18 based on the context it is used
func (m *InlineResponse20018) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20018) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20018) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20018) UnmarshalBinary(b []byte) error {
	var res InlineResponse20018
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
