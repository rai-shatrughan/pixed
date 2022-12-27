// Code generated by go-swagger; DO NOT EDIT.

package aspecttype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// GetAspectTypeOKCode is the HTTP code returned for type GetAspectTypeOK
const GetAspectTypeOKCode int = 200

/*
GetAspectTypeOK Returns the aspect type

swagger:response getAspectTypeOK
*/
type GetAspectTypeOK struct {

	/*
	  In: Body
	*/
	Payload *models.AspectTypeResource `json:"body,omitempty"`
}

// NewGetAspectTypeOK creates GetAspectTypeOK with default headers values
func NewGetAspectTypeOK() *GetAspectTypeOK {

	return &GetAspectTypeOK{}
}

// WithPayload adds the payload to the get aspect type o k response
func (o *GetAspectTypeOK) WithPayload(payload *models.AspectTypeResource) *GetAspectTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get aspect type o k response
func (o *GetAspectTypeOK) SetPayload(payload *models.AspectTypeResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAspectTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAspectTypeNotModifiedCode is the HTTP code returned for type GetAspectTypeNotModified
const GetAspectTypeNotModifiedCode int = 304

/*
GetAspectTypeNotModified AspectType not changed

swagger:response getAspectTypeNotModified
*/
type GetAspectTypeNotModified struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAspectTypeNotModified creates GetAspectTypeNotModified with default headers values
func NewGetAspectTypeNotModified() *GetAspectTypeNotModified {

	return &GetAspectTypeNotModified{}
}

// WithPayload adds the payload to the get aspect type not modified response
func (o *GetAspectTypeNotModified) WithPayload(payload *models.Errors) *GetAspectTypeNotModified {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get aspect type not modified response
func (o *GetAspectTypeNotModified) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAspectTypeNotModified) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(304)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAspectTypeBadRequestCode is the HTTP code returned for type GetAspectTypeBadRequest
const GetAspectTypeBadRequestCode int = 400

/*
GetAspectTypeBadRequest Invalid request

swagger:response getAspectTypeBadRequest
*/
type GetAspectTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAspectTypeBadRequest creates GetAspectTypeBadRequest with default headers values
func NewGetAspectTypeBadRequest() *GetAspectTypeBadRequest {

	return &GetAspectTypeBadRequest{}
}

// WithPayload adds the payload to the get aspect type bad request response
func (o *GetAspectTypeBadRequest) WithPayload(payload *models.Errors) *GetAspectTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get aspect type bad request response
func (o *GetAspectTypeBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAspectTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAspectTypeUnauthorizedCode is the HTTP code returned for type GetAspectTypeUnauthorized
const GetAspectTypeUnauthorizedCode int = 401

/*
GetAspectTypeUnauthorized User is not authenticated

swagger:response getAspectTypeUnauthorized
*/
type GetAspectTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAspectTypeUnauthorized creates GetAspectTypeUnauthorized with default headers values
func NewGetAspectTypeUnauthorized() *GetAspectTypeUnauthorized {

	return &GetAspectTypeUnauthorized{}
}

// WithPayload adds the payload to the get aspect type unauthorized response
func (o *GetAspectTypeUnauthorized) WithPayload(payload *models.Errors) *GetAspectTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get aspect type unauthorized response
func (o *GetAspectTypeUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAspectTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAspectTypeForbiddenCode is the HTTP code returned for type GetAspectTypeForbidden
const GetAspectTypeForbiddenCode int = 403

/*
GetAspectTypeForbidden User is not authorized for request

swagger:response getAspectTypeForbidden
*/
type GetAspectTypeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAspectTypeForbidden creates GetAspectTypeForbidden with default headers values
func NewGetAspectTypeForbidden() *GetAspectTypeForbidden {

	return &GetAspectTypeForbidden{}
}

// WithPayload adds the payload to the get aspect type forbidden response
func (o *GetAspectTypeForbidden) WithPayload(payload *models.Errors) *GetAspectTypeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get aspect type forbidden response
func (o *GetAspectTypeForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAspectTypeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAspectTypeNotFoundCode is the HTTP code returned for type GetAspectTypeNotFound
const GetAspectTypeNotFoundCode int = 404

/*
GetAspectTypeNotFound AspectType not found

swagger:response getAspectTypeNotFound
*/
type GetAspectTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAspectTypeNotFound creates GetAspectTypeNotFound with default headers values
func NewGetAspectTypeNotFound() *GetAspectTypeNotFound {

	return &GetAspectTypeNotFound{}
}

// WithPayload adds the payload to the get aspect type not found response
func (o *GetAspectTypeNotFound) WithPayload(payload *models.Errors) *GetAspectTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get aspect type not found response
func (o *GetAspectTypeNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAspectTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAspectTypeInternalServerErrorCode is the HTTP code returned for type GetAspectTypeInternalServerError
const GetAspectTypeInternalServerErrorCode int = 500

/*
GetAspectTypeInternalServerError Server error, for more information see errorcode and message

swagger:response getAspectTypeInternalServerError
*/
type GetAspectTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAspectTypeInternalServerError creates GetAspectTypeInternalServerError with default headers values
func NewGetAspectTypeInternalServerError() *GetAspectTypeInternalServerError {

	return &GetAspectTypeInternalServerError{}
}

// WithPayload adds the payload to the get aspect type internal server error response
func (o *GetAspectTypeInternalServerError) WithPayload(payload *models.Errors) *GetAspectTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get aspect type internal server error response
func (o *GetAspectTypeInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAspectTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}