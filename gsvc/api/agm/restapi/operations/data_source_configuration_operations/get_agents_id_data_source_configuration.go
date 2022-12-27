// Code generated by go-swagger; DO NOT EDIT.

package data_source_configuration_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/agm/models"
)

// GetAgentsIDDataSourceConfigurationHandlerFunc turns a function with the right signature into a get agents ID data source configuration handler
type GetAgentsIDDataSourceConfigurationHandlerFunc func(GetAgentsIDDataSourceConfigurationParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAgentsIDDataSourceConfigurationHandlerFunc) Handle(params GetAgentsIDDataSourceConfigurationParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetAgentsIDDataSourceConfigurationHandler interface for that can handle valid get agents ID data source configuration params
type GetAgentsIDDataSourceConfigurationHandler interface {
	Handle(GetAgentsIDDataSourceConfigurationParams, *models.Principal) middleware.Responder
}

// NewGetAgentsIDDataSourceConfiguration creates a new http.Handler for the get agents ID data source configuration operation
func NewGetAgentsIDDataSourceConfiguration(ctx *middleware.Context, handler GetAgentsIDDataSourceConfigurationHandler) *GetAgentsIDDataSourceConfiguration {
	return &GetAgentsIDDataSourceConfiguration{Context: ctx, Handler: handler}
}

/*
	GetAgentsIDDataSourceConfiguration swagger:route GET /agents/{id}/dataSourceConfiguration Data Source Configuration Operations getAgentsIdDataSourceConfiguration

# Get Data Source Configuration

Data source configuration is needed for interpreting the data received from an agent. It is crucial for diot to interpret the data sent by an agent. The data source configuration contains data sources and data points. Data sources are logical groups of data points, e.g. a sensor or a machine. Fetches data source configuration object.
*/
type GetAgentsIDDataSourceConfiguration struct {
	Context *middleware.Context
	Handler GetAgentsIDDataSourceConfigurationHandler
}

func (o *GetAgentsIDDataSourceConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAgentsIDDataSourceConfigurationParams()
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