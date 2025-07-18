package worker

import (
	"context"
	"github.com/airsss993/email-notification-service/internal/model"
	"github.com/airsss993/email-notification-service/internal/service"
	"github.com/airsss993/email-notification-service/internal/store"
	"github.com/rs/zerolog/log"
)

type Processor struct {
	Store       *store.TemplateStore
	EmailSender *service.EmailSender
}

func NewProcessor(store *store.TemplateStore, sender *service.EmailSender) *Processor {
	return &Processor{Store: store, EmailSender: sender}
}

func (p *Processor) Process(ctx context.Context, task *model.Task) error {
	templateId := task.TemplateID
	template, err := p.Store.GetTemplateById(ctx, templateId)
	if err != nil {
		log.Err(err).Msg("failed to get template by ID")
		return err
	}

	outputText, err := service.Render(template.Body, task.Params)
	if err != nil {
		log.Err(err).Msg("failed to render template with provided parameters")
		return err
	}

	// TODO: необходимо изменить структуру БД и убрать от туда имя шаблона
	// TODO: изменить структуру SendRequest
	// TODO: передавать в SendMail subject из SendRequest
	err = p.EmailSender.SendEmail(task.To, template.Name, outputText)
	if err != nil {
		log.Err(err).Msg("failed to send email after render")
		return err
	}

	return nil
}
