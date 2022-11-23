// Code generated by go-swagger; DO NOT EDIT.

package agm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TooManyRequests too many requests
//
// swagger:model TooManyRequests
type TooManyRequests struct {

	// error
	// Example: too_many_requests
	Error string `json:"error,omitempty"`

	// An error message with Correlation-ID value.
	// Example: [d6270fa4-f8f2-46d7-8370-1fbcacb37c52] There is already an ongoing registration process for the agent.
	ErrorDescription string `json:"error_description,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this too many requests
func (m *TooManyRequests) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this too many requests based on context it is used
func (m *TooManyRequests) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TooManyRequests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TooManyRequests) UnmarshalBinary(b []byte) error {
	var res TooManyRequests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
