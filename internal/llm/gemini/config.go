package gemini

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config contains the configuration for the Gemini API client.
type Config struct {
	ApiKey      string  `envconfig:"GEMINI_API_KEY"`
	Temperature float32 `envconfig:"GEMINI_TEMPERATURE" default:"0.9"`
}

// ConfigFromEnv loads the configuration from the environment variables.
func ConfigFromEnv() *Config {
	cfg := &Config{}
	envconfig.MustProcess("", cfg)
	cfg.validate()

	return cfg
}

func (c *Config) validate() {
	if len(c.ApiKey) == 0 {
		log.Fatal("GEMINI_API_KEY must be set")
	}
}
