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

// DiagnosticInformationMessage diagnostic information message
//
// swagger:model DiagnosticInformationMessage
type DiagnosticInformationMessage struct {

	// correlation Id
	// Example: 3fcf2a5ecc7611e7abc4cec278b6b50a
	// Max Length: 36
	CorrelationID string `json:"correlationId,omitempty"`

	// message
	// Example: [Finished] TimeSeries upload completed, \u003c1\u003e samplings and \u003c1\u003e properties in \u003c1\u003e requests\u003e. \u003c0\u003e values are dropped.
	// Max Length: 4096
	Message string `json:"message,omitempty"`

	// severity
	// Enum: [INFO WARN ERROR]
	Severity string `json:"severity,omitempty"`

	// Source of diagnostic information.
	// Example: TIMESERIES
	Source string `json:"source,omitempty"`

	// State of diagnostic information.
	// Enum: [ACCEPTED RETRYING DROPPED PROCESSING FINISHED]
	State string `json:"state,omitempty"`

	// Diagnostic information creation date.
	// Example: 2018-08-27T16:40:11.235Z
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this diagnostic information message
func (m *DiagnosticInformationMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCorrelationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiagnosticInformationMessage) validateCorrelationID(formats strfmt.Registry) error {
	if swag.IsZero(m.CorrelationID) { // not required
		return nil
	}

	if err := validate.MaxLength("correlationId", "body", m.CorrelationID, 36); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticInformationMessage) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if err := validate.MaxLength("message", "body", m.Message, 4096); err != nil {
		return err
	}

	return nil
}

var diagnosticInformationMessageTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INFO","WARN","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diagnosticInformationMessageTypeSeverityPropEnum = append(diagnosticInformationMessageTypeSeverityPropEnum, v)
	}
}

const (

	// DiagnosticInformationMessageSeverityINFO captures enum value "INFO"
	DiagnosticInformationMessageSeverityINFO string = "INFO"

	// DiagnosticInformationMessageSeverityWARN captures enum value "WARN"
	DiagnosticInformationMessageSeverityWARN string = "WARN"

	// DiagnosticInformationMessageSeverityERROR captures enum value "ERROR"
	DiagnosticInformationMessageSeverityERROR string = "ERROR"
)

// prop value enum
func (m *DiagnosticInformationMessage) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, diagnosticInformationMessageTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DiagnosticInformationMessage) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

var diagnosticInformationMessageTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCEPTED","RETRYING","DROPPED","PROCESSING","FINISHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diagnosticInformationMessageTypeStatePropEnum = append(diagnosticInformationMessageTypeStatePropEnum, v)
	}
}

const (

	// DiagnosticInformationMessageStateACCEPTED captures enum value "ACCEPTED"
	DiagnosticInformationMessageStateACCEPTED string = "ACCEPTED"

	// DiagnosticInformationMessageStateRETRYING captures enum value "RETRYING"
	DiagnosticInformationMessageStateRETRYING string = "RETRYING"

	// DiagnosticInformationMessageStateDROPPED captures enum value "DROPPED"
	DiagnosticInformationMessageStateDROPPED string = "DROPPED"

	// DiagnosticInformationMessageStatePROCESSING captures enum value "PROCESSING"
	DiagnosticInformationMessageStatePROCESSING string = "PROCESSING"

	// DiagnosticInformationMessageStateFINISHED captures enum value "FINISHED"
	DiagnosticInformationMessageStateFINISHED string = "FINISHED"
)

// prop value enum
func (m *DiagnosticInformationMessage) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, diagnosticInformationMessageTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DiagnosticInformationMessage) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticInformationMessage) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this diagnostic information message based on context it is used
func (m *DiagnosticInformationMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DiagnosticInformationMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiagnosticInformationMessage) UnmarshalBinary(b []byte) error {
	var res DiagnosticInformationMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
