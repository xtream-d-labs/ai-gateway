// Code generated by go-swagger; DO NOT EDIT.

package notebook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModifyNotebookHandlerFunc turns a function with the right signature into a modify notebook handler
type ModifyNotebookHandlerFunc func(ModifyNotebookParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ModifyNotebookHandlerFunc) Handle(params ModifyNotebookParams) middleware.Responder {
	return fn(params)
}

// ModifyNotebookHandler interface for that can handle valid modify notebook params
type ModifyNotebookHandler interface {
	Handle(ModifyNotebookParams) middleware.Responder
}

// NewModifyNotebook creates a new http.Handler for the modify notebook operation
func NewModifyNotebook(ctx *middleware.Context, handler ModifyNotebookHandler) *ModifyNotebook {
	return &ModifyNotebook{Context: ctx, Handler: handler}
}

/*ModifyNotebook swagger:route PATCH /notebooks/{id} notebook modifyNotebook

modify the notebook status


*/
type ModifyNotebook struct {
	Context *middleware.Context
	Handler ModifyNotebookHandler
}

func (o *ModifyNotebook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewModifyNotebookParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ModifyNotebookBody NotebookAttrs
//
// swagger:model ModifyNotebookBody
type ModifyNotebookBody struct {

	// status
	// Enum: [started stopped]
	Status string `json:"status,omitempty"`
}

// Validate validates this modify notebook body
func (o *ModifyNotebookBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var modifyNotebookBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["started","stopped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modifyNotebookBodyTypeStatusPropEnum = append(modifyNotebookBodyTypeStatusPropEnum, v)
	}
}

const (

	// ModifyNotebookBodyStatusStarted captures enum value "started"
	ModifyNotebookBodyStatusStarted string = "started"

	// ModifyNotebookBodyStatusStopped captures enum value "stopped"
	ModifyNotebookBodyStatusStopped string = "stopped"
)

// prop value enum
func (o *ModifyNotebookBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modifyNotebookBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ModifyNotebookBody) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("body"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ModifyNotebookBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyNotebookBody) UnmarshalBinary(b []byte) error {
	var res ModifyNotebookBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
