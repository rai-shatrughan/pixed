// Code generated by go-swagger; DO NOT EDIT.

package mappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/mc/models"
)

// PostDataPointMappingsCreatedCode is the HTTP code returned for type PostDataPointMappingsCreated
const PostDataPointMappingsCreatedCode int = 201

/*
PostDataPointMappingsCreated Created

swagger:response postDataPointMappingsCreated
*/
type PostDataPointMappingsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Mapping `json:"body,omitempty"`
}

// NewPostDataPointMappingsCreated creates PostDataPointMappingsCreated with default headers values
func NewPostDataPointMappingsCreated() *PostDataPointMappingsCreated {

	return &PostDataPointMappingsCreated{}
}

// WithPayload adds the payload to the post data point mappings created response
func (o *PostDataPointMappingsCreated) WithPayload(payload *models.Mapping) *PostDataPointMappingsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post data point mappings created response
func (o *PostDataPointMappingsCreated) SetPayload(payload *models.Mapping) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDataPointMappingsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostDataPointMappingsBadRequestCode is the HTTP code returned for type PostDataPointMappingsBadRequest
const PostDataPointMappingsBadRequestCode int = 400

/*
PostDataPointMappingsBadRequest Bad Request. Request validations failed.

swagger:response postDataPointMappingsBadRequest
*/
type PostDataPointMappingsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Badrequest `json:"body,omitempty"`
}

// NewPostDataPointMappingsBadRequest creates PostDataPointMappingsBadRequest with default headers values
func NewPostDataPointMappingsBadRequest() *PostDataPointMappingsBadRequest {

	return &PostDataPointMappingsBadRequest{}
}

// WithPayload adds the payload to the post data point mappings bad request response
func (o *PostDataPointMappingsBadRequest) WithPayload(payload *models.Badrequest) *PostDataPointMappingsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post data point mappings bad request response
func (o *PostDataPointMappingsBadRequest) SetPayload(payload *models.Badrequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDataPointMappingsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostDataPointMappingsUnauthorizedCode is the HTTP code returned for type PostDataPointMappingsUnauthorized
const PostDataPointMappingsUnauthorizedCode int = 401

/*
PostDataPointMappingsUnauthorized Unauthorized

swagger:response postDataPointMappingsUnauthorized
*/
type PostDataPointMappingsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewPostDataPointMappingsUnauthorized creates PostDataPointMappingsUnauthorized with default headers values
func NewPostDataPointMappingsUnauthorized() *PostDataPointMappingsUnauthorized {

	return &PostDataPointMappingsUnauthorized{}
}

// WithPayload adds the payload to the post data point mappings unauthorized response
func (o *PostDataPointMappingsUnauthorized) WithPayload(payload *models.Unauthorized) *PostDataPointMappingsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post data point mappings unauthorized response
func (o *PostDataPointMappingsUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDataPointMappingsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostDataPointMappingsForbiddenCode is the HTTP code returned for type PostDataPointMappingsForbidden
const PostDataPointMappingsForbiddenCode int = 403

/*
PostDataPointMappingsForbidden Forbidden

swagger:response postDataPointMappingsForbidden
*/
type PostDataPointMappingsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Forbidden `json:"body,omitempty"`
}

// NewPostDataPointMappingsForbidden creates PostDataPointMappingsForbidden with default headers values
func NewPostDataPointMappingsForbidden() *PostDataPointMappingsForbidden {

	return &PostDataPointMappingsForbidden{}
}

// WithPayload adds the payload to the post data point mappings forbidden response
func (o *PostDataPointMappingsForbidden) WithPayload(payload *models.Forbidden) *PostDataPointMappingsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post data point mappings forbidden response
func (o *PostDataPointMappingsForbidden) SetPayload(payload *models.Forbidden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDataPointMappingsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostDataPointMappingsConflictCode is the HTTP code returned for type PostDataPointMappingsConflict
const PostDataPointMappingsConflictCode int = 409

/*
PostDataPointMappingsConflict Conflict

swagger:response postDataPointMappingsConflict
*/
type PostDataPointMappingsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Conflict `json:"body,omitempty"`
}

// NewPostDataPointMappingsConflict creates PostDataPointMappingsConflict with default headers values
func NewPostDataPointMappingsConflict() *PostDataPointMappingsConflict {

	return &PostDataPointMappingsConflict{}
}

// WithPayload adds the payload to the post data point mappings conflict response
func (o *PostDataPointMappingsConflict) WithPayload(payload *models.Conflict) *PostDataPointMappingsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post data point mappings conflict response
func (o *PostDataPointMappingsConflict) SetPayload(payload *models.Conflict) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDataPointMappingsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostDataPointMappingsDefault unexpected error

swagger:response postDataPointMappingsDefault
*/
type PostDataPointMappingsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostDataPointMappingsDefault creates PostDataPointMappingsDefault with default headers values
func NewPostDataPointMappingsDefault(code int) *PostDataPointMappingsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostDataPointMappingsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post data point mappings default response
func (o *PostDataPointMappingsDefault) WithStatusCode(code int) *PostDataPointMappingsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post data point mappings default response
func (o *PostDataPointMappingsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post data point mappings default response
func (o *PostDataPointMappingsDefault) WithPayload(payload *models.Error) *PostDataPointMappingsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post data point mappings default response
func (o *PostDataPointMappingsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDataPointMappingsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
