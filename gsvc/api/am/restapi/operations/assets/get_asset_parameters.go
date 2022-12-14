// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetAssetParams creates a new GetAssetParams object
// with the default values initialized.
func NewGetAssetParams() GetAssetParams {

	var (
		// initialize parameters with default values

		includeSharedDefault = bool(false)
	)

	return GetAssetParams{
		IncludeShared: &includeSharedDefault,
	}
}

// GetAssetParams contains all the bound params for the get asset operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAsset
type GetAssetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ETag hash of previous request to allow caching
	  In: header
	*/
	IfNoneMatch *string
	/*Unique identifier
	  Required: true
	  Pattern: [0-9A-F]{32}
	  In: path
	*/
	ID string
	/*Specifies if the operation should take into account shared (received) assets, aspects and asset types. Received resources are not visible in case includeShared=false. For query operations, received resources are not returned. Endpoints addressing individual resources respond with 404. In case received resources are referenced in a request parameter or property, they are treated as not existing.
	  In: query
	  Default: false
	*/
	IncludeShared *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAssetParams() beforehand.
func (o *GetAssetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindIfNoneMatch(r.Header[http.CanonicalHeaderKey("If-None-Match")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qIncludeShared, qhkIncludeShared, _ := qs.GetOK("includeShared")
	if err := o.bindIncludeShared(qIncludeShared, qhkIncludeShared, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfNoneMatch binds and validates parameter IfNoneMatch from header.
func (o *GetAssetParams) bindIfNoneMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.IfNoneMatch = &raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetAssetParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	if err := o.validateID(formats); err != nil {
		return err
	}

	return nil
}

// validateID carries on validations for parameter ID
func (o *GetAssetParams) validateID(formats strfmt.Registry) error {

	if err := validate.Pattern("id", "path", o.ID, `[0-9A-F]{32}`); err != nil {
		return err
	}

	return nil
}

// bindIncludeShared binds and validates parameter IncludeShared from query.
func (o *GetAssetParams) bindIncludeShared(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAssetParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("includeShared", "query", "bool", raw)
	}
	o.IncludeShared = &value

	return nil
}
