// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// DeleteAssetFileAssigmentURL generates an URL for the delete asset file assigment operation
type DeleteAssetFileAssigmentURL struct {
	ID  string
	Key string

	IncludeShared *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteAssetFileAssigmentURL) WithBasePath(bp string) *DeleteAssetFileAssigmentURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteAssetFileAssigmentURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteAssetFileAssigmentURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/assets/{id}/fileAssignments/{key}"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on DeleteAssetFileAssigmentURL")
	}

	key := o.Key
	if key != "" {
		_path = strings.Replace(_path, "{key}", key, -1)
	} else {
		return nil, errors.New("key is required on DeleteAssetFileAssigmentURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/assetmanagement/v3"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var includeSharedQ string
	if o.IncludeShared != nil {
		includeSharedQ = swag.FormatBool(*o.IncludeShared)
	}
	if includeSharedQ != "" {
		qs.Set("includeShared", includeSharedQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteAssetFileAssigmentURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteAssetFileAssigmentURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteAssetFileAssigmentURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteAssetFileAssigmentURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteAssetFileAssigmentURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteAssetFileAssigmentURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}