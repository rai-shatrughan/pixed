// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TimeSeriesItem time series item
//
// swagger:model TimeSeriesItem
type TimeSeriesItem struct {

	// Timeseries data for corresponsding asset (entity) Id and aspect (property set) name.
	// Required: true
	Data []*TimeSeriesDataItem `json:"data"`

	// asset (entity) Id of timeseries data.
	// Required: true
	EntityID *string `json:"entityId"`

	// aspect (property set) name of timeseries data.
	// Required: true
	PropertySetName *string `json:"propertySetName"`
}

// Validate validates this time series item
func (m *TimeSeriesItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertySetName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeSeriesItem) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeSeriesItem) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entityId", "body", m.EntityID); err != nil {
		return err
	}

	return nil
}

func (m *TimeSeriesItem) validatePropertySetName(formats strfmt.Registry) error {

	if err := validate.Required("propertySetName", "body", m.PropertySetName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this time series item based on the context it is used
func (m *TimeSeriesItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeSeriesItem) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeSeriesItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeSeriesItem) UnmarshalBinary(b []byte) error {
	var res TimeSeriesItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
