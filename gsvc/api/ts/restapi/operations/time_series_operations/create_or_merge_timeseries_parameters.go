// Code generated by go-swagger; DO NOT EDIT.

package time_series_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"gsvc/api/ts/models"
)

// NewCreateOrMergeTimeseriesParams creates a new CreateOrMergeTimeseriesParams object
//
// There are no default values defined in the spec.
func NewCreateOrMergeTimeseriesParams() CreateOrMergeTimeseriesParams {

	return CreateOrMergeTimeseriesParams{}
}

// CreateOrMergeTimeseriesParams contains all the bound params for the create or merge timeseries operation
// typically these are obtained from a http.Request
//
// swagger:parameters createOrMergeTimeseries
type CreateOrMergeTimeseriesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*unique identifier of the asset (entity)
	  Required: true
	  Max Length: 32
	  Min Length: 32
	  Pattern: [a-f0-9]{32}
	  In: path
	*/
	EntityID string
	/*Name of the aspect (property set).
	  Required: true
	  In: path
	*/
	PropertySetName string
	/*time series data array
	  Required: true
	  In: body
	*/
	Timeseries []*models.TimeSeriesDataItem
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateOrMergeTimeseriesParams() beforehand.
func (o *CreateOrMergeTimeseriesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rEntityID, rhkEntityID, _ := route.Params.GetOK("entityId")
	if err := o.bindEntityID(rEntityID, rhkEntityID, route.Formats); err != nil {
		res = append(res, err)
	}

	rPropertySetName, rhkPropertySetName, _ := route.Params.GetOK("propertySetName")
	if err := o.bindPropertySetName(rPropertySetName, rhkPropertySetName, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.TimeSeriesDataItem
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("timeseries", "body", ""))
			} else {
				res = append(res, errors.NewParseError("timeseries", "body", "", err))
			}
		} else {

			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Timeseries = body
			}
		}
	} else {
		res = append(res, errors.Required("timeseries", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEntityID binds and validates parameter EntityID from path.
func (o *CreateOrMergeTimeseriesParams) bindEntityID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.EntityID = raw

	if err := o.validateEntityID(formats); err != nil {
		return err
	}

	return nil
}

// validateEntityID carries on validations for parameter EntityID
func (o *CreateOrMergeTimeseriesParams) validateEntityID(formats strfmt.Registry) error {

	if err := validate.MinLength("entityId", "path", o.EntityID, 32); err != nil {
		return err
	}

	if err := validate.MaxLength("entityId", "path", o.EntityID, 32); err != nil {
		return err
	}

	if err := validate.Pattern("entityId", "path", o.EntityID, `[a-f0-9]{32}`); err != nil {
		return err
	}

	return nil
}

// bindPropertySetName binds and validates parameter PropertySetName from path.
func (o *CreateOrMergeTimeseriesParams) bindPropertySetName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.PropertySetName = raw

	return nil
}
