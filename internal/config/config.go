package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TgToken    string `envconfig:"TG_TOKEN" required:"true"`
	ServerHost string `envconfig:"VRABBER_HOST" default:"localhost"`
	ServerPort int    `envconfig:"VRABBER_PORT" required:"true"`
}

func MustLoad() *Config {
	var config Config
	envconfig.MustProcess("", &config)

	if config.TgToken == "" {
		panic("TG_TOKEN env variable must not be empty")
	}

	return &config
}
