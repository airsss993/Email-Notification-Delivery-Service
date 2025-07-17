package handler

import (
	"github.com/airsss993/email-notification-service/internal/model"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
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

func (h *TemplateHandler) GetTemplateById(c *gin.Context) {
	tmplIdStr, ok := c.Params.Get("id")
	if !ok {
		log.Error().Msg("Failed to parse template ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid template ID"})
		return
	}

	templateId, err := strconv.ParseInt(tmplIdStr, 10, 64)
	if err != nil {
		log.Err(err).Msg("Failed to convert ID from path")
		return
	}

	template, err := h.Store.GetTemplateById(c.Request.Context(), templateId)
	if err != nil {
		log.Err(err).Msg("Failed to select template")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to select template"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"template": template})
}
