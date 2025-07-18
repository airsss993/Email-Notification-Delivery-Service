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

func (q *TaskQueue) PopTask(ctx context.Context) (*model.Task, error) {
	res, err := q.client.BRPop(ctx, 0, q.queueName).Result()
	if err != nil {
		log.Err(err).Msg("failed to BRPOP task from Redis")
		return nil, err
	}

	var task model.Task
	err = json.Unmarshal([]byte(res[1]), &task)
	if err != nil {
		log.Err(err).Msg("failed to unmarshal data")
		return nil, err
	}

	return &task, nil
}
