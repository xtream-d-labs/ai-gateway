// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rescale-labs/scaleshift/api/src/auth"
)

// ModifyJobHandlerFunc turns a function with the right signature into a modify job handler
type ModifyJobHandlerFunc func(ModifyJobParams, *auth.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ModifyJobHandlerFunc) Handle(params ModifyJobParams, principal *auth.Principal) middleware.Responder {
	return fn(params, principal)
}

// ModifyJobHandler interface for that can handle valid modify job params
type ModifyJobHandler interface {
	Handle(ModifyJobParams, *auth.Principal) middleware.Responder
}

// NewModifyJob creates a new http.Handler for the modify job operation
func NewModifyJob(ctx *middleware.Context, handler ModifyJobHandler) *ModifyJob {
	return &ModifyJob{Context: ctx, Handler: handler}
}

/*ModifyJob swagger:route PATCH /jobs/{id} job modifyJob

modify the job status


*/
type ModifyJob struct {
	Context *middleware.Context
	Handler ModifyJobHandler
}

func (o *ModifyJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewModifyJobParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *auth.Principal
	if uprinc != nil {
		principal = uprinc.(*auth.Principal) // this is really a auth.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
