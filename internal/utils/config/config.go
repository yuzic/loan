package config

import (
	"errors"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type HTTP struct {
	Port int
}

type DB struct {
	Host            string
	Port            int
	Name            string
	User            string
	Password        string
	Debug           bool
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}

func DBConfig() DB {
	return DB{
		Host:            viper.GetString("DATABASE_HOST"),
		Port:            viper.GetInt("DATABASE_PORT"),
		Name:            viper.GetString("DATABASE_NAME"),
		User:            viper.GetString("DATABASE_USER"),
		Password:        viper.GetString("DATABASE_PASSWORD"),
		Debug:           viper.GetBool("DATABASE_DEBUG"),
		MaxOpenConns:    viper.GetInt("DATABASE_MAX_OPEN_CONNS"),
		MaxIdleConns:    viper.GetInt("DATABASE_MAX_IDLE_CONNS"),
		ConnMaxLifetime: viper.GetDuration("DATABASE_CONN_MAX_LIFE_TIME") * time.Second,
	}
}

func HTTPConfig() HTTP {
	return HTTP{
		Port: viper.GetInt("HTTP_PORT"),
	}
}

func InitViperByEnv(configPath string) {
	if configPath == "" {
		zap.S().Info("Config path not provided, using environment variables")
		viper.AutomaticEnv()
		return
	}

	if _, err := os.Stat(configPath); err != nil {
		if os.IsNotExist(err) {
			zap.S().Errorf("Config file does not exist: %s", configPath)
		} else {
			zap.S().Errorf("Error checking config file: %v", err)
		}
		viper.AutomaticEnv()
		return
	}

	viper.SetConfigFile(configPath)

	readInConfig()
}

func readInConfig() {
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			zap.S().Panicf("Config file not found, using defaults settings: %v", err)
		}
	}
}
