// Code generated by go-swagger; DO NOT EDIT.

package aspecttype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// DeleteAspectTypeNoContentCode is the HTTP code returned for type DeleteAspectTypeNoContent
const DeleteAspectTypeNoContentCode int = 204

/*
DeleteAspectTypeNoContent Aspect type deletion was successful

swagger:response deleteAspectTypeNoContent
*/
type DeleteAspectTypeNoContent struct {
}

// NewDeleteAspectTypeNoContent creates DeleteAspectTypeNoContent with default headers values
func NewDeleteAspectTypeNoContent() *DeleteAspectTypeNoContent {

	return &DeleteAspectTypeNoContent{}
}

// WriteResponse to the client
func (o *DeleteAspectTypeNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteAspectTypeBadRequestCode is the HTTP code returned for type DeleteAspectTypeBadRequest
const DeleteAspectTypeBadRequestCode int = 400

/*
DeleteAspectTypeBadRequest Invalid request

swagger:response deleteAspectTypeBadRequest
*/
type DeleteAspectTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteAspectTypeBadRequest creates DeleteAspectTypeBadRequest with default headers values
func NewDeleteAspectTypeBadRequest() *DeleteAspectTypeBadRequest {

	return &DeleteAspectTypeBadRequest{}
}

// WithPayload adds the payload to the delete aspect type bad request response
func (o *DeleteAspectTypeBadRequest) WithPayload(payload *models.Errors) *DeleteAspectTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete aspect type bad request response
func (o *DeleteAspectTypeBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAspectTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAspectTypeUnauthorizedCode is the HTTP code returned for type DeleteAspectTypeUnauthorized
const DeleteAspectTypeUnauthorizedCode int = 401

/*
DeleteAspectTypeUnauthorized User is not authenticated

swagger:response deleteAspectTypeUnauthorized
*/
type DeleteAspectTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteAspectTypeUnauthorized creates DeleteAspectTypeUnauthorized with default headers values
func NewDeleteAspectTypeUnauthorized() *DeleteAspectTypeUnauthorized {

	return &DeleteAspectTypeUnauthorized{}
}

// WithPayload adds the payload to the delete aspect type unauthorized response
func (o *DeleteAspectTypeUnauthorized) WithPayload(payload *models.Errors) *DeleteAspectTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete aspect type unauthorized response
func (o *DeleteAspectTypeUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAspectTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAspectTypeForbiddenCode is the HTTP code returned for type DeleteAspectTypeForbidden
const DeleteAspectTypeForbiddenCode int = 403

/*
DeleteAspectTypeForbidden User is not authorized for request

swagger:response deleteAspectTypeForbidden
*/
type DeleteAspectTypeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteAspectTypeForbidden creates DeleteAspectTypeForbidden with default headers values
func NewDeleteAspectTypeForbidden() *DeleteAspectTypeForbidden {

	return &DeleteAspectTypeForbidden{}
}

// WithPayload adds the payload to the delete aspect type forbidden response
func (o *DeleteAspectTypeForbidden) WithPayload(payload *models.Errors) *DeleteAspectTypeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete aspect type forbidden response
func (o *DeleteAspectTypeForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAspectTypeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAspectTypeNotFoundCode is the HTTP code returned for type DeleteAspectTypeNotFound
const DeleteAspectTypeNotFoundCode int = 404

/*
DeleteAspectTypeNotFound AspectType not found

swagger:response deleteAspectTypeNotFound
*/
type DeleteAspectTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteAspectTypeNotFound creates DeleteAspectTypeNotFound with default headers values
func NewDeleteAspectTypeNotFound() *DeleteAspectTypeNotFound {

	return &DeleteAspectTypeNotFound{}
}

// WithPayload adds the payload to the delete aspect type not found response
func (o *DeleteAspectTypeNotFound) WithPayload(payload *models.Errors) *DeleteAspectTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete aspect type not found response
func (o *DeleteAspectTypeNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAspectTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAspectTypePreconditionFailedCode is the HTTP code returned for type DeleteAspectTypePreconditionFailed
const DeleteAspectTypePreconditionFailedCode int = 412

/*
DeleteAspectTypePreconditionFailed AspectType is changed in the background

swagger:response deleteAspectTypePreconditionFailed
*/
type DeleteAspectTypePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteAspectTypePreconditionFailed creates DeleteAspectTypePreconditionFailed with default headers values
func NewDeleteAspectTypePreconditionFailed() *DeleteAspectTypePreconditionFailed {

	return &DeleteAspectTypePreconditionFailed{}
}

// WithPayload adds the payload to the delete aspect type precondition failed response
func (o *DeleteAspectTypePreconditionFailed) WithPayload(payload *models.Errors) *DeleteAspectTypePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete aspect type precondition failed response
func (o *DeleteAspectTypePreconditionFailed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAspectTypePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAspectTypeInternalServerErrorCode is the HTTP code returned for type DeleteAspectTypeInternalServerError
const DeleteAspectTypeInternalServerErrorCode int = 500

/*
DeleteAspectTypeInternalServerError Server error, for more information see errorcode and message

swagger:response deleteAspectTypeInternalServerError
*/
type DeleteAspectTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteAspectTypeInternalServerError creates DeleteAspectTypeInternalServerError with default headers values
func NewDeleteAspectTypeInternalServerError() *DeleteAspectTypeInternalServerError {

	return &DeleteAspectTypeInternalServerError{}
}

// WithPayload adds the payload to the delete aspect type internal server error response
func (o *DeleteAspectTypeInternalServerError) WithPayload(payload *models.Errors) *DeleteAspectTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete aspect type internal server error response
func (o *DeleteAspectTypeInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAspectTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
