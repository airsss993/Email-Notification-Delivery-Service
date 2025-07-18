package handler

import (
	"github.com/airsss993/email-notification-service/internal/model"
	"github.com/airsss993/email-notification-service/internal/queue"
	"github.com/airsss993/email-notification-service/internal/service"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type SendHandler struct {
	Store       *store.TemplateStore
	EmailSender *service.EmailSender
	TaskQueue   *queue.TaskQueue
}

func NewSendHandler(templateStore *store.TemplateStore, emailSender *service.EmailSender, taskQueue *queue.TaskQueue) *SendHandler {
	return &SendHandler{Store: templateStore, EmailSender: emailSender, TaskQueue: taskQueue}
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

	task := &model.Task{
		TemplateID: sendRequest.TemplateID,
		To:         sendRequest.To,
		Params:     sendRequest.Params,
		CreatedAt:  time.Now(),
		RetryCount: 0,
	}

	err := h.TaskQueue.PushTask(c, task)
	if err != nil {
		log.Err(err).Msg("failed to enqueue task to Redis")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enqueue task"})
		return
	}

	templateId := sendRequest.TemplateID
	template, err := h.Store.GetTemplateById(c.Request.Context(), templateId)
	if err != nil {
		log.Err(err).Msg("failed to get template by ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get template by ID"})
		return
	}

	task, err = h.TaskQueue.PopTask(c)
	if err != nil {
		log.Err(err).Msg("failed to pop task from Redis queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve task from queue"})
		return
	}

	outputText, err := service.Render(template.Body, sendRequest.Params)
	if err != nil {
		log.Err(err).Msg("failed to render template with provided parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to render template with provided parameters"})
		return
	}

	// TODO: необходимо изменить структуру БД и убрать от туда имя шаблона
	// TODO: изменить структуру SendRequest
	// TODO: передавать в SendMail subject из SendRequest
	err = h.EmailSender.SendEmail(sendRequest.To, "Test Subject", outputText)
	if err != nil {
		log.Err(err).Msg("failed to send email after render")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send email after render"})
		return
	}

	log.Info().Msg("successfully send message")
	c.JSON(http.StatusOK, gin.H{"sended message": outputText})
}
