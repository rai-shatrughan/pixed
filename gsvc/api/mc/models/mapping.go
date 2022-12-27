// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Mapping mapping
//
// swagger:model Mapping
type Mapping struct {

	// Unique identifier of the agent
	// Example: 11961bc396cd4a87a9b26b723f5b7ba0
	// Required: true
	// Max Length: 36
	// Format: uuid
	AgentID *strfmt.UUID `json:"agentId"`

	// Unique identifier of the data point
	// Example: DP0001
	// Required: true
	// Max Length: 36
	DataPointID *string `json:"dataPointId"`

	// Type of the data point
	// Read Only: true
	// Enum: [INT LONG DOUBLE BOOLEAN STRING BIG_STRING TIMESTAMP]
	DataPointType string `json:"dataPointType,omitempty"`

	// Unit of the data point
	// Example: %
	// Read Only: true
	// Max Length: 32
	DataPointUnit string `json:"dataPointUnit,omitempty"`

	// Unique identifier of the entity
	// Example: 83e78008eadf453bae4f5c7bef3db550
	// Required: true
	// Max Length: 36
	// Format: uuid
	EntityID *strfmt.UUID `json:"entityId"`

	// Unique identifier of the mapping resource
	// Example: 4fad6258-5def-4d84-a4c2-1481b209c116
	// Read Only: true
	// Max Length: 36
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Identifies auto deleting mapping or keeping mapping.
	// Example: false
	KeepMapping *bool `json:"keepMapping,omitempty"`

	// property name
	// Example: Voltage
	// Required: true
	// Max Length: 256
	PropertyName *string `json:"propertyName"`

	// property set name
	// Example: ElectricalProperties
	// Required: true
	// Max Length: 256
	PropertySetName *string `json:"propertySetName"`

	// property type
	// Read Only: true
	// Enum: [INT LONG DOUBLE BOOLEAN STRING BIG_STRING TIMESTAMP]
	PropertyType string `json:"propertyType,omitempty"`

	// property unit
	// Example: %
	// Read Only: true
	// Max Length: 32
	PropertyUnit string `json:"propertyUnit,omitempty"`

	// quality enabled
	// Read Only: true
	QualityEnabled *bool `json:"qualityEnabled,omitempty"`

	// validity
	// Read Only: true
	Validity struct {
		Validity
	} `json:"validity,omitempty"`
}

// Validate validates this mapping
func (m *Mapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPointID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPointType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPointUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertySetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mapping) validateAgentID(formats strfmt.Registry) error {

	if err := validate.Required("agentId", "body", m.AgentID); err != nil {
		return err
	}

	if err := validate.MaxLength("agentId", "body", m.AgentID.String(), 36); err != nil {
		return err
	}

	if err := validate.FormatOf("agentId", "body", "uuid", m.AgentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validateDataPointID(formats strfmt.Registry) error {

	if err := validate.Required("dataPointId", "body", m.DataPointID); err != nil {
		return err
	}

	if err := validate.MaxLength("dataPointId", "body", *m.DataPointID, 36); err != nil {
		return err
	}

	return nil
}

var mappingTypeDataPointTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INT","LONG","DOUBLE","BOOLEAN","STRING","BIG_STRING","TIMESTAMP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mappingTypeDataPointTypePropEnum = append(mappingTypeDataPointTypePropEnum, v)
	}
}

const (

	// MappingDataPointTypeINT captures enum value "INT"
	MappingDataPointTypeINT string = "INT"

	// MappingDataPointTypeLONG captures enum value "LONG"
	MappingDataPointTypeLONG string = "LONG"

	// MappingDataPointTypeDOUBLE captures enum value "DOUBLE"
	MappingDataPointTypeDOUBLE string = "DOUBLE"

	// MappingDataPointTypeBOOLEAN captures enum value "BOOLEAN"
	MappingDataPointTypeBOOLEAN string = "BOOLEAN"

	// MappingDataPointTypeSTRING captures enum value "STRING"
	MappingDataPointTypeSTRING string = "STRING"

	// MappingDataPointTypeBIGSTRING captures enum value "BIG_STRING"
	MappingDataPointTypeBIGSTRING string = "BIG_STRING"

	// MappingDataPointTypeTIMESTAMP captures enum value "TIMESTAMP"
	MappingDataPointTypeTIMESTAMP string = "TIMESTAMP"
)

// prop value enum
func (m *Mapping) validateDataPointTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mappingTypeDataPointTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Mapping) validateDataPointType(formats strfmt.Registry) error {
	if swag.IsZero(m.DataPointType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataPointTypeEnum("dataPointType", "body", m.DataPointType); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validateDataPointUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.DataPointUnit) { // not required
		return nil
	}

	if err := validate.MaxLength("dataPointUnit", "body", m.DataPointUnit, 32); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entityId", "body", m.EntityID); err != nil {
		return err
	}

	if err := validate.MaxLength("entityId", "body", m.EntityID.String(), 36); err != nil {
		return err
	}

	if err := validate.FormatOf("entityId", "body", "uuid", m.EntityID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", m.ID.String(), 36); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validatePropertyName(formats strfmt.Registry) error {

	if err := validate.Required("propertyName", "body", m.PropertyName); err != nil {
		return err
	}

	if err := validate.MaxLength("propertyName", "body", *m.PropertyName, 256); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validatePropertySetName(formats strfmt.Registry) error {

	if err := validate.Required("propertySetName", "body", m.PropertySetName); err != nil {
		return err
	}

	if err := validate.MaxLength("propertySetName", "body", *m.PropertySetName, 256); err != nil {
		return err
	}

	return nil
}

var mappingTypePropertyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INT","LONG","DOUBLE","BOOLEAN","STRING","BIG_STRING","TIMESTAMP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mappingTypePropertyTypePropEnum = append(mappingTypePropertyTypePropEnum, v)
	}
}

const (

	// MappingPropertyTypeINT captures enum value "INT"
	MappingPropertyTypeINT string = "INT"

	// MappingPropertyTypeLONG captures enum value "LONG"
	MappingPropertyTypeLONG string = "LONG"

	// MappingPropertyTypeDOUBLE captures enum value "DOUBLE"
	MappingPropertyTypeDOUBLE string = "DOUBLE"

	// MappingPropertyTypeBOOLEAN captures enum value "BOOLEAN"
	MappingPropertyTypeBOOLEAN string = "BOOLEAN"

	// MappingPropertyTypeSTRING captures enum value "STRING"
	MappingPropertyTypeSTRING string = "STRING"

	// MappingPropertyTypeBIGSTRING captures enum value "BIG_STRING"
	MappingPropertyTypeBIGSTRING string = "BIG_STRING"

	// MappingPropertyTypeTIMESTAMP captures enum value "TIMESTAMP"
	MappingPropertyTypeTIMESTAMP string = "TIMESTAMP"
)

// prop value enum
func (m *Mapping) validatePropertyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mappingTypePropertyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Mapping) validatePropertyType(formats strfmt.Registry) error {
	if swag.IsZero(m.PropertyType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePropertyTypeEnum("propertyType", "body", m.PropertyType); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validatePropertyUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.PropertyUnit) { // not required
		return nil
	}

	if err := validate.MaxLength("propertyUnit", "body", m.PropertyUnit, 32); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validateValidity(formats strfmt.Registry) error {
	if swag.IsZero(m.Validity) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this mapping based on the context it is used
func (m *Mapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataPointType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataPointUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePropertyType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePropertyUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQualityEnabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mapping) contextValidateDataPointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataPointType", "body", string(m.DataPointType)); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) contextValidateDataPointUnit(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataPointUnit", "body", string(m.DataPointUnit)); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) contextValidatePropertyType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "propertyType", "body", string(m.PropertyType)); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) contextValidatePropertyUnit(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "propertyUnit", "body", string(m.PropertyUnit)); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) contextValidateQualityEnabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "qualityEnabled", "body", m.QualityEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) contextValidateValidity(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *Mapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mapping) UnmarshalBinary(b []byte) error {
	var res Mapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
