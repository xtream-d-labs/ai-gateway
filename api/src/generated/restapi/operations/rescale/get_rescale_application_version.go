// Code generated by go-swagger; DO NOT EDIT.

package rescale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/xtream-d-labs/ai-gateway/api/src/auth"
)

// GetRescaleApplicationVersionHandlerFunc turns a function with the right signature into a get rescale application version handler
type GetRescaleApplicationVersionHandlerFunc func(GetRescaleApplicationVersionParams, *auth.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRescaleApplicationVersionHandlerFunc) Handle(params GetRescaleApplicationVersionParams, principal *auth.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetRescaleApplicationVersionHandler interface for that can handle valid get rescale application version params
type GetRescaleApplicationVersionHandler interface {
	Handle(GetRescaleApplicationVersionParams, *auth.Principal) middleware.Responder
}

// NewGetRescaleApplicationVersion creates a new http.Handler for the get rescale application version operation
func NewGetRescaleApplicationVersion(ctx *middleware.Context, handler GetRescaleApplicationVersionHandler) *GetRescaleApplicationVersion {
	return &GetRescaleApplicationVersion{Context: ctx, Handler: handler}
}

/*GetRescaleApplicationVersion swagger:route GET /rescale/applications/{code}/{version}/ rescale getRescaleApplicationVersion

returns version information of a specified Rescale application


*/
type GetRescaleApplicationVersion struct {
	Context *middleware.Context
	Handler GetRescaleApplicationVersionHandler
}

func (o *GetRescaleApplicationVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRescaleApplicationVersionParams()

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
