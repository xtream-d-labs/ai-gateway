// Code generated by go-swagger; DO NOT EDIT.

package notebook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/xtream-d-labs/ai-gateway/api/src/generated/models"
)

// GetNotebooksOKCode is the HTTP code returned for type GetNotebooksOK
const GetNotebooksOKCode int = 200

/*GetNotebooksOK OK

swagger:response getNotebooksOK
*/
type GetNotebooksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Notebook `json:"body,omitempty"`
}

// NewGetNotebooksOK creates GetNotebooksOK with default headers values
func NewGetNotebooksOK() *GetNotebooksOK {

	return &GetNotebooksOK{}
}

// WithPayload adds the payload to the get notebooks o k response
func (o *GetNotebooksOK) WithPayload(payload []*models.Notebook) *GetNotebooksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get notebooks o k response
func (o *GetNotebooksOK) SetPayload(payload []*models.Notebook) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNotebooksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Notebook, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetNotebooksDefault Unexpected error

swagger:response getNotebooksDefault
*/
type GetNotebooksDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNotebooksDefault creates GetNotebooksDefault with default headers values
func NewGetNotebooksDefault(code int) *GetNotebooksDefault {
	if code <= 0 {
		code = 500
	}

	return &GetNotebooksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get notebooks default response
func (o *GetNotebooksDefault) WithStatusCode(code int) *GetNotebooksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get notebooks default response
func (o *GetNotebooksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get notebooks default response
func (o *GetNotebooksDefault) WithPayload(payload *models.Error) *GetNotebooksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get notebooks default response
func (o *GetNotebooksDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNotebooksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
