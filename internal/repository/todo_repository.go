package repository

import (
	"context"

	"github.com/hienvl125/todo-api/internal/domain"
	"github.com/jmoiron/sqlx"
)

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) domain.TodoRepository {
	return &todoRepository{db: db}
}

func (r todoRepository) Select(ctx context.Context) ([]*domain.Todo, error) {
	todos := []*domain.Todo{}
	if err := r.db.Select(&todos, "SELECT * FROM todos ORDER BY id DESC"); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r todoRepository) Insert(ctx context.Context, body string) (*domain.Todo, error) {
	lastInsertId := 0
	if err := r.db.
		QueryRow("INSERT INTO todos(body) VALUES($1) RETURNING id", body).
		Scan(&lastInsertId); err != nil {
		return nil, err
	}

	var todo domain.Todo
	if err := r.db.Get(&todo, "SELECT * FROM todos WHERE id = $1", lastInsertId); err != nil {
		return nil, err
	}

	return &todo, nil
}
