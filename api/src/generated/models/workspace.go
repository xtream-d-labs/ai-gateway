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

// Workspace User's workspace
//
// swagger:model Workspace
type Workspace struct {

	// absolute path
	AbsolutePath string `json:"absolute_path,omitempty"`

	// Jobs which are mounting the workspace
	Jobs []string `json:"jobs"`

	// Notebooks which are mounting the workspace
	Notebooks []string `json:"notebooks"`

	// path
	// Required: true
	Path *string `json:"path"`
}

// Validate validates this workspace
func (m *Workspace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workspace) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workspace) UnmarshalBinary(b []byte) error {
	var res Workspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
