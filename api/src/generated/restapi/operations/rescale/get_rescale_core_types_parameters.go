// Code generated by go-swagger; DO NOT EDIT.

package rescale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRescaleCoreTypesParams creates a new GetRescaleCoreTypesParams object
// no default values defined in spec.
func NewGetRescaleCoreTypesParams() GetRescaleCoreTypesParams {

	return GetRescaleCoreTypesParams{}
}

// GetRescaleCoreTypesParams contains all the bound params for the get rescale core types operation
// typically these are obtained from a http.Request
//
// swagger:parameters getRescaleCoreTypes
type GetRescaleCoreTypesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Rescale Application version
	  In: query
	*/
	AppVer *string
	/*Required number of GPUs
	  In: query
	*/
	MinGpus *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRescaleCoreTypesParams() beforehand.
func (o *GetRescaleCoreTypesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAppVer, qhkAppVer, _ := qs.GetOK("app_ver")
	if err := o.bindAppVer(qAppVer, qhkAppVer, route.Formats); err != nil {
		res = append(res, err)
	}

	qMinGpus, qhkMinGpus, _ := qs.GetOK("min_gpus")
	if err := o.bindMinGpus(qMinGpus, qhkMinGpus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAppVer binds and validates parameter AppVer from query.
func (o *GetRescaleCoreTypesParams) bindAppVer(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.AppVer = &raw

	return nil
}

// bindMinGpus binds and validates parameter MinGpus from query.
func (o *GetRescaleCoreTypesParams) bindMinGpus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("min_gpus", "query", "int64", raw)
	}
	o.MinGpus = &value

	return nil
}
