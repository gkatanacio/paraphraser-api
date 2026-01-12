package chatgpt

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config contains the configuration for the ChatGPT API client.
type Config struct {
	ApiKey string `envconfig:"CHATGPT_API_KEY"`
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
		log.Fatal("CHATGPT_API_KEY must be set")
	}
}
