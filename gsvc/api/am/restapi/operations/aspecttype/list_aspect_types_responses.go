// Code generated by go-swagger; DO NOT EDIT.

package aspecttype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// ListAspectTypesOKCode is the HTTP code returned for type ListAspectTypesOK
const ListAspectTypesOKCode int = 200

/*
ListAspectTypesOK Array of aspect types matched the tenant scope and filter criterias

swagger:response listAspectTypesOK
*/
type ListAspectTypesOK struct {
	/*Shallow ETag of the resource

	 */
	ETag string `json:"ETag"`

	/*
	  In: Body
	*/
	Payload *models.AspectTypeListResource `json:"body,omitempty"`
}

// NewListAspectTypesOK creates ListAspectTypesOK with default headers values
func NewListAspectTypesOK() *ListAspectTypesOK {

	return &ListAspectTypesOK{}
}

// WithETag adds the eTag to the list aspect types o k response
func (o *ListAspectTypesOK) WithETag(eTag string) *ListAspectTypesOK {
	o.ETag = eTag
	return o
}

// SetETag sets the eTag to the list aspect types o k response
func (o *ListAspectTypesOK) SetETag(eTag string) {
	o.ETag = eTag
}

// WithPayload adds the payload to the list aspect types o k response
func (o *ListAspectTypesOK) WithPayload(payload *models.AspectTypeListResource) *ListAspectTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list aspect types o k response
func (o *ListAspectTypesOK) SetPayload(payload *models.AspectTypeListResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAspectTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header ETag

	eTag := o.ETag
	if eTag != "" {
		rw.Header().Set("ETag", eTag)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAspectTypesNotModifiedCode is the HTTP code returned for type ListAspectTypesNotModified
const ListAspectTypesNotModifiedCode int = 304

/*
ListAspectTypesNotModified Aspect-types page not changed

swagger:response listAspectTypesNotModified
*/
type ListAspectTypesNotModified struct {
}

// NewListAspectTypesNotModified creates ListAspectTypesNotModified with default headers values
func NewListAspectTypesNotModified() *ListAspectTypesNotModified {

	return &ListAspectTypesNotModified{}
}

// WriteResponse to the client
func (o *ListAspectTypesNotModified) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(304)
}

// ListAspectTypesBadRequestCode is the HTTP code returned for type ListAspectTypesBadRequest
const ListAspectTypesBadRequestCode int = 400

/*
ListAspectTypesBadRequest Invalid request

swagger:response listAspectTypesBadRequest
*/
type ListAspectTypesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListAspectTypesBadRequest creates ListAspectTypesBadRequest with default headers values
func NewListAspectTypesBadRequest() *ListAspectTypesBadRequest {

	return &ListAspectTypesBadRequest{}
}

// WithPayload adds the payload to the list aspect types bad request response
func (o *ListAspectTypesBadRequest) WithPayload(payload *models.Errors) *ListAspectTypesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list aspect types bad request response
func (o *ListAspectTypesBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAspectTypesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAspectTypesUnauthorizedCode is the HTTP code returned for type ListAspectTypesUnauthorized
const ListAspectTypesUnauthorizedCode int = 401

/*
ListAspectTypesUnauthorized User is not authenticated

swagger:response listAspectTypesUnauthorized
*/
type ListAspectTypesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListAspectTypesUnauthorized creates ListAspectTypesUnauthorized with default headers values
func NewListAspectTypesUnauthorized() *ListAspectTypesUnauthorized {

	return &ListAspectTypesUnauthorized{}
}

// WithPayload adds the payload to the list aspect types unauthorized response
func (o *ListAspectTypesUnauthorized) WithPayload(payload *models.Errors) *ListAspectTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list aspect types unauthorized response
func (o *ListAspectTypesUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAspectTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAspectTypesForbiddenCode is the HTTP code returned for type ListAspectTypesForbidden
const ListAspectTypesForbiddenCode int = 403

/*
ListAspectTypesForbidden User is not authorized for request

swagger:response listAspectTypesForbidden
*/
type ListAspectTypesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListAspectTypesForbidden creates ListAspectTypesForbidden with default headers values
func NewListAspectTypesForbidden() *ListAspectTypesForbidden {

	return &ListAspectTypesForbidden{}
}

// WithPayload adds the payload to the list aspect types forbidden response
func (o *ListAspectTypesForbidden) WithPayload(payload *models.Errors) *ListAspectTypesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list aspect types forbidden response
func (o *ListAspectTypesForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAspectTypesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAspectTypesInternalServerErrorCode is the HTTP code returned for type ListAspectTypesInternalServerError
const ListAspectTypesInternalServerErrorCode int = 500

/*
ListAspectTypesInternalServerError Server error, for more information see errorcode and message

swagger:response listAspectTypesInternalServerError
*/
type ListAspectTypesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListAspectTypesInternalServerError creates ListAspectTypesInternalServerError with default headers values
func NewListAspectTypesInternalServerError() *ListAspectTypesInternalServerError {

	return &ListAspectTypesInternalServerError{}
}

// WithPayload adds the payload to the list aspect types internal server error response
func (o *ListAspectTypesInternalServerError) WithPayload(payload *models.Errors) *ListAspectTypesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list aspect types internal server error response
func (o *ListAspectTypesInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAspectTypesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
