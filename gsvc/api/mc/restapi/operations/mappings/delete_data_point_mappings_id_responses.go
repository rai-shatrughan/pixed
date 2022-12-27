// Code generated by go-swagger; DO NOT EDIT.

package mappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/mc/models"
)

// DeleteDataPointMappingsIDNoContentCode is the HTTP code returned for type DeleteDataPointMappingsIDNoContent
const DeleteDataPointMappingsIDNoContentCode int = 204

/*
DeleteDataPointMappingsIDNoContent Deleted

swagger:response deleteDataPointMappingsIdNoContent
*/
type DeleteDataPointMappingsIDNoContent struct {
}

// NewDeleteDataPointMappingsIDNoContent creates DeleteDataPointMappingsIDNoContent with default headers values
func NewDeleteDataPointMappingsIDNoContent() *DeleteDataPointMappingsIDNoContent {

	return &DeleteDataPointMappingsIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteDataPointMappingsIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteDataPointMappingsIDBadRequestCode is the HTTP code returned for type DeleteDataPointMappingsIDBadRequest
const DeleteDataPointMappingsIDBadRequestCode int = 400

/*
DeleteDataPointMappingsIDBadRequest Bad Request

swagger:response deleteDataPointMappingsIdBadRequest
*/
type DeleteDataPointMappingsIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Badrequest `json:"body,omitempty"`
}

// NewDeleteDataPointMappingsIDBadRequest creates DeleteDataPointMappingsIDBadRequest with default headers values
func NewDeleteDataPointMappingsIDBadRequest() *DeleteDataPointMappingsIDBadRequest {

	return &DeleteDataPointMappingsIDBadRequest{}
}

// WithPayload adds the payload to the delete data point mappings Id bad request response
func (o *DeleteDataPointMappingsIDBadRequest) WithPayload(payload *models.Badrequest) *DeleteDataPointMappingsIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete data point mappings Id bad request response
func (o *DeleteDataPointMappingsIDBadRequest) SetPayload(payload *models.Badrequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDataPointMappingsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDataPointMappingsIDUnauthorizedCode is the HTTP code returned for type DeleteDataPointMappingsIDUnauthorized
const DeleteDataPointMappingsIDUnauthorizedCode int = 401

/*
DeleteDataPointMappingsIDUnauthorized unauthorized

swagger:response deleteDataPointMappingsIdUnauthorized
*/
type DeleteDataPointMappingsIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewDeleteDataPointMappingsIDUnauthorized creates DeleteDataPointMappingsIDUnauthorized with default headers values
func NewDeleteDataPointMappingsIDUnauthorized() *DeleteDataPointMappingsIDUnauthorized {

	return &DeleteDataPointMappingsIDUnauthorized{}
}

// WithPayload adds the payload to the delete data point mappings Id unauthorized response
func (o *DeleteDataPointMappingsIDUnauthorized) WithPayload(payload *models.Unauthorized) *DeleteDataPointMappingsIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete data point mappings Id unauthorized response
func (o *DeleteDataPointMappingsIDUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDataPointMappingsIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDataPointMappingsIDForbiddenCode is the HTTP code returned for type DeleteDataPointMappingsIDForbidden
const DeleteDataPointMappingsIDForbiddenCode int = 403

/*
DeleteDataPointMappingsIDForbidden Forbidden

swagger:response deleteDataPointMappingsIdForbidden
*/
type DeleteDataPointMappingsIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Forbidden `json:"body,omitempty"`
}

// NewDeleteDataPointMappingsIDForbidden creates DeleteDataPointMappingsIDForbidden with default headers values
func NewDeleteDataPointMappingsIDForbidden() *DeleteDataPointMappingsIDForbidden {

	return &DeleteDataPointMappingsIDForbidden{}
}

// WithPayload adds the payload to the delete data point mappings Id forbidden response
func (o *DeleteDataPointMappingsIDForbidden) WithPayload(payload *models.Forbidden) *DeleteDataPointMappingsIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete data point mappings Id forbidden response
func (o *DeleteDataPointMappingsIDForbidden) SetPayload(payload *models.Forbidden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDataPointMappingsIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDataPointMappingsIDNotFoundCode is the HTTP code returned for type DeleteDataPointMappingsIDNotFound
const DeleteDataPointMappingsIDNotFoundCode int = 404

/*
DeleteDataPointMappingsIDNotFound Mapping not found

swagger:response deleteDataPointMappingsIdNotFound
*/
type DeleteDataPointMappingsIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Notfound `json:"body,omitempty"`
}

// NewDeleteDataPointMappingsIDNotFound creates DeleteDataPointMappingsIDNotFound with default headers values
func NewDeleteDataPointMappingsIDNotFound() *DeleteDataPointMappingsIDNotFound {

	return &DeleteDataPointMappingsIDNotFound{}
}

// WithPayload adds the payload to the delete data point mappings Id not found response
func (o *DeleteDataPointMappingsIDNotFound) WithPayload(payload *models.Notfound) *DeleteDataPointMappingsIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete data point mappings Id not found response
func (o *DeleteDataPointMappingsIDNotFound) SetPayload(payload *models.Notfound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDataPointMappingsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
DeleteDataPointMappingsIDDefault unexpected error

swagger:response deleteDataPointMappingsIdDefault
*/
type DeleteDataPointMappingsIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDataPointMappingsIDDefault creates DeleteDataPointMappingsIDDefault with default headers values
func NewDeleteDataPointMappingsIDDefault(code int) *DeleteDataPointMappingsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteDataPointMappingsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete data point mappings ID default response
func (o *DeleteDataPointMappingsIDDefault) WithStatusCode(code int) *DeleteDataPointMappingsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete data point mappings ID default response
func (o *DeleteDataPointMappingsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete data point mappings ID default response
func (o *DeleteDataPointMappingsIDDefault) WithPayload(payload *models.Error) *DeleteDataPointMappingsIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete data point mappings ID default response
func (o *DeleteDataPointMappingsIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDataPointMappingsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
