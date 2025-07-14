package api

import (
	"github.com/airsss993/email-notification-service/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/template", handler.TemplateHandler)
}
