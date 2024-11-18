package config

import (
	"log"
	"os"
)

type Config struct {
	Port           string
	DatabaseURL    string
	FirebaseConfig string
	FirebaseAPIKey string
}

func LoadConfig() *Config {
	cfg := &Config{
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		FirebaseConfig: getEnv("FIREBASE_CONFIG", ""),
		FirebaseAPIKey: getEnv("FIREBASE_API_KEY", ""),
	}

	if cfg.FirebaseConfig == "" {
		log.Println("WARNING: FIREBASE_CONFIG is not set. Firebase authentication is disabled.")
	}

	if cfg.FirebaseAPIKey == "" {
		log.Println("WARNING: FIREBASE_API_KEY is not set. Sign-In endpoint will not function.")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
