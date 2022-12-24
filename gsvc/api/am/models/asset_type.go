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

// AssetType asset type
//
// swagger:model AssetType
type AssetType struct {
	AssetTypeBase

	// Aspects of the asset-type.
	Aspects []*AssetTypeAspectsItems0 `json:"aspects"`

	// file assignments
	FileAssignments []*FileAssignment `json:"fileAssignments"`

	// Direct variables of the asset-type. Variable names has to be unique inside the whole type-family (ancestors and descendants). Once added variables cannot be changed or removed.
	Variables []*VariableDefinition `json:"variables"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetType) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AssetTypeBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AssetTypeBase = aO0

	// AO1
	var dataAO1 struct {
		Aspects []*AssetTypeAspectsItems0 `json:"aspects"`

		FileAssignments []*FileAssignment `json:"fileAssignments"`

		Variables []*VariableDefinition `json:"variables"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Aspects = dataAO1.Aspects

	m.FileAssignments = dataAO1.FileAssignments

	m.Variables = dataAO1.Variables

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetType) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AssetTypeBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Aspects []*AssetTypeAspectsItems0 `json:"aspects"`

		FileAssignments []*FileAssignment `json:"fileAssignments"`

		Variables []*VariableDefinition `json:"variables"`
	}

	dataAO1.Aspects = m.Aspects

	dataAO1.FileAssignments = m.FileAssignments

	dataAO1.Variables = m.Variables

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset type
func (m *AssetType) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AssetTypeBase
	if err := m.AssetTypeBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAspects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetType) validateAspects(formats strfmt.Registry) error {

	if swag.IsZero(m.Aspects) { // not required
		return nil
	}

	for i := 0; i < len(m.Aspects); i++ {
		if swag.IsZero(m.Aspects[i]) { // not required
			continue
		}

		if m.Aspects[i] != nil {
			if err := m.Aspects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aspects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aspects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssetType) validateFileAssignments(formats strfmt.Registry) error {

	if swag.IsZero(m.FileAssignments) { // not required
		return nil
	}

	for i := 0; i < len(m.FileAssignments); i++ {
		if swag.IsZero(m.FileAssignments[i]) { // not required
			continue
		}

		if m.FileAssignments[i] != nil {
			if err := m.FileAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssetType) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this asset type based on the context it is used
func (m *AssetType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AssetTypeBase
	if err := m.AssetTypeBase.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAspects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetType) contextValidateAspects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Aspects); i++ {

		if m.Aspects[i] != nil {
			if err := m.Aspects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aspects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aspects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssetType) contextValidateFileAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileAssignments); i++ {

		if m.FileAssignments[i] != nil {
			if err := m.FileAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssetType) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variables); i++ {

		if m.Variables[i] != nil {
			if err := m.Variables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetType) UnmarshalBinary(b []byte) error {
	var res AssetType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AssetTypeAspectsItems0 asset type aspects items0
//
// swagger:model AssetTypeAspectsItems0
type AssetTypeAspectsItems0 struct {

	// aspect Id
	AspectID AspectID `json:"aspectId,omitempty"`

	// aspect type Id
	// Required: true
	AspectTypeID *AspectTypeID `json:"aspectTypeId"`

	// Name of the aspect. It has to be unique inside the type-family (ancestors and descendants). Reserved words (id, name, description, tenant, etag, scope, properties, propertySets, extends, variables, aspects, parentTypeId) cannot be used as aspect names.
	// Example: leftWing
	// Required: true
	// Max Length: 64
	// Min Length: 1
	// Pattern: [a-zA-Z0-9_]+
	Name *string `json:"name"`
}

// Validate validates this asset type aspects items0
func (m *AssetTypeAspectsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAspectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAspectTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetTypeAspectsItems0) validateAspectID(formats strfmt.Registry) error {
	if swag.IsZero(m.AspectID) { // not required
		return nil
	}

	if err := m.AspectID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aspectId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("aspectId")
		}
		return err
	}

	return nil
}

func (m *AssetTypeAspectsItems0) validateAspectTypeID(formats strfmt.Registry) error {

	if err := validate.Required("aspectTypeId", "body", m.AspectTypeID); err != nil {
		return err
	}

	if err := validate.Required("aspectTypeId", "body", m.AspectTypeID); err != nil {
		return err
	}

	if m.AspectTypeID != nil {
		if err := m.AspectTypeID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aspectTypeId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aspectTypeId")
			}
			return err
		}
	}

	return nil
}

func (m *AssetTypeAspectsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[a-zA-Z0-9_]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this asset type aspects items0 based on the context it is used
func (m *AssetTypeAspectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAspectID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAspectTypeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetTypeAspectsItems0) contextValidateAspectID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AspectID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aspectId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("aspectId")
		}
		return err
	}

	return nil
}

func (m *AssetTypeAspectsItems0) contextValidateAspectTypeID(ctx context.Context, formats strfmt.Registry) error {

	if m.AspectTypeID != nil {
		if err := m.AspectTypeID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aspectTypeId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aspectTypeId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetTypeAspectsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetTypeAspectsItems0) UnmarshalBinary(b []byte) error {
	var res AssetTypeAspectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
