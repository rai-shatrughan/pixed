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

// AspectVariable aspect variable
//
// swagger:model AspectVariable
type AspectVariable struct {
	VariableDefinition

	// Indicates whether the variable has quality code. Cannot be changed.
	// Example: true
	QualityCode *bool `json:"qualityCode,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AspectVariable) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VariableDefinition
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VariableDefinition = aO0

	// AO1
	var dataAO1 struct {
		QualityCode *bool `json:"qualityCode,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.QualityCode = dataAO1.QualityCode

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AspectVariable) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.VariableDefinition)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		QualityCode *bool `json:"qualityCode,omitempty"`
	}

	dataAO1.QualityCode = m.QualityCode

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this aspect variable
func (m *AspectVariable) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VariableDefinition
	if err := m.VariableDefinition.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this aspect variable based on the context it is used
func (m *AspectVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VariableDefinition
	if err := m.VariableDefinition.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AspectVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AspectVariable) UnmarshalBinary(b []byte) error {
	var res AspectVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}