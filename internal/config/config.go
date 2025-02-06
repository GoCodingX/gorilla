package config

import (
	"os"
)

type Config struct {
	Port                 string
	AppleSigningKey      string
	FeatureFlagAPIURL    string
	AppleRootCertificate string
}

func Load() (*Config, error) {
	return &Config{
		Port:                 getEnvOrDefault("PORT", ":8080"),
		FeatureFlagAPIURL:    os.Getenv("FEATURE_FLAG_API_URL"),
		AppleRootCertificate: os.Getenv("APPLE_ROOT_CERTIFICATE"),
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
