// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"gsvc/api/ts/models"
	"gsvc/api/ts/restapi/operations/time_series_operations"
)

// NewTsAPI creates a new Ts instance
func NewTsAPI(spec *loads.Document) *TsAPI {
	return &TsAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		TimeSeriesOperationsCreateOrMergeTimeseriesHandler: time_series_operations.CreateOrMergeTimeseriesHandlerFunc(func(params time_series_operations.CreateOrMergeTimeseriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.CreateOrMergeTimeseries has not yet been implemented")
		}),
		TimeSeriesOperationsCreateOrUpdateTimeseriesHandler: time_series_operations.CreateOrUpdateTimeseriesHandlerFunc(func(params time_series_operations.CreateOrUpdateTimeseriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.CreateOrUpdateTimeseries has not yet been implemented")
		}),
		TimeSeriesOperationsCreateOrUpdateTimeseriesDataHandler: time_series_operations.CreateOrUpdateTimeseriesDataHandlerFunc(func(params time_series_operations.CreateOrUpdateTimeseriesDataParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.CreateOrUpdateTimeseriesData has not yet been implemented")
		}),
		TimeSeriesOperationsDeleteTimeseriesHandler: time_series_operations.DeleteTimeseriesHandlerFunc(func(params time_series_operations.DeleteTimeseriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.DeleteTimeseries has not yet been implemented")
		}),
		TimeSeriesOperationsRetrieveTimeseriesHandler: time_series_operations.RetrieveTimeseriesHandlerFunc(func(params time_series_operations.RetrieveTimeseriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.RetrieveTimeseries has not yet been implemented")
		}),

		TimeseriesAuth: func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (timeseries) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*TsAPI Create, update, and query time series data with a precision of 1 millisecond. */
type TsAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// TimeseriesAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	TimeseriesAuth func(string, []string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// TimeSeriesOperationsCreateOrMergeTimeseriesHandler sets the operation handler for the create or merge timeseries operation
	TimeSeriesOperationsCreateOrMergeTimeseriesHandler time_series_operations.CreateOrMergeTimeseriesHandler
	// TimeSeriesOperationsCreateOrUpdateTimeseriesHandler sets the operation handler for the create or update timeseries operation
	TimeSeriesOperationsCreateOrUpdateTimeseriesHandler time_series_operations.CreateOrUpdateTimeseriesHandler
	// TimeSeriesOperationsCreateOrUpdateTimeseriesDataHandler sets the operation handler for the create or update timeseries data operation
	TimeSeriesOperationsCreateOrUpdateTimeseriesDataHandler time_series_operations.CreateOrUpdateTimeseriesDataHandler
	// TimeSeriesOperationsDeleteTimeseriesHandler sets the operation handler for the delete timeseries operation
	TimeSeriesOperationsDeleteTimeseriesHandler time_series_operations.DeleteTimeseriesHandler
	// TimeSeriesOperationsRetrieveTimeseriesHandler sets the operation handler for the retrieve timeseries operation
	TimeSeriesOperationsRetrieveTimeseriesHandler time_series_operations.RetrieveTimeseriesHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *TsAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *TsAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *TsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TsAPI
func (o *TsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.TimeseriesAuth == nil {
		unregistered = append(unregistered, "TimeseriesAuth")
	}

	if o.TimeSeriesOperationsCreateOrMergeTimeseriesHandler == nil {
		unregistered = append(unregistered, "time_series_operations.CreateOrMergeTimeseriesHandler")
	}
	if o.TimeSeriesOperationsCreateOrUpdateTimeseriesHandler == nil {
		unregistered = append(unregistered, "time_series_operations.CreateOrUpdateTimeseriesHandler")
	}
	if o.TimeSeriesOperationsCreateOrUpdateTimeseriesDataHandler == nil {
		unregistered = append(unregistered, "time_series_operations.CreateOrUpdateTimeseriesDataHandler")
	}
	if o.TimeSeriesOperationsDeleteTimeseriesHandler == nil {
		unregistered = append(unregistered, "time_series_operations.DeleteTimeseriesHandler")
	}
	if o.TimeSeriesOperationsRetrieveTimeseriesHandler == nil {
		unregistered = append(unregistered, "time_series_operations.RetrieveTimeseriesHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "timeseries":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.TimeseriesAuth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *TsAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *TsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *TsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the ts API
func (o *TsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/timeseries/{entityId}/{propertySetName}"] = time_series_operations.NewCreateOrMergeTimeseries(o.context, o.TimeSeriesOperationsCreateOrMergeTimeseriesHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/timeseries"] = time_series_operations.NewCreateOrUpdateTimeseries(o.context, o.TimeSeriesOperationsCreateOrUpdateTimeseriesHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/timeseries/{entityId}/{propertySetName}"] = time_series_operations.NewCreateOrUpdateTimeseriesData(o.context, o.TimeSeriesOperationsCreateOrUpdateTimeseriesDataHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/timeseries/{entityId}/{propertySetName}"] = time_series_operations.NewDeleteTimeseries(o.context, o.TimeSeriesOperationsDeleteTimeseriesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/timeseries/{entityId}/{propertySetName}"] = time_series_operations.NewRetrieveTimeseries(o.context, o.TimeSeriesOperationsRetrieveTimeseriesHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *TsAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
