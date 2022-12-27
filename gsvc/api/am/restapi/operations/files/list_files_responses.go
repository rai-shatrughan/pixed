// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// ListFilesOKCode is the HTTP code returned for type ListFilesOK
const ListFilesOKCode int = 200

/*
ListFilesOK Metadata of files visible for the tenant

swagger:response listFilesOK
*/
type ListFilesOK struct {

	/*
	  In: Body
	*/
	Payload *models.FileMetadataListResource `json:"body,omitempty"`
}

// NewListFilesOK creates ListFilesOK with default headers values
func NewListFilesOK() *ListFilesOK {

	return &ListFilesOK{}
}

// WithPayload adds the payload to the list files o k response
func (o *ListFilesOK) WithPayload(payload *models.FileMetadataListResource) *ListFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list files o k response
func (o *ListFilesOK) SetPayload(payload *models.FileMetadataListResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListFilesNotModifiedCode is the HTTP code returned for type ListFilesNotModified
const ListFilesNotModifiedCode int = 304

/*
ListFilesNotModified Resource have not been modified

swagger:response listFilesNotModified
*/
type ListFilesNotModified struct {
}

// NewListFilesNotModified creates ListFilesNotModified with default headers values
func NewListFilesNotModified() *ListFilesNotModified {

	return &ListFilesNotModified{}
}

// WriteResponse to the client
func (o *ListFilesNotModified) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(304)
}

// ListFilesBadRequestCode is the HTTP code returned for type ListFilesBadRequest
const ListFilesBadRequestCode int = 400

/*
ListFilesBadRequest Invalid request

swagger:response listFilesBadRequest
*/
type ListFilesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListFilesBadRequest creates ListFilesBadRequest with default headers values
func NewListFilesBadRequest() *ListFilesBadRequest {

	return &ListFilesBadRequest{}
}

// WithPayload adds the payload to the list files bad request response
func (o *ListFilesBadRequest) WithPayload(payload *models.Errors) *ListFilesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list files bad request response
func (o *ListFilesBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFilesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListFilesUnauthorizedCode is the HTTP code returned for type ListFilesUnauthorized
const ListFilesUnauthorizedCode int = 401

/*
ListFilesUnauthorized User is not authenticated

swagger:response listFilesUnauthorized
*/
type ListFilesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListFilesUnauthorized creates ListFilesUnauthorized with default headers values
func NewListFilesUnauthorized() *ListFilesUnauthorized {

	return &ListFilesUnauthorized{}
}

// WithPayload adds the payload to the list files unauthorized response
func (o *ListFilesUnauthorized) WithPayload(payload *models.Errors) *ListFilesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list files unauthorized response
func (o *ListFilesUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFilesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListFilesForbiddenCode is the HTTP code returned for type ListFilesForbidden
const ListFilesForbiddenCode int = 403

/*
ListFilesForbidden User is not authorized for request

swagger:response listFilesForbidden
*/
type ListFilesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListFilesForbidden creates ListFilesForbidden with default headers values
func NewListFilesForbidden() *ListFilesForbidden {

	return &ListFilesForbidden{}
}

// WithPayload adds the payload to the list files forbidden response
func (o *ListFilesForbidden) WithPayload(payload *models.Errors) *ListFilesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list files forbidden response
func (o *ListFilesForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFilesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListFilesInternalServerErrorCode is the HTTP code returned for type ListFilesInternalServerError
const ListFilesInternalServerErrorCode int = 500

/*
ListFilesInternalServerError Server error, for more information see errorcode and message

swagger:response listFilesInternalServerError
*/
type ListFilesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListFilesInternalServerError creates ListFilesInternalServerError with default headers values
func NewListFilesInternalServerError() *ListFilesInternalServerError {

	return &ListFilesInternalServerError{}
}

// WithPayload adds the payload to the list files internal server error response
func (o *ListFilesInternalServerError) WithPayload(payload *models.Errors) *ListFilesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list files internal server error response
func (o *ListFilesInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFilesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}