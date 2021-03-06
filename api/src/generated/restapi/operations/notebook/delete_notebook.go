// Code generated by go-swagger; DO NOT EDIT.

package notebook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteNotebookHandlerFunc turns a function with the right signature into a delete notebook handler
type DeleteNotebookHandlerFunc func(DeleteNotebookParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNotebookHandlerFunc) Handle(params DeleteNotebookParams) middleware.Responder {
	return fn(params)
}

// DeleteNotebookHandler interface for that can handle valid delete notebook params
type DeleteNotebookHandler interface {
	Handle(DeleteNotebookParams) middleware.Responder
}

// NewDeleteNotebook creates a new http.Handler for the delete notebook operation
func NewDeleteNotebook(ctx *middleware.Context, handler DeleteNotebookHandler) *DeleteNotebook {
	return &DeleteNotebook{Context: ctx, Handler: handler}
}

/*DeleteNotebook swagger:route DELETE /notebooks/{id} notebook deleteNotebook

delete a specified notebook


*/
type DeleteNotebook struct {
	Context *middleware.Context
	Handler DeleteNotebookHandler
}

func (o *DeleteNotebook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNotebookParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
