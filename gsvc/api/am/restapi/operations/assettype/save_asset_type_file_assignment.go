// Code generated by go-swagger; DO NOT EDIT.

package assettype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// SaveAssetTypeFileAssignmentHandlerFunc turns a function with the right signature into a save asset type file assignment handler
type SaveAssetTypeFileAssignmentHandlerFunc func(SaveAssetTypeFileAssignmentParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveAssetTypeFileAssignmentHandlerFunc) Handle(params SaveAssetTypeFileAssignmentParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SaveAssetTypeFileAssignmentHandler interface for that can handle valid save asset type file assignment params
type SaveAssetTypeFileAssignmentHandler interface {
	Handle(SaveAssetTypeFileAssignmentParams, *models.Principal) middleware.Responder
}

// NewSaveAssetTypeFileAssignment creates a new http.Handler for the save asset type file assignment operation
func NewSaveAssetTypeFileAssignment(ctx *middleware.Context, handler SaveAssetTypeFileAssignmentHandler) *SaveAssetTypeFileAssignment {
	return &SaveAssetTypeFileAssignment{Context: ctx, Handler: handler}
}

/*
	SaveAssetTypeFileAssignment swagger:route PUT /assettypes/{id}/fileAssignments/{key} assettype saveAssetTypeFileAssignment

# Add a new file assignment to an asset type

Add a new file assignment to a given asset type. All asset which extends these types will have its file by default.
*/
type SaveAssetTypeFileAssignment struct {
	Context *middleware.Context
	Handler SaveAssetTypeFileAssignmentHandler
}

func (o *SaveAssetTypeFileAssignment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSaveAssetTypeFileAssignmentParams()
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