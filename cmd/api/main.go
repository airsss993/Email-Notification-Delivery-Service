package main

import (
	"github.com/airsss993/email-notification-service/internal/api"
	"github.com/airsss993/email-notification-service/internal/database"
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize logger
	logger.Init()

	r := gin.New()
	r.Use(gin.Recovery(), logger.CustomLogger())

	api.SetupRoutes(r)

	_, err := database.DBConn()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
		return
	} else {
		log.Info().Msg("Database connection established successfully")
	}
	
	err = r.Run()
	if err != nil {
		log.Fatal().Msg("Failed to start the service: " + err.Error())
	}

}
