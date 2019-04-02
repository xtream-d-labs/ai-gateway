// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetWorkspacesHandlerFunc turns a function with the right signature into a get workspaces handler
type GetWorkspacesHandlerFunc func(GetWorkspacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWorkspacesHandlerFunc) Handle(params GetWorkspacesParams) middleware.Responder {
	return fn(params)
}

// GetWorkspacesHandler interface for that can handle valid get workspaces params
type GetWorkspacesHandler interface {
	Handle(GetWorkspacesParams) middleware.Responder
}

// NewGetWorkspaces creates a new http.Handler for the get workspaces operation
func NewGetWorkspaces(ctx *middleware.Context, handler GetWorkspacesHandler) *GetWorkspaces {
	return &GetWorkspaces{Context: ctx, Handler: handler}
}

/*GetWorkspaces swagger:route GET /workspaces workspace getWorkspaces

returns user's workspaces


*/
type GetWorkspaces struct {
	Context *middleware.Context
	Handler GetWorkspacesHandler
}

func (o *GetWorkspaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetWorkspacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
