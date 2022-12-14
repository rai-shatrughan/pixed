// Code generated by go-swagger; DO NOT EDIT.

package time_series_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/ts/models"
)

// DeleteTimeseriesHandlerFunc turns a function with the right signature into a delete timeseries handler
type DeleteTimeseriesHandlerFunc func(DeleteTimeseriesParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTimeseriesHandlerFunc) Handle(params DeleteTimeseriesParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteTimeseriesHandler interface for that can handle valid delete timeseries params
type DeleteTimeseriesHandler interface {
	Handle(DeleteTimeseriesParams, *models.Principal) middleware.Responder
}

// NewDeleteTimeseries creates a new http.Handler for the delete timeseries operation
func NewDeleteTimeseries(ctx *middleware.Context, handler DeleteTimeseriesHandler) *DeleteTimeseries {
	return &DeleteTimeseries{Context: ctx, Handler: handler}
}

/*
	DeleteTimeseries swagger:route DELETE /timeseries/{entityId}/{propertySetName} Time Series Operations deleteTimeseries

# Delete time series data

Delete time series data for one combination of an asset (entity) and an(a) aspect (property set). All property values within the given time range are deleted.
*/
type DeleteTimeseries struct {
	Context *middleware.Context
	Handler DeleteTimeseriesHandler
}

func (o *DeleteTimeseries) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTimeseriesParams()
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
