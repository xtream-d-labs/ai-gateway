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

// Notebook Jupyter notebook information
//
// swagger:model Notebook
type Notebook struct {

	// the number of NVIDIA GPUs
	Gpus string `json:"gpus,omitempty"`

	// the container ID
	// Required: true
	ID *string `json:"id"`

	// the image ID
	// Required: true
	Image *string `json:"image"`

	// the container name
	// Required: true
	Name *string `json:"name"`

	// the container published port
	Port int64 `json:"port,omitempty"`

	// started unix timestamp
	// Format: date-time
	Started strfmt.DateTime `json:"started,omitempty"`

	// state of the container
	State string `json:"state,omitempty"`
}

// Validate validates this notebook
func (m *Notebook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notebook) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Notebook) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *Notebook) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Notebook) validateStarted(formats strfmt.Registry) error {

	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notebook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notebook) UnmarshalBinary(b []byte) error {
	var res Notebook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
