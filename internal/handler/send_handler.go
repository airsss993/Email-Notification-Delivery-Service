package handler

import (
	"github.com/airsss993/email-notification-service/internal/model"
	"github.com/airsss993/email-notification-service/internal/service"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type SendHandler struct {
	Store       *store.TemplateStore
	EmailSender *service.EmailSender
}

func NewSendHandler(templateStore *store.TemplateStore, emailSender *service.EmailSender) *SendHandler {
	return &SendHandler{Store: templateStore, EmailSender: emailSender}
}

func (h *SendHandler) SendEmail(c *gin.Context) {
	var sendRequest *model.SendRequest

	if err := c.BindJSON(&sendRequest); err != nil {
		log.Err(err).Msg("failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON"})
		return
	}

	if len(sendRequest.Params) < 1 {
		log.Warn().Msg("no parameters provided for template rendering")
		c.JSON(http.StatusBadRequest, gin.H{"error": "no parameters provided"})
		return
	}

	templateId := sendRequest.TemplateID
	template, err := h.Store.GetTemplateById(c.Request.Context(), templateId)
	if err != nil {
		log.Err(err).Msg("failed to get template by ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get template by ID"})
		return
	}

	outputText, err := service.Render(template.Body, sendRequest.Params)
	if err != nil {
		log.Err(err).Msg("failed to render template with provided parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to render template with provided parameters"})
		return
	}

	err = h.EmailSender.SendEmail(sendRequest.To, "Test", outputText)
	if err != nil {
		log.Err(err).Msg("failed to send email after render")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send email after render"})
		return
	}

	log.Info().Msg("successfully send message")
	c.JSON(http.StatusOK, gin.H{"sended message": outputText})
}
