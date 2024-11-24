package config

import (
	"log"
	"os"
)

type Config struct {
	Port                  string
	FirebaseAPIKey        string
	GoogleCredentialsPath string
}

func GetConfig() *Config {
	cfg := &Config{
		Port:                  getEnv("PORT", "8080"),
		FirebaseAPIKey:        getEnv("FIREBASE_API_KEY", ""),
		GoogleCredentialsPath: getEnv("GOOGLE_APPLICATION_CREDENTIALS", ""),
	}

	if cfg.FirebaseAPIKey == "" {
		log.Println("WARNING: FIREBASE_API_KEY is not set. Sign-In endpoint will not function.")
	}

	if cfg.GoogleCredentialsPath == "" {
		log.Println("ERROR: GOOGLE_APPLICATION_CREDENTIALS is not set. Firebase services will not work.")
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
