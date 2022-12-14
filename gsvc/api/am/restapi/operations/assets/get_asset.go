// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// GetAssetHandlerFunc turns a function with the right signature into a get asset handler
type GetAssetHandlerFunc func(GetAssetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAssetHandlerFunc) Handle(params GetAssetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetAssetHandler interface for that can handle valid get asset params
type GetAssetHandler interface {
	Handle(GetAssetParams, *models.Principal) middleware.Responder
}

// NewGetAsset creates a new http.Handler for the get asset operation
func NewGetAsset(ctx *middleware.Context, handler GetAssetHandler) *GetAsset {
	return &GetAsset{Context: ctx, Handler: handler}
}

/*
	GetAsset swagger:route GET /assets/{id} assets getAsset

Returns an asset.

Read a single asset. All static properties of asset are returned.
*/
type GetAsset struct {
	Context *middleware.Context
	Handler GetAssetHandler
}

func (o *GetAsset) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAssetParams()
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
