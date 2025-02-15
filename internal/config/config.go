package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TgToken         string `envconfig:"TG_TOKEN" required:"true"`
	ServerHost      string `envconfig:"VRABBER_HOST" default:"localhost"`
	ServerPort      int    `envconfig:"VRABBER_PORT" required:"true"`
	MessagesBuffer  int    `envconfig:"MESSAGES_BUFFER" default:"100"`
	ResponsesBuffer int    `envconfig:"RESPONSES_BUFFER" default:"100"`
	LogLevel        string `envconfig:"LOG_LEVEL" default:"INFO"`
}

func MustLoad() *Config {
	var config Config
	envconfig.MustProcess("", &config)

	if config.TgToken == "" {
		panic("TG_TOKEN env variable must not be empty")
	}
	if config.ServerHost == "" {
		panic("VRABBER_HOST env variable must not be empty")
	}
	if config.ServerPort < 1 || config.ServerPort > 65535 {
		panic("VRABBER_PORT env variable must be between 1 and 65535")
	}
	if config.MessagesBuffer < 1 || config.MessagesBuffer > 1000 {
		panic("MESSAGES_BUFFER env variable must be between 1 and 1000")
	}
	if config.ResponsesBuffer < 1 || config.ResponsesBuffer > 1000 {
		panic("RESPONSES_BUFFER env variable must be between 1 and 1000")
	}
	return &config
}
