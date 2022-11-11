package config

import (
	"demi-anchor/pkg/errtrace"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	LogLevel  string `envconfig:"LOG_LEVEL" required:"true"`
	IsJsonLog bool   `envconfig:"IS_JSON_LOG" required:"true"`
	Port      string `envconfig:"PORT" required:"true"`

	// Repo *repo.Config
}

func Load() (*config, error) {
	cfg := &config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, errtrace.AddTrace(err)
	}

	return cfg, nil
}
