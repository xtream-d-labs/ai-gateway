// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/xtream-d-labs/ai-gateway/api/src/generated/models"
)

// GetImagesOKCode is the HTTP code returned for type GetImagesOK
const GetImagesOKCode int = 200

/*GetImagesOK OK

swagger:response getImagesOK
*/
type GetImagesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Image `json:"body,omitempty"`
}

// NewGetImagesOK creates GetImagesOK with default headers values
func NewGetImagesOK() *GetImagesOK {

	return &GetImagesOK{}
}

// WithPayload adds the payload to the get images o k response
func (o *GetImagesOK) WithPayload(payload []*models.Image) *GetImagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get images o k response
func (o *GetImagesOK) SetPayload(payload []*models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Image, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetImagesDefault Unexpected error

swagger:response getImagesDefault
*/
type GetImagesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetImagesDefault creates GetImagesDefault with default headers values
func NewGetImagesDefault(code int) *GetImagesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetImagesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get images default response
func (o *GetImagesDefault) WithStatusCode(code int) *GetImagesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get images default response
func (o *GetImagesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get images default response
func (o *GetImagesDefault) WithPayload(payload *models.Error) *GetImagesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get images default response
func (o *GetImagesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImagesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
