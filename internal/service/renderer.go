package service

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"text/template"
)

func Render(templateText string, data map[string]string) (string, error) {
	tpl, err := template.New("my-template").Parse(templateText)
	if err != nil {
		log.Err(err).Msg("failed to parse template")
		return "", err
	}

	var out bytes.Buffer
	err = tpl.Execute(&out, data)
	if err != nil {
		log.Err(err).Msg("failed to execute template")
		return "", err
	}

	out1 := out.String()
	return out1, nil
}
