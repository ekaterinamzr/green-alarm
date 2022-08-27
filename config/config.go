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
	URL      string
}

type ServerConfig struct {
	Port int
}

type Config struct {
	Logger     LoggerConfig
	Database   DatabaseConfig
	Server     ServerConfig
	ExampleVar string
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
