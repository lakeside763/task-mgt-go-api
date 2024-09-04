package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL  string
	RedisURL     string
	DatabaseType string
}

func LoadConfig() *Config {
	databaseType := os.Getenv("DATABASE_TYPE")

	var databaseURL string
	if databaseType == "mongodb" {
		databaseURL = os.Getenv("DATABASE_URL")
	} else if databaseType == "postgresql" {
		databaseURL = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	}

	return &Config{
		DatabaseURL:  databaseURL,
		RedisURL:     os.Getenv("REDIS_URL"),
		DatabaseType: databaseType,
	}
}
