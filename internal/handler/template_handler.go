package handler

import (
	"github.com/airsss993/email-notification-service/internal/model"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type TemplateHandler struct {
	Store *store.TemplateStore
}

func NewTemplateHandler(s *store.TemplateStore) *TemplateHandler {
	return &TemplateHandler{Store: s}
}

func (h *TemplateHandler) CreateTemplate(c *gin.Context) {
	var template model.Template

	if err := c.Bind(&template); err != nil {
		log.Err(err).Msg("Failed to bind JSON")
		return
	}

	if len(template.Body) < 3 || len(template.Name) < 3 {
		log.Error().Msg("Empty JSON fields")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to validate JSON",
		})
		return
	}

	templateID, err := h.Store.CreateTemplate(c.Request.Context(), template)
	if err != nil {
		log.Err(err).Msg("Failed to create template")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create template",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Created ID Template": templateID,
	})
}
