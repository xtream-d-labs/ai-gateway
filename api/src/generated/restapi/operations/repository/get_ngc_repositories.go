// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/xtream-d-labs/ai-gateway/api/src/auth"
)

// GetNgcRepositoriesHandlerFunc turns a function with the right signature into a get ngc repositories handler
type GetNgcRepositoriesHandlerFunc func(GetNgcRepositoriesParams, *auth.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNgcRepositoriesHandlerFunc) Handle(params GetNgcRepositoriesParams, principal *auth.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetNgcRepositoriesHandler interface for that can handle valid get ngc repositories params
type GetNgcRepositoriesHandler interface {
	Handle(GetNgcRepositoriesParams, *auth.Principal) middleware.Responder
}

// NewGetNgcRepositories creates a new http.Handler for the get ngc repositories operation
func NewGetNgcRepositories(ctx *middleware.Context, handler GetNgcRepositoriesHandler) *GetNgcRepositories {
	return &GetNgcRepositories{Context: ctx, Handler: handler}
}

/*GetNgcRepositories swagger:route GET /nvidia/repositories repository getNgcRepositories

returns NGC repositories


*/
type GetNgcRepositories struct {
	Context *middleware.Context
	Handler GetNgcRepositoriesHandler
}

func (o *GetNgcRepositories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNgcRepositoriesParams()

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
