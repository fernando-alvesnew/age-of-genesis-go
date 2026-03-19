package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv              string
	AppPort             string
	MySQLDSN            string
	JWTSecret           string
	JWTExpiresInMinutes int
	PagSeguroBaseURL    string
	PagSeguroToken      string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		AppEnv:              getEnv("APP_ENV", "local"),
		AppPort:             getEnv("APP_PORT", "8080"),
		MySQLDSN:            os.Getenv("MYSQL_DSN"),
		JWTSecret:           os.Getenv("JWT_SECRET"),
		PagSeguroBaseURL:    getEnv("PAGSEGURO_BASE_URL", "https://sandbox.api.pagseguro.com"),
		PagSeguroToken:      os.Getenv("PAGSEGURO_TOKEN"),
		JWTExpiresInMinutes: 60,
	}

	if cfg.MySQLDSN == "" {
		return Config{}, fmt.Errorf("MYSQL_DSN is required")
	}
	if cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("JWT_SECRET is required")
	}
	if cfg.PagSeguroToken == "" {
		return Config{}, fmt.Errorf("PAGSEGURO_TOKEN is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
