// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"gsvc/api/am/models"
	"gsvc/api/am/restapi/operations"
	"gsvc/api/am/restapi/operations/aspecttype"
	"gsvc/api/am/restapi/operations/assetmodellock"
	"gsvc/api/am/restapi/operations/assets"
	"gsvc/api/am/restapi/operations/assettype"
	"gsvc/api/am/restapi/operations/billboard"
	"gsvc/api/am/restapi/operations/files"
	"gsvc/api/am/restapi/operations/locations"
	"gsvc/api/am/restapi/operations/structure"
)

//go:generate swagger generate server --target ../../am --name Am --spec ../../../yaml/assetmanagement-v3-18-3.yaml --principal models.Principal

func configureFlags(api *operations.AmAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AmAPI) http.Handler {
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
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.ApplicationBase64Producer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("applicationBase64 producer has not yet been implemented")
	})
	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()

	if api.AspecttypeAuth == nil {
		api.AspecttypeAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (aspecttype) has not yet been implemented")
		}
	}
	if api.AssetAuth == nil {
		api.AssetAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (asset) has not yet been implemented")
		}
	}
	if api.AssetmodelAuth == nil {
		api.AssetmodelAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (assetmodel) has not yet been implemented")
		}
	}
	if api.AssettypeAuth == nil {
		api.AssettypeAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (assettype) has not yet been implemented")
		}
	}
	if api.FileAuth == nil {
		api.FileAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (file) has not yet been implemented")
		}
	}
	if api.LocationAuth == nil {
		api.LocationAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (location) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// files.ReplaceFileMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// files.UploadFileMaxParseMemory = 32 << 20

	if api.AssetsAddAssetHandler == nil {
		api.AssetsAddAssetHandler = assets.AddAssetHandlerFunc(func(params assets.AddAssetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.AddAsset has not yet been implemented")
		})
	}
	if api.AspecttypeDeleteAspectTypeHandler == nil {
		api.AspecttypeDeleteAspectTypeHandler = aspecttype.DeleteAspectTypeHandlerFunc(func(params aspecttype.DeleteAspectTypeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation aspecttype.DeleteAspectType has not yet been implemented")
		})
	}
	if api.AssetsDeleteAssetHandler == nil {
		api.AssetsDeleteAssetHandler = assets.DeleteAssetHandlerFunc(func(params assets.DeleteAssetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.DeleteAsset has not yet been implemented")
		})
	}
	if api.AssetsDeleteAssetFileAssigmentHandler == nil {
		api.AssetsDeleteAssetFileAssigmentHandler = assets.DeleteAssetFileAssigmentHandlerFunc(func(params assets.DeleteAssetFileAssigmentParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.DeleteAssetFileAssigment has not yet been implemented")
		})
	}
	if api.LocationsDeleteAssetLocationHandler == nil {
		api.LocationsDeleteAssetLocationHandler = locations.DeleteAssetLocationHandlerFunc(func(params locations.DeleteAssetLocationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation locations.DeleteAssetLocation has not yet been implemented")
		})
	}
	if api.AssettypeDeleteAssetTypeHandler == nil {
		api.AssettypeDeleteAssetTypeHandler = assettype.DeleteAssetTypeHandlerFunc(func(params assettype.DeleteAssetTypeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assettype.DeleteAssetType has not yet been implemented")
		})
	}
	if api.AssettypeDeleteAssetTypeFileAssignmentHandler == nil {
		api.AssettypeDeleteAssetTypeFileAssignmentHandler = assettype.DeleteAssetTypeFileAssignmentHandlerFunc(func(params assettype.DeleteAssetTypeFileAssignmentParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assettype.DeleteAssetTypeFileAssignment has not yet been implemented")
		})
	}
	if api.FilesDeleteFileHandler == nil {
		api.FilesDeleteFileHandler = files.DeleteFileHandlerFunc(func(params files.DeleteFileParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation files.DeleteFile has not yet been implemented")
		})
	}
	if api.FilesDownloadFileHandler == nil {
		api.FilesDownloadFileHandler = files.DownloadFileHandlerFunc(func(params files.DownloadFileParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation files.DownloadFile has not yet been implemented")
		})
	}
	if api.AspecttypeGetAspectTypeHandler == nil {
		api.AspecttypeGetAspectTypeHandler = aspecttype.GetAspectTypeHandlerFunc(func(params aspecttype.GetAspectTypeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation aspecttype.GetAspectType has not yet been implemented")
		})
	}
	if api.AssetsGetAssetHandler == nil {
		api.AssetsGetAssetHandler = assets.GetAssetHandlerFunc(func(params assets.GetAssetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.GetAsset has not yet been implemented")
		})
	}
	if api.AssetmodellockGetAssetModelLockHandler == nil {
		api.AssetmodellockGetAssetModelLockHandler = assetmodellock.GetAssetModelLockHandlerFunc(func(params assetmodellock.GetAssetModelLockParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assetmodellock.GetAssetModelLock has not yet been implemented")
		})
	}
	if api.AssettypeGetAssetTypeHandler == nil {
		api.AssettypeGetAssetTypeHandler = assettype.GetAssetTypeHandlerFunc(func(params assettype.GetAssetTypeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assettype.GetAssetType has not yet been implemented")
		})
	}
	if api.BillboardGetBillboardHandler == nil {
		api.BillboardGetBillboardHandler = billboard.GetBillboardHandlerFunc(func(params billboard.GetBillboardParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation billboard.GetBillboard has not yet been implemented")
		})
	}
	if api.FilesGetFileHandler == nil {
		api.FilesGetFileHandler = files.GetFileHandlerFunc(func(params files.GetFileParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation files.GetFile has not yet been implemented")
		})
	}
	if api.AssetsGetRootAssetHandler == nil {
		api.AssetsGetRootAssetHandler = assets.GetRootAssetHandlerFunc(func(params assets.GetRootAssetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.GetRootAsset has not yet been implemented")
		})
	}
	if api.AspecttypeListAspectTypesHandler == nil {
		api.AspecttypeListAspectTypesHandler = aspecttype.ListAspectTypesHandlerFunc(func(params aspecttype.ListAspectTypesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation aspecttype.ListAspectTypes has not yet been implemented")
		})
	}
	if api.StructureListAssetAspectsHandler == nil {
		api.StructureListAssetAspectsHandler = structure.ListAssetAspectsHandlerFunc(func(params structure.ListAssetAspectsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation structure.ListAssetAspects has not yet been implemented")
		})
	}
	if api.AssettypeListAssetTypesHandler == nil {
		api.AssettypeListAssetTypesHandler = assettype.ListAssetTypesHandlerFunc(func(params assettype.ListAssetTypesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assettype.ListAssetTypes has not yet been implemented")
		})
	}
	if api.StructureListAssetVariablesHandler == nil {
		api.StructureListAssetVariablesHandler = structure.ListAssetVariablesHandlerFunc(func(params structure.ListAssetVariablesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation structure.ListAssetVariables has not yet been implemented")
		})
	}
	if api.AssetsListAssetsHandler == nil {
		api.AssetsListAssetsHandler = assets.ListAssetsHandlerFunc(func(params assets.ListAssetsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.ListAssets has not yet been implemented")
		})
	}
	if api.FilesListFilesHandler == nil {
		api.FilesListFilesHandler = files.ListFilesHandlerFunc(func(params files.ListFilesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation files.ListFiles has not yet been implemented")
		})
	}
	if api.AssetsMoveAssetHandler == nil {
		api.AssetsMoveAssetHandler = assets.MoveAssetHandlerFunc(func(params assets.MoveAssetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.MoveAsset has not yet been implemented")
		})
	}
	if api.AssetmodellockPutAssetModelLockHandler == nil {
		api.AssetmodellockPutAssetModelLockHandler = assetmodellock.PutAssetModelLockHandlerFunc(func(params assetmodellock.PutAssetModelLockParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assetmodellock.PutAssetModelLock has not yet been implemented")
		})
	}
	if api.AssetsReplaceAssetHandler == nil {
		api.AssetsReplaceAssetHandler = assets.ReplaceAssetHandlerFunc(func(params assets.ReplaceAssetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.ReplaceAsset has not yet been implemented")
		})
	}
	if api.FilesReplaceFileHandler == nil {
		api.FilesReplaceFileHandler = files.ReplaceFileHandlerFunc(func(params files.ReplaceFileParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation files.ReplaceFile has not yet been implemented")
		})
	}
	if api.AspecttypeSaveAspectTypeHandler == nil {
		api.AspecttypeSaveAspectTypeHandler = aspecttype.SaveAspectTypeHandlerFunc(func(params aspecttype.SaveAspectTypeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation aspecttype.SaveAspectType has not yet been implemented")
		})
	}
	if api.AssetsSaveAssetFileAssignmentHandler == nil {
		api.AssetsSaveAssetFileAssignmentHandler = assets.SaveAssetFileAssignmentHandlerFunc(func(params assets.SaveAssetFileAssignmentParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.SaveAssetFileAssignment has not yet been implemented")
		})
	}
	if api.LocationsSaveAssetLocationHandler == nil {
		api.LocationsSaveAssetLocationHandler = locations.SaveAssetLocationHandlerFunc(func(params locations.SaveAssetLocationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation locations.SaveAssetLocation has not yet been implemented")
		})
	}
	if api.AssettypeSaveAssetTypeHandler == nil {
		api.AssettypeSaveAssetTypeHandler = assettype.SaveAssetTypeHandlerFunc(func(params assettype.SaveAssetTypeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assettype.SaveAssetType has not yet been implemented")
		})
	}
	if api.AssettypeSaveAssetTypeFileAssignmentHandler == nil {
		api.AssettypeSaveAssetTypeFileAssignmentHandler = assettype.SaveAssetTypeFileAssignmentHandlerFunc(func(params assettype.SaveAssetTypeFileAssignmentParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assettype.SaveAssetTypeFileAssignment has not yet been implemented")
		})
	}
	if api.AspecttypeUpdateAspectTypeHandler == nil {
		api.AspecttypeUpdateAspectTypeHandler = aspecttype.UpdateAspectTypeHandlerFunc(func(params aspecttype.UpdateAspectTypeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation aspecttype.UpdateAspectType has not yet been implemented")
		})
	}
	if api.AspecttypeUpdateAspectTypeVariableHandler == nil {
		api.AspecttypeUpdateAspectTypeVariableHandler = aspecttype.UpdateAspectTypeVariableHandlerFunc(func(params aspecttype.UpdateAspectTypeVariableParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation aspecttype.UpdateAspectTypeVariable has not yet been implemented")
		})
	}
	if api.AssetsUpdateAssetHandler == nil {
		api.AssetsUpdateAssetHandler = assets.UpdateAssetHandlerFunc(func(params assets.UpdateAssetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assets.UpdateAsset has not yet been implemented")
		})
	}
	if api.AssettypeUpdateAssetTypeHandler == nil {
		api.AssettypeUpdateAssetTypeHandler = assettype.UpdateAssetTypeHandlerFunc(func(params assettype.UpdateAssetTypeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assettype.UpdateAssetType has not yet been implemented")
		})
	}
	if api.AssettypeUpdateAssetTypeVariableHandler == nil {
		api.AssettypeUpdateAssetTypeVariableHandler = assettype.UpdateAssetTypeVariableHandlerFunc(func(params assettype.UpdateAssetTypeVariableParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation assettype.UpdateAssetTypeVariable has not yet been implemented")
		})
	}
	if api.FilesUploadFileHandler == nil {
		api.FilesUploadFileHandler = files.UploadFileHandlerFunc(func(params files.UploadFileParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation files.UploadFile has not yet been implemented")
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
