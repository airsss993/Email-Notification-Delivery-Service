package model

import "time"

type Task struct {
	TemplateID int64             `json:"template_id"`
	To         string            `json:"to"`
	Params     map[string]string `json:"params"`
	CreatedAt  time.Time         `json:"created_at"`
	RetryCount int               `json:"retry_count"`
}
