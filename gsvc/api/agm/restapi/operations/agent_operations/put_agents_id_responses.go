// Code generated by go-swagger; DO NOT EDIT.

package agent_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/agm/models"
)

// PutAgentsIDOKCode is the HTTP code returned for type PutAgentsIDOK
const PutAgentsIDOKCode int = 200

/*
PutAgentsIDOK OK

swagger:response putAgentsIdOK
*/
type PutAgentsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Agent `json:"body,omitempty"`
}

// NewPutAgentsIDOK creates PutAgentsIDOK with default headers values
func NewPutAgentsIDOK() *PutAgentsIDOK {

	return &PutAgentsIDOK{}
}

// WithPayload adds the payload to the put agents Id o k response
func (o *PutAgentsIDOK) WithPayload(payload *models.Agent) *PutAgentsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put agents Id o k response
func (o *PutAgentsIDOK) SetPayload(payload *models.Agent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAgentsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAgentsIDBadRequestCode is the HTTP code returned for type PutAgentsIDBadRequest
const PutAgentsIDBadRequestCode int = 400

/*
PutAgentsIDBadRequest Bad Request

swagger:response putAgentsIdBadRequest
*/
type PutAgentsIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Badrequest `json:"body,omitempty"`
}

// NewPutAgentsIDBadRequest creates PutAgentsIDBadRequest with default headers values
func NewPutAgentsIDBadRequest() *PutAgentsIDBadRequest {

	return &PutAgentsIDBadRequest{}
}

// WithPayload adds the payload to the put agents Id bad request response
func (o *PutAgentsIDBadRequest) WithPayload(payload *models.Badrequest) *PutAgentsIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put agents Id bad request response
func (o *PutAgentsIDBadRequest) SetPayload(payload *models.Badrequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAgentsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAgentsIDUnauthorizedCode is the HTTP code returned for type PutAgentsIDUnauthorized
const PutAgentsIDUnauthorizedCode int = 401

/*
PutAgentsIDUnauthorized Unauthorized

swagger:response putAgentsIdUnauthorized
*/
type PutAgentsIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewPutAgentsIDUnauthorized creates PutAgentsIDUnauthorized with default headers values
func NewPutAgentsIDUnauthorized() *PutAgentsIDUnauthorized {

	return &PutAgentsIDUnauthorized{}
}

// WithPayload adds the payload to the put agents Id unauthorized response
func (o *PutAgentsIDUnauthorized) WithPayload(payload *models.Unauthorized) *PutAgentsIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put agents Id unauthorized response
func (o *PutAgentsIDUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAgentsIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAgentsIDForbiddenCode is the HTTP code returned for type PutAgentsIDForbidden
const PutAgentsIDForbiddenCode int = 403

/*
PutAgentsIDForbidden Forbidden

swagger:response putAgentsIdForbidden
*/
type PutAgentsIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Forbidden `json:"body,omitempty"`
}

// NewPutAgentsIDForbidden creates PutAgentsIDForbidden with default headers values
func NewPutAgentsIDForbidden() *PutAgentsIDForbidden {

	return &PutAgentsIDForbidden{}
}

// WithPayload adds the payload to the put agents Id forbidden response
func (o *PutAgentsIDForbidden) WithPayload(payload *models.Forbidden) *PutAgentsIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put agents Id forbidden response
func (o *PutAgentsIDForbidden) SetPayload(payload *models.Forbidden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAgentsIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAgentsIDNotFoundCode is the HTTP code returned for type PutAgentsIDNotFound
const PutAgentsIDNotFoundCode int = 404

/*
PutAgentsIDNotFound Not Found

swagger:response putAgentsIdNotFound
*/
type PutAgentsIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Notfound `json:"body,omitempty"`
}

// NewPutAgentsIDNotFound creates PutAgentsIDNotFound with default headers values
func NewPutAgentsIDNotFound() *PutAgentsIDNotFound {

	return &PutAgentsIDNotFound{}
}

// WithPayload adds the payload to the put agents Id not found response
func (o *PutAgentsIDNotFound) WithPayload(payload *models.Notfound) *PutAgentsIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put agents Id not found response
func (o *PutAgentsIDNotFound) SetPayload(payload *models.Notfound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAgentsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAgentsIDConflictCode is the HTTP code returned for type PutAgentsIDConflict
const PutAgentsIDConflictCode int = 409

/*
PutAgentsIDConflict Resource is already available.

swagger:response putAgentsIdConflict
*/
type PutAgentsIDConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Conflict `json:"body,omitempty"`
}

// NewPutAgentsIDConflict creates PutAgentsIDConflict with default headers values
func NewPutAgentsIDConflict() *PutAgentsIDConflict {

	return &PutAgentsIDConflict{}
}

// WithPayload adds the payload to the put agents Id conflict response
func (o *PutAgentsIDConflict) WithPayload(payload *models.Conflict) *PutAgentsIDConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put agents Id conflict response
func (o *PutAgentsIDConflict) SetPayload(payload *models.Conflict) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAgentsIDConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAgentsIDPreconditionFailedCode is the HTTP code returned for type PutAgentsIDPreconditionFailed
const PutAgentsIDPreconditionFailedCode int = 412

/*
PutAgentsIDPreconditionFailed Precondition Failed

swagger:response putAgentsIdPreconditionFailed
*/
type PutAgentsIDPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Preconditionfailed `json:"body,omitempty"`
}

// NewPutAgentsIDPreconditionFailed creates PutAgentsIDPreconditionFailed with default headers values
func NewPutAgentsIDPreconditionFailed() *PutAgentsIDPreconditionFailed {

	return &PutAgentsIDPreconditionFailed{}
}

// WithPayload adds the payload to the put agents Id precondition failed response
func (o *PutAgentsIDPreconditionFailed) WithPayload(payload *models.Preconditionfailed) *PutAgentsIDPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put agents Id precondition failed response
func (o *PutAgentsIDPreconditionFailed) SetPayload(payload *models.Preconditionfailed) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAgentsIDPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAgentsIDInternalServerErrorCode is the HTTP code returned for type PutAgentsIDInternalServerError
const PutAgentsIDInternalServerErrorCode int = 500

/*
PutAgentsIDInternalServerError unexpected error

swagger:response putAgentsIdInternalServerError
*/
type PutAgentsIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutAgentsIDInternalServerError creates PutAgentsIDInternalServerError with default headers values
func NewPutAgentsIDInternalServerError() *PutAgentsIDInternalServerError {

	return &PutAgentsIDInternalServerError{}
}

// WithPayload adds the payload to the put agents Id internal server error response
func (o *PutAgentsIDInternalServerError) WithPayload(payload *models.Error) *PutAgentsIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put agents Id internal server error response
func (o *PutAgentsIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAgentsIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PutAgentsIDDefault Other error with any status code and response body format.

swagger:response putAgentsIdDefault
*/
type PutAgentsIDDefault struct {
	_statusCode int
}

// NewPutAgentsIDDefault creates PutAgentsIDDefault with default headers values
func NewPutAgentsIDDefault(code int) *PutAgentsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PutAgentsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put agents ID default response
func (o *PutAgentsIDDefault) WithStatusCode(code int) *PutAgentsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put agents ID default response
func (o *PutAgentsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PutAgentsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}