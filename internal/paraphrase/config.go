package paraphrase

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config contains the configuration for the Paraphrase Service.
type Config struct {
	Timeout time.Duration `envconfig:"PARAPHRASE_TIMEOUT" default:"5000ms"`
}

// ConfigFromEnv loads the configuration from the environment variables.
func ConfigFromEnv() *Config {
	cfg := &Config{}
	envconfig.MustProcess("", cfg)

	return cfg
}
