package service

import (
	"bytes"
	"text/template"
)

type TemplateRenderer struct {
}

func Render(templateText string, data map[string]string) (string, error) {
	tpl, _ := template.New("my-template").Parse(templateText)

	var out bytes.Buffer
	_ = tpl.Execute(&out, data)

	out1 := out.String()
	return out1, nil
}
