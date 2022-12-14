// Code generated by go-swagger; DO NOT EDIT.

package record_recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/mc/models"
)

// GetRecoverableRecordsHandlerFunc turns a function with the right signature into a get recoverable records handler
type GetRecoverableRecordsHandlerFunc func(GetRecoverableRecordsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRecoverableRecordsHandlerFunc) Handle(params GetRecoverableRecordsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetRecoverableRecordsHandler interface for that can handle valid get recoverable records params
type GetRecoverableRecordsHandler interface {
	Handle(GetRecoverableRecordsParams, *models.Principal) middleware.Responder
}

// NewGetRecoverableRecords creates a new http.Handler for the get recoverable records operation
func NewGetRecoverableRecords(ctx *middleware.Context, handler GetRecoverableRecordsHandler) *GetRecoverableRecords {
	return &GetRecoverableRecords{Context: ctx, Handler: handler}
}

/*
	GetRecoverableRecords swagger:route GET /recoverableRecords Record Recovery getRecoverableRecords

# Get all recoverable records

Gets all recoverable records
*/
type GetRecoverableRecords struct {
	Context *middleware.Context
	Handler GetRecoverableRecordsHandler
}

func (o *GetRecoverableRecords) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRecoverableRecordsParams()
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
