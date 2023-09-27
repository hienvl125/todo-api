package domain

import "context"

type TodoID int

type Todo struct {
	ID        TodoID `db:"id" json:"id"`
	Body      string `db:"body" json:"body"`
	Completed bool   `db:"completed" json:"completed"`
}

type TodoRepository interface {
	Select(ctx context.Context) ([]*Todo, error)
	Insert(ctx context.Context, body string) (*Todo, error)
}

type TodoUsecase interface {
	List(ctx context.Context) ([]*Todo, error)
	Create(ctx context.Context, body string) (*Todo, error)
}
