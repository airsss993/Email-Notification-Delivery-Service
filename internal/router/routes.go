package router

import (
	"github.com/airsss993/email-notification-service/internal/handler"
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), logger.CustomLogger())
	r.GET("/template", handler.TemplateHandler)
	return r
}
