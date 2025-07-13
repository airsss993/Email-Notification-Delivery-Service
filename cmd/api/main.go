package main

import (
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init()
	//cfg := config.Load()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := router.Run()
	if err != nil {
		log.Fatal().Msg("Failed to start the service: " + err.Error())
	} else {
		log.Info().Msg("Service started on port :8080")
	}

}
