// Code generated by go-swagger; DO NOT EDIT.

package agm

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

// CreateAgentRequest create agent request
//
// swagger:model CreateAgentRequest
type CreateAgentRequest struct {

	// Unique identifier of the entity
	// Example: 3b27818ea09a46b48c7eb3fbd878349f
	// Required: true
	// Max Length: 36
	EntityID *string `json:"entityId"`

	// Name must be unique per tenant.
	// Example: Nanobox Agent
	// Required: true
	// Max Length: 128
	Name *string `json:"name"`

	// security profile
	// Required: true
	// Enum: [SHARED_SECRET RSA_3072]
	SecurityProfile *string `json:"securityProfile"`
}

// Validate validates this create agent request
func (m *CreateAgentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAgentRequest) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entityId", "body", m.EntityID); err != nil {
		return err
	}

	if err := validate.MaxLength("entityId", "body", *m.EntityID, 36); err != nil {
		return err
	}

	return nil
}

func (m *CreateAgentRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 128); err != nil {
		return err
	}

	return nil
}

var createAgentRequestTypeSecurityProfilePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SHARED_SECRET","RSA_3072"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAgentRequestTypeSecurityProfilePropEnum = append(createAgentRequestTypeSecurityProfilePropEnum, v)
	}
}

const (

	// CreateAgentRequestSecurityProfileSHAREDSECRET captures enum value "SHARED_SECRET"
	CreateAgentRequestSecurityProfileSHAREDSECRET string = "SHARED_SECRET"

	// CreateAgentRequestSecurityProfileRSA3072 captures enum value "RSA_3072"
	CreateAgentRequestSecurityProfileRSA3072 string = "RSA_3072"
)

// prop value enum
func (m *CreateAgentRequest) validateSecurityProfileEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAgentRequestTypeSecurityProfilePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateAgentRequest) validateSecurityProfile(formats strfmt.Registry) error {

	if err := validate.Required("securityProfile", "body", m.SecurityProfile); err != nil {
		return err
	}

	// value enum
	if err := m.validateSecurityProfileEnum("securityProfile", "body", *m.SecurityProfile); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create agent request based on context it is used
func (m *CreateAgentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAgentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAgentRequest) UnmarshalBinary(b []byte) error {
	var res CreateAgentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
