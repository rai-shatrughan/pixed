// Code generated by go-swagger; DO NOT EDIT.

package agm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Notfound notfound
//
// swagger:model Notfound
type Notfound struct {

	// id
	ID string `json:"id,omitempty"`

	// message
	// Example: Resource not found.
	Message string `json:"message,omitempty"`
}

// Validate validates this notfound
func (m *Notfound) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notfound based on context it is used
func (m *Notfound) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Notfound) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notfound) UnmarshalBinary(b []byte) error {
	var res Notfound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
