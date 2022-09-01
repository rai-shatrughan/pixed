// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"ts/models"
)

// GetTimeseriesVersionHandlerFunc turns a function with the right signature into a get timeseries version handler
type GetTimeseriesVersionHandlerFunc func(GetTimeseriesVersionParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTimeseriesVersionHandlerFunc) Handle(params GetTimeseriesVersionParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetTimeseriesVersionHandler interface for that can handle valid get timeseries version params
type GetTimeseriesVersionHandler interface {
	Handle(GetTimeseriesVersionParams, *models.Principal) middleware.Responder
}

// NewGetTimeseriesVersion creates a new http.Handler for the get timeseries version operation
func NewGetTimeseriesVersion(ctx *middleware.Context, handler GetTimeseriesVersionHandler) *GetTimeseriesVersion {
	return &GetTimeseriesVersion{Context: ctx, Handler: handler}
}

/* GetTimeseriesVersion swagger:route GET /timeseries/version getTimeseriesVersion

Returns current api version.

*/
type GetTimeseriesVersion struct {
	Context *middleware.Context
	Handler GetTimeseriesVersionHandler
}

func (o *GetTimeseriesVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTimeseriesVersionParams()
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
