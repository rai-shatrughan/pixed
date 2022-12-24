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

// MultiStatusErrorItem multi status error item
//
// swagger:model MultiStatusErrorItem
type MultiStatusErrorItem struct {

	// entity Id
	// Example: entityId2
	EntityID string `json:"entityId,omitempty"`

	// error
	// Example: Not Found
	Error string `json:"error,omitempty"`

	// exception
	// Example: com.diot.iot.exception.IotException
	Exception string `json:"exception,omitempty"`

	// message
	// Example: [6007] Invalid value ' for request param - propertysetname
	Message string `json:"message,omitempty"`

	// path
	// Example: /api/iottimeseries/v3/timeseries
	Path string `json:"path,omitempty"`

	// property set name
	// Example: propertySetName2
	PropertySetName string `json:"propertySetName,omitempty"`

	// status
	// Example: 404
	Status float64 `json:"status,omitempty"`

	// timestamp
	// Example: 2019-04-08T06:32:01.261+0000
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this multi status error item
func (m *MultiStatusErrorItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiStatusErrorItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this multi status error item based on context it is used
func (m *MultiStatusErrorItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MultiStatusErrorItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiStatusErrorItem) UnmarshalBinary(b []byte) error {
	var res MultiStatusErrorItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
