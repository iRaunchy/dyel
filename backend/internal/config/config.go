package config

import (
	"fmt"
	"os"
)

// Config holds all application settings.
type Config struct {
	Port        string
	DatabaseURL string
}

// Load reads from the environment (or defaults) and constructs
func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		user := os.Getenv("POSTGRES_USER")
		pass := os.Getenv("POSTGRES_PASSWORD")
		name := os.Getenv("POSTGRES_DB")
		host := os.Getenv("POSTGRES_HOST")
		if host == "" {
			host = "db" // Docker Compose service name
		}
		portDB := os.Getenv("POSTGRES_PORT")
		if portDB == "" {
			portDB = "5432"
		}

		if user == "" || pass == "" || name == "" {
			return nil, fmt.Errorf(
				"either DATABASE_URL or all POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB must be set",
			)
		}

		dbURL = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			user, pass, host, portDB, name,
		)
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}, nil
}
