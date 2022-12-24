// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/ts/models"
	"gsvc/api/ts/restapi/operations"
	"gsvc/api/ts/restapi/operations/time_series_operations"
)

//go:generate swagger generate server --target ../../ts --name Ts --spec ../../../yaml/iottimeseries-v3-7-0.yaml --principal models.Principal

func configureFlags(api *operations.TsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.TimeseriesAuth == nil {
		api.TimeseriesAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (timeseries) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.TimeSeriesOperationsCreateOrMergeTimeseriesHandler == nil {
		api.TimeSeriesOperationsCreateOrMergeTimeseriesHandler = time_series_operations.CreateOrMergeTimeseriesHandlerFunc(func(params time_series_operations.CreateOrMergeTimeseriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.CreateOrMergeTimeseries has not yet been implemented")
		})
	}
	if api.TimeSeriesOperationsCreateOrUpdateTimeseriesHandler == nil {
		api.TimeSeriesOperationsCreateOrUpdateTimeseriesHandler = time_series_operations.CreateOrUpdateTimeseriesHandlerFunc(func(params time_series_operations.CreateOrUpdateTimeseriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.CreateOrUpdateTimeseries has not yet been implemented")
		})
	}
	if api.TimeSeriesOperationsCreateOrUpdateTimeseriesDataHandler == nil {
		api.TimeSeriesOperationsCreateOrUpdateTimeseriesDataHandler = time_series_operations.CreateOrUpdateTimeseriesDataHandlerFunc(func(params time_series_operations.CreateOrUpdateTimeseriesDataParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.CreateOrUpdateTimeseriesData has not yet been implemented")
		})
	}
	if api.TimeSeriesOperationsDeleteTimeseriesHandler == nil {
		api.TimeSeriesOperationsDeleteTimeseriesHandler = time_series_operations.DeleteTimeseriesHandlerFunc(func(params time_series_operations.DeleteTimeseriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.DeleteTimeseries has not yet been implemented")
		})
	}
	if api.TimeSeriesOperationsRetrieveTimeseriesHandler == nil {
		api.TimeSeriesOperationsRetrieveTimeseriesHandler = time_series_operations.RetrieveTimeseriesHandlerFunc(func(params time_series_operations.RetrieveTimeseriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation time_series_operations.RetrieveTimeseries has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
