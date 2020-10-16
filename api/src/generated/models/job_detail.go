// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobDetail the details of a job
//
// swagger:model JobDetail
type JobDetail struct {
	Job

	JobLogs

	JobFiles
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *JobDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Job
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Job = aO0

	// AO1
	var aO1 JobLogs
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.JobLogs = aO1

	// AO2
	var aO2 JobFiles
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.JobFiles = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m JobDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.Job)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.JobLogs)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.JobFiles)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this job detail
func (m *JobDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Job
	if err := m.Job.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with JobLogs
	if err := m.JobLogs.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with JobFiles
	if err := m.JobFiles.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JobDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobDetail) UnmarshalBinary(b []byte) error {
	var res JobDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
