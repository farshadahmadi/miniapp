package features

import (
	log "github.com/sirupsen/logrus"
	"github.com/splitio/go-client/v6/splitio/client"
	"github.com/splitio/go-client/v6/splitio/conf"

	"github.com/farshadahmadi/miniapp/internal/configuration"
)

const timer = 10

// Service provides a struct to access the split.io client
type Service struct {
	client *client.SplitClient
}

// GetFlagParams contains all the params required for getting a flag
type GetFlagParams struct {
	Feature    string
	TrafficID  interface{}
	Attributes map[string]interface{}
}

// NewService returns a new Service struct and initializes the connection
func NewService(ffConf *configuration.FeatureFlagConfiguration) (*Service, error) {
	cfg := conf.Default()
	cfg.OperationMode = ffConf.OperationMode

	// Change split SDK logging level:
	// cfg.LoggerConfig.LogLevel = logging.LevelDebug

	factory, err := client.NewSplitFactory(ffConf.APIKey, cfg)
	if err != nil {
		log.WithError(err).Error("error initializing split.io")

		return nil, err
	}

	splitClient := factory.Client()
	err = splitClient.BlockUntilReady(timer)

	if err != nil {
		log.WithError(err).Error("timed out initializing split.io")

		return nil, err
	}

	return &Service{
		client: splitClient,
	}, nil
}

// Get gets a flag with given params
func (s *Service) Get(params GetFlagParams) (treatment string, err error) {
	treatment = s.client.Treatment(params.TrafficID, params.Feature, params.Attributes)

	return
}

// ShutDown shuts down the service
func (s *Service) ShutDown() {
	log.Info("Shutting down split.io service")
	s.client.Destroy()
}
