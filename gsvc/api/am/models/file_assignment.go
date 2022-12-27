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

// FileAssignment file assignment
//
// swagger:model FileAssignment
type FileAssignment struct {

	// The id of the file to be assigned
	// Example: c27a28b6eb16b507fabc11e75da8b4ce
	FileID string `json:"fileId,omitempty"`

	// Keyword for the file to be assigned to an asset.
	// Example: logo_small
	// Pattern: [\w]+
	Key string `json:"key,omitempty"`
}

// Validate validates this file assignment
func (m *FileAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileAssignment) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := validate.Pattern("key", "body", m.Key, `[\w]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file assignment based on context it is used
func (m *FileAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileAssignment) UnmarshalBinary(b []byte) error {
	var res FileAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}