// Code generated by go-swagger; DO NOT EDIT.

package assettype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// SaveAssetTypeFileAssignmentOKCode is the HTTP code returned for type SaveAssetTypeFileAssignmentOK
const SaveAssetTypeFileAssignmentOKCode int = 200

/*
SaveAssetTypeFileAssignmentOK The assignment has been updated or a new assignment has been added

swagger:response saveAssetTypeFileAssignmentOK
*/
type SaveAssetTypeFileAssignmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.AssetTypeResource `json:"body,omitempty"`
}

// NewSaveAssetTypeFileAssignmentOK creates SaveAssetTypeFileAssignmentOK with default headers values
func NewSaveAssetTypeFileAssignmentOK() *SaveAssetTypeFileAssignmentOK {

	return &SaveAssetTypeFileAssignmentOK{}
}

// WithPayload adds the payload to the save asset type file assignment o k response
func (o *SaveAssetTypeFileAssignmentOK) WithPayload(payload *models.AssetTypeResource) *SaveAssetTypeFileAssignmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type file assignment o k response
func (o *SaveAssetTypeFileAssignmentOK) SetPayload(payload *models.AssetTypeResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeFileAssignmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeFileAssignmentUnauthorizedCode is the HTTP code returned for type SaveAssetTypeFileAssignmentUnauthorized
const SaveAssetTypeFileAssignmentUnauthorizedCode int = 401

/*
SaveAssetTypeFileAssignmentUnauthorized User is not authenticated

swagger:response saveAssetTypeFileAssignmentUnauthorized
*/
type SaveAssetTypeFileAssignmentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeFileAssignmentUnauthorized creates SaveAssetTypeFileAssignmentUnauthorized with default headers values
func NewSaveAssetTypeFileAssignmentUnauthorized() *SaveAssetTypeFileAssignmentUnauthorized {

	return &SaveAssetTypeFileAssignmentUnauthorized{}
}

// WithPayload adds the payload to the save asset type file assignment unauthorized response
func (o *SaveAssetTypeFileAssignmentUnauthorized) WithPayload(payload *models.Errors) *SaveAssetTypeFileAssignmentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type file assignment unauthorized response
func (o *SaveAssetTypeFileAssignmentUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeFileAssignmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeFileAssignmentForbiddenCode is the HTTP code returned for type SaveAssetTypeFileAssignmentForbidden
const SaveAssetTypeFileAssignmentForbiddenCode int = 403

/*
SaveAssetTypeFileAssignmentForbidden User is not authorized for request

swagger:response saveAssetTypeFileAssignmentForbidden
*/
type SaveAssetTypeFileAssignmentForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeFileAssignmentForbidden creates SaveAssetTypeFileAssignmentForbidden with default headers values
func NewSaveAssetTypeFileAssignmentForbidden() *SaveAssetTypeFileAssignmentForbidden {

	return &SaveAssetTypeFileAssignmentForbidden{}
}

// WithPayload adds the payload to the save asset type file assignment forbidden response
func (o *SaveAssetTypeFileAssignmentForbidden) WithPayload(payload *models.Errors) *SaveAssetTypeFileAssignmentForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type file assignment forbidden response
func (o *SaveAssetTypeFileAssignmentForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeFileAssignmentForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeFileAssignmentNotFoundCode is the HTTP code returned for type SaveAssetTypeFileAssignmentNotFound
const SaveAssetTypeFileAssignmentNotFoundCode int = 404

/*
SaveAssetTypeFileAssignmentNotFound Asset type not found

swagger:response saveAssetTypeFileAssignmentNotFound
*/
type SaveAssetTypeFileAssignmentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeFileAssignmentNotFound creates SaveAssetTypeFileAssignmentNotFound with default headers values
func NewSaveAssetTypeFileAssignmentNotFound() *SaveAssetTypeFileAssignmentNotFound {

	return &SaveAssetTypeFileAssignmentNotFound{}
}

// WithPayload adds the payload to the save asset type file assignment not found response
func (o *SaveAssetTypeFileAssignmentNotFound) WithPayload(payload *models.Errors) *SaveAssetTypeFileAssignmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type file assignment not found response
func (o *SaveAssetTypeFileAssignmentNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeFileAssignmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeFileAssignmentPreconditionFailedCode is the HTTP code returned for type SaveAssetTypeFileAssignmentPreconditionFailed
const SaveAssetTypeFileAssignmentPreconditionFailedCode int = 412

/*
SaveAssetTypeFileAssignmentPreconditionFailed Asset type or the file assignment has changed in the background

swagger:response saveAssetTypeFileAssignmentPreconditionFailed
*/
type SaveAssetTypeFileAssignmentPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeFileAssignmentPreconditionFailed creates SaveAssetTypeFileAssignmentPreconditionFailed with default headers values
func NewSaveAssetTypeFileAssignmentPreconditionFailed() *SaveAssetTypeFileAssignmentPreconditionFailed {

	return &SaveAssetTypeFileAssignmentPreconditionFailed{}
}

// WithPayload adds the payload to the save asset type file assignment precondition failed response
func (o *SaveAssetTypeFileAssignmentPreconditionFailed) WithPayload(payload *models.Errors) *SaveAssetTypeFileAssignmentPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type file assignment precondition failed response
func (o *SaveAssetTypeFileAssignmentPreconditionFailed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeFileAssignmentPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeFileAssignmentInternalServerErrorCode is the HTTP code returned for type SaveAssetTypeFileAssignmentInternalServerError
const SaveAssetTypeFileAssignmentInternalServerErrorCode int = 500

/*
SaveAssetTypeFileAssignmentInternalServerError Server error, for more information see errorcode and message

swagger:response saveAssetTypeFileAssignmentInternalServerError
*/
type SaveAssetTypeFileAssignmentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeFileAssignmentInternalServerError creates SaveAssetTypeFileAssignmentInternalServerError with default headers values
func NewSaveAssetTypeFileAssignmentInternalServerError() *SaveAssetTypeFileAssignmentInternalServerError {

	return &SaveAssetTypeFileAssignmentInternalServerError{}
}

// WithPayload adds the payload to the save asset type file assignment internal server error response
func (o *SaveAssetTypeFileAssignmentInternalServerError) WithPayload(payload *models.Errors) *SaveAssetTypeFileAssignmentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type file assignment internal server error response
func (o *SaveAssetTypeFileAssignmentInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeFileAssignmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}