package config

import (
	"time"

	"github.com/spf13/viper"
)

const configPath = "../config/config.yaml"

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Auth     AuthConfig     `yaml:"auth"`
	HTTP     HTTPConfig     `yaml:"http"`
	Database DatabaseConfig `yaml:"database"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type AuthConfig struct {
	SigningKey      string        `yaml:"signingKey"`
	AccessTokenTTL  time.Duration `yaml:"accessTokenTTL"`
	RefreshTokenTTL time.Duration `yaml:"refreshTokenTTL"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewConfig() (*Config, error) {
	var err error
	var config Config

	viper.SetConfigFile(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
