package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	GinMode string
	DBUrl   string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	return &Config{
		Port:    getEnv("PORT", "8080"),
		GinMode: getEnv("GIN_MODE", "debug"),
		DBUrl:   getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
