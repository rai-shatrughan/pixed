// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// TwinType Indicates that the asset is a real asset (performance) or for simulation. If omitted on creation then it defaults to performance. Setting the twinType to simulation allows high resolution timestamps (microsecond precision).
//
// swagger:model twinType
type TwinType string

func NewTwinType(value TwinType) *TwinType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TwinType.
func (m TwinType) Pointer() *TwinType {
	return &m
}

const (

	// TwinTypePerformance captures enum value "performance"
	TwinTypePerformance TwinType = "performance"

	// TwinTypeSimulation captures enum value "simulation"
	TwinTypeSimulation TwinType = "simulation"
)

// for schema
var twinTypeEnum []interface{}

func init() {
	var res []TwinType
	if err := json.Unmarshal([]byte(`["performance","simulation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		twinTypeEnum = append(twinTypeEnum, v)
	}
}

func (m TwinType) validateTwinTypeEnum(path, location string, value TwinType) error {
	if err := validate.EnumCase(path, location, value, twinTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this twin type
func (m TwinType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTwinTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this twin type based on context it is used
func (m TwinType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
