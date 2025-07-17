package model

type SendRequest struct {
	TemplateID int64             `json:"template_id"`
	To         string            `json:"to"`
	Params     map[string]string `json:"params"`
}
