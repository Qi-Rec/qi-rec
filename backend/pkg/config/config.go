package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`

	HTTPPort string `envconfig:"HTTP_PORT" required:"true"`

	MLHost string `envconfig:"ML_HOST" required:"true"`
	MLPort string `envconfig:"ML_PORT" required:"true"`

	JWTSecret string `envconfig:"JWT_SECRET" required:"true"`

	DbURL         string `envconfig:"DB_URL" required:"true"`
	MigrationPath string `envconfig:"MIGRATION_PATH" required:"true"`
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
