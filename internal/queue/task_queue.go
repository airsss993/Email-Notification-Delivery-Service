package queue

import (
	"context"
	"encoding/json"
	"github.com/airsss993/email-notification-service/internal/model"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type TaskQueue struct {
	client    *redis.Client
	queueName string
}

func NewTaskQueue(client *redis.Client, queueName string) *TaskQueue {
	return &TaskQueue{
		client:    client,
		queueName: queueName,
	}
}

func (q *TaskQueue) PushTask(ctx context.Context, task *model.Task) error {
	taskJson, err := json.Marshal(task)
	if err != nil {
		log.Err(err).Msg("failed to marshal task")
		return err
	}

	err = q.client.LPush(ctx, q.queueName, taskJson).Err()
	if err != nil {
		log.Err(err).Msg("failed to LPUSH task to Redis")
		return err
	}

	return nil
}
