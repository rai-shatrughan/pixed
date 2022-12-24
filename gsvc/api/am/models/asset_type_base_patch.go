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

// AssetTypeBasePatch asset type base patch
//
// swagger:model AssetTypeBasePatch
type AssetTypeBasePatch struct {

	// description
	// Example: Hyperspace jump capable space ship
	// Max Length: 255
	Description string `json:"description,omitempty"`

	// If instances can be created from this type. A non-instantiable type could be changed to be instantiable but not the other way around.
	// Example: true
	Instantiable *bool `json:"instantiable,omitempty"`

	// The type's name.
	// Example: X_Wing
	// Max Length: 128
	// Min Length: 1
	// Pattern: [\p{L}_0-9_\. ]+
	Name string `json:"name,omitempty"`

	// parent type Id
	ParentTypeID AssetTypeID `json:"parentTypeId,omitempty"`

	// Visibility of the assettype. Setting this property to public makes it available to other tenants. Private types are only visible to the user's own tenant. Currently only private types could be created.
	// Example: private
	// Enum: [public private]
	Scope *string `json:"scope,omitempty"`
}

// Validate validates this asset type base patch
func (m *AssetTypeBasePatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetTypeBasePatch) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 255); err != nil {
		return err
	}

	return nil
}

func (m *AssetTypeBasePatch) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 128); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `[\p{L}_0-9_\. ]+`); err != nil {
		return err
	}

	return nil
}

func (m *AssetTypeBasePatch) validateParentTypeID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentTypeID) { // not required
		return nil
	}

	if err := m.ParentTypeID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("parentTypeId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("parentTypeId")
		}
		return err
	}

	return nil
}

var assetTypeBasePatchTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetTypeBasePatchTypeScopePropEnum = append(assetTypeBasePatchTypeScopePropEnum, v)
	}
}

const (

	// AssetTypeBasePatchScopePublic captures enum value "public"
	AssetTypeBasePatchScopePublic string = "public"

	// AssetTypeBasePatchScopePrivate captures enum value "private"
	AssetTypeBasePatchScopePrivate string = "private"
)

// prop value enum
func (m *AssetTypeBasePatch) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetTypeBasePatchTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AssetTypeBasePatch) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this asset type base patch based on the context it is used
func (m *AssetTypeBasePatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParentTypeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetTypeBasePatch) contextValidateParentTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ParentTypeID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("parentTypeId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("parentTypeId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetTypeBasePatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetTypeBasePatch) UnmarshalBinary(b []byte) error {
	var res AssetTypeBasePatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
