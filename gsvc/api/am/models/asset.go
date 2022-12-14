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

// Asset asset
//
// swagger:model Asset
type Asset struct {

	// aspects
	Aspects []*Aspect `json:"aspects"`

	// The description of the asset
	// Example: The ship of Han Solo and Chewbacca
	// Max Length: 255
	// Pattern: [^']*
	Description string `json:"description,omitempty"`

	// The id given by the user
	// Example: SN 123456-123-123456
	// Max Length: 255
	ExternalID string `json:"externalId,omitempty"`

	// file assignments
	FileAssignments []*FileAssignment `json:"fileAssignments"`

	// location
	Location *Location `json:"location,omitempty"`

	// Name of the asset
	// Example: Millenium Falcon
	// Max Length: 128
	// Min Length: 1
	// Pattern: [^\/\\]*
	Name string `json:"name,omitempty"`

	// parent Id
	// Required: true
	ParentID *UniqueID `json:"parentId"`

	// timezone
	Timezone Timezone `json:"timezone,omitempty"`

	// twin type
	TwinType *TwinType `json:"twinType,omitempty"`

	// type Id
	// Required: true
	TypeID *AssetTypeID `json:"typeId"`

	// variables
	Variables []*Variable `json:"variables"`
}

// Validate validates this asset
func (m *Asset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAspects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwinType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
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

func (m *Asset) validateAspects(formats strfmt.Registry) error {
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

func (m *Asset) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 255); err != nil {
		return err
	}

	if err := validate.Pattern("description", "body", m.Description, `[^']*`); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateExternalID(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalID) { // not required
		return nil
	}

	if err := validate.MaxLength("externalId", "body", m.ExternalID, 255); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateFileAssignments(formats strfmt.Registry) error {
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

func (m *Asset) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 128); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `[^\/\\]*`); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parentId", "body", m.ParentID); err != nil {
		return err
	}

	if err := validate.Required("parentId", "body", m.ParentID); err != nil {
		return err
	}

	if m.ParentID != nil {
		if err := m.ParentID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentId")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) validateTimezone(formats strfmt.Registry) error {
	if swag.IsZero(m.Timezone) { // not required
		return nil
	}

	if err := m.Timezone.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timezone")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("timezone")
		}
		return err
	}

	return nil
}

func (m *Asset) validateTwinType(formats strfmt.Registry) error {
	if swag.IsZero(m.TwinType) { // not required
		return nil
	}

	if m.TwinType != nil {
		if err := m.TwinType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("twinType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("twinType")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("typeId", "body", m.TypeID); err != nil {
		return err
	}

	if err := validate.Required("typeId", "body", m.TypeID); err != nil {
		return err
	}

	if m.TypeID != nil {
		if err := m.TypeID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typeId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typeId")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) validateVariables(formats strfmt.Registry) error {
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

// ContextValidate validate this asset based on the context it is used
func (m *Asset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAspects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimezone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTwinType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeID(ctx, formats); err != nil {
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

func (m *Asset) contextValidateAspects(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Asset) contextValidateFileAssignments(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Asset) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) contextValidateParentID(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentID != nil {
		if err := m.ParentID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentId")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) contextValidateTimezone(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Timezone.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timezone")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("timezone")
		}
		return err
	}

	return nil
}

func (m *Asset) contextValidateTwinType(ctx context.Context, formats strfmt.Registry) error {

	if m.TwinType != nil {
		if err := m.TwinType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("twinType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("twinType")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) contextValidateTypeID(ctx context.Context, formats strfmt.Registry) error {

	if m.TypeID != nil {
		if err := m.TypeID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typeId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typeId")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Asset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Asset) UnmarshalBinary(b []byte) error {
	var res Asset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
