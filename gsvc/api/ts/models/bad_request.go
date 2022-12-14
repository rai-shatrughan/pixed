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

// BadRequest bad request
// Example: {"error":"Bad Request","exception":"com.diot.iot.exception.IotException","message":"[6008] Invalid date format '2018-10-16T00:00:00Z111' for request param - from","path":"/api/iottimeseries/v3/timeseries/entityId/propertySetName","status":400,"timestamp":"2019-04-08T06:32:01.261+0000"}
//
// swagger:model BadRequest
type BadRequest struct {

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

// Validate validates this bad request
func (m *BadRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BadRequest) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bad request based on context it is used
func (m *BadRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BadRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BadRequest) UnmarshalBinary(b []byte) error {
	var res BadRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
