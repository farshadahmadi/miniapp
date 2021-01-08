package app

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-openapi/loads"

	logger "github.com/sirupsen/logrus"

	"github.com/farshadahmadi/miniapp/internal/configuration"
	"github.com/farshadahmadi/miniapp/internal/constants/envs"
	"github.com/farshadahmadi/miniapp/internal/feature-flag/decisions"
	"github.com/farshadahmadi/miniapp/internal/feature-flag/features"
	"github.com/farshadahmadi/miniapp/internal/generated/swagger/server/miniapp"
	miniappoperation "github.com/farshadahmadi/miniapp/internal/generated/swagger/server/miniapp/operation"
	"github.com/farshadahmadi/miniapp/internal/generated/swagger/server/miniapp/operation/feature"
	featurehandler "github.com/farshadahmadi/miniapp/internal/handler/feature"
	"github.com/farshadahmadi/miniapp/internal/projectpath"
)

// Start point of the application.
func Start(env envs.Env) {
	// Prepare configuration struct
	conf := configuration.Configuration{}
	// Load configuration file config.yml from env-specific folder
	if err := configuration.Load(&conf, projectpath.Root+"/config/"+env.String()+"/"); err != nil {
		logger.WithError(err).Error("Configuration loading failed")

		return
	}

	// Setup & launch the servers (miniapp backend)
	setupServers(&conf, env)
}

// Launches the server
func setupServers(config *configuration.Configuration, env envs.Env) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Initialize Miniapp server
	go func() {
		miniappServer, err := setupMiniapp(config, env)
		if err != nil {
			logger.WithError(err).Error("Setting up Miniapp server failed")

			return
		}

		defer func() {
			if miniappServer != nil {
				err = miniappServer.Shutdown()
				if err != nil {
					logger.WithError(err).Error("Shutting down Miniapp failed")
				}
			}

			wg.Done()
		}()

		err = miniappServer.Serve()
		if err != nil {
			logger.WithError(err).Error("Serving Miniapp failed")

			return
		}
	}()

	wg.Wait()
}

// sets up Miniapp server
func setupMiniapp(conf *configuration.Configuration, env envs.Env) (server *miniapp.Server, err error) {
	var (
		splitService *features.Service
		swaggerSpec  *loads.Document
	)

	// Load Swagger spec
	swaggerSpec, err = loads.Embedded(miniapp.SwaggerJSON, miniapp.FlatSwaggerJSON)
	if err != nil {
		logger.WithError(err).Error("Swagger Spec Loading Failed")

		return
	}

	api := miniappoperation.NewMiniappAPI(swaggerSpec)

	// Initialize Feature Flag service
	splitService, err = features.NewService(&conf.FeatureFlag)
	if err != nil {
		logger.WithError(err).Error("Split service initialization failed")

		return
	}

	decisionService := decisions.NewDecisions(splitService)

	// Assign endpoint handlers
	api.FeatureFeatureFlagHandler = feature.FeatureFlagHandlerFunc(featurehandler.NewGetFeatureFlag(decisionService).GetFeatureFlag)

	// Initialize Miniapp server
	server = miniapp.NewServer(api)
	server.Port = conf.App.MiniappPort
	server.ConfigureAPI()

	// disable server timeout error in debug mode
	if env != envs.DEBUG {
		// Add a middleware that sends back timeout error if API handlers take more than the configured time
		server.SetHandler(
			http.TimeoutHandler(
				server.GetHandler(),
				time.Millisecond*500,
				`{"code": 503, "mesaage": "Timeout"}`,
			),
		)
	}

	// This needs to be set after server.ConfigureAPI()
	api.ServerShutdown = func() {
		splitService.ShutDown()
	}

	return
}
