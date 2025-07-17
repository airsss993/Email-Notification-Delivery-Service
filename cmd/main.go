package main

import (
	"database/sql"
	"github.com/airsss993/email-notification-service/internal/config"
	"github.com/airsss993/email-notification-service/internal/handler"
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/airsss993/email-notification-service/internal/routes"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Err(err).Msg("failed to connect to database")
	}

	// Initialize logger and configuration

	logger.Init()
	cfg := config.Load()

	// Initialize database connection

	DB, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open DB connection")
	}

	if err := DB.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	templateStore := store.TemplateStore{DB: DB}
	templateHandler := handler.TemplateHandler{Store: &templateStore}
	sendHandler := handler.SendHandler{Store: &templateStore}

	// Setup routes

	r := routes.InitRouter(&templateHandler, &sendHandler)

	// Start the server

	err = r.Run()
	if err != nil {
		log.Fatal().Msg("failed to start the service: " + err.Error())
	}
}
