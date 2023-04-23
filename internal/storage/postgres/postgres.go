package postgres

import (
	"certification/internal/config"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitConnect(cfg config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf("user=%s port=%s password=%s host=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Port, cfg.Password, cfg.Host, cfg.Dbname, cfg.Sslmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
