// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"gsvc/api/am/models"
)

// AddAssetCreatedCode is the HTTP code returned for type AddAssetCreated
const AddAssetCreatedCode int = 201

/*
AddAssetCreated The asset is created

swagger:response addAssetCreated
*/
type AddAssetCreated struct {
	/*URL of the created resource

	 */
	Location strfmt.URI `json:"Location"`

	/*
	  In: Body
	*/
	Payload *models.AssetResourceWithHierarchyPath `json:"body,omitempty"`
}

// NewAddAssetCreated creates AddAssetCreated with default headers values
func NewAddAssetCreated() *AddAssetCreated {

	return &AddAssetCreated{}
}

// WithLocation adds the location to the add asset created response
func (o *AddAssetCreated) WithLocation(location strfmt.URI) *AddAssetCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the add asset created response
func (o *AddAssetCreated) SetLocation(location strfmt.URI) {
	o.Location = location
}

// WithPayload adds the payload to the add asset created response
func (o *AddAssetCreated) WithPayload(payload *models.AssetResourceWithHierarchyPath) *AddAssetCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add asset created response
func (o *AddAssetCreated) SetPayload(payload *models.AssetResourceWithHierarchyPath) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAssetCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location.String()
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAssetBadRequestCode is the HTTP code returned for type AddAssetBadRequest
const AddAssetBadRequestCode int = 400

/*
AddAssetBadRequest Invalid request

swagger:response addAssetBadRequest
*/
type AddAssetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewAddAssetBadRequest creates AddAssetBadRequest with default headers values
func NewAddAssetBadRequest() *AddAssetBadRequest {

	return &AddAssetBadRequest{}
}

// WithPayload adds the payload to the add asset bad request response
func (o *AddAssetBadRequest) WithPayload(payload *models.Errors) *AddAssetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add asset bad request response
func (o *AddAssetBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAssetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAssetUnauthorizedCode is the HTTP code returned for type AddAssetUnauthorized
const AddAssetUnauthorizedCode int = 401

/*
AddAssetUnauthorized User is not authenticated

swagger:response addAssetUnauthorized
*/
type AddAssetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewAddAssetUnauthorized creates AddAssetUnauthorized with default headers values
func NewAddAssetUnauthorized() *AddAssetUnauthorized {

	return &AddAssetUnauthorized{}
}

// WithPayload adds the payload to the add asset unauthorized response
func (o *AddAssetUnauthorized) WithPayload(payload *models.Errors) *AddAssetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add asset unauthorized response
func (o *AddAssetUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAssetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAssetForbiddenCode is the HTTP code returned for type AddAssetForbidden
const AddAssetForbiddenCode int = 403

/*
AddAssetForbidden User is not authorized for request

swagger:response addAssetForbidden
*/
type AddAssetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewAddAssetForbidden creates AddAssetForbidden with default headers values
func NewAddAssetForbidden() *AddAssetForbidden {

	return &AddAssetForbidden{}
}

// WithPayload adds the payload to the add asset forbidden response
func (o *AddAssetForbidden) WithPayload(payload *models.Errors) *AddAssetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add asset forbidden response
func (o *AddAssetForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAssetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAssetInternalServerErrorCode is the HTTP code returned for type AddAssetInternalServerError
const AddAssetInternalServerErrorCode int = 500

/*
AddAssetInternalServerError Server error, for more information see errorcode and message

swagger:response addAssetInternalServerError
*/
type AddAssetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewAddAssetInternalServerError creates AddAssetInternalServerError with default headers values
func NewAddAssetInternalServerError() *AddAssetInternalServerError {

	return &AddAssetInternalServerError{}
}

// WithPayload adds the payload to the add asset internal server error response
func (o *AddAssetInternalServerError) WithPayload(payload *models.Errors) *AddAssetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add asset internal server error response
func (o *AddAssetInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAssetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAssetServiceUnavailableCode is the HTTP code returned for type AddAssetServiceUnavailable
const AddAssetServiceUnavailableCode int = 503

/*
AddAssetServiceUnavailable Service unavailable, for more information see errorcode and message

swagger:response addAssetServiceUnavailable
*/
type AddAssetServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewAddAssetServiceUnavailable creates AddAssetServiceUnavailable with default headers values
func NewAddAssetServiceUnavailable() *AddAssetServiceUnavailable {

	return &AddAssetServiceUnavailable{}
}

// WithPayload adds the payload to the add asset service unavailable response
func (o *AddAssetServiceUnavailable) WithPayload(payload *models.Errors) *AddAssetServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add asset service unavailable response
func (o *AddAssetServiceUnavailable) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAssetServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}