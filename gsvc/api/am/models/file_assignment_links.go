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

// FileAssignmentLinks file assignment links
//
// swagger:model FileAssignmentLinks
type FileAssignmentLinks struct {

	// download
	Download *FileAssignmentLinksDownload `json:"download,omitempty"`

	// metadata
	Metadata *FileAssignmentLinksMetadata `json:"metadata,omitempty"`

	// origin
	Origin *FileAssignmentLinksOrigin `json:"origin,omitempty"`
}

// Validate validates this file assignment links
func (m *FileAssignmentLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileAssignmentLinks) validateDownload(formats strfmt.Registry) error {
	if swag.IsZero(m.Download) { // not required
		return nil
	}

	if m.Download != nil {
		if err := m.Download.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("download")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("download")
			}
			return err
		}
	}

	return nil
}

func (m *FileAssignmentLinks) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *FileAssignmentLinks) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file assignment links based on the context it is used
func (m *FileAssignmentLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDownload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileAssignmentLinks) contextValidateDownload(ctx context.Context, formats strfmt.Registry) error {

	if m.Download != nil {
		if err := m.Download.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("download")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("download")
			}
			return err
		}
	}

	return nil
}

func (m *FileAssignmentLinks) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *FileAssignmentLinks) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.Origin != nil {
		if err := m.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileAssignmentLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileAssignmentLinks) UnmarshalBinary(b []byte) error {
	var res FileAssignmentLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileAssignmentLinksDownload file assignment links download
//
// swagger:model FileAssignmentLinksDownload
type FileAssignmentLinksDownload struct {

	// Link to download the file
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`
}

// Validate validates this file assignment links download
func (m *FileAssignmentLinksDownload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileAssignmentLinksDownload) validateHref(formats strfmt.Registry) error {
	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("download"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file assignment links download based on context it is used
func (m *FileAssignmentLinksDownload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileAssignmentLinksDownload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileAssignmentLinksDownload) UnmarshalBinary(b []byte) error {
	var res FileAssignmentLinksDownload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileAssignmentLinksMetadata file assignment links metadata
//
// swagger:model FileAssignmentLinksMetadata
type FileAssignmentLinksMetadata struct {

	// Link to get metadata of the file
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`
}

// Validate validates this file assignment links metadata
func (m *FileAssignmentLinksMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileAssignmentLinksMetadata) validateHref(formats strfmt.Registry) error {
	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("metadata"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file assignment links metadata based on context it is used
func (m *FileAssignmentLinksMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileAssignmentLinksMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileAssignmentLinksMetadata) UnmarshalBinary(b []byte) error {
	var res FileAssignmentLinksMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileAssignmentLinksOrigin file assignment links origin
//
// swagger:model FileAssignmentLinksOrigin
type FileAssignmentLinksOrigin struct {

	// Link to access the file assignment. Only visible if assignment is inherited
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`
}

// Validate validates this file assignment links origin
func (m *FileAssignmentLinksOrigin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileAssignmentLinksOrigin) validateHref(formats strfmt.Registry) error {
	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("origin"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file assignment links origin based on context it is used
func (m *FileAssignmentLinksOrigin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileAssignmentLinksOrigin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileAssignmentLinksOrigin) UnmarshalBinary(b []byte) error {
	var res FileAssignmentLinksOrigin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}