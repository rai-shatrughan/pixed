// Code generated by go-swagger; DO NOT EDIT.

package time_series_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/ts/models"
)

// RetrieveTimeseriesHandlerFunc turns a function with the right signature into a retrieve timeseries handler
type RetrieveTimeseriesHandlerFunc func(RetrieveTimeseriesParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveTimeseriesHandlerFunc) Handle(params RetrieveTimeseriesParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// RetrieveTimeseriesHandler interface for that can handle valid retrieve timeseries params
type RetrieveTimeseriesHandler interface {
	Handle(RetrieveTimeseriesParams, *models.Principal) middleware.Responder
}

// NewRetrieveTimeseries creates a new http.Handler for the retrieve timeseries operation
func NewRetrieveTimeseries(ctx *middleware.Context, handler RetrieveTimeseriesHandler) *RetrieveTimeseries {
	return &RetrieveTimeseries{Context: ctx, Handler: handler}
}

/*
	RetrieveTimeseries swagger:route GET /timeseries/{entityId}/{propertySetName} Time Series Operations retrieveTimeseries

# Retrieve time series data

Retrieve time series data for one combination of an asset (entity) and an(a) aspect (property set). The maximum number of time series data items returned per request is defined by parameter <i>limit</i>. In case more time series data items are present in the requested time range, only a subset of data items will be returned and a header <i>link</i> is added to the response. The header value contains the request URL to fetch the next set of  time series data items, by increasing the <i>from</i> parameter accordingly. Returns the latest available timestamp record if no range is provided.
Note:

	> If no range is provided, Based on Customer data ingestion timestamp record, it may not contain all variables of aspect. For example, an Aspect has 5 variables; but latest timeseries record has only 3 variables with values. Only latest timeseries record will be returned in response with values of 3 variables only. It will not fetch latest value of all 5 variables of aspect.
*/
type RetrieveTimeseries struct {
	Context *middleware.Context
	Handler RetrieveTimeseriesHandler
}

func (o *RetrieveTimeseries) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRetrieveTimeseriesParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}