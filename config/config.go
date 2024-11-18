package config

import (
	"log"
	"os"
)

type Config struct {
	Port           string
	DatabaseURL    string
	FirebaseConfig string
}

func LoadConfig() *Config {
	cfg := &Config{
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		FirebaseConfig: getEnv("FIREBASE_CONFIG", ""),
	}

	if cfg.DatabaseURL == "" {
		log.Println("WARNING: DATABASE_URL is not set. Database integration is not enabled.")
	}

	if cfg.FirebaseConfig == "" {
		log.Println("WARNING: FIREBASE_CONFIG is not set. Firebase authentication is not enabled.")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
