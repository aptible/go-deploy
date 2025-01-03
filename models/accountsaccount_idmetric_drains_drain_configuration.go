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

// AccountsaccountIdmetricDrainsDrainConfiguration For external metric drain hosts
// swagger:model accountsaccount_idmetric_drains_drain_configuration
type AccountsaccountIdmetricDrainsDrainConfiguration struct {

	// address
	// Format: uri
	Address strfmt.URI `json:"address,omitempty"`

	// api key
	APIKey string `json:"api_key,omitempty"`

	// database
	Database string `json:"database,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// series url
	// Format: uri
	SeriesURL strfmt.URI `json:"series_url,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	AuthToken string `json:"authToken,omitempty"`

	Bucket string `json:"bucket,omitempty"`

	Org string `json:"org,omitempty"`
}

// Validate validates this accountsaccount idmetric drains drain configuration
func (m *AccountsaccountIdmetricDrainsDrainConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeriesURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsaccountIdmetricDrainsDrainConfiguration) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.FormatOf("address", "body", "uri", m.Address.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountsaccountIdmetricDrainsDrainConfiguration) validateSeriesURL(formats strfmt.Registry) error {

	if swag.IsZero(m.SeriesURL) { // not required
		return nil
	}

	if err := validate.FormatOf("series_url", "body", "uri", m.SeriesURL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountsaccountIdmetricDrainsDrainConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountsaccountIdmetricDrainsDrainConfiguration) UnmarshalBinary(b []byte) error {
	var res AccountsaccountIdmetricDrainsDrainConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
