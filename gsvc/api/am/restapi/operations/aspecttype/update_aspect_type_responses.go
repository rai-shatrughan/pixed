// Code generated by go-swagger; DO NOT EDIT.

package aspecttype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// UpdateAspectTypeOKCode is the HTTP code returned for type UpdateAspectTypeOK
const UpdateAspectTypeOKCode int = 200

/*
UpdateAspectTypeOK The aspect type has been updated

swagger:response updateAspectTypeOK
*/
type UpdateAspectTypeOK struct {

	/*
	  In: Body
	*/
	Payload *models.AspectTypeResource `json:"body,omitempty"`
}

// NewUpdateAspectTypeOK creates UpdateAspectTypeOK with default headers values
func NewUpdateAspectTypeOK() *UpdateAspectTypeOK {

	return &UpdateAspectTypeOK{}
}

// WithPayload adds the payload to the update aspect type o k response
func (o *UpdateAspectTypeOK) WithPayload(payload *models.AspectTypeResource) *UpdateAspectTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update aspect type o k response
func (o *UpdateAspectTypeOK) SetPayload(payload *models.AspectTypeResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAspectTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAspectTypeBadRequestCode is the HTTP code returned for type UpdateAspectTypeBadRequest
const UpdateAspectTypeBadRequestCode int = 400

/*
UpdateAspectTypeBadRequest Invalid request

swagger:response updateAspectTypeBadRequest
*/
type UpdateAspectTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAspectTypeBadRequest creates UpdateAspectTypeBadRequest with default headers values
func NewUpdateAspectTypeBadRequest() *UpdateAspectTypeBadRequest {

	return &UpdateAspectTypeBadRequest{}
}

// WithPayload adds the payload to the update aspect type bad request response
func (o *UpdateAspectTypeBadRequest) WithPayload(payload *models.Errors) *UpdateAspectTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update aspect type bad request response
func (o *UpdateAspectTypeBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAspectTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAspectTypeUnauthorizedCode is the HTTP code returned for type UpdateAspectTypeUnauthorized
const UpdateAspectTypeUnauthorizedCode int = 401

/*
UpdateAspectTypeUnauthorized User is not authenticated

swagger:response updateAspectTypeUnauthorized
*/
type UpdateAspectTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAspectTypeUnauthorized creates UpdateAspectTypeUnauthorized with default headers values
func NewUpdateAspectTypeUnauthorized() *UpdateAspectTypeUnauthorized {

	return &UpdateAspectTypeUnauthorized{}
}

// WithPayload adds the payload to the update aspect type unauthorized response
func (o *UpdateAspectTypeUnauthorized) WithPayload(payload *models.Errors) *UpdateAspectTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update aspect type unauthorized response
func (o *UpdateAspectTypeUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAspectTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAspectTypeForbiddenCode is the HTTP code returned for type UpdateAspectTypeForbidden
const UpdateAspectTypeForbiddenCode int = 403

/*
UpdateAspectTypeForbidden User is not authorized for request

swagger:response updateAspectTypeForbidden
*/
type UpdateAspectTypeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAspectTypeForbidden creates UpdateAspectTypeForbidden with default headers values
func NewUpdateAspectTypeForbidden() *UpdateAspectTypeForbidden {

	return &UpdateAspectTypeForbidden{}
}

// WithPayload adds the payload to the update aspect type forbidden response
func (o *UpdateAspectTypeForbidden) WithPayload(payload *models.Errors) *UpdateAspectTypeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update aspect type forbidden response
func (o *UpdateAspectTypeForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAspectTypeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAspectTypeNotFoundCode is the HTTP code returned for type UpdateAspectTypeNotFound
const UpdateAspectTypeNotFoundCode int = 404

/*
UpdateAspectTypeNotFound AspectType not found

swagger:response updateAspectTypeNotFound
*/
type UpdateAspectTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAspectTypeNotFound creates UpdateAspectTypeNotFound with default headers values
func NewUpdateAspectTypeNotFound() *UpdateAspectTypeNotFound {

	return &UpdateAspectTypeNotFound{}
}

// WithPayload adds the payload to the update aspect type not found response
func (o *UpdateAspectTypeNotFound) WithPayload(payload *models.Errors) *UpdateAspectTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update aspect type not found response
func (o *UpdateAspectTypeNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAspectTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAspectTypePreconditionFailedCode is the HTTP code returned for type UpdateAspectTypePreconditionFailed
const UpdateAspectTypePreconditionFailedCode int = 412

/*
UpdateAspectTypePreconditionFailed AspectType is changed in the background

swagger:response updateAspectTypePreconditionFailed
*/
type UpdateAspectTypePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAspectTypePreconditionFailed creates UpdateAspectTypePreconditionFailed with default headers values
func NewUpdateAspectTypePreconditionFailed() *UpdateAspectTypePreconditionFailed {

	return &UpdateAspectTypePreconditionFailed{}
}

// WithPayload adds the payload to the update aspect type precondition failed response
func (o *UpdateAspectTypePreconditionFailed) WithPayload(payload *models.Errors) *UpdateAspectTypePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update aspect type precondition failed response
func (o *UpdateAspectTypePreconditionFailed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAspectTypePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAspectTypeInternalServerErrorCode is the HTTP code returned for type UpdateAspectTypeInternalServerError
const UpdateAspectTypeInternalServerErrorCode int = 500

/*
UpdateAspectTypeInternalServerError Server error, for more information see errorcode and message

swagger:response updateAspectTypeInternalServerError
*/
type UpdateAspectTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAspectTypeInternalServerError creates UpdateAspectTypeInternalServerError with default headers values
func NewUpdateAspectTypeInternalServerError() *UpdateAspectTypeInternalServerError {

	return &UpdateAspectTypeInternalServerError{}
}

// WithPayload adds the payload to the update aspect type internal server error response
func (o *UpdateAspectTypeInternalServerError) WithPayload(payload *models.Errors) *UpdateAspectTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update aspect type internal server error response
func (o *UpdateAspectTypeInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAspectTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
