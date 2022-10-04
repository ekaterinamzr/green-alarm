package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type LoggerConfig struct {
	Level string
}

type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	URI      string
}

type ServerConfig struct {
	Port string
}

type AuthConfig struct {
	Salt       string
	TokenTTL   int
	SigningKey string
}

type MongoDBConfig struct {
	URI string
}

type Config struct {
	Logger   LoggerConfig
	DBType string
	Database DatabaseConfig
	Server   ServerConfig
	Auth     AuthConfig
	MongoDB  MongoDBConfig
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read the configuration file: %w", err)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &c, nil
}
