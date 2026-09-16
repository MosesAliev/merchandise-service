// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"merchandise-service/domain"
	"merchandise-service/models"
	"merchandise-service/postgres"
	"merchandise-service/repository"
	"merchandise-service/restapi/operations"
	"merchandise-service/service"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//go:generate swagger generate server --target ..\..\merchandise-service --name Merch --spec ..\schema.json --principal any

func configureFlags(api *operations.MerchAPI) {
	// api.CommandLineOptionsGroups = []cmdutils.CommandLineOptionsGroup{ ... }
	_ = api
}

func configureAPI(api *operations.MerchAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...any)
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuthAuth == nil {
		api.BearerAuthAuth = func(token string) (any, error) {
			_ = token

			return nil, errors.NotImplemented("api key auth (BearerAuth) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.GetAPIBuyItemHandler == nil {
		api.GetAPIBuyItemHandler = operations.GetAPIBuyItemHandlerFunc(func(params operations.GetAPIBuyItemParams, principal any) middleware.Responder {
			_ = params
			_ = principal

			return middleware.NotImplemented("operation operations.GetAPIBuyItem has not yet been implemented")
		})
	}
	if api.GetAPIInfoHandler == nil {
		api.GetAPIInfoHandler = operations.GetAPIInfoHandlerFunc(func(params operations.GetAPIInfoParams, principal any) middleware.Responder {
			_ = params
			_ = principal

			return middleware.NotImplemented("operation operations.GetAPIInfo has not yet been implemented")
		})
	}
	api.PostAPIAuthHandler = operations.PostAPIAuthHandlerFunc(func(params operations.PostAPIAuthParams) middleware.Responder {
		user := domain.NewUser(*params.Body.Username, params.Body.Password.String())
		repo := repository.NewPostgresUserRepository(postgres.DB)
		authService := service.NewAuthService(repo)
		token, err := authService.Auth(user)
		if err != nil {
			response := operations.NewPostAPIAuthBadRequest()
			log.Println(err.Error())
			response.Payload = &models.ErrorResponse{
				Errors: err.Error(),
			}
			return response
		}

		response := operations.NewPostAPIAuthOK()
		response.Payload = &models.AuthResponse{
			Token: token,
		}
		return response
	})

	if api.PostAPISendCoinHandler == nil {
		api.PostAPISendCoinHandler = operations.PostAPISendCoinHandlerFunc(func(params operations.PostAPISendCoinParams, principal any) middleware.Responder {
			_ = params
			_ = principal

			return middleware.NotImplemented("operation operations.PostAPISendCoin has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
	_ = tlsConfig
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(server *http.Server, scheme, addr string) {
	_ = server
	_ = scheme
	_ = addr
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
