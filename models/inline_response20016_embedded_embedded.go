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

// InlineResponse20016EmbeddedEmbedded inline response 200 16 embedded embedded
//
// swagger:model inline_response_200_16__embedded__embedded
type InlineResponse20016EmbeddedEmbedded struct {

	// database credentials
	DatabaseCredentials []*InlineResponse20012EmbeddedDatabaseCredentials `json:"database_credentials"`

	// disk
	Disk *InlineResponse20016EmbeddedEmbeddedDisk `json:"disk,omitempty"`

	// last operation
	LastOperation *InlineResponse2003EmbeddedEmbeddedLastOperation `json:"last_operation,omitempty"`
}

// Validate validates this inline response 200 16 embedded embedded
func (m *InlineResponse20016EmbeddedEmbedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20016EmbeddedEmbedded) validateDatabaseCredentials(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("database_credentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InlineResponse20016EmbeddedEmbedded) validateDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if m.Disk != nil {
		if err := m.Disk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20016EmbeddedEmbedded) validateLastOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.LastOperation) { // not required
		return nil
	}

	if m.LastOperation != nil {
		if err := m.LastOperation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_operation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this inline response 200 16 embedded embedded based on the context it is used
func (m *InlineResponse20016EmbeddedEmbedded) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatabaseCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20016EmbeddedEmbedded) contextValidateDatabaseCredentials(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DatabaseCredentials); i++ {

		if m.DatabaseCredentials[i] != nil {
			if err := m.DatabaseCredentials[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("database_credentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("database_credentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InlineResponse20016EmbeddedEmbedded) contextValidateDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Disk != nil {
		if err := m.Disk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20016EmbeddedEmbedded) contextValidateLastOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.LastOperation != nil {
		if err := m.LastOperation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20016EmbeddedEmbedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20016EmbeddedEmbedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20016EmbeddedEmbedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
