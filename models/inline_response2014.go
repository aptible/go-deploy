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

// InlineResponse2014 inline response 201 4
//
// swagger:model inline_response_201_4
type InlineResponse2014 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// embedded
	// Required: true
	Embedded *InlineResponse20016EmbeddedEmbedded `json:"_embedded"`

	// links
	// Required: true
	Links *InlineResponse20016EmbeddedLinks `json:"_links"`

	// connection url
	// Required: true
	ConnectionURL *string `json:"connection_url"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// current kms arn
	// Required: true
	CurrentKmsArn *string `json:"current_kms_arn"`

	// docker repo
	// Required: true
	DockerRepo *string `json:"docker_repo"`

	// handle
	// Required: true
	Handle *string `json:"handle"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// initial container size
	// Required: true
	InitialContainerSize *int64 `json:"initial_container_size"`

	// initial disk size
	// Required: true
	InitialDiskSize *int64 `json:"initial_disk_size"`

	// passphrase
	// Required: true
	Passphrase *string `json:"passphrase"`

	// port mapping
	// Required: true
	PortMapping [][]int64 `json:"port_mapping"`

	// provisioned
	// Required: true
	Provisioned *bool `json:"provisioned"`

	// status
	// Required: true
	Status *string `json:"status"`

	// type
	// Required: true
	Type *string `json:"type"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this inline response 201 4
func (m *InlineResponse2014) Validate(formats strfmt.Registry) error {
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

	if err := m.validateConnectionURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentKmsArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialContainerSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialDiskSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassphrase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisioned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *InlineResponse2014) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateEmbedded(formats strfmt.Registry) error {

	if err := validate.Required("_embedded", "body", m.Embedded); err != nil {
		return err
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2014) validateLinks(formats strfmt.Registry) error {

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

func (m *InlineResponse2014) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.Required("connection_url", "body", m.ConnectionURL); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateCurrentKmsArn(formats strfmt.Registry) error {

	if err := validate.Required("current_kms_arn", "body", m.CurrentKmsArn); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateDockerRepo(formats strfmt.Registry) error {

	if err := validate.Required("docker_repo", "body", m.DockerRepo); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateInitialContainerSize(formats strfmt.Registry) error {

	if err := validate.Required("initial_container_size", "body", m.InitialContainerSize); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateInitialDiskSize(formats strfmt.Registry) error {

	if err := validate.Required("initial_disk_size", "body", m.InitialDiskSize); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validatePassphrase(formats strfmt.Registry) error {

	if err := validate.Required("passphrase", "body", m.Passphrase); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validatePortMapping(formats strfmt.Registry) error {

	if err := validate.Required("port_mapping", "body", m.PortMapping); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateProvisioned(formats strfmt.Registry) error {

	if err := validate.Required("provisioned", "body", m.Provisioned); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2014) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this inline response 201 4 based on the context it is used
func (m *InlineResponse2014) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmbedded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2014) contextValidateEmbedded(ctx context.Context, formats strfmt.Registry) error {

	if m.Embedded != nil {
		if err := m.Embedded.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2014) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *InlineResponse2014) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2014) UnmarshalBinary(b []byte) error {
	var res InlineResponse2014
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
