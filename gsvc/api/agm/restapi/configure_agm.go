// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/agm/models"
	"gsvc/api/agm/restapi/operations"
	"gsvc/api/agm/restapi/operations/agent_operations"
	"gsvc/api/agm/restapi/operations/boarding_operations"
	"gsvc/api/agm/restapi/operations/data_source_configuration_operations"
	"gsvc/api/agm/restapi/operations/registration_operations"
	"gsvc/api/agm/restapi/operations/token_operations"
)

//go:generate swagger generate server --target ../../agm --name Agm --spec ../../../yaml/agentmanagement-v3-4-0.yaml --principal models.Principal

func configureFlags(api *operations.AgmAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AgmAPI) http.Handler {
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
	api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	if api.AgentsAuth == nil {
		api.AgentsAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (agents) has not yet been implemented")
		}
	}
	if api.BoardingAuth == nil {
		api.BoardingAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (boarding) has not yet been implemented")
		}
	}
	if api.DataSourcesAuth == nil {
		api.DataSourcesAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (dataSources) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// token_operations.PostOauthTokenMaxParseMemory = 32 << 20

	if api.AgentOperationsDeleteAgentsIDHandler == nil {
		api.AgentOperationsDeleteAgentsIDHandler = agent_operations.DeleteAgentsIDHandlerFunc(func(params agent_operations.DeleteAgentsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation agent_operations.DeleteAgentsID has not yet been implemented")
		})
	}
	if api.AgentOperationsGetAgentsHandler == nil {
		api.AgentOperationsGetAgentsHandler = agent_operations.GetAgentsHandlerFunc(func(params agent_operations.GetAgentsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation agent_operations.GetAgents has not yet been implemented")
		})
	}
	if api.AgentOperationsGetAgentsIDHandler == nil {
		api.AgentOperationsGetAgentsIDHandler = agent_operations.GetAgentsIDHandlerFunc(func(params agent_operations.GetAgentsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation agent_operations.GetAgentsID has not yet been implemented")
		})
	}
	if api.BoardingOperationsGetAgentsIDBoardingConfigurationHandler == nil {
		api.BoardingOperationsGetAgentsIDBoardingConfigurationHandler = boarding_operations.GetAgentsIDBoardingConfigurationHandlerFunc(func(params boarding_operations.GetAgentsIDBoardingConfigurationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation boarding_operations.GetAgentsIDBoardingConfiguration has not yet been implemented")
		})
	}
	if api.BoardingOperationsGetAgentsIDBoardingStatusHandler == nil {
		api.BoardingOperationsGetAgentsIDBoardingStatusHandler = boarding_operations.GetAgentsIDBoardingStatusHandlerFunc(func(params boarding_operations.GetAgentsIDBoardingStatusParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation boarding_operations.GetAgentsIDBoardingStatus has not yet been implemented")
		})
	}
	if api.DataSourceConfigurationOperationsGetAgentsIDDataSourceConfigurationHandler == nil {
		api.DataSourceConfigurationOperationsGetAgentsIDDataSourceConfigurationHandler = data_source_configuration_operations.GetAgentsIDDataSourceConfigurationHandlerFunc(func(params data_source_configuration_operations.GetAgentsIDDataSourceConfigurationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation data_source_configuration_operations.GetAgentsIDDataSourceConfiguration has not yet been implemented")
		})
	}
	if api.AgentOperationsGetAgentsIDStatusHandler == nil {
		api.AgentOperationsGetAgentsIDStatusHandler = agent_operations.GetAgentsIDStatusHandlerFunc(func(params agent_operations.GetAgentsIDStatusParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation agent_operations.GetAgentsIDStatus has not yet been implemented")
		})
	}
	if api.TokenOperationsGetOauthTokenKeyHandler == nil {
		api.TokenOperationsGetOauthTokenKeyHandler = token_operations.GetOauthTokenKeyHandlerFunc(func(params token_operations.GetOauthTokenKeyParams) middleware.Responder {
			return middleware.NotImplemented("operation token_operations.GetOauthTokenKey has not yet been implemented")
		})
	}
	if api.TokenOperationsGetOauthTokenKeysHandler == nil {
		api.TokenOperationsGetOauthTokenKeysHandler = token_operations.GetOauthTokenKeysHandlerFunc(func(params token_operations.GetOauthTokenKeysParams) middleware.Responder {
			return middleware.NotImplemented("operation token_operations.GetOauthTokenKeys has not yet been implemented")
		})
	}
	if api.AgentOperationsPostAgentsHandler == nil {
		api.AgentOperationsPostAgentsHandler = agent_operations.PostAgentsHandlerFunc(func(params agent_operations.PostAgentsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation agent_operations.PostAgents has not yet been implemented")
		})
	}
	if api.BoardingOperationsPostAgentsIDBoardingOffboardHandler == nil {
		api.BoardingOperationsPostAgentsIDBoardingOffboardHandler = boarding_operations.PostAgentsIDBoardingOffboardHandlerFunc(func(params boarding_operations.PostAgentsIDBoardingOffboardParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation boarding_operations.PostAgentsIDBoardingOffboard has not yet been implemented")
		})
	}
	if api.TokenOperationsPostOauthTokenHandler == nil {
		api.TokenOperationsPostOauthTokenHandler = token_operations.PostOauthTokenHandlerFunc(func(params token_operations.PostOauthTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation token_operations.PostOauthToken has not yet been implemented")
		})
	}
	if api.RegistrationOperationsPostRegisterHandler == nil {
		api.RegistrationOperationsPostRegisterHandler = registration_operations.PostRegisterHandlerFunc(func(params registration_operations.PostRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation registration_operations.PostRegister has not yet been implemented")
		})
	}
	if api.AgentOperationsPutAgentsIDHandler == nil {
		api.AgentOperationsPutAgentsIDHandler = agent_operations.PutAgentsIDHandlerFunc(func(params agent_operations.PutAgentsIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation agent_operations.PutAgentsID has not yet been implemented")
		})
	}
	if api.DataSourceConfigurationOperationsPutAgentsIDDataSourceConfigurationHandler == nil {
		api.DataSourceConfigurationOperationsPutAgentsIDDataSourceConfigurationHandler = data_source_configuration_operations.PutAgentsIDDataSourceConfigurationHandlerFunc(func(params data_source_configuration_operations.PutAgentsIDDataSourceConfigurationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation data_source_configuration_operations.PutAgentsIDDataSourceConfiguration has not yet been implemented")
		})
	}
	if api.RegistrationOperationsPutRegisterClientIDHandler == nil {
		api.RegistrationOperationsPutRegisterClientIDHandler = registration_operations.PutRegisterClientIDHandlerFunc(func(params registration_operations.PutRegisterClientIDParams) middleware.Responder {
			return middleware.NotImplemented("operation registration_operations.PutRegisterClientID has not yet been implemented")
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
