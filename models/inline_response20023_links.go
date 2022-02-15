// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InlineResponse20023Links inline response 200 23 links
//
// swagger:model inline_response_200_23__links
type InlineResponse20023Links struct {

	// accounts
	Accounts *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"accounts,omitempty"`

	// apps
	Apps *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"apps,omitempty"`

	// database images
	DatabaseImages *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"database_images,omitempty"`

	// databases
	Databases *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"databases,omitempty"`

	// stacks
	Stacks *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"stacks,omitempty"`
}

// Validate validates this inline response 200 23 links
func (m *InlineResponse20023Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20023Links) validateAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	if m.Accounts != nil {
		if err := m.Accounts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accounts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accounts")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20023Links) validateApps(formats strfmt.Registry) error {
	if swag.IsZero(m.Apps) { // not required
		return nil
	}

	if m.Apps != nil {
		if err := m.Apps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apps")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20023Links) validateDatabaseImages(formats strfmt.Registry) error {
	if swag.IsZero(m.DatabaseImages) { // not required
		return nil
	}

	if m.DatabaseImages != nil {
		if err := m.DatabaseImages.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database_images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("database_images")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20023Links) validateDatabases(formats strfmt.Registry) error {
	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	if m.Databases != nil {
		if err := m.Databases.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databases")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databases")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20023Links) validateStacks(formats strfmt.Registry) error {
	if swag.IsZero(m.Stacks) { // not required
		return nil
	}

	if m.Stacks != nil {
		if err := m.Stacks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stacks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stacks")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this inline response 200 23 links based on the context it is used
func (m *InlineResponse20023Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatabaseImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatabases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20023Links) contextValidateAccounts(ctx context.Context, formats strfmt.Registry) error {

	if m.Accounts != nil {
		if err := m.Accounts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accounts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accounts")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20023Links) contextValidateApps(ctx context.Context, formats strfmt.Registry) error {

	if m.Apps != nil {
		if err := m.Apps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apps")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20023Links) contextValidateDatabaseImages(ctx context.Context, formats strfmt.Registry) error {

	if m.DatabaseImages != nil {
		if err := m.DatabaseImages.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database_images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("database_images")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20023Links) contextValidateDatabases(ctx context.Context, formats strfmt.Registry) error {

	if m.Databases != nil {
		if err := m.Databases.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databases")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databases")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20023Links) contextValidateStacks(ctx context.Context, formats strfmt.Registry) error {

	if m.Stacks != nil {
		if err := m.Stacks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stacks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stacks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20023Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20023Links) UnmarshalBinary(b []byte) error {
	var res InlineResponse20023Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
