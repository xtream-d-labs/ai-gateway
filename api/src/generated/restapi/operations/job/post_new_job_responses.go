// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rescale-labs/scaleshift/api/src/generated/models"
)

// PostNewJobCreatedCode is the HTTP code returned for type PostNewJobCreated
const PostNewJobCreatedCode int = 201

/*PostNewJobCreated OK

swagger:response postNewJobCreated
*/
type PostNewJobCreated struct {

	/*
	  In: Body
	*/
	Payload *models.PostNewJobCreatedBody `json:"body,omitempty"`
}

// NewPostNewJobCreated creates PostNewJobCreated with default headers values
func NewPostNewJobCreated() *PostNewJobCreated {

	return &PostNewJobCreated{}
}

// WithPayload adds the payload to the post new job created response
func (o *PostNewJobCreated) WithPayload(payload *models.PostNewJobCreatedBody) *PostNewJobCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post new job created response
func (o *PostNewJobCreated) SetPayload(payload *models.PostNewJobCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostNewJobCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostNewJobDefault Unexpected error

swagger:response postNewJobDefault
*/
type PostNewJobDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostNewJobDefault creates PostNewJobDefault with default headers values
func NewPostNewJobDefault(code int) *PostNewJobDefault {
	if code <= 0 {
		code = 500
	}

	return &PostNewJobDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post new job default response
func (o *PostNewJobDefault) WithStatusCode(code int) *PostNewJobDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post new job default response
func (o *PostNewJobDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post new job default response
func (o *PostNewJobDefault) WithPayload(payload *models.Error) *PostNewJobDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post new job default response
func (o *PostNewJobDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostNewJobDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
