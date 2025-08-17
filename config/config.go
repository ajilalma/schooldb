package config

import (
	"os"
)

type Config struct {
	Port string
	DbURL string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "localhost:5432" // default database URL
	}

	return &Config{
		Port: port,
		DbURL: dbURL,
	}, nil
}