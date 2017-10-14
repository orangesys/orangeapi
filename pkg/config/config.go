package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// KongConfiguration init kong configuration struct
type KongConfiguration struct {
	KongAdminHost string `envconfig:"kong_host" default:"localhost"`
	KongAdminPort string `envconfig:"kong_port" default:"8001"`
	KongAdminURL  string `envconfig:"kong_url" default:"http://127.0.0.1:8001/"`
}

// FirebaseConfiguration init firebase configuration struct
type FirebaseConfiguration struct {
	FirebaseURL  string `envconfig:"firebase_url" default:"https://saas-orangesys-io.firebaseio.com/"`
	FirebaseAuth string `envconfig:"firebase_auth" default:""`
}

const (
	// ConfigPrefix is envconfig prefix path
	ConfigPrefix = "orangeapi"
)

// LoadKongConfig load kong configuration with environment values
func LoadKongConfig() (*KongConfiguration, error) {
	var config KongConfiguration
	err := envconfig.Process(ConfigPrefix, &config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config from envs.")
	}
	if config.KongAdminURL == "" {
		config.KongAdminURL = "http://" + config.KongAdminHost + ":" + config.KongAdminPort + "/"
	}
	return &config, nil
}

// LoadFirebaseConfig load firebase configuration with environment values
func LoadFirebaseConfig() (*FirebaseConfiguration, error) {
	var config FirebaseConfiguration
	err := envconfig.Process(ConfigPrefix, &config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config from envs.")
	}
	return &config, nil
}
