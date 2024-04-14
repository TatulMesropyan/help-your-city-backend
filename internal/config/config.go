package config

import (
	"os"

	"github.com/rs/zerolog/log"
)

type PGConfig struct {
	// Address  string
	Password string
	Name     string
	Username string
	Host     string
}

type Config struct {
	PostgresConfig *PGConfig
}

func getEnv(envKey, fallback string) string {
	if value, ok := os.LookupEnv(envKey); ok {
		return value
	}
	return fallback
}
func LoadConfig() (*Config, error) {
	pgAddress := getEnv("DB_ADDRESS", "5432")
	if pgAddress == "" {
		log.Fatal().Msg("DB_ADDRESS is not set")
	}
	pgPassword := getEnv("DB_PASSWORD", "adrine2009")
	if pgPassword == "" {
		log.Fatal().Msg("DB_PASSWORD is not set")
	}
	pgName := getEnv("DB_NAME", "help_city_data")
	if pgName == "" {
		log.Fatal().Msg("DB_NAME is not set")
	}
	pgUsername := getEnv("DB_USERNAME", "postgres")
	if pgUsername == "" {
		log.Fatal().Msg("DB_USERNAME is not set")
	}
	pgHost := getEnv("DB_HOST", "0.0.0.0")
	if pgHost == "" {
		log.Fatal().Msg("DB_HOST is not set")
	}

	return &Config{
		PostgresConfig: &PGConfig{
			Password: pgPassword,
			Name:     pgName,
			Username: pgUsername,
			Host:     pgHost,
		}}, nil
}
