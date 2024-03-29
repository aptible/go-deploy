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

// InlineResponse20030 inline response 200 30
// swagger:model inline_response_200_30
type InlineResponse20030 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// links
	// Required: true
	Links *InlineResponse2003EmbeddedEmbeddedLastOperationLinks `json:"_links"`

	// aborted
	// Required: true
	Aborted *bool `json:"aborted"`

	// cancelled
	// Required: true
	Cancelled *bool `json:"cancelled"`

	// certificate
	// Required: true
	Certificate *string `json:"certificate"`

	// command
	// Required: true
	Command *string `json:"command"`

	// container count
	// Required: true
	ContainerCount *int64 `json:"container_count"`

	// container size
	// Required: true
	ContainerSize *int64 `json:"container_size"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// destination region
	// Required: true
	DestinationRegion *string `json:"destination_region"`

	// disk size
	// Required: true
	DiskSize *int64 `json:"disk_size"`

	// docker ref
	// Required: true
	DockerRef *string `json:"docker_ref"`

	// env
	// Required: true
	Env interface{} `json:"env"`

	// git ref
	// Required: true
	GitRef *string `json:"git_ref"`

	// handle
	// Required: true
	Handle *string `json:"handle"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// instance profile
	// Required: true
	InstanceProfile *string `json:"instance_profile"`

	// interactive
	// Required: true
	Interactive *bool `json:"interactive"`

	// private key
	// Required: true
	PrivateKey *string `json:"private_key"`

	// status
	// Required: true
	Status *string `json:"status"`

	// type
	// Required: true
	Type *string `json:"type"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`

	// user email
	// Required: true
	UserEmail *string `json:"user_email"`

	// user name
	// Required: true
	UserName *string `json:"user_name"`
}

// Validate validates this inline response 200 30
func (m *InlineResponse20030) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAborted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCancelled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInteractive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
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

	if err := m.validateUserEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20030) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateLinks(formats strfmt.Registry) error {

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

func (m *InlineResponse20030) validateAborted(formats strfmt.Registry) error {

	if err := validate.Required("aborted", "body", m.Aborted); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateCancelled(formats strfmt.Registry) error {

	if err := validate.Required("cancelled", "body", m.Cancelled); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateContainerCount(formats strfmt.Registry) error {

	if err := validate.Required("container_count", "body", m.ContainerCount); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateContainerSize(formats strfmt.Registry) error {

	if err := validate.Required("container_size", "body", m.ContainerSize); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateDestinationRegion(formats strfmt.Registry) error {

	if err := validate.Required("destination_region", "body", m.DestinationRegion); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateDiskSize(formats strfmt.Registry) error {

	if err := validate.Required("disk_size", "body", m.DiskSize); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateDockerRef(formats strfmt.Registry) error {

	if err := validate.Required("docker_ref", "body", m.DockerRef); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateEnv(formats strfmt.Registry) error {

	if err := validate.Required("env", "body", m.Env); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateGitRef(formats strfmt.Registry) error {

	if err := validate.Required("git_ref", "body", m.GitRef); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateInstanceProfile(formats strfmt.Registry) error {

	if err := validate.Required("instance_profile", "body", m.InstanceProfile); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateInteractive(formats strfmt.Registry) error {

	if err := validate.Required("interactive", "body", m.Interactive); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("private_key", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateUserEmail(formats strfmt.Registry) error {

	if err := validate.Required("user_email", "body", m.UserEmail); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20030) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("user_name", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20030) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20030) UnmarshalBinary(b []byte) error {
	var res InlineResponse20030
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
