package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TgToken string `envconfig:"TG_TOKEN" required:"true"`
}

func MustLoad() *Config {
	var config Config
	envconfig.MustProcess("", &config)
	return &config
}
