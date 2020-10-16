// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Versions application versions
//
// swagger:model Versions
type Versions struct {

	// Current running service version
	// Required: true
	Current *Version `json:"current"`

	// The latest application version which can be installed
	// Required: true
	Latest *Version `json:"latest"`
}

// Validate validates this versions
func (m *Versions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Versions) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *Versions) validateLatest(formats strfmt.Registry) error {

	if err := validate.Required("latest", "body", m.Latest); err != nil {
		return err
	}

	if m.Latest != nil {
		if err := m.Latest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Versions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Versions) UnmarshalBinary(b []byte) error {
	var res Versions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
