package config

import (
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL string

	SMTPHost  string
	SMTPPort  int
	SMTPUser  string
	SMTPPass  string
	SMTPEmail string
}

func Load() Config {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	return Config{
		DatabaseURL: os.Getenv("DB_URL"),
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    port,
		SMTPUser:    os.Getenv("SMTP_USER"),
		SMTPPass:    os.Getenv("SMTP_PASS"),
		SMTPEmail:   os.Getenv("SMTP_EMAIL"),
	}
}
