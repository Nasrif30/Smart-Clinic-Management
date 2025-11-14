package config

import (
	"log"
	"os"

	"[github.com/joho/godotenv](https://github.com/joho/godotenv)"
)

// Config holds all configuration for the application
type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

// LoadConfig reads configuration from .env file or environment variables
func LoadConfig() (*Config, error) {
	// Attempt to load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "26257"), // Default CockroachDB port
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "smartclinic"),
		JWTSecret:  getEnv("JWT_SECRET", "a_very_secret_key_that_should_be_changed"),
	}

	if cfg.JWTSecret == "a_very_secret_key_that_should_be_changed" {
		log.Println("WARNING: JWT_SECRET is not set, using default insecure key.")
	}

	return cfg, nil
}

// getEnv helper function to read an environment variable or return a default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}