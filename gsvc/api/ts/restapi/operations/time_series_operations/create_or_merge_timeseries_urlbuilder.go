// Code generated by go-swagger; DO NOT EDIT.

package time_series_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// CreateOrMergeTimeseriesURL generates an URL for the create or merge timeseries operation
type CreateOrMergeTimeseriesURL struct {
	EntityID        string
	PropertySetName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreateOrMergeTimeseriesURL) WithBasePath(bp string) *CreateOrMergeTimeseriesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreateOrMergeTimeseriesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *CreateOrMergeTimeseriesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/timeseries/{entityId}/{propertySetName}"

	entityID := o.EntityID
	if entityID != "" {
		_path = strings.Replace(_path, "{entityId}", entityID, -1)
	} else {
		return nil, errors.New("entityId is required on CreateOrMergeTimeseriesURL")
	}

	propertySetName := o.PropertySetName
	if propertySetName != "" {
		_path = strings.Replace(_path, "{propertySetName}", propertySetName, -1)
	} else {
		return nil, errors.New("propertySetName is required on CreateOrMergeTimeseriesURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/iottimeseries/v3"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *CreateOrMergeTimeseriesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *CreateOrMergeTimeseriesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *CreateOrMergeTimeseriesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on CreateOrMergeTimeseriesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on CreateOrMergeTimeseriesURL")
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
func (o *CreateOrMergeTimeseriesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
