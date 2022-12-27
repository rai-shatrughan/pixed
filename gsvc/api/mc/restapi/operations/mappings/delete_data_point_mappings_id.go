// Code generated by go-swagger; DO NOT EDIT.

package mappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/mc/models"
)

// DeleteDataPointMappingsIDHandlerFunc turns a function with the right signature into a delete data point mappings ID handler
type DeleteDataPointMappingsIDHandlerFunc func(DeleteDataPointMappingsIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDataPointMappingsIDHandlerFunc) Handle(params DeleteDataPointMappingsIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteDataPointMappingsIDHandler interface for that can handle valid delete data point mappings ID params
type DeleteDataPointMappingsIDHandler interface {
	Handle(DeleteDataPointMappingsIDParams, *models.Principal) middleware.Responder
}

// NewDeleteDataPointMappingsID creates a new http.Handler for the delete data point mappings ID operation
func NewDeleteDataPointMappingsID(ctx *middleware.Context, handler DeleteDataPointMappingsIDHandler) *DeleteDataPointMappingsID {
	return &DeleteDataPointMappingsID{Context: ctx, Handler: handler}
}

/*
	DeleteDataPointMappingsID swagger:route DELETE /dataPointMappings/{id} Mappings deleteDataPointMappingsId

# Delete a mapping

Deletes a mapping.
*/
type DeleteDataPointMappingsID struct {
	Context *middleware.Context
	Handler DeleteDataPointMappingsIDHandler
}

func (o *DeleteDataPointMappingsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteDataPointMappingsIDParams()
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