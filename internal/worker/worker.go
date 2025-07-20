package worker

import (
	"context"
	"github.com/airsss993/email-notification-service/internal/queue"
	"github.com/rs/zerolog/log"
	"time"
)

type Worker struct {
	Queue     *queue.TaskQueue
	Processor *Processor
}

func NewWorker(queue *queue.TaskQueue, processor *Processor) *Worker {
	return &Worker{
		Queue:     queue,
		Processor: processor,
	}
}

func (w *Worker) Start(ctx context.Context) {
	for {
		log.Info().Msg("worker started")

		select {
		case <-ctx.Done():
			return
		default:
			task, err := w.Queue.PopTask(ctx)
			if err != nil {
				log.Err(err).Msg("error reading task")
				time.Sleep(1 * time.Second)
				continue
			}

			if task == nil {
				time.Sleep(500 * time.Millisecond)
				continue
			}

			err = w.Processor.Process(ctx, task)
			if err != nil {
				log.Err(err).Msg("error during processing")
				task.RetryCount++
				if task.RetryCount < 3 {
					_ = w.Queue.PushTask(ctx, task)
				}
			}
		}
	}
}
