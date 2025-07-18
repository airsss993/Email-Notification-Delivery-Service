package routes

import (
	"github.com/airsss993/email-notification-service/internal/handler"
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/gin-gonic/gin"
)

func InitRouter(templateHandler *handler.TemplateHandler, sendHandler *handler.EnqueuHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), logger.CustomLogger())

	r.POST("/templates", templateHandler.CreateTemplate)
	r.POST("/send", sendHandler.EnqueueEmail)
	r.GET("/templates/:id", templateHandler.GetTemplateById)

	return r
}
