package config

import (
	"os"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	DatabasePath  string
	ServerPort    string
	Environment   string
	MaxOpenConns  int
	MaxIdleConns  int
	OpenAIAPIKey  string
}

// Load reads configuration from environment variables with sensible defaults
func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	return &Config{
		DatabasePath: getEnv("DB_PATH", "./adrian.db"),
		ServerPort:   getEnv("PORT", "8080"),
		Environment:  getEnv("ENV", "development"),
		MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	godotenv.Load()

	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt retrieves an environment variable as int or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	godotenv.Load()

	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
