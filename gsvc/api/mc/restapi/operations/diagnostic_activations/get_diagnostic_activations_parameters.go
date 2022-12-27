// Code generated by go-swagger; DO NOT EDIT.

package diagnostic_activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetDiagnosticActivationsParams creates a new GetDiagnosticActivationsParams object
// with the default values initialized.
func NewGetDiagnosticActivationsParams() GetDiagnosticActivationsParams {

	var (
		// initialize parameters with default values

		pageDefault = int32(0)
		sizeDefault = int32(20)
	)

	return GetDiagnosticActivationsParams{
		Page: &pageDefault,

		Size: &sizeDefault,
	}
}

// GetDiagnosticActivationsParams contains all the bound params for the get diagnostic activations operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetDiagnosticActivations
type GetDiagnosticActivationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The (0-based) index of page.
	  In: query
	  Default: 0
	*/
	Page *int32
	/*The maximum number of elements in a page.
	  In: query
	  Default: 20
	*/
	Size *int32
	/*The order of returned elements.
	Multiple fields could be used separated by commas (e.g. ''field1,field2'').
	Descending order could be requested by appending '',desc'' at the end of parameter.(e.g. ''field1,field2,desc'')'

	  In: query
	*/
	Sort *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDiagnosticActivationsParams() beforehand.
func (o *GetDiagnosticActivationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qSize, qhkSize, _ := qs.GetOK("size")
	if err := o.bindSize(qSize, qhkSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qSort, qhkSort, _ := qs.GetOK("sort")
	if err := o.bindSort(qSort, qhkSort, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetDiagnosticActivationsParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetDiagnosticActivationsParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int32", raw)
	}
	o.Page = &value

	return nil
}

// bindSize binds and validates parameter Size from query.
func (o *GetDiagnosticActivationsParams) bindSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetDiagnosticActivationsParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("size", "query", "int32", raw)
	}
	o.Size = &value

	return nil
}

// bindSort binds and validates parameter Sort from query.
func (o *GetDiagnosticActivationsParams) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Sort = &raw

	return nil
}