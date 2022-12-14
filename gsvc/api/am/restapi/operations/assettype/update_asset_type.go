// Code generated by go-swagger; DO NOT EDIT.

package assettype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// UpdateAssetTypeHandlerFunc turns a function with the right signature into a update asset type handler
type UpdateAssetTypeHandlerFunc func(UpdateAssetTypeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateAssetTypeHandlerFunc) Handle(params UpdateAssetTypeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateAssetTypeHandler interface for that can handle valid update asset type params
type UpdateAssetTypeHandler interface {
	Handle(UpdateAssetTypeParams, *models.Principal) middleware.Responder
}

// NewUpdateAssetType creates a new http.Handler for the update asset type operation
func NewUpdateAssetType(ctx *middleware.Context, handler UpdateAssetTypeHandler) *UpdateAssetType {
	return &UpdateAssetType{Context: ctx, Handler: handler}
}

/*
	UpdateAssetType swagger:route PATCH /assettypes/{id} assettype updateAssetType

# Patch an asset type

Patch an asset type. Conforms to RFC 7396 - JSON merge Patch.
<b> Following are the updates allowed</b>
* Asset type description can be updated.
* Asset type variable's can be created, updated and removed. Varaible length, default Value and Unit can be changed. The unit changes from the API does not compute any value changes derived after the unit changes, the values will remain as it is and only the unit will be updated. The length can only be increased of a string variable and it cannot be decreased.
* File assignments can be updated and removed.

* Aspects can be created, updated and removed, please refer sample payloads below

  - Create aspect
    {
    "name": "leftWing",
    "aspectTypeId": "mdsp.wing"
    }
  - Update aspect, <b><u>aspectId should be provided in payload</u></b> (only name can be updated)
    {
    "name": "rightWing",
    "aspectId": "b9cbfc7073be4530887cdb1e71c932b8",
    "aspectTypeId": "mdsp.wing"
    }
  - To delete an aspect, aspect should not be part of aspects payload
*/
type UpdateAssetType struct {
	Context *middleware.Context
	Handler UpdateAssetTypeHandler
}

func (o *UpdateAssetType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateAssetTypeParams()
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
