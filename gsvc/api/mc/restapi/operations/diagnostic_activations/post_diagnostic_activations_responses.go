// Code generated by go-swagger; DO NOT EDIT.

package diagnostic_activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/mc/models"
)

// PostDiagnosticActivationsCreatedCode is the HTTP code returned for type PostDiagnosticActivationsCreated
const PostDiagnosticActivationsCreatedCode int = 201

/*
PostDiagnosticActivationsCreated Created

swagger:response postDiagnosticActivationsCreated
*/
type PostDiagnosticActivationsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.DiagnosticActivation `json:"body,omitempty"`
}

// NewPostDiagnosticActivationsCreated creates PostDiagnosticActivationsCreated with default headers values
func NewPostDiagnosticActivationsCreated() *PostDiagnosticActivationsCreated {

	return &PostDiagnosticActivationsCreated{}
}

// WithPayload adds the payload to the post diagnostic activations created response
func (o *PostDiagnosticActivationsCreated) WithPayload(payload *models.DiagnosticActivation) *PostDiagnosticActivationsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post diagnostic activations created response
func (o *PostDiagnosticActivationsCreated) SetPayload(payload *models.DiagnosticActivation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDiagnosticActivationsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostDiagnosticActivationsBadRequestCode is the HTTP code returned for type PostDiagnosticActivationsBadRequest
const PostDiagnosticActivationsBadRequestCode int = 400

/*
PostDiagnosticActivationsBadRequest Bad Request

swagger:response postDiagnosticActivationsBadRequest
*/
type PostDiagnosticActivationsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Badrequest `json:"body,omitempty"`
}

// NewPostDiagnosticActivationsBadRequest creates PostDiagnosticActivationsBadRequest with default headers values
func NewPostDiagnosticActivationsBadRequest() *PostDiagnosticActivationsBadRequest {

	return &PostDiagnosticActivationsBadRequest{}
}

// WithPayload adds the payload to the post diagnostic activations bad request response
func (o *PostDiagnosticActivationsBadRequest) WithPayload(payload *models.Badrequest) *PostDiagnosticActivationsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post diagnostic activations bad request response
func (o *PostDiagnosticActivationsBadRequest) SetPayload(payload *models.Badrequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDiagnosticActivationsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostDiagnosticActivationsUnauthorizedCode is the HTTP code returned for type PostDiagnosticActivationsUnauthorized
const PostDiagnosticActivationsUnauthorizedCode int = 401

/*
PostDiagnosticActivationsUnauthorized Unauthorized

swagger:response postDiagnosticActivationsUnauthorized
*/
type PostDiagnosticActivationsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewPostDiagnosticActivationsUnauthorized creates PostDiagnosticActivationsUnauthorized with default headers values
func NewPostDiagnosticActivationsUnauthorized() *PostDiagnosticActivationsUnauthorized {

	return &PostDiagnosticActivationsUnauthorized{}
}

// WithPayload adds the payload to the post diagnostic activations unauthorized response
func (o *PostDiagnosticActivationsUnauthorized) WithPayload(payload *models.Unauthorized) *PostDiagnosticActivationsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post diagnostic activations unauthorized response
func (o *PostDiagnosticActivationsUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDiagnosticActivationsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostDiagnosticActivationsForbiddenCode is the HTTP code returned for type PostDiagnosticActivationsForbidden
const PostDiagnosticActivationsForbiddenCode int = 403

/*
PostDiagnosticActivationsForbidden Forbidden

swagger:response postDiagnosticActivationsForbidden
*/
type PostDiagnosticActivationsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Forbidden `json:"body,omitempty"`
}

// NewPostDiagnosticActivationsForbidden creates PostDiagnosticActivationsForbidden with default headers values
func NewPostDiagnosticActivationsForbidden() *PostDiagnosticActivationsForbidden {

	return &PostDiagnosticActivationsForbidden{}
}

// WithPayload adds the payload to the post diagnostic activations forbidden response
func (o *PostDiagnosticActivationsForbidden) WithPayload(payload *models.Forbidden) *PostDiagnosticActivationsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post diagnostic activations forbidden response
func (o *PostDiagnosticActivationsForbidden) SetPayload(payload *models.Forbidden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDiagnosticActivationsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostDiagnosticActivationsConflictCode is the HTTP code returned for type PostDiagnosticActivationsConflict
const PostDiagnosticActivationsConflictCode int = 409

/*
PostDiagnosticActivationsConflict Resource is already available.

swagger:response postDiagnosticActivationsConflict
*/
type PostDiagnosticActivationsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Conflict `json:"body,omitempty"`
}

// NewPostDiagnosticActivationsConflict creates PostDiagnosticActivationsConflict with default headers values
func NewPostDiagnosticActivationsConflict() *PostDiagnosticActivationsConflict {

	return &PostDiagnosticActivationsConflict{}
}

// WithPayload adds the payload to the post diagnostic activations conflict response
func (o *PostDiagnosticActivationsConflict) WithPayload(payload *models.Conflict) *PostDiagnosticActivationsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post diagnostic activations conflict response
func (o *PostDiagnosticActivationsConflict) SetPayload(payload *models.Conflict) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDiagnosticActivationsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostDiagnosticActivationsDefault unexpected error

swagger:response postDiagnosticActivationsDefault
*/
type PostDiagnosticActivationsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostDiagnosticActivationsDefault creates PostDiagnosticActivationsDefault with default headers values
func NewPostDiagnosticActivationsDefault(code int) *PostDiagnosticActivationsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostDiagnosticActivationsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post diagnostic activations default response
func (o *PostDiagnosticActivationsDefault) WithStatusCode(code int) *PostDiagnosticActivationsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post diagnostic activations default response
func (o *PostDiagnosticActivationsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post diagnostic activations default response
func (o *PostDiagnosticActivationsDefault) WithPayload(payload *models.Error) *PostDiagnosticActivationsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post diagnostic activations default response
func (o *PostDiagnosticActivationsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDiagnosticActivationsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
