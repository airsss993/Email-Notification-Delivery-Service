package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	DbPort string
	DbHost string
	DbName string
	DbUser string
}

func Load() *Config {
	cfg := &Config{
		DbPort: checkEnv("DB_PORT"),
		DbHost: checkEnv("DB_HOST"),
		DbName: checkEnv("DB_NAME"),
		DbUser: checkEnv("DB_USER"),
	}

	log.Info().Msg("Configuration loaded successfully")
	return cfg
}

func checkEnv(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Err(err)
	}

	val := os.Getenv(key)
	if val == "" {
		log.Fatal().Msg("Environment variable " + key + " is not set")
	}
	return val
}
