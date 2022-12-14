// Code generated by go-swagger; DO NOT EDIT.

package diagnostic_activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/mc/models"
)

// PostDiagnosticActivationsHandlerFunc turns a function with the right signature into a post diagnostic activations handler
type PostDiagnosticActivationsHandlerFunc func(PostDiagnosticActivationsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostDiagnosticActivationsHandlerFunc) Handle(params PostDiagnosticActivationsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostDiagnosticActivationsHandler interface for that can handle valid post diagnostic activations params
type PostDiagnosticActivationsHandler interface {
	Handle(PostDiagnosticActivationsParams, *models.Principal) middleware.Responder
}

// NewPostDiagnosticActivations creates a new http.Handler for the post diagnostic activations operation
func NewPostDiagnosticActivations(ctx *middleware.Context, handler PostDiagnosticActivationsHandler) *PostDiagnosticActivations {
	return &PostDiagnosticActivations{Context: ctx, Handler: handler}
}

/*
	PostDiagnosticActivations swagger:route POST /diagnosticActivations Diagnostic Activations postDiagnosticActivations

# Creates a new diagnostic activation

Create a new diagnostic activation.
Agents are allowed to create activation for itself only. Users with sufficient scopes are allowed to create activations of the agents in the same tenant as in the token.
*/
type PostDiagnosticActivations struct {
	Context *middleware.Context
	Handler PostDiagnosticActivationsHandler
}

func (o *PostDiagnosticActivations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostDiagnosticActivationsParams()
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
