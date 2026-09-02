package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost, DBPort, DBUser, DBPass, DBName string
	AppPort string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5433"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", "postgres"),
		DBName: getEnv("DB_NAME", "restful-api"),
		AppPort: getEnv("APP_PORT", "3000"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

