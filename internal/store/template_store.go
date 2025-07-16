package store

import (
	"context"
	"database/sql"
	"github.com/airsss993/email-notification-service/internal/model"
)

type TemplateStore struct {
	DB *sql.DB
}

func NewTemplateHandler(db *sql.DB) *TemplateStore {
	return &TemplateStore{DB: db}
}

func (s *TemplateStore) CreateTemplate(ctx context.Context, template model.Template) (int64, error) {
	var id int64

	query := `INSERT INTO templates(type, name, body) VALUES($1, $2, $3) RETURNING id;`
	err := s.DB.QueryRowContext(ctx, query, template.Type, template.Name, template.Body).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
