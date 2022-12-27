// Code generated by go-swagger; DO NOT EDIT.

package diagnostic_activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/mc/models"
)

// DeleteDiagnosticActivationsIDNoContentCode is the HTTP code returned for type DeleteDiagnosticActivationsIDNoContent
const DeleteDiagnosticActivationsIDNoContentCode int = 204

/*
DeleteDiagnosticActivationsIDNoContent No Content

swagger:response deleteDiagnosticActivationsIdNoContent
*/
type DeleteDiagnosticActivationsIDNoContent struct {
}

// NewDeleteDiagnosticActivationsIDNoContent creates DeleteDiagnosticActivationsIDNoContent with default headers values
func NewDeleteDiagnosticActivationsIDNoContent() *DeleteDiagnosticActivationsIDNoContent {

	return &DeleteDiagnosticActivationsIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteDiagnosticActivationsIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteDiagnosticActivationsIDBadRequestCode is the HTTP code returned for type DeleteDiagnosticActivationsIDBadRequest
const DeleteDiagnosticActivationsIDBadRequestCode int = 400

/*
DeleteDiagnosticActivationsIDBadRequest Bad Request

swagger:response deleteDiagnosticActivationsIdBadRequest
*/
type DeleteDiagnosticActivationsIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Badrequest `json:"body,omitempty"`
}

// NewDeleteDiagnosticActivationsIDBadRequest creates DeleteDiagnosticActivationsIDBadRequest with default headers values
func NewDeleteDiagnosticActivationsIDBadRequest() *DeleteDiagnosticActivationsIDBadRequest {

	return &DeleteDiagnosticActivationsIDBadRequest{}
}

// WithPayload adds the payload to the delete diagnostic activations Id bad request response
func (o *DeleteDiagnosticActivationsIDBadRequest) WithPayload(payload *models.Badrequest) *DeleteDiagnosticActivationsIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete diagnostic activations Id bad request response
func (o *DeleteDiagnosticActivationsIDBadRequest) SetPayload(payload *models.Badrequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDiagnosticActivationsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDiagnosticActivationsIDUnauthorizedCode is the HTTP code returned for type DeleteDiagnosticActivationsIDUnauthorized
const DeleteDiagnosticActivationsIDUnauthorizedCode int = 401

/*
DeleteDiagnosticActivationsIDUnauthorized Unauthorized

swagger:response deleteDiagnosticActivationsIdUnauthorized
*/
type DeleteDiagnosticActivationsIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewDeleteDiagnosticActivationsIDUnauthorized creates DeleteDiagnosticActivationsIDUnauthorized with default headers values
func NewDeleteDiagnosticActivationsIDUnauthorized() *DeleteDiagnosticActivationsIDUnauthorized {

	return &DeleteDiagnosticActivationsIDUnauthorized{}
}

// WithPayload adds the payload to the delete diagnostic activations Id unauthorized response
func (o *DeleteDiagnosticActivationsIDUnauthorized) WithPayload(payload *models.Unauthorized) *DeleteDiagnosticActivationsIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete diagnostic activations Id unauthorized response
func (o *DeleteDiagnosticActivationsIDUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDiagnosticActivationsIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDiagnosticActivationsIDForbiddenCode is the HTTP code returned for type DeleteDiagnosticActivationsIDForbidden
const DeleteDiagnosticActivationsIDForbiddenCode int = 403

/*
DeleteDiagnosticActivationsIDForbidden Forbidden

swagger:response deleteDiagnosticActivationsIdForbidden
*/
type DeleteDiagnosticActivationsIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Forbidden `json:"body,omitempty"`
}

// NewDeleteDiagnosticActivationsIDForbidden creates DeleteDiagnosticActivationsIDForbidden with default headers values
func NewDeleteDiagnosticActivationsIDForbidden() *DeleteDiagnosticActivationsIDForbidden {

	return &DeleteDiagnosticActivationsIDForbidden{}
}

// WithPayload adds the payload to the delete diagnostic activations Id forbidden response
func (o *DeleteDiagnosticActivationsIDForbidden) WithPayload(payload *models.Forbidden) *DeleteDiagnosticActivationsIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete diagnostic activations Id forbidden response
func (o *DeleteDiagnosticActivationsIDForbidden) SetPayload(payload *models.Forbidden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDiagnosticActivationsIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDiagnosticActivationsIDNotFoundCode is the HTTP code returned for type DeleteDiagnosticActivationsIDNotFound
const DeleteDiagnosticActivationsIDNotFoundCode int = 404

/*
DeleteDiagnosticActivationsIDNotFound Not Found

swagger:response deleteDiagnosticActivationsIdNotFound
*/
type DeleteDiagnosticActivationsIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Notfound `json:"body,omitempty"`
}

// NewDeleteDiagnosticActivationsIDNotFound creates DeleteDiagnosticActivationsIDNotFound with default headers values
func NewDeleteDiagnosticActivationsIDNotFound() *DeleteDiagnosticActivationsIDNotFound {

	return &DeleteDiagnosticActivationsIDNotFound{}
}

// WithPayload adds the payload to the delete diagnostic activations Id not found response
func (o *DeleteDiagnosticActivationsIDNotFound) WithPayload(payload *models.Notfound) *DeleteDiagnosticActivationsIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete diagnostic activations Id not found response
func (o *DeleteDiagnosticActivationsIDNotFound) SetPayload(payload *models.Notfound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDiagnosticActivationsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
DeleteDiagnosticActivationsIDDefault unexpected error

swagger:response deleteDiagnosticActivationsIdDefault
*/
type DeleteDiagnosticActivationsIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDiagnosticActivationsIDDefault creates DeleteDiagnosticActivationsIDDefault with default headers values
func NewDeleteDiagnosticActivationsIDDefault(code int) *DeleteDiagnosticActivationsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteDiagnosticActivationsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete diagnostic activations ID default response
func (o *DeleteDiagnosticActivationsIDDefault) WithStatusCode(code int) *DeleteDiagnosticActivationsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete diagnostic activations ID default response
func (o *DeleteDiagnosticActivationsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete diagnostic activations ID default response
func (o *DeleteDiagnosticActivationsIDDefault) WithPayload(payload *models.Error) *DeleteDiagnosticActivationsIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete diagnostic activations ID default response
func (o *DeleteDiagnosticActivationsIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDiagnosticActivationsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
