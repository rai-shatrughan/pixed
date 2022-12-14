// Code generated by go-swagger; DO NOT EDIT.

package billboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// GetBillboardHandlerFunc turns a function with the right signature into a get billboard handler
type GetBillboardHandlerFunc func(GetBillboardParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBillboardHandlerFunc) Handle(params GetBillboardParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetBillboardHandler interface for that can handle valid get billboard params
type GetBillboardHandler interface {
	Handle(GetBillboardParams, *models.Principal) middleware.Responder
}

// NewGetBillboard creates a new http.Handler for the get billboard operation
func NewGetBillboard(ctx *middleware.Context, handler GetBillboardHandler) *GetBillboard {
	return &GetBillboard{Context: ctx, Handler: handler}
}

/*
	GetBillboard swagger:route GET / billboard getBillboard

# List all links for available resources

List all links for available resources
*/
type GetBillboard struct {
	Context *middleware.Context
	Handler GetBillboardHandler
}

func (o *GetBillboard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBillboardParams()
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
