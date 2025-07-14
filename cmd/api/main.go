package main

import (
	"database/sql"
	"github.com/airsss993/email-notification-service/internal/config"
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/airsss993/email-notification-service/internal/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Load .env
	_ = godotenv.Load()

	// Initialize logger and configuration
	logger.Init()
	config.Load()

	// Initialize database connection
	connStr := os.Getenv("DB_URL")
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	} else {
		log.Info().Msg("Database connection established successfully")
	}

	// Setup router
	r := router.SetupRouter()

	// Start the server
	err = r.Run()
	if err != nil {
		log.Fatal().Msg("Failed to start the service: " + err.Error())
	}

}
