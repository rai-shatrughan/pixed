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

// PagedMapping paged mapping
//
// swagger:model PagedMapping
type PagedMapping struct {

	// content
	// Required: true
	Content []*Mapping `json:"content"`

	// Whether the current item is the first one.
	// Example: true
	// Required: true
	First *bool `json:"first"`

	// Whether the current item is the last one.
	// Example: true
	// Required: true
	Last *bool `json:"last"`

	// The number of the current item.
	// Example: 0
	// Required: true
	Number *int64 `json:"number"`

	// The number of elements currently on this page.
	// Example: 1
	// Required: true
	NumberOfElements *int64 `json:"numberOfElements"`

	// The size of the page.
	// Example: 20
	// Required: true
	Size *int64 `json:"size"`

	// The sorting parameters for the page.
	// Required: true
	Sort []*Order `json:"sort"`

	// The total amount of elements.
	// Example: 1
	// Required: true
	TotalElements *int64 `json:"totalElements"`

	// The number of total pages.
	// Example: 1
	// Required: true
	TotalPages *int64 `json:"totalPages"`
}

// Validate validates this paged mapping
func (m *PagedMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalPages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagedMapping) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagedMapping) validateFirst(formats strfmt.Registry) error {

	if err := validate.Required("first", "body", m.First); err != nil {
		return err
	}

	return nil
}

func (m *PagedMapping) validateLast(formats strfmt.Registry) error {

	if err := validate.Required("last", "body", m.Last); err != nil {
		return err
	}

	return nil
}

func (m *PagedMapping) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *PagedMapping) validateNumberOfElements(formats strfmt.Registry) error {

	if err := validate.Required("numberOfElements", "body", m.NumberOfElements); err != nil {
		return err
	}

	return nil
}

func (m *PagedMapping) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *PagedMapping) validateSort(formats strfmt.Registry) error {

	if err := validate.Required("sort", "body", m.Sort); err != nil {
		return err
	}

	for i := 0; i < len(m.Sort); i++ {
		if swag.IsZero(m.Sort[i]) { // not required
			continue
		}

		if m.Sort[i] != nil {
			if err := m.Sort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagedMapping) validateTotalElements(formats strfmt.Registry) error {

	if err := validate.Required("totalElements", "body", m.TotalElements); err != nil {
		return err
	}

	return nil
}

func (m *PagedMapping) validateTotalPages(formats strfmt.Registry) error {

	if err := validate.Required("totalPages", "body", m.TotalPages); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this paged mapping based on the context it is used
func (m *PagedMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagedMapping) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Content); i++ {

		if m.Content[i] != nil {
			if err := m.Content[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagedMapping) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sort); i++ {

		if m.Sort[i] != nil {
			if err := m.Sort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PagedMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagedMapping) UnmarshalBinary(b []byte) error {
	var res PagedMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
