// Code generated by go-swagger; DO NOT EDIT.

package assettype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// SaveAssetTypeOKCode is the HTTP code returned for type SaveAssetTypeOK
const SaveAssetTypeOKCode int = 200

/*
SaveAssetTypeOK The asset type has been updated

swagger:response saveAssetTypeOK
*/
type SaveAssetTypeOK struct {

	/*
	  In: Body
	*/
	Payload *models.AssetTypeResource `json:"body,omitempty"`
}

// NewSaveAssetTypeOK creates SaveAssetTypeOK with default headers values
func NewSaveAssetTypeOK() *SaveAssetTypeOK {

	return &SaveAssetTypeOK{}
}

// WithPayload adds the payload to the save asset type o k response
func (o *SaveAssetTypeOK) WithPayload(payload *models.AssetTypeResource) *SaveAssetTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type o k response
func (o *SaveAssetTypeOK) SetPayload(payload *models.AssetTypeResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeCreatedCode is the HTTP code returned for type SaveAssetTypeCreated
const SaveAssetTypeCreatedCode int = 201

/*
SaveAssetTypeCreated The asset type has been created

swagger:response saveAssetTypeCreated
*/
type SaveAssetTypeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.AssetTypeResource `json:"body,omitempty"`
}

// NewSaveAssetTypeCreated creates SaveAssetTypeCreated with default headers values
func NewSaveAssetTypeCreated() *SaveAssetTypeCreated {

	return &SaveAssetTypeCreated{}
}

// WithPayload adds the payload to the save asset type created response
func (o *SaveAssetTypeCreated) WithPayload(payload *models.AssetTypeResource) *SaveAssetTypeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type created response
func (o *SaveAssetTypeCreated) SetPayload(payload *models.AssetTypeResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeBadRequestCode is the HTTP code returned for type SaveAssetTypeBadRequest
const SaveAssetTypeBadRequestCode int = 400

/*
SaveAssetTypeBadRequest Invalid request

swagger:response saveAssetTypeBadRequest
*/
type SaveAssetTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeBadRequest creates SaveAssetTypeBadRequest with default headers values
func NewSaveAssetTypeBadRequest() *SaveAssetTypeBadRequest {

	return &SaveAssetTypeBadRequest{}
}

// WithPayload adds the payload to the save asset type bad request response
func (o *SaveAssetTypeBadRequest) WithPayload(payload *models.Errors) *SaveAssetTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type bad request response
func (o *SaveAssetTypeBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeUnauthorizedCode is the HTTP code returned for type SaveAssetTypeUnauthorized
const SaveAssetTypeUnauthorizedCode int = 401

/*
SaveAssetTypeUnauthorized User is not authenticated

swagger:response saveAssetTypeUnauthorized
*/
type SaveAssetTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeUnauthorized creates SaveAssetTypeUnauthorized with default headers values
func NewSaveAssetTypeUnauthorized() *SaveAssetTypeUnauthorized {

	return &SaveAssetTypeUnauthorized{}
}

// WithPayload adds the payload to the save asset type unauthorized response
func (o *SaveAssetTypeUnauthorized) WithPayload(payload *models.Errors) *SaveAssetTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type unauthorized response
func (o *SaveAssetTypeUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeForbiddenCode is the HTTP code returned for type SaveAssetTypeForbidden
const SaveAssetTypeForbiddenCode int = 403

/*
SaveAssetTypeForbidden User is not authorized for request

swagger:response saveAssetTypeForbidden
*/
type SaveAssetTypeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeForbidden creates SaveAssetTypeForbidden with default headers values
func NewSaveAssetTypeForbidden() *SaveAssetTypeForbidden {

	return &SaveAssetTypeForbidden{}
}

// WithPayload adds the payload to the save asset type forbidden response
func (o *SaveAssetTypeForbidden) WithPayload(payload *models.Errors) *SaveAssetTypeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type forbidden response
func (o *SaveAssetTypeForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypePreconditionFailedCode is the HTTP code returned for type SaveAssetTypePreconditionFailed
const SaveAssetTypePreconditionFailedCode int = 412

/*
SaveAssetTypePreconditionFailed AssetType is changed in the background

swagger:response saveAssetTypePreconditionFailed
*/
type SaveAssetTypePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypePreconditionFailed creates SaveAssetTypePreconditionFailed with default headers values
func NewSaveAssetTypePreconditionFailed() *SaveAssetTypePreconditionFailed {

	return &SaveAssetTypePreconditionFailed{}
}

// WithPayload adds the payload to the save asset type precondition failed response
func (o *SaveAssetTypePreconditionFailed) WithPayload(payload *models.Errors) *SaveAssetTypePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type precondition failed response
func (o *SaveAssetTypePreconditionFailed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveAssetTypeInternalServerErrorCode is the HTTP code returned for type SaveAssetTypeInternalServerError
const SaveAssetTypeInternalServerErrorCode int = 500

/*
SaveAssetTypeInternalServerError Server error, for more information see errorcode and message

swagger:response saveAssetTypeInternalServerError
*/
type SaveAssetTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSaveAssetTypeInternalServerError creates SaveAssetTypeInternalServerError with default headers values
func NewSaveAssetTypeInternalServerError() *SaveAssetTypeInternalServerError {

	return &SaveAssetTypeInternalServerError{}
}

// WithPayload adds the payload to the save asset type internal server error response
func (o *SaveAssetTypeInternalServerError) WithPayload(payload *models.Errors) *SaveAssetTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save asset type internal server error response
func (o *SaveAssetTypeInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveAssetTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
