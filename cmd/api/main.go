package main

import (
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init()
	//cfg := config.Load()
	log.Info().Msg("Service started successfully")
}
