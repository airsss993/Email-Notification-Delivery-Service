package model

type Template struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Body string `json:"body"`
}
