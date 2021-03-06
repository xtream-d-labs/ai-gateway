// Code generated by go-swagger; DO NOT EDIT.

package notebook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetNotebookDetailsHandlerFunc turns a function with the right signature into a get notebook details handler
type GetNotebookDetailsHandlerFunc func(GetNotebookDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNotebookDetailsHandlerFunc) Handle(params GetNotebookDetailsParams) middleware.Responder {
	return fn(params)
}

// GetNotebookDetailsHandler interface for that can handle valid get notebook details params
type GetNotebookDetailsHandler interface {
	Handle(GetNotebookDetailsParams) middleware.Responder
}

// NewGetNotebookDetails creates a new http.Handler for the get notebook details operation
func NewGetNotebookDetails(ctx *middleware.Context, handler GetNotebookDetailsHandler) *GetNotebookDetails {
	return &GetNotebookDetails{Context: ctx, Handler: handler}
}

/*GetNotebookDetails swagger:route GET /notebooks/{id} notebook getNotebookDetails

returns Jupyter notebook detail information


*/
type GetNotebookDetails struct {
	Context *middleware.Context
	Handler GetNotebookDetailsHandler
}

func (o *GetNotebookDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNotebookDetailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
