package service

import (
	"crypto/tls"
	"github.com/airsss993/email-notification-service/internal/config"
	gomail "github.com/go-mail/mail/v2"
	"github.com/rs/zerolog/log"
)

type EmailSender struct {
	From   string
	Config config.Config
}

func (s *EmailSender) SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(s.Config.SMTPHost, s.Config.SMTPPort, s.Config.SMTPUser, s.Config.SMTPPass)
	d.SSL = false
	d.StartTLSPolicy = gomail.MandatoryStartTLS
	d.TLSConfig = &tls.Config{ServerName: s.Config.SMTPHost}

	if err := d.DialAndSend(m); err != nil {
		log.Err(err).Msg("failed to send email")
		return err
	}

	return nil
}
