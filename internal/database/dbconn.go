package database

import (
	"database/sql"
	"fmt"
	"github.com/airsss993/email-notification-service/internal/config"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func DBConn() (*sql.DB, error) {
	cfg := config.Load()
	connStr := fmt.Sprintf("user=%s dbname=%s port=%s host=%s sslmode=disable", cfg.DbUser, cfg.DbName, cfg.DbPort, cfg.DbHost)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to the database")
		return nil, err
	}

	return db, nil
}
