// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// DeleteAssetHandlerFunc turns a function with the right signature into a delete asset handler
type DeleteAssetHandlerFunc func(DeleteAssetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAssetHandlerFunc) Handle(params DeleteAssetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteAssetHandler interface for that can handle valid delete asset params
type DeleteAssetHandler interface {
	Handle(DeleteAssetParams, *models.Principal) middleware.Responder
}

// NewDeleteAsset creates a new http.Handler for the delete asset operation
func NewDeleteAsset(ctx *middleware.Context, handler DeleteAssetHandler) *DeleteAsset {
	return &DeleteAsset{Context: ctx, Handler: handler}
}

/*
	DeleteAsset swagger:route DELETE /assets/{id} assets deleteAsset

# Delete an asset

Deletes the given asset. It's not possible to delete an asset if it has children.
*/
type DeleteAsset struct {
	Context *middleware.Context
	Handler DeleteAssetHandler
}

func (o *DeleteAsset) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteAssetParams()
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
