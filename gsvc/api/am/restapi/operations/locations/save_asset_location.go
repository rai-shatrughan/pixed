// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// SaveAssetLocationHandlerFunc turns a function with the right signature into a save asset location handler
type SaveAssetLocationHandlerFunc func(SaveAssetLocationParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveAssetLocationHandlerFunc) Handle(params SaveAssetLocationParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SaveAssetLocationHandler interface for that can handle valid save asset location params
type SaveAssetLocationHandler interface {
	Handle(SaveAssetLocationParams, *models.Principal) middleware.Responder
}

// NewSaveAssetLocation creates a new http.Handler for the save asset location operation
func NewSaveAssetLocation(ctx *middleware.Context, handler SaveAssetLocationHandler) *SaveAssetLocation {
	return &SaveAssetLocation{Context: ctx, Handler: handler}
}

/*
	SaveAssetLocation swagger:route PUT /assets/{id}/location locations saveAssetLocation

# Create or Update location assigned to given asset

* If the given asset has own location, this endpoint will update that location. * If the given asset has no location, this endpoint will create a new location and update the given asset. * If the given asset has inherited location, this endpoint will create a new location and update the given asset. If you wanted to update the inherited location you have to use the 'location' url in AssetResource object (with PUT method).
*/
type SaveAssetLocation struct {
	Context *middleware.Context
	Handler SaveAssetLocationHandler
}

func (o *SaveAssetLocation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSaveAssetLocationParams()
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