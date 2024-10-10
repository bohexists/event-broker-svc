package config

import (
	"os"
)

// Config holds the service's configuration values.
type Config struct {
	NatsURL string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() Config {
	return Config{
		NatsURL: os.Getenv("NATS_URL"),
	}
}
