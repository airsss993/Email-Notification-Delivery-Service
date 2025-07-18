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

type EnqueuHandler struct {
	Store       *store.TemplateStore
	EmailSender *service.EmailSender
	TaskQueue   *queue.TaskQueue
}

func NewEnqueueHandler(templateStore *store.TemplateStore, emailSender *service.EmailSender, taskQueue *queue.TaskQueue) *EnqueuHandler {
	return &EnqueuHandler{Store: templateStore, EmailSender: emailSender, TaskQueue: taskQueue}
}

func (h *EnqueuHandler) EnqueueEmail(c *gin.Context) {
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

	log.Info().Msg("task sended to Redis")
	c.JSON(http.StatusOK, gin.H{"message": "task sended to Redis"})
}
