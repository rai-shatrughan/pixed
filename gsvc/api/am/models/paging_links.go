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

// PagingLinks paging links
//
// swagger:model PagingLinks
type PagingLinks struct {

	// first
	First *PagingLinksFirst `json:"first,omitempty"`

	// last
	Last *PagingLinksLast `json:"last,omitempty"`

	// next
	Next *PagingLinksNext `json:"next,omitempty"`

	// prev
	Prev *PagingLinksPrev `json:"prev,omitempty"`

	// self
	Self *RelSelf `json:"self,omitempty"`
}

// Validate validates this paging links
func (m *PagingLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagingLinks) validateFirst(formats strfmt.Registry) error {
	if swag.IsZero(m.First) { // not required
		return nil
	}

	if m.First != nil {
		if err := m.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *PagingLinks) validateLast(formats strfmt.Registry) error {
	if swag.IsZero(m.Last) { // not required
		return nil
	}

	if m.Last != nil {
		if err := m.Last.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last")
			}
			return err
		}
	}

	return nil
}

func (m *PagingLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *PagingLinks) validatePrev(formats strfmt.Registry) error {
	if swag.IsZero(m.Prev) { // not required
		return nil
	}

	if m.Prev != nil {
		if err := m.Prev.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *PagingLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this paging links based on the context it is used
func (m *PagingLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFirst(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLast(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrev(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagingLinks) contextValidateFirst(ctx context.Context, formats strfmt.Registry) error {

	if m.First != nil {
		if err := m.First.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *PagingLinks) contextValidateLast(ctx context.Context, formats strfmt.Registry) error {

	if m.Last != nil {
		if err := m.Last.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last")
			}
			return err
		}
	}

	return nil
}

func (m *PagingLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *PagingLinks) contextValidatePrev(ctx context.Context, formats strfmt.Registry) error {

	if m.Prev != nil {
		if err := m.Prev.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *PagingLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PagingLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagingLinks) UnmarshalBinary(b []byte) error {
	var res PagingLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PagingLinksFirst link to first page
//
// swagger:model PagingLinksFirst
type PagingLinksFirst struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this paging links first
func (m *PagingLinksFirst) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this paging links first based on context it is used
func (m *PagingLinksFirst) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PagingLinksFirst) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagingLinksFirst) UnmarshalBinary(b []byte) error {
	var res PagingLinksFirst
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PagingLinksLast link to last page
//
// swagger:model PagingLinksLast
type PagingLinksLast struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this paging links last
func (m *PagingLinksLast) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this paging links last based on context it is used
func (m *PagingLinksLast) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PagingLinksLast) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagingLinksLast) UnmarshalBinary(b []byte) error {
	var res PagingLinksLast
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PagingLinksNext link to next page
//
// swagger:model PagingLinksNext
type PagingLinksNext struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this paging links next
func (m *PagingLinksNext) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this paging links next based on context it is used
func (m *PagingLinksNext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PagingLinksNext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagingLinksNext) UnmarshalBinary(b []byte) error {
	var res PagingLinksNext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PagingLinksPrev link to previous page
//
// swagger:model PagingLinksPrev
type PagingLinksPrev struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this paging links prev
func (m *PagingLinksPrev) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this paging links prev based on context it is used
func (m *PagingLinksPrev) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PagingLinksPrev) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagingLinksPrev) UnmarshalBinary(b []byte) error {
	var res PagingLinksPrev
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
