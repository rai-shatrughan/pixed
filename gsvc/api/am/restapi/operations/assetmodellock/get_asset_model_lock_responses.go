// Code generated by go-swagger; DO NOT EDIT.

package assetmodellock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gsvc/api/am/models"
)

// GetAssetModelLockOKCode is the HTTP code returned for type GetAssetModelLockOK
const GetAssetModelLockOKCode int = 200

/*
GetAssetModelLockOK Returns lock state of an asset model and associated jobs.

swagger:response getAssetModelLockOK
*/
type GetAssetModelLockOK struct {

	/*
	  In: Body
	*/
	Payload *models.AssetModelLock `json:"body,omitempty"`
}

// NewGetAssetModelLockOK creates GetAssetModelLockOK with default headers values
func NewGetAssetModelLockOK() *GetAssetModelLockOK {

	return &GetAssetModelLockOK{}
}

// WithPayload adds the payload to the get asset model lock o k response
func (o *GetAssetModelLockOK) WithPayload(payload *models.AssetModelLock) *GetAssetModelLockOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get asset model lock o k response
func (o *GetAssetModelLockOK) SetPayload(payload *models.AssetModelLock) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAssetModelLockOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAssetModelLockUnauthorizedCode is the HTTP code returned for type GetAssetModelLockUnauthorized
const GetAssetModelLockUnauthorizedCode int = 401

/*
GetAssetModelLockUnauthorized User is not authenticated

swagger:response getAssetModelLockUnauthorized
*/
type GetAssetModelLockUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAssetModelLockUnauthorized creates GetAssetModelLockUnauthorized with default headers values
func NewGetAssetModelLockUnauthorized() *GetAssetModelLockUnauthorized {

	return &GetAssetModelLockUnauthorized{}
}

// WithPayload adds the payload to the get asset model lock unauthorized response
func (o *GetAssetModelLockUnauthorized) WithPayload(payload *models.Errors) *GetAssetModelLockUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get asset model lock unauthorized response
func (o *GetAssetModelLockUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAssetModelLockUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAssetModelLockForbiddenCode is the HTTP code returned for type GetAssetModelLockForbidden
const GetAssetModelLockForbiddenCode int = 403

/*
GetAssetModelLockForbidden User is not authorized for request

swagger:response getAssetModelLockForbidden
*/
type GetAssetModelLockForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAssetModelLockForbidden creates GetAssetModelLockForbidden with default headers values
func NewGetAssetModelLockForbidden() *GetAssetModelLockForbidden {

	return &GetAssetModelLockForbidden{}
}

// WithPayload adds the payload to the get asset model lock forbidden response
func (o *GetAssetModelLockForbidden) WithPayload(payload *models.Errors) *GetAssetModelLockForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get asset model lock forbidden response
func (o *GetAssetModelLockForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAssetModelLockForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAssetModelLockTooManyRequestsCode is the HTTP code returned for type GetAssetModelLockTooManyRequests
const GetAssetModelLockTooManyRequestsCode int = 429

/*
GetAssetModelLockTooManyRequests Too Many Requests

swagger:response getAssetModelLockTooManyRequests
*/
type GetAssetModelLockTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAssetModelLockTooManyRequests creates GetAssetModelLockTooManyRequests with default headers values
func NewGetAssetModelLockTooManyRequests() *GetAssetModelLockTooManyRequests {

	return &GetAssetModelLockTooManyRequests{}
}

// WithPayload adds the payload to the get asset model lock too many requests response
func (o *GetAssetModelLockTooManyRequests) WithPayload(payload *models.Errors) *GetAssetModelLockTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get asset model lock too many requests response
func (o *GetAssetModelLockTooManyRequests) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAssetModelLockTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAssetModelLockInternalServerErrorCode is the HTTP code returned for type GetAssetModelLockInternalServerError
const GetAssetModelLockInternalServerErrorCode int = 500

/*
GetAssetModelLockInternalServerError Server error, for more information see errorcode and message

swagger:response getAssetModelLockInternalServerError
*/
type GetAssetModelLockInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetAssetModelLockInternalServerError creates GetAssetModelLockInternalServerError with default headers values
func NewGetAssetModelLockInternalServerError() *GetAssetModelLockInternalServerError {

	return &GetAssetModelLockInternalServerError{}
}

// WithPayload adds the payload to the get asset model lock internal server error response
func (o *GetAssetModelLockInternalServerError) WithPayload(payload *models.Errors) *GetAssetModelLockInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get asset model lock internal server error response
func (o *GetAssetModelLockInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAssetModelLockInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
