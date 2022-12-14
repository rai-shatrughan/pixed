// Code generated by go-swagger; DO NOT EDIT.

package token_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetOauthTokenKeyHandlerFunc turns a function with the right signature into a get oauth token key handler
type GetOauthTokenKeyHandlerFunc func(GetOauthTokenKeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOauthTokenKeyHandlerFunc) Handle(params GetOauthTokenKeyParams) middleware.Responder {
	return fn(params)
}

// GetOauthTokenKeyHandler interface for that can handle valid get oauth token key params
type GetOauthTokenKeyHandler interface {
	Handle(GetOauthTokenKeyParams) middleware.Responder
}

// NewGetOauthTokenKey creates a new http.Handler for the get oauth token key operation
func NewGetOauthTokenKey(ctx *middleware.Context, handler GetOauthTokenKeyHandler) *GetOauthTokenKey {
	return &GetOauthTokenKey{Context: ctx, Handler: handler}
}

/*
	GetOauthTokenKey swagger:route GET /oauth/token_key Token Operations getOauthTokenKey

# Returns the JWT signing key currently employed by Agent IAM

Provides current RSA public key of AgentIAM. AgentIAM uses corresponding RSA private key to sign access tokens, RATs and IATs.
*/
type GetOauthTokenKey struct {
	Context *middleware.Context
	Handler GetOauthTokenKeyHandler
}

func (o *GetOauthTokenKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetOauthTokenKeyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
