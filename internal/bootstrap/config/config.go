package config

import (
	"loan/internal/utils/config"
)

type AppConfig struct {
	db   config.DB
	HTTP config.HTTP
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		db:   config.DBConfig(),
		HTTP: config.HTTPConfig(),
	}
}
