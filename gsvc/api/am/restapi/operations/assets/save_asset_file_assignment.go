// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// SaveAssetFileAssignmentHandlerFunc turns a function with the right signature into a save asset file assignment handler
type SaveAssetFileAssignmentHandlerFunc func(SaveAssetFileAssignmentParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveAssetFileAssignmentHandlerFunc) Handle(params SaveAssetFileAssignmentParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SaveAssetFileAssignmentHandler interface for that can handle valid save asset file assignment params
type SaveAssetFileAssignmentHandler interface {
	Handle(SaveAssetFileAssignmentParams, *models.Principal) middleware.Responder
}

// NewSaveAssetFileAssignment creates a new http.Handler for the save asset file assignment operation
func NewSaveAssetFileAssignment(ctx *middleware.Context, handler SaveAssetFileAssignmentHandler) *SaveAssetFileAssignment {
	return &SaveAssetFileAssignment{Context: ctx, Handler: handler}
}

/*
	SaveAssetFileAssignment swagger:route PUT /assets/{id}/fileAssignments/{key} assets saveAssetFileAssignment

# Save an file assignment to an asset

Save a file assignment to a given asset
*/
type SaveAssetFileAssignment struct {
	Context *middleware.Context
	Handler SaveAssetFileAssignmentHandler
}

func (o *SaveAssetFileAssignment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSaveAssetFileAssignmentParams()
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
