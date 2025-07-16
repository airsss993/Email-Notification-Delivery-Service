package routes

import (
	"github.com/airsss993/email-notification-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(templateHandler *handler.TemplateHandler) *gin.Engine {
	r := gin.New()

	r.POST("/templates", templateHandler.CreateTemplate)

	return r
}
