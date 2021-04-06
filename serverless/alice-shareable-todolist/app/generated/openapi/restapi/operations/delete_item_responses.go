// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yandex-cloud/examples/serverless/alice-shareable-todolist/app/generated/openapi/models"
)

// DeleteItemNoContentCode is the HTTP code returned for type DeleteItemNoContent
const DeleteItemNoContentCode int = 204

/*DeleteItemNoContent OK

swagger:response deleteItemNoContent
*/
type DeleteItemNoContent struct {
}

// NewDeleteItemNoContent creates DeleteItemNoContent with default headers values
func NewDeleteItemNoContent() *DeleteItemNoContent {

	return &DeleteItemNoContent{}
}

// WriteResponse to the client
func (o *DeleteItemNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteItemDefault error

swagger:response deleteItemDefault
*/
type DeleteItemDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteItemDefault creates DeleteItemDefault with default headers values
func NewDeleteItemDefault(code int) *DeleteItemDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteItemDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete item default response
func (o *DeleteItemDefault) WithStatusCode(code int) *DeleteItemDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete item default response
func (o *DeleteItemDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete item default response
func (o *DeleteItemDefault) WithPayload(payload *models.Error) *DeleteItemDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete item default response
func (o *DeleteItemDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteItemDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}