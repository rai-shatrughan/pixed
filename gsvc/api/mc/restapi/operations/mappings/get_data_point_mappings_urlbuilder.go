// Code generated by go-swagger; DO NOT EDIT.

package mappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetDataPointMappingsURL generates an URL for the get data point mappings operation
type GetDataPointMappingsURL struct {
	Filter *string
	Page   *int32
	Size   *int32
	Sort   *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetDataPointMappingsURL) WithBasePath(bp string) *GetDataPointMappingsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetDataPointMappingsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetDataPointMappingsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/dataPointMappings"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/mindconnect/v3"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var filterQ string
	if o.Filter != nil {
		filterQ = *o.Filter
	}
	if filterQ != "" {
		qs.Set("filter", filterQ)
	}

	var pageQ string
	if o.Page != nil {
		pageQ = swag.FormatInt32(*o.Page)
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	var sizeQ string
	if o.Size != nil {
		sizeQ = swag.FormatInt32(*o.Size)
	}
	if sizeQ != "" {
		qs.Set("size", sizeQ)
	}

	var sortQ string
	if o.Sort != nil {
		sortQ = *o.Sort
	}
	if sortQ != "" {
		qs.Set("sort", sortQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetDataPointMappingsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetDataPointMappingsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetDataPointMappingsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetDataPointMappingsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetDataPointMappingsURL")
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
func (o *GetDataPointMappingsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
