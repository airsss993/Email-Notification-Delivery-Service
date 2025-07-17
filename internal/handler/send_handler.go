package handler

import (
	"github.com/airsss993/email-notification-service/internal/model"
	"github.com/airsss993/email-notification-service/internal/service"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	// TODO: подключить пакет рендерера, когда будет готов
	// "github.com/airsss993/email-notification-service/internal/service"
)

type SendHandler struct {
	Store *store.TemplateStore
	// TODO: добавить поле TemplateRenderer, чтобы рендерить шаблоны
	//Renderer

	// TODO: позже — добавить EmailSender для отправки писем
	// EmailSender *service.EmailSender
}

func NewSendHandler(templateStore *store.TemplateStore) *SendHandler {
	return &SendHandler{Store: templateStore}
}

func (h *SendHandler) SendNotification(c *gin.Context) {
	var sendRequest *model.SendRequest

	if err := c.BindJSON(&sendRequest); err != nil {
		log.Err(err).Msg("failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON"})
		return
	}

	// TODO: сделать валидацию JSON на пустые поля

	templateId := sendRequest.TemplateID
	template, err := h.Store.GetTemplateById(c.Request.Context(), templateId)
	if err != nil {
		log.Err(err).Msg("failed to get template by ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get template by ID"})
		return
	}

	// TODO: обработать ошибку
	// TODO: отрендерить body через TemplateRenderer
	outputText, _ := service.Render(template.Body, sendRequest.Params)

	// - вернуть 200 OK или ошибку

	c.JSON(http.StatusOK, gin.H{"sended message": outputText})
}
