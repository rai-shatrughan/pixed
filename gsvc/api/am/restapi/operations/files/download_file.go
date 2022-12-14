// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
)

// DownloadFileHandlerFunc turns a function with the right signature into a download file handler
type DownloadFileHandlerFunc func(DownloadFileParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DownloadFileHandlerFunc) Handle(params DownloadFileParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DownloadFileHandler interface for that can handle valid download file params
type DownloadFileHandler interface {
	Handle(DownloadFileParams, *models.Principal) middleware.Responder
}

// NewDownloadFile creates a new http.Handler for the download file operation
func NewDownloadFile(ctx *middleware.Context, handler DownloadFileHandler) *DownloadFile {
	return &DownloadFile{Context: ctx, Handler: handler}
}

/*
	DownloadFile swagger:route GET /files/{fileId}/file files downloadFile

# Returns a file by its id

Returns a file by its id
*/
type DownloadFile struct {
	Context *middleware.Context
	Handler DownloadFileHandler
}

func (o *DownloadFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDownloadFileParams()
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
