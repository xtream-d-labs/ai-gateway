// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostConfigurationsParamsBody AccountInfo
// swagger:model postConfigurationsParamsBody
type PostConfigurationsParamsBody struct {

	// Hostname for the private Docker registry
	DockerHostname string `json:"docker_hostname,omitempty"`

	// Password for the private Docker registry
	// Format: password
	DockerPassword strfmt.Password `json:"docker_password,omitempty"`

	// Docker Registry endpoint
	DockerRegistry string `json:"docker_registry,omitempty"`

	// Username for the private Docker registry
	DockerUsername string `json:"docker_username,omitempty"`

	// kubecfg
	K8sConfig string `json:"k8s_config,omitempty"`

	// NGC - API Key
	NgcApikey string `json:"ngc_apikey,omitempty"`

	// E-mail address for NGC console
	// Format: email
	NgcEmail strfmt.Email `json:"ngc_email,omitempty"`

	// Password for NGC console
	// Format: password
	NgcPassword strfmt.Password `json:"ngc_password,omitempty"`

	// Rescale - API Key
	RescaleKey string `json:"rescale_key,omitempty"`

	// Rescale platform endopoint
	// Enum: [https://platform.rescale.com https://platform.rescale.jp https://kr.rescale.com https://eu.rescale.com]
	RescalePlatform string `json:"rescale_platform,omitempty"`
}

// Validate validates this post configurations params body
func (m *PostConfigurationsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDockerPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNgcEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNgcPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescalePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostConfigurationsParamsBody) validateDockerPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.DockerPassword) { // not required
		return nil
	}

	if err := validate.FormatOf("docker_password", "body", "password", m.DockerPassword.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostConfigurationsParamsBody) validateNgcEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.NgcEmail) { // not required
		return nil
	}

	if err := validate.FormatOf("ngc_email", "body", "email", m.NgcEmail.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostConfigurationsParamsBody) validateNgcPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.NgcPassword) { // not required
		return nil
	}

	if err := validate.FormatOf("ngc_password", "body", "password", m.NgcPassword.String(), formats); err != nil {
		return err
	}

	return nil
}

var postConfigurationsParamsBodyTypeRescalePlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["https://platform.rescale.com","https://platform.rescale.jp","https://kr.rescale.com","https://eu.rescale.com"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postConfigurationsParamsBodyTypeRescalePlatformPropEnum = append(postConfigurationsParamsBodyTypeRescalePlatformPropEnum, v)
	}
}

const (

	// PostConfigurationsParamsBodyRescalePlatformHTTPSPlatformRescaleCom captures enum value "https://platform.rescale.com"
	PostConfigurationsParamsBodyRescalePlatformHTTPSPlatformRescaleCom string = "https://platform.rescale.com"

	// PostConfigurationsParamsBodyRescalePlatformHTTPSPlatformRescaleJp captures enum value "https://platform.rescale.jp"
	PostConfigurationsParamsBodyRescalePlatformHTTPSPlatformRescaleJp string = "https://platform.rescale.jp"

	// PostConfigurationsParamsBodyRescalePlatformHTTPSKrRescaleCom captures enum value "https://kr.rescale.com"
	PostConfigurationsParamsBodyRescalePlatformHTTPSKrRescaleCom string = "https://kr.rescale.com"

	// PostConfigurationsParamsBodyRescalePlatformHTTPSEuRescaleCom captures enum value "https://eu.rescale.com"
	PostConfigurationsParamsBodyRescalePlatformHTTPSEuRescaleCom string = "https://eu.rescale.com"
)

// prop value enum
func (m *PostConfigurationsParamsBody) validateRescalePlatformEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postConfigurationsParamsBodyTypeRescalePlatformPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostConfigurationsParamsBody) validateRescalePlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.RescalePlatform) { // not required
		return nil
	}

	// value enum
	if err := m.validateRescalePlatformEnum("rescale_platform", "body", m.RescalePlatform); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostConfigurationsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostConfigurationsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostConfigurationsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
