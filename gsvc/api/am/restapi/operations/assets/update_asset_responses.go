// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// UpdateAssetOKCode is the HTTP code returned for type UpdateAssetOK
const UpdateAssetOKCode int = 200

/*
UpdateAssetOK The asset is updated

swagger:response updateAssetOK
*/
type UpdateAssetOK struct {

	/*
	  In: Body
	*/
	Payload *models.AssetResourceWithHierarchyPath `json:"body,omitempty"`
}

// NewUpdateAssetOK creates UpdateAssetOK with default headers values
func NewUpdateAssetOK() *UpdateAssetOK {

	return &UpdateAssetOK{}
}

// WithPayload adds the payload to the update asset o k response
func (o *UpdateAssetOK) WithPayload(payload *models.AssetResourceWithHierarchyPath) *UpdateAssetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update asset o k response
func (o *UpdateAssetOK) SetPayload(payload *models.AssetResourceWithHierarchyPath) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAssetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAssetBadRequestCode is the HTTP code returned for type UpdateAssetBadRequest
const UpdateAssetBadRequestCode int = 400

/*
UpdateAssetBadRequest Invalid request

swagger:response updateAssetBadRequest
*/
type UpdateAssetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAssetBadRequest creates UpdateAssetBadRequest with default headers values
func NewUpdateAssetBadRequest() *UpdateAssetBadRequest {

	return &UpdateAssetBadRequest{}
}

// WithPayload adds the payload to the update asset bad request response
func (o *UpdateAssetBadRequest) WithPayload(payload *models.Errors) *UpdateAssetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update asset bad request response
func (o *UpdateAssetBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAssetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAssetUnauthorizedCode is the HTTP code returned for type UpdateAssetUnauthorized
const UpdateAssetUnauthorizedCode int = 401

/*
UpdateAssetUnauthorized User is not authenticated

swagger:response updateAssetUnauthorized
*/
type UpdateAssetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAssetUnauthorized creates UpdateAssetUnauthorized with default headers values
func NewUpdateAssetUnauthorized() *UpdateAssetUnauthorized {

	return &UpdateAssetUnauthorized{}
}

// WithPayload adds the payload to the update asset unauthorized response
func (o *UpdateAssetUnauthorized) WithPayload(payload *models.Errors) *UpdateAssetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update asset unauthorized response
func (o *UpdateAssetUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAssetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAssetForbiddenCode is the HTTP code returned for type UpdateAssetForbidden
const UpdateAssetForbiddenCode int = 403

/*
UpdateAssetForbidden User is not authorized for request

swagger:response updateAssetForbidden
*/
type UpdateAssetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAssetForbidden creates UpdateAssetForbidden with default headers values
func NewUpdateAssetForbidden() *UpdateAssetForbidden {

	return &UpdateAssetForbidden{}
}

// WithPayload adds the payload to the update asset forbidden response
func (o *UpdateAssetForbidden) WithPayload(payload *models.Errors) *UpdateAssetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update asset forbidden response
func (o *UpdateAssetForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAssetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAssetNotFoundCode is the HTTP code returned for type UpdateAssetNotFound
const UpdateAssetNotFoundCode int = 404

/*
UpdateAssetNotFound Asset not found

swagger:response updateAssetNotFound
*/
type UpdateAssetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAssetNotFound creates UpdateAssetNotFound with default headers values
func NewUpdateAssetNotFound() *UpdateAssetNotFound {

	return &UpdateAssetNotFound{}
}

// WithPayload adds the payload to the update asset not found response
func (o *UpdateAssetNotFound) WithPayload(payload *models.Errors) *UpdateAssetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update asset not found response
func (o *UpdateAssetNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAssetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAssetPreconditionFailedCode is the HTTP code returned for type UpdateAssetPreconditionFailed
const UpdateAssetPreconditionFailedCode int = 412

/*
UpdateAssetPreconditionFailed Asset is changed in the background

swagger:response updateAssetPreconditionFailed
*/
type UpdateAssetPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAssetPreconditionFailed creates UpdateAssetPreconditionFailed with default headers values
func NewUpdateAssetPreconditionFailed() *UpdateAssetPreconditionFailed {

	return &UpdateAssetPreconditionFailed{}
}

// WithPayload adds the payload to the update asset precondition failed response
func (o *UpdateAssetPreconditionFailed) WithPayload(payload *models.Errors) *UpdateAssetPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update asset precondition failed response
func (o *UpdateAssetPreconditionFailed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAssetPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAssetInternalServerErrorCode is the HTTP code returned for type UpdateAssetInternalServerError
const UpdateAssetInternalServerErrorCode int = 500

/*
UpdateAssetInternalServerError Server error, for more information see errorcode and message

swagger:response updateAssetInternalServerError
*/
type UpdateAssetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAssetInternalServerError creates UpdateAssetInternalServerError with default headers values
func NewUpdateAssetInternalServerError() *UpdateAssetInternalServerError {

	return &UpdateAssetInternalServerError{}
}

// WithPayload adds the payload to the update asset internal server error response
func (o *UpdateAssetInternalServerError) WithPayload(payload *models.Errors) *UpdateAssetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update asset internal server error response
func (o *UpdateAssetInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAssetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAssetServiceUnavailableCode is the HTTP code returned for type UpdateAssetServiceUnavailable
const UpdateAssetServiceUnavailableCode int = 503

/*
UpdateAssetServiceUnavailable Service unavailable, for more information see errorcode and message

swagger:response updateAssetServiceUnavailable
*/
type UpdateAssetServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateAssetServiceUnavailable creates UpdateAssetServiceUnavailable with default headers values
func NewUpdateAssetServiceUnavailable() *UpdateAssetServiceUnavailable {

	return &UpdateAssetServiceUnavailable{}
}

// WithPayload adds the payload to the update asset service unavailable response
func (o *UpdateAssetServiceUnavailable) WithPayload(payload *models.Errors) *UpdateAssetServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update asset service unavailable response
func (o *UpdateAssetServiceUnavailable) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAssetServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
