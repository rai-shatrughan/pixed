// Code generated by go-swagger; DO NOT EDIT.

package assettype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// DeleteAssetTypeHandlerFunc turns a function with the right signature into a delete asset type handler
type DeleteAssetTypeHandlerFunc func(DeleteAssetTypeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAssetTypeHandlerFunc) Handle(params DeleteAssetTypeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteAssetTypeHandler interface for that can handle valid delete asset type params
type DeleteAssetTypeHandler interface {
	Handle(DeleteAssetTypeParams, *models.Principal) middleware.Responder
}

// NewDeleteAssetType creates a new http.Handler for the delete asset type operation
func NewDeleteAssetType(ctx *middleware.Context, handler DeleteAssetTypeHandler) *DeleteAssetType {
	return &DeleteAssetType{Context: ctx, Handler: handler}
}

/*
	DeleteAssetType swagger:route DELETE /assettypes/{id} assettype deleteAssetType

# Delete an asset type

Deletes an asset type. Deletion only possible when the type has no child-type and there is no asset that instantiate it.
*/
type DeleteAssetType struct {
	Context *middleware.Context
	Handler DeleteAssetTypeHandler
}

func (o *DeleteAssetType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteAssetTypeParams()
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
