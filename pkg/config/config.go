package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Address string
}

// DatabaseConfig stores MySQL connection details.
type DatabaseConfig struct {
	DSN string
}

// AuthConfig stores authentication related settings.
type AuthConfig struct {
	TokenSecret string
	TokenTTL    time.Duration
}

// Config aggregates all application configuration.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

const (
	defaultServerAddress = ":7070"
	defaultTokenTTL      = time.Hour * 24
)

// Load reads configuration from environment variables with sensible defaults.
func Load() (*Config, error) {
	cfg := &Config{}

	cfg.Server.Address = valueOrDefault(os.Getenv("ADMIN_SERVER_ADDR"), defaultServerAddress)
	cfg.Database.DSN = os.Getenv("ADMIN_MYSQL_DSN")
	if cfg.Database.DSN == "" {
		return nil, fmt.Errorf("environment variable ADMIN_MYSQL_DSN is required")
	}

	cfg.Auth.TokenSecret = os.Getenv("ADMIN_AUTH_SECRET")
	if cfg.Auth.TokenSecret == "" {
		return nil, fmt.Errorf("environment variable ADMIN_AUTH_SECRET is required")
	}

	cfg.Auth.TokenTTL = parseDurationOrDefault(os.Getenv("ADMIN_AUTH_TOKEN_TTL"), defaultTokenTTL)

	return cfg, nil
}

func valueOrDefault(value, def string) string {
	if value == "" {
		return def
	}
	return value
}

func parseDurationOrDefault(value string, def time.Duration) time.Duration {
	if value == "" {
		return def
	}
	if ttl, err := time.ParseDuration(value); err == nil {
		return ttl
	}
	if hours, err := strconv.Atoi(value); err == nil {
		return time.Duration(hours) * time.Hour
	}
	return def
}
