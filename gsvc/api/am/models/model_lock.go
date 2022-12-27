// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelLock model lock
//
// swagger:model ModelLock
type ModelLock struct {

	// enabled
	Enabled Enabled `json:"enabled,omitempty"`
}

// Validate validates this model lock
func (m *ModelLock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelLock) validateEnabled(formats strfmt.Registry) error {
	if swag.IsZero(m.Enabled) { // not required
		return nil
	}

	if err := m.Enabled.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("enabled")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("enabled")
		}
		return err
	}

	return nil
}

// ContextValidate validate this model lock based on the context it is used
func (m *ModelLock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelLock) contextValidateEnabled(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Enabled.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("enabled")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("enabled")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelLock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelLock) UnmarshalBinary(b []byte) error {
	var res ModelLock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}