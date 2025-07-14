package handler

import (
	"github.com/gin-gonic/gin"
)

func TemplateHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is a template handler",
	})
}
