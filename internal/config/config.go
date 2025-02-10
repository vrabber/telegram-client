package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TgToken string
}

func Load() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("failed to process env vars: %v", err)
	}
	return &config, nil
}
