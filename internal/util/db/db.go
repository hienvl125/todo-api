package db

import (
	"github.com/hienvl125/todo-api/internal/util/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(conf *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", conf.Dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
