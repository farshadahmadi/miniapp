package configuration

import (
	"errors"

	"github.com/spf13/viper"
)

// ErrMissingFeatureFlagAPIKey is used when FeatureFlag api key is missing
var ErrMissingFeatureFlagAPIKey = errors.New("Missing FeatureFlag.APIKey value in config")

// Load tries to load config.yml file from specified paths and sets default values for configuration
func Load(config *Configuration, paths ...string) (err error) {
	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Create a struct from config using Unmarshal
	if err = viper.Unmarshal(&config); err != nil {
		return err
	}

	// Validate the config
	if config.FeatureFlag.APIKey == "" {
		err = ErrMissingFeatureFlagAPIKey

		return
	}

	return
}

// Configuration provides main level configuration container
type Configuration struct {
	FeatureFlag FeatureFlagConfiguration
	App         Application
}

// Application provides application specific configs
type Application struct {
	Host        string
	MiniappPort int
	Scheme      string
}

// FeatureFlagConfiguration provides split.io specific configs
type FeatureFlagConfiguration struct {
	OperationMode string
	APIKey        string
	RedisPrefix   string
}
