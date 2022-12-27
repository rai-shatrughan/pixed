// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/mc/models"
	"gsvc/api/mc/restapi/operations"
	"gsvc/api/mc/restapi/operations/diagnostic_activations"
	"gsvc/api/mc/restapi/operations/diagnostic_information"
	"gsvc/api/mc/restapi/operations/exchange"
	"gsvc/api/mc/restapi/operations/mappings"
	"gsvc/api/mc/restapi/operations/record_recovery"
)

//go:generate swagger generate server --target ../../mc --name Mc --spec ../../../yaml/mindconnect-v3-5-0.yaml --principal models.Principal

func configureFlags(api *operations.McAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.McAPI) http.Handler {
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
	api.MultipartMixedConsumer = runtime.ConsumerFunc(func(r io.Reader, target interface{}) error {
		return errors.NotImplemented("multipartMixed consumer has not yet been implemented")
	})

	api.JSONProducer = runtime.JSONProducer()

	if api.DiagnosticActivationsAuth == nil {
		api.DiagnosticActivationsAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (diagnosticActivations) has not yet been implemented")
		}
	}
	if api.DiagnosticInformationAuth == nil {
		api.DiagnosticInformationAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (diagnosticInformation) has not yet been implemented")
		}
	}
	if api.ExchangeAuth == nil {
		api.ExchangeAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (exchange) has not yet been implemented")
		}
	}
	if api.MappingsAuth == nil {
		api.MappingsAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (mappings) has not yet been implemented")
		}
	}
	if api.RecordRecoveryAuth == nil {
		api.RecordRecoveryAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (recordRecovery) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.MappingsDeleteDataPointMappingsIDHandler == nil {
		api.MappingsDeleteDataPointMappingsIDHandler = mappings.DeleteDataPointMappingsIDHandlerFunc(func(params mappings.DeleteDataPointMappingsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation mappings.DeleteDataPointMappingsID has not yet been implemented")
		})
	}
	if api.DiagnosticActivationsDeleteDiagnosticActivationsIDHandler == nil {
		api.DiagnosticActivationsDeleteDiagnosticActivationsIDHandler = diagnostic_activations.DeleteDiagnosticActivationsIDHandlerFunc(func(params diagnostic_activations.DeleteDiagnosticActivationsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation diagnostic_activations.DeleteDiagnosticActivationsID has not yet been implemented")
		})
	}
	if api.RecordRecoveryDeleteRecoverableRecordsIDHandler == nil {
		api.RecordRecoveryDeleteRecoverableRecordsIDHandler = record_recovery.DeleteRecoverableRecordsIDHandlerFunc(func(params record_recovery.DeleteRecoverableRecordsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation record_recovery.DeleteRecoverableRecordsID has not yet been implemented")
		})
	}
	if api.MappingsGetDataPointMappingsHandler == nil {
		api.MappingsGetDataPointMappingsHandler = mappings.GetDataPointMappingsHandlerFunc(func(params mappings.GetDataPointMappingsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation mappings.GetDataPointMappings has not yet been implemented")
		})
	}
	if api.MappingsGetDataPointMappingsIDHandler == nil {
		api.MappingsGetDataPointMappingsIDHandler = mappings.GetDataPointMappingsIDHandlerFunc(func(params mappings.GetDataPointMappingsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation mappings.GetDataPointMappingsID has not yet been implemented")
		})
	}
	if api.DiagnosticActivationsGetDiagnosticActivationsHandler == nil {
		api.DiagnosticActivationsGetDiagnosticActivationsHandler = diagnostic_activations.GetDiagnosticActivationsHandlerFunc(func(params diagnostic_activations.GetDiagnosticActivationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation diagnostic_activations.GetDiagnosticActivations has not yet been implemented")
		})
	}
	if api.DiagnosticActivationsGetDiagnosticActivationsIDHandler == nil {
		api.DiagnosticActivationsGetDiagnosticActivationsIDHandler = diagnostic_activations.GetDiagnosticActivationsIDHandlerFunc(func(params diagnostic_activations.GetDiagnosticActivationsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation diagnostic_activations.GetDiagnosticActivationsID has not yet been implemented")
		})
	}
	if api.DiagnosticActivationsGetDiagnosticActivationsIDMessagesHandler == nil {
		api.DiagnosticActivationsGetDiagnosticActivationsIDMessagesHandler = diagnostic_activations.GetDiagnosticActivationsIDMessagesHandlerFunc(func(params diagnostic_activations.GetDiagnosticActivationsIDMessagesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation diagnostic_activations.GetDiagnosticActivationsIDMessages has not yet been implemented")
		})
	}
	if api.DiagnosticInformationGetDiagnosticInformationHandler == nil {
		api.DiagnosticInformationGetDiagnosticInformationHandler = diagnostic_information.GetDiagnosticInformationHandlerFunc(func(params diagnostic_information.GetDiagnosticInformationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation diagnostic_information.GetDiagnosticInformation has not yet been implemented")
		})
	}
	if api.RecordRecoveryGetRecoverableRecordsHandler == nil {
		api.RecordRecoveryGetRecoverableRecordsHandler = record_recovery.GetRecoverableRecordsHandlerFunc(func(params record_recovery.GetRecoverableRecordsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation record_recovery.GetRecoverableRecords has not yet been implemented")
		})
	}
	if api.RecordRecoveryGetRecoverableRecordsIDDownloadLinkHandler == nil {
		api.RecordRecoveryGetRecoverableRecordsIDDownloadLinkHandler = record_recovery.GetRecoverableRecordsIDDownloadLinkHandlerFunc(func(params record_recovery.GetRecoverableRecordsIDDownloadLinkParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation record_recovery.GetRecoverableRecordsIDDownloadLink has not yet been implemented")
		})
	}
	if api.MappingsPostDataPointMappingsHandler == nil {
		api.MappingsPostDataPointMappingsHandler = mappings.PostDataPointMappingsHandlerFunc(func(params mappings.PostDataPointMappingsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation mappings.PostDataPointMappings has not yet been implemented")
		})
	}
	if api.DiagnosticActivationsPostDiagnosticActivationsHandler == nil {
		api.DiagnosticActivationsPostDiagnosticActivationsHandler = diagnostic_activations.PostDiagnosticActivationsHandlerFunc(func(params diagnostic_activations.PostDiagnosticActivationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation diagnostic_activations.PostDiagnosticActivations has not yet been implemented")
		})
	}
	if api.ExchangePostExchangeHandler == nil {
		api.ExchangePostExchangeHandler = exchange.PostExchangeHandlerFunc(func(params exchange.PostExchangeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation exchange.PostExchange has not yet been implemented")
		})
	}
	if api.RecordRecoveryPostRecoverableRecordsIDReplayHandler == nil {
		api.RecordRecoveryPostRecoverableRecordsIDReplayHandler = record_recovery.PostRecoverableRecordsIDReplayHandlerFunc(func(params record_recovery.PostRecoverableRecordsIDReplayParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation record_recovery.PostRecoverableRecordsIDReplay has not yet been implemented")
		})
	}
	if api.DiagnosticActivationsPutDiagnosticActivationsIDHandler == nil {
		api.DiagnosticActivationsPutDiagnosticActivationsIDHandler = diagnostic_activations.PutDiagnosticActivationsIDHandlerFunc(func(params diagnostic_activations.PutDiagnosticActivationsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation diagnostic_activations.PutDiagnosticActivationsID has not yet been implemented")
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
