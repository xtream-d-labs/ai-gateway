// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rescale-labs/scaleshift/api/src/generated/models"
)

// GetJobsOKCode is the HTTP code returned for type GetJobsOK
const GetJobsOKCode int = 200

/*GetJobsOK OK

swagger:response getJobsOK
*/
type GetJobsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Job `json:"body,omitempty"`
}

// NewGetJobsOK creates GetJobsOK with default headers values
func NewGetJobsOK() *GetJobsOK {

	return &GetJobsOK{}
}

// WithPayload adds the payload to the get jobs o k response
func (o *GetJobsOK) WithPayload(payload []*models.Job) *GetJobsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get jobs o k response
func (o *GetJobsOK) SetPayload(payload []*models.Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Job, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetJobsDefault Unexpected error

swagger:response getJobsDefault
*/
type GetJobsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetJobsDefault creates GetJobsDefault with default headers values
func NewGetJobsDefault(code int) *GetJobsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetJobsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get jobs default response
func (o *GetJobsDefault) WithStatusCode(code int) *GetJobsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get jobs default response
func (o *GetJobsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get jobs default response
func (o *GetJobsDefault) WithPayload(payload *models.Error) *GetJobsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get jobs default response
func (o *GetJobsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
