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

// VariableUpdate variable update
//
// swagger:model VariableUpdate
type VariableUpdate struct {

	// The default value of the variable. It must be compatible with the dataType! The default value will be inherited by the asset type's child types and by the asset instantiating it. It can be defined in aspect types and asset types.
	// Example: 25/77
	// Max Length: 255
	// Min Length: 1
	DefaultValue string `json:"defaultValue,omitempty"`

	// The max length of the variable's value. The length field is only used for variables of string or big_string dataType. Max length for string is 255 and max length for big_string 100000. Cannot be changed.
	// Example: 5
	// Maximum: 100000
	// Minimum: 1
	Length int64 `json:"length,omitempty"`

	// Name of the variable. Reserved words (id, name, description, tenant, etag, scope, properties, propertySets, extends, variables, aspects, parentTypeId) cannot be used as variable names.
	// Example: temperature
	// Max Length: 64
	// Min Length: 1
	// Pattern: [a-zA-Z_][a-zA-Z0-9_]*
	Name string `json:"name,omitempty"`

	// Unit of measurement. Can be changed
	// Example: C/F
	// Max Length: 32
	// Pattern: [^\']*
	Unit string `json:"unit,omitempty"`
}

// Validate validates this variable update
func (m *VariableUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableUpdate) validateDefaultValue(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultValue) { // not required
		return nil
	}

	if err := validate.MinLength("defaultValue", "body", m.DefaultValue, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("defaultValue", "body", m.DefaultValue, 255); err != nil {
		return err
	}

	return nil
}

func (m *VariableUpdate) validateLength(formats strfmt.Registry) error {
	if swag.IsZero(m.Length) { // not required
		return nil
	}

	if err := validate.MinimumInt("length", "body", m.Length, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("length", "body", m.Length, 100000, false); err != nil {
		return err
	}

	return nil
}

func (m *VariableUpdate) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 64); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `[a-zA-Z_][a-zA-Z0-9_]*`); err != nil {
		return err
	}

	return nil
}

func (m *VariableUpdate) validateUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if err := validate.MaxLength("unit", "body", m.Unit, 32); err != nil {
		return err
	}

	if err := validate.Pattern("unit", "body", m.Unit, `[^\']*`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this variable update based on context it is used
func (m *VariableUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VariableUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableUpdate) UnmarshalBinary(b []byte) error {
	var res VariableUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
