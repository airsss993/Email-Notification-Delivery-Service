package main

import (
	"context"
	"database/sql"
	"github.com/airsss993/email-notification-service/internal/config"
	"github.com/airsss993/email-notification-service/internal/handler"
	"github.com/airsss993/email-notification-service/internal/logger"
	"github.com/airsss993/email-notification-service/internal/queue"
	"github.com/airsss993/email-notification-service/internal/routes"
	"github.com/airsss993/email-notification-service/internal/service"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/airsss993/email-notification-service/internal/worker"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Err(err).Msg("failed to connect to database")
	}

	logger.Init()
	cfg := config.Load()

	DB, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open DB connection")
	}

	if err := DB.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	taskQueue := queue.NewTaskQueue(rdb, "send_tasks")
	templateStore := store.NewTemplateHandler(DB)
	templateHandler := handler.NewTemplateHandler(templateStore)
	emailSender := service.EmailSender{
		From:   cfg.SMTPEmail,
		Config: cfg,
	}
	sendHandler := handler.NewEnqueueHandler(templateStore, &emailSender, taskQueue)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	processor := worker.NewProcessor(templateStore, &emailSender)
	worker := worker.NewWorker(taskQueue, processor)

	for i := 0; i < 3; i++ {
		go worker.Start(ctx)
	}

	r := routes.InitRouter(templateHandler, sendHandler)

	err = r.Run()
	if err != nil {
		log.Fatal().Msg("failed to start the service: " + err.Error())
	}
}
