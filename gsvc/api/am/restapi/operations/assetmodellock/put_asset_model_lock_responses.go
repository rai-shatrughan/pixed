// Code generated by go-swagger; DO NOT EDIT.

package assetmodellock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// PutAssetModelLockOKCode is the HTTP code returned for type PutAssetModelLockOK
const PutAssetModelLockOKCode int = 200

/*
PutAssetModelLockOK Returns lock state of an asset model applied on a tenant.

swagger:response putAssetModelLockOK
*/
type PutAssetModelLockOK struct {

	/*
	  In: Body
	*/
	Payload *models.ModelLock `json:"body,omitempty"`
}

// NewPutAssetModelLockOK creates PutAssetModelLockOK with default headers values
func NewPutAssetModelLockOK() *PutAssetModelLockOK {

	return &PutAssetModelLockOK{}
}

// WithPayload adds the payload to the put asset model lock o k response
func (o *PutAssetModelLockOK) WithPayload(payload *models.ModelLock) *PutAssetModelLockOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put asset model lock o k response
func (o *PutAssetModelLockOK) SetPayload(payload *models.ModelLock) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAssetModelLockOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAssetModelLockUnauthorizedCode is the HTTP code returned for type PutAssetModelLockUnauthorized
const PutAssetModelLockUnauthorizedCode int = 401

/*
PutAssetModelLockUnauthorized User is not authenticated

swagger:response putAssetModelLockUnauthorized
*/
type PutAssetModelLockUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewPutAssetModelLockUnauthorized creates PutAssetModelLockUnauthorized with default headers values
func NewPutAssetModelLockUnauthorized() *PutAssetModelLockUnauthorized {

	return &PutAssetModelLockUnauthorized{}
}

// WithPayload adds the payload to the put asset model lock unauthorized response
func (o *PutAssetModelLockUnauthorized) WithPayload(payload *models.Errors) *PutAssetModelLockUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put asset model lock unauthorized response
func (o *PutAssetModelLockUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAssetModelLockUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAssetModelLockForbiddenCode is the HTTP code returned for type PutAssetModelLockForbidden
const PutAssetModelLockForbiddenCode int = 403

/*
PutAssetModelLockForbidden User is not authorized for request

swagger:response putAssetModelLockForbidden
*/
type PutAssetModelLockForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewPutAssetModelLockForbidden creates PutAssetModelLockForbidden with default headers values
func NewPutAssetModelLockForbidden() *PutAssetModelLockForbidden {

	return &PutAssetModelLockForbidden{}
}

// WithPayload adds the payload to the put asset model lock forbidden response
func (o *PutAssetModelLockForbidden) WithPayload(payload *models.Errors) *PutAssetModelLockForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put asset model lock forbidden response
func (o *PutAssetModelLockForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAssetModelLockForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAssetModelLockPreconditionFailedCode is the HTTP code returned for type PutAssetModelLockPreconditionFailed
const PutAssetModelLockPreconditionFailedCode int = 412

/*
PutAssetModelLockPreconditionFailed Precondition Failed

swagger:response putAssetModelLockPreconditionFailed
*/
type PutAssetModelLockPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewPutAssetModelLockPreconditionFailed creates PutAssetModelLockPreconditionFailed with default headers values
func NewPutAssetModelLockPreconditionFailed() *PutAssetModelLockPreconditionFailed {

	return &PutAssetModelLockPreconditionFailed{}
}

// WithPayload adds the payload to the put asset model lock precondition failed response
func (o *PutAssetModelLockPreconditionFailed) WithPayload(payload *models.Errors) *PutAssetModelLockPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put asset model lock precondition failed response
func (o *PutAssetModelLockPreconditionFailed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAssetModelLockPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAssetModelLockTooManyRequestsCode is the HTTP code returned for type PutAssetModelLockTooManyRequests
const PutAssetModelLockTooManyRequestsCode int = 429

/*
PutAssetModelLockTooManyRequests Too Many Requests

swagger:response putAssetModelLockTooManyRequests
*/
type PutAssetModelLockTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewPutAssetModelLockTooManyRequests creates PutAssetModelLockTooManyRequests with default headers values
func NewPutAssetModelLockTooManyRequests() *PutAssetModelLockTooManyRequests {

	return &PutAssetModelLockTooManyRequests{}
}

// WithPayload adds the payload to the put asset model lock too many requests response
func (o *PutAssetModelLockTooManyRequests) WithPayload(payload *models.Errors) *PutAssetModelLockTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put asset model lock too many requests response
func (o *PutAssetModelLockTooManyRequests) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAssetModelLockTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutAssetModelLockInternalServerErrorCode is the HTTP code returned for type PutAssetModelLockInternalServerError
const PutAssetModelLockInternalServerErrorCode int = 500

/*
PutAssetModelLockInternalServerError Server error, for more information see errorcode and message

swagger:response putAssetModelLockInternalServerError
*/
type PutAssetModelLockInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewPutAssetModelLockInternalServerError creates PutAssetModelLockInternalServerError with default headers values
func NewPutAssetModelLockInternalServerError() *PutAssetModelLockInternalServerError {

	return &PutAssetModelLockInternalServerError{}
}

// WithPayload adds the payload to the put asset model lock internal server error response
func (o *PutAssetModelLockInternalServerError) WithPayload(payload *models.Errors) *PutAssetModelLockInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put asset model lock internal server error response
func (o *PutAssetModelLockInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutAssetModelLockInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
