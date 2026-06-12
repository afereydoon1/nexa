package config

import (
	"fmt"
	"os"
)

// Config holds all the configuration for the application
type Config struct {
	AppPort      string
	AppEnv       string
	DBConnection string
	DBUser       string
	DBPassword   string
	DBHost       string
	DBPort       string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
}

// LoadConfig loads environment variables into a Config struct
// Returns error if any required config (like JWT_SECRET) is missing
func LoadConfig() (*Config, error) {
	cfg := &Config{
		AppPort:      getEnv("APP_PORT", "3000"),
		AppEnv:       getEnv("APP_ENV", "development"),
		DBConnection: getEnv("DB_CONNECTION", "pgsql"),
		DBUser:       getEnv("DB_USERNAME", ""),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_DATABASE", ""),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBSSLMode:    getEnv("DB_SSLMODE", "disable"), // configurable SSL mode for production
		JWTSecret:    getEnv("JWT_SECRET", ""),
	}

	// Validate critical config values
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	if cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("database credentials must be set")
	}

	return cfg, nil
}

// getEnv returns the environment variable value if exists, otherwise returns the fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
