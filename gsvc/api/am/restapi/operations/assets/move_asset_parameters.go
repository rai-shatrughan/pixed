// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"gsvc/api/am/models"
)

// NewMoveAssetParams creates a new MoveAssetParams object
// with the default values initialized.
func NewMoveAssetParams() MoveAssetParams {

	var (
		// initialize parameters with default values

		includeSharedDefault = bool(false)
	)

	return MoveAssetParams{
		IncludeShared: &includeSharedDefault,
	}
}

// MoveAssetParams contains all the bound params for the move asset operation
// typically these are obtained from a http.Request
//
// swagger:parameters moveAsset
type MoveAssetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Last known version to facilitate optimistic locking
	  Required: true
	  In: header
	*/
	IfMatch string
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
	/*parameters for the asset move
	  Required: true
	  In: body
	*/
	MoveParameters *models.AssetMove
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewMoveAssetParams() beforehand.
func (o *MoveAssetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindIfMatch(r.Header[http.CanonicalHeaderKey("If-Match")], true, route.Formats); err != nil {
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

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AssetMove
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("moveParameters", "body", ""))
			} else {
				res = append(res, errors.NewParseError("moveParameters", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.MoveParameters = &body
			}
		}
	} else {
		res = append(res, errors.Required("moveParameters", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfMatch binds and validates parameter IfMatch from header.
func (o *MoveAssetParams) bindIfMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("If-Match", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("If-Match", "header", raw); err != nil {
		return err
	}
	o.IfMatch = raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *MoveAssetParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *MoveAssetParams) validateID(formats strfmt.Registry) error {

	if err := validate.Pattern("id", "path", o.ID, `[0-9A-F]{32}`); err != nil {
		return err
	}

	return nil
}

// bindIncludeShared binds and validates parameter IncludeShared from query.
func (o *MoveAssetParams) bindIncludeShared(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewMoveAssetParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("includeShared", "query", "bool", raw)
	}
	o.IncludeShared = &value

	return nil
}
