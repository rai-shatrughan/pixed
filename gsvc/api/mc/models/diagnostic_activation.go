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

// DiagnosticActivation diagnostic activation
//
// swagger:model DiagnosticActivation
type DiagnosticActivation struct {

	// Unique identifier of the agent
	// Example: 3b27818ea09a46b48c7eb3fbd878349f
	// Required: true
	// Max Length: 36
	AgentID *string `json:"agentId"`

	// Unique identifier of diagnostic activation resource
	// Example: 8f273818-e975-11e8-9f32-f2801f1b9fd1
	// Read Only: true
	// Max Length: 36
	ID string `json:"id,omitempty"`

	// Status of the activation
	// Example: ACTIVE
	// Enum: [ACTIVE INACTIVE]
	Status string `json:"status,omitempty"`
}

// Validate validates this diagnostic activation
func (m *DiagnosticActivation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiagnosticActivation) validateAgentID(formats strfmt.Registry) error {

	if err := validate.Required("agentId", "body", m.AgentID); err != nil {
		return err
	}

	if err := validate.MaxLength("agentId", "body", *m.AgentID, 36); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticActivation) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", m.ID, 36); err != nil {
		return err
	}

	return nil
}

var diagnosticActivationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diagnosticActivationTypeStatusPropEnum = append(diagnosticActivationTypeStatusPropEnum, v)
	}
}

const (

	// DiagnosticActivationStatusACTIVE captures enum value "ACTIVE"
	DiagnosticActivationStatusACTIVE string = "ACTIVE"

	// DiagnosticActivationStatusINACTIVE captures enum value "INACTIVE"
	DiagnosticActivationStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *DiagnosticActivation) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, diagnosticActivationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DiagnosticActivation) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this diagnostic activation based on the context it is used
func (m *DiagnosticActivation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiagnosticActivation) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiagnosticActivation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiagnosticActivation) UnmarshalBinary(b []byte) error {
	var res DiagnosticActivation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}