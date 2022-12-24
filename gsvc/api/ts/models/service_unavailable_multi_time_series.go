// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceUnavailableMultiTimeSeries service unavailable multi time series
// Example: {"error":"Service Unavailable","message":"Service not available.","path":"/api/iottimeseries/v3/timeseries","status":503,"timestamp":"2019-04-08T06:32:01.261+0000"}
//
// swagger:model ServiceUnavailableMultiTimeSeries
type ServiceUnavailableMultiTimeSeries struct {

	// error
	Error string `json:"error,omitempty"`

	// exception
	Exception string `json:"exception,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// status
	Status float64 `json:"status,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this service unavailable multi time series
func (m *ServiceUnavailableMultiTimeSeries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceUnavailableMultiTimeSeries) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service unavailable multi time series based on context it is used
func (m *ServiceUnavailableMultiTimeSeries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceUnavailableMultiTimeSeries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceUnavailableMultiTimeSeries) UnmarshalBinary(b []byte) error {
	var res ServiceUnavailableMultiTimeSeries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
