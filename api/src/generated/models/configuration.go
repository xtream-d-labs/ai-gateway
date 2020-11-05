// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Configuration app's configurations
//
// swagger:model Configuration
type Configuration struct {

	// Hostname for the private Docker registry
	DockerHostname string `json:"docker_hostname,omitempty"`

	// Fist 3 chars of the password of the private Docker registry
	DockerPassword string `json:"docker_password,omitempty"`

	// Docker Registry endpoint
	DockerRegistry string `json:"docker_registry,omitempty"`

	// Username for the private Docker registry
	DockerUsername string `json:"docker_username,omitempty"`

	// kubecfg
	K8sConfig string `json:"k8s_config,omitempty"`

	// Number of the host GPUs
	// Required: true
	LocalGpus *int64 `json:"local_gpus"`

	// Number of GPUs per container
	// Required: true
	LocalGpusPerContainer *int64 `json:"local_gpus_per_container"`

	// Users should be signed in
	// Required: true
	// Enum: [yes no]
	MustSignedIn *string `json:"must_signed_in"`

	// Fist 5 chars of NGC API Key
	NgcApikey string `json:"ngc_apikey,omitempty"`

	// E-mail address for NGC console
	// Format: email
	NgcEmail strfmt.Email `json:"ngc_email,omitempty"`

	// Fist 3 chars of the password for NGC console
	NgcPassword string `json:"ngc_password,omitempty"`

	// Fist 5 chars of Rescal API Key
	RescaleKey string `json:"rescale_key,omitempty"`

	// rescale platform
	// Enum: [https://platform.rescale.com https://platform.rescale.jp https://kr.rescale.com https://eu.rescale.com]
	RescalePlatform string `json:"rescale_platform,omitempty"`

	// Kubernetes will be used or not
	// Enum: [yes no]
	UseK8s string `json:"use_k8s,omitempty"`

	// NGC will be used or not
	// Enum: [yes no]
	UseNgc string `json:"use_ngc,omitempty"`

	// Private registry will be used or not
	// Enum: [yes no]
	UsePrivateRegistry string `json:"use_private_registry,omitempty"`

	// Rescale will be used or not
	// Enum: [yes no]
	UseRescale string `json:"use_rescale,omitempty"`
}

// Validate validates this configuration
func (m *Configuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalGpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalGpusPerContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMustSignedIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNgcEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescalePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseK8s(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseNgc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsePrivateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseRescale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configuration) validateLocalGpus(formats strfmt.Registry) error {

	if err := validate.Required("local_gpus", "body", m.LocalGpus); err != nil {
		return err
	}

	return nil
}

func (m *Configuration) validateLocalGpusPerContainer(formats strfmt.Registry) error {

	if err := validate.Required("local_gpus_per_container", "body", m.LocalGpusPerContainer); err != nil {
		return err
	}

	return nil
}

var configurationTypeMustSignedInPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["yes","no"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationTypeMustSignedInPropEnum = append(configurationTypeMustSignedInPropEnum, v)
	}
}

const (

	// ConfigurationMustSignedInYes captures enum value "yes"
	ConfigurationMustSignedInYes string = "yes"

	// ConfigurationMustSignedInNo captures enum value "no"
	ConfigurationMustSignedInNo string = "no"
)

// prop value enum
func (m *Configuration) validateMustSignedInEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configurationTypeMustSignedInPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Configuration) validateMustSignedIn(formats strfmt.Registry) error {

	if err := validate.Required("must_signed_in", "body", m.MustSignedIn); err != nil {
		return err
	}

	// value enum
	if err := m.validateMustSignedInEnum("must_signed_in", "body", *m.MustSignedIn); err != nil {
		return err
	}

	return nil
}

func (m *Configuration) validateNgcEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.NgcEmail) { // not required
		return nil
	}

	if err := validate.FormatOf("ngc_email", "body", "email", m.NgcEmail.String(), formats); err != nil {
		return err
	}

	return nil
}

var configurationTypeRescalePlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["https://platform.rescale.com","https://platform.rescale.jp","https://kr.rescale.com","https://eu.rescale.com"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationTypeRescalePlatformPropEnum = append(configurationTypeRescalePlatformPropEnum, v)
	}
}

const (

	// ConfigurationRescalePlatformHTTPSPlatformRescaleCom captures enum value "https://platform.rescale.com"
	ConfigurationRescalePlatformHTTPSPlatformRescaleCom string = "https://platform.rescale.com"

	// ConfigurationRescalePlatformHTTPSPlatformRescaleJp captures enum value "https://platform.rescale.jp"
	ConfigurationRescalePlatformHTTPSPlatformRescaleJp string = "https://platform.rescale.jp"

	// ConfigurationRescalePlatformHTTPSKrRescaleCom captures enum value "https://kr.rescale.com"
	ConfigurationRescalePlatformHTTPSKrRescaleCom string = "https://kr.rescale.com"

	// ConfigurationRescalePlatformHTTPSEuRescaleCom captures enum value "https://eu.rescale.com"
	ConfigurationRescalePlatformHTTPSEuRescaleCom string = "https://eu.rescale.com"
)

// prop value enum
func (m *Configuration) validateRescalePlatformEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configurationTypeRescalePlatformPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Configuration) validateRescalePlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.RescalePlatform) { // not required
		return nil
	}

	// value enum
	if err := m.validateRescalePlatformEnum("rescale_platform", "body", m.RescalePlatform); err != nil {
		return err
	}

	return nil
}

var configurationTypeUseK8sPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["yes","no"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationTypeUseK8sPropEnum = append(configurationTypeUseK8sPropEnum, v)
	}
}

const (

	// ConfigurationUseK8sYes captures enum value "yes"
	ConfigurationUseK8sYes string = "yes"

	// ConfigurationUseK8sNo captures enum value "no"
	ConfigurationUseK8sNo string = "no"
)

// prop value enum
func (m *Configuration) validateUseK8sEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configurationTypeUseK8sPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Configuration) validateUseK8s(formats strfmt.Registry) error {

	if swag.IsZero(m.UseK8s) { // not required
		return nil
	}

	// value enum
	if err := m.validateUseK8sEnum("use_k8s", "body", m.UseK8s); err != nil {
		return err
	}

	return nil
}

var configurationTypeUseNgcPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["yes","no"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationTypeUseNgcPropEnum = append(configurationTypeUseNgcPropEnum, v)
	}
}

const (

	// ConfigurationUseNgcYes captures enum value "yes"
	ConfigurationUseNgcYes string = "yes"

	// ConfigurationUseNgcNo captures enum value "no"
	ConfigurationUseNgcNo string = "no"
)

// prop value enum
func (m *Configuration) validateUseNgcEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configurationTypeUseNgcPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Configuration) validateUseNgc(formats strfmt.Registry) error {

	if swag.IsZero(m.UseNgc) { // not required
		return nil
	}

	// value enum
	if err := m.validateUseNgcEnum("use_ngc", "body", m.UseNgc); err != nil {
		return err
	}

	return nil
}

var configurationTypeUsePrivateRegistryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["yes","no"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationTypeUsePrivateRegistryPropEnum = append(configurationTypeUsePrivateRegistryPropEnum, v)
	}
}

const (

	// ConfigurationUsePrivateRegistryYes captures enum value "yes"
	ConfigurationUsePrivateRegistryYes string = "yes"

	// ConfigurationUsePrivateRegistryNo captures enum value "no"
	ConfigurationUsePrivateRegistryNo string = "no"
)

// prop value enum
func (m *Configuration) validateUsePrivateRegistryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configurationTypeUsePrivateRegistryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Configuration) validateUsePrivateRegistry(formats strfmt.Registry) error {

	if swag.IsZero(m.UsePrivateRegistry) { // not required
		return nil
	}

	// value enum
	if err := m.validateUsePrivateRegistryEnum("use_private_registry", "body", m.UsePrivateRegistry); err != nil {
		return err
	}

	return nil
}

var configurationTypeUseRescalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["yes","no"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationTypeUseRescalePropEnum = append(configurationTypeUseRescalePropEnum, v)
	}
}

const (

	// ConfigurationUseRescaleYes captures enum value "yes"
	ConfigurationUseRescaleYes string = "yes"

	// ConfigurationUseRescaleNo captures enum value "no"
	ConfigurationUseRescaleNo string = "no"
)

// prop value enum
func (m *Configuration) validateUseRescaleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configurationTypeUseRescalePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Configuration) validateUseRescale(formats strfmt.Registry) error {

	if swag.IsZero(m.UseRescale) { // not required
		return nil
	}

	// value enum
	if err := m.validateUseRescaleEnum("use_rescale", "body", m.UseRescale); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configuration) UnmarshalBinary(b []byte) error {
	var res Configuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
