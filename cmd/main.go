package main

import (
	"database/sql"
	"github.com/airsss993/email-notification-service/internal/config"
	"github.com/airsss993/email-notification-service/internal/handler"
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/airsss993/email-notification-service/internal/routes"
	"github.com/airsss993/email-notification-service/internal/service"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Err(err).Msg("failed to connect to database")
	}

	logger.Init()
	cfg := config.Load()

	DB, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open DB connection")
	}

	if err := DB.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	templateStore := store.TemplateStore{DB: DB}
	templateHandler := handler.TemplateHandler{Store: &templateStore}
	emailSender := service.EmailSender{
		From:   cfg.SMTPEmail,
		Config: cfg,
	}

	sendHandler := handler.SendHandler{Store: &templateStore, EmailSender: &emailSender}
	
	log.Info().Msgf("SMTP host=%s port=%d user=%s", cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser)

	r := routes.InitRouter(&templateHandler, &sendHandler)

	err = r.Run()
	if err != nil {
		log.Fatal().Msg("failed to start the service: " + err.Error())
	}
}
