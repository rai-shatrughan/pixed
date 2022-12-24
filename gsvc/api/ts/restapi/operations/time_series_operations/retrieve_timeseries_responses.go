// Code generated by go-swagger; DO NOT EDIT.

package time_series_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"gsvc/api/ts/models"
)

// RetrieveTimeseriesOKCode is the HTTP code returned for type RetrieveTimeseriesOK
const RetrieveTimeseriesOKCode int = 200

/*
RetrieveTimeseriesOK array of time series

swagger:response retrieveTimeseriesOK
*/
type RetrieveTimeseriesOK struct {
	/*Contains a link to the next page, if there is a next page

	 */
	Link string `json:"link"`

	/*
	  In: Body
	*/
	Payload []*models.TimeSeriesDataItem `json:"body,omitempty"`
}

// NewRetrieveTimeseriesOK creates RetrieveTimeseriesOK with default headers values
func NewRetrieveTimeseriesOK() *RetrieveTimeseriesOK {

	return &RetrieveTimeseriesOK{}
}

// WithLink adds the link to the retrieve timeseries o k response
func (o *RetrieveTimeseriesOK) WithLink(link string) *RetrieveTimeseriesOK {
	o.Link = link
	return o
}

// SetLink sets the link to the retrieve timeseries o k response
func (o *RetrieveTimeseriesOK) SetLink(link string) {
	o.Link = link
}

// WithPayload adds the payload to the retrieve timeseries o k response
func (o *RetrieveTimeseriesOK) WithPayload(payload []*models.TimeSeriesDataItem) *RetrieveTimeseriesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve timeseries o k response
func (o *RetrieveTimeseriesOK) SetPayload(payload []*models.TimeSeriesDataItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTimeseriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header link

	link := o.Link
	if link != "" {
		rw.Header().Set("link", link)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.TimeSeriesDataItem, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// RetrieveTimeseriesBadRequestCode is the HTTP code returned for type RetrieveTimeseriesBadRequest
const RetrieveTimeseriesBadRequestCode int = 400

/*
RetrieveTimeseriesBadRequest bad request

swagger:response retrieveTimeseriesBadRequest
*/
type RetrieveTimeseriesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequest `json:"body,omitempty"`
}

// NewRetrieveTimeseriesBadRequest creates RetrieveTimeseriesBadRequest with default headers values
func NewRetrieveTimeseriesBadRequest() *RetrieveTimeseriesBadRequest {

	return &RetrieveTimeseriesBadRequest{}
}

// WithPayload adds the payload to the retrieve timeseries bad request response
func (o *RetrieveTimeseriesBadRequest) WithPayload(payload *models.BadRequest) *RetrieveTimeseriesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve timeseries bad request response
func (o *RetrieveTimeseriesBadRequest) SetPayload(payload *models.BadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTimeseriesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveTimeseriesUnauthorizedCode is the HTTP code returned for type RetrieveTimeseriesUnauthorized
const RetrieveTimeseriesUnauthorizedCode int = 401

/*
RetrieveTimeseriesUnauthorized unauthorized

swagger:response retrieveTimeseriesUnauthorized
*/
type RetrieveTimeseriesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewRetrieveTimeseriesUnauthorized creates RetrieveTimeseriesUnauthorized with default headers values
func NewRetrieveTimeseriesUnauthorized() *RetrieveTimeseriesUnauthorized {

	return &RetrieveTimeseriesUnauthorized{}
}

// WithPayload adds the payload to the retrieve timeseries unauthorized response
func (o *RetrieveTimeseriesUnauthorized) WithPayload(payload *models.Unauthorized) *RetrieveTimeseriesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve timeseries unauthorized response
func (o *RetrieveTimeseriesUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTimeseriesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveTimeseriesForbiddenCode is the HTTP code returned for type RetrieveTimeseriesForbidden
const RetrieveTimeseriesForbiddenCode int = 403

/*
RetrieveTimeseriesForbidden Forbidden

swagger:response retrieveTimeseriesForbidden
*/
type RetrieveTimeseriesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Forbidden `json:"body,omitempty"`
}

// NewRetrieveTimeseriesForbidden creates RetrieveTimeseriesForbidden with default headers values
func NewRetrieveTimeseriesForbidden() *RetrieveTimeseriesForbidden {

	return &RetrieveTimeseriesForbidden{}
}

// WithPayload adds the payload to the retrieve timeseries forbidden response
func (o *RetrieveTimeseriesForbidden) WithPayload(payload *models.Forbidden) *RetrieveTimeseriesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve timeseries forbidden response
func (o *RetrieveTimeseriesForbidden) SetPayload(payload *models.Forbidden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTimeseriesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveTimeseriesNotFoundCode is the HTTP code returned for type RetrieveTimeseriesNotFound
const RetrieveTimeseriesNotFoundCode int = 404

/*
RetrieveTimeseriesNotFound asset (entity) not found

swagger:response retrieveTimeseriesNotFound
*/
type RetrieveTimeseriesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewRetrieveTimeseriesNotFound creates RetrieveTimeseriesNotFound with default headers values
func NewRetrieveTimeseriesNotFound() *RetrieveTimeseriesNotFound {

	return &RetrieveTimeseriesNotFound{}
}

// WithPayload adds the payload to the retrieve timeseries not found response
func (o *RetrieveTimeseriesNotFound) WithPayload(payload *models.NotFound) *RetrieveTimeseriesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve timeseries not found response
func (o *RetrieveTimeseriesNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTimeseriesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveTimeseriesTooManyRequestsCode is the HTTP code returned for type RetrieveTimeseriesTooManyRequests
const RetrieveTimeseriesTooManyRequestsCode int = 429

/*
RetrieveTimeseriesTooManyRequests Too many requests that can be arised based on subsequent case:
- User has sent too many requests in a given amount of time.

A Retry-After header might be included to this response indicating how long in 'seconds' to wait before making a new request.

swagger:response retrieveTimeseriesTooManyRequests
*/
type RetrieveTimeseriesTooManyRequests struct {
	/*Time in seconds to wait before making a new request

	 */
	RetryAfter int64 `json:"Retry-After"`

	/*
	  In: Body
	*/
	Payload *models.RateLimiter `json:"body,omitempty"`
}

// NewRetrieveTimeseriesTooManyRequests creates RetrieveTimeseriesTooManyRequests with default headers values
func NewRetrieveTimeseriesTooManyRequests() *RetrieveTimeseriesTooManyRequests {

	return &RetrieveTimeseriesTooManyRequests{}
}

// WithRetryAfter adds the retryAfter to the retrieve timeseries too many requests response
func (o *RetrieveTimeseriesTooManyRequests) WithRetryAfter(retryAfter int64) *RetrieveTimeseriesTooManyRequests {
	o.RetryAfter = retryAfter
	return o
}

// SetRetryAfter sets the retryAfter to the retrieve timeseries too many requests response
func (o *RetrieveTimeseriesTooManyRequests) SetRetryAfter(retryAfter int64) {
	o.RetryAfter = retryAfter
}

// WithPayload adds the payload to the retrieve timeseries too many requests response
func (o *RetrieveTimeseriesTooManyRequests) WithPayload(payload *models.RateLimiter) *RetrieveTimeseriesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve timeseries too many requests response
func (o *RetrieveTimeseriesTooManyRequests) SetPayload(payload *models.RateLimiter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTimeseriesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Retry-After

	retryAfter := swag.FormatInt64(o.RetryAfter)
	if retryAfter != "" {
		rw.Header().Set("Retry-After", retryAfter)
	}

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveTimeseriesServiceUnavailableCode is the HTTP code returned for type RetrieveTimeseriesServiceUnavailable
const RetrieveTimeseriesServiceUnavailableCode int = 503

/*
RetrieveTimeseriesServiceUnavailable Service unavailable

swagger:response retrieveTimeseriesServiceUnavailable
*/
type RetrieveTimeseriesServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceUnavailable `json:"body,omitempty"`
}

// NewRetrieveTimeseriesServiceUnavailable creates RetrieveTimeseriesServiceUnavailable with default headers values
func NewRetrieveTimeseriesServiceUnavailable() *RetrieveTimeseriesServiceUnavailable {

	return &RetrieveTimeseriesServiceUnavailable{}
}

// WithPayload adds the payload to the retrieve timeseries service unavailable response
func (o *RetrieveTimeseriesServiceUnavailable) WithPayload(payload *models.ServiceUnavailable) *RetrieveTimeseriesServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve timeseries service unavailable response
func (o *RetrieveTimeseriesServiceUnavailable) SetPayload(payload *models.ServiceUnavailable) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTimeseriesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
RetrieveTimeseriesDefault Other error with any status code and response body format.

swagger:response retrieveTimeseriesDefault
*/
type RetrieveTimeseriesDefault struct {
	_statusCode int
}

// NewRetrieveTimeseriesDefault creates RetrieveTimeseriesDefault with default headers values
func NewRetrieveTimeseriesDefault(code int) *RetrieveTimeseriesDefault {
	if code <= 0 {
		code = 500
	}

	return &RetrieveTimeseriesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the retrieve timeseries default response
func (o *RetrieveTimeseriesDefault) WithStatusCode(code int) *RetrieveTimeseriesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the retrieve timeseries default response
func (o *RetrieveTimeseriesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *RetrieveTimeseriesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
