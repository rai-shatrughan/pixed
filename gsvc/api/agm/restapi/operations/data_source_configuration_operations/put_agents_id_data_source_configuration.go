// Code generated by go-swagger; DO NOT EDIT.

package data_source_configuration_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/agm/models"
)

// PutAgentsIDDataSourceConfigurationHandlerFunc turns a function with the right signature into a put agents ID data source configuration handler
type PutAgentsIDDataSourceConfigurationHandlerFunc func(PutAgentsIDDataSourceConfigurationParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutAgentsIDDataSourceConfigurationHandlerFunc) Handle(params PutAgentsIDDataSourceConfigurationParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutAgentsIDDataSourceConfigurationHandler interface for that can handle valid put agents ID data source configuration params
type PutAgentsIDDataSourceConfigurationHandler interface {
	Handle(PutAgentsIDDataSourceConfigurationParams, *models.Principal) middleware.Responder
}

// NewPutAgentsIDDataSourceConfiguration creates a new http.Handler for the put agents ID data source configuration operation
func NewPutAgentsIDDataSourceConfiguration(ctx *middleware.Context, handler PutAgentsIDDataSourceConfigurationHandler) *PutAgentsIDDataSourceConfiguration {
	return &PutAgentsIDDataSourceConfiguration{Context: ctx, Handler: handler}
}

/*
	PutAgentsIDDataSourceConfiguration swagger:route PUT /agents/{id}/dataSourceConfiguration Data Source Configuration Operations putAgentsIdDataSourceConfiguration

# Update Data Source Configuration

Creates or updates data source configuration object.
*/
type PutAgentsIDDataSourceConfiguration struct {
	Context *middleware.Context
	Handler PutAgentsIDDataSourceConfigurationHandler
}

func (o *PutAgentsIDDataSourceConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutAgentsIDDataSourceConfigurationParams()
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
