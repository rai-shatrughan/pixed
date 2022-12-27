// Code generated by go-swagger; DO NOT EDIT.

package token_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PostOauthTokenMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var PostOauthTokenMaxParseMemory int64 = 32 << 20

// NewPostOauthTokenParams creates a new PostOauthTokenParams object
//
// There are no default values defined in the spec.
func NewPostOauthTokenParams() PostOauthTokenParams {

	return PostOauthTokenParams{}
}

// PostOauthTokenParams contains all the bound params for the post oauth token operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostOauthToken
type PostOauthTokenParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A jwt which is signed with client secret

	Signing keys (client secret) can vary depending on agent security profile:<br>
	    - __SHARED_SECRET__: Client secret is provided by '/register' | ‘/register/<client_id>’ endpoint
	    - __RSA_3072__: Private part of the client's RSA key which the public part was provided at '/register' | '/register/<client_id>’
	    - __CACertifiedX509__: Private counterpart of device certificate’s public key.

	  Required: true
	  In: formData
	*/
	ClientAssertion string
	/*Defines the assertion type, only urn:ietf:params:oauth:client-assertion-type:jwt-bearer is supported.
	  Required: true
	  In: formData
	*/
	ClientAssertionType string
	/*The type of authentication being used to obtain the token, only  client_credentials is supported.
	  Required: true
	  In: formData
	*/
	GrantType string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostOauthTokenParams() beforehand.
func (o *PostOauthTokenParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(PostOauthTokenMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdClientAssertion, fdhkClientAssertion, _ := fds.GetOK("client_assertion")
	if err := o.bindClientAssertion(fdClientAssertion, fdhkClientAssertion, route.Formats); err != nil {
		res = append(res, err)
	}

	fdClientAssertionType, fdhkClientAssertionType, _ := fds.GetOK("client_assertion_type")
	if err := o.bindClientAssertionType(fdClientAssertionType, fdhkClientAssertionType, route.Formats); err != nil {
		res = append(res, err)
	}

	fdGrantType, fdhkGrantType, _ := fds.GetOK("grant_type")
	if err := o.bindGrantType(fdGrantType, fdhkGrantType, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClientAssertion binds and validates parameter ClientAssertion from formData.
func (o *PostOauthTokenParams) bindClientAssertion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("client_assertion", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("client_assertion", "formData", raw); err != nil {
		return err
	}
	o.ClientAssertion = raw

	return nil
}

// bindClientAssertionType binds and validates parameter ClientAssertionType from formData.
func (o *PostOauthTokenParams) bindClientAssertionType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("client_assertion_type", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("client_assertion_type", "formData", raw); err != nil {
		return err
	}
	o.ClientAssertionType = raw

	return nil
}

// bindGrantType binds and validates parameter GrantType from formData.
func (o *PostOauthTokenParams) bindGrantType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("grant_type", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("grant_type", "formData", raw); err != nil {
		return err
	}
	o.GrantType = raw

	return nil
}