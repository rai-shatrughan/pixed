// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// ListAssetsHandlerFunc turns a function with the right signature into a list assets handler
type ListAssetsHandlerFunc func(ListAssetsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAssetsHandlerFunc) Handle(params ListAssetsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListAssetsHandler interface for that can handle valid list assets params
type ListAssetsHandler interface {
	Handle(ListAssetsParams, *models.Principal) middleware.Responder
}

// NewListAssets creates a new http.Handler for the list assets operation
func NewListAssets(ctx *middleware.Context, handler ListAssetsHandler) *ListAssets {
	return &ListAssets{Context: ctx, Handler: handler}
}

/*
	ListAssets swagger:route GET /assets assets listAssets

# List all available assets

List all assets available for the authenticated user. ## Filter Supports all basic fields and the 'hasType' filter which search for the assets with type that originated from the given asset type. # Examples Example: This filter searches for assets which type is exactly the given type: ``` { "typeId": "mandal.ship" } ``` Beyond the basic fields we can search for assets which type is inherited from the given type: ``` { "hasType": "mandal.ship" } ```
*/
type ListAssets struct {
	Context *middleware.Context
	Handler ListAssetsHandler
}

func (o *ListAssets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListAssetsParams()
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
