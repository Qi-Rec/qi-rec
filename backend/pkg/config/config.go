package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load env variables: %w", err)
	}

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to process env variables: %w", err)
	}

	return &cfg, nil
}
