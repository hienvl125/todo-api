package service

import (
	"context"

	"github.com/hienvl125/todo-api/internal/domain"
)

type todoService struct {
	todoRepository domain.TodoRepository
}

func NewTodoService(todoRepository domain.TodoRepository) domain.TodoUsecase {
	return &todoService{todoRepository: todoRepository}
}

func (s todoService) List(ctx context.Context) ([]*domain.Todo, error) {
	todos, err := s.todoRepository.Select(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil

}

func (s todoService) Create(ctx context.Context, body string) (*domain.Todo, error) {
	todo, err := s.todoRepository.Insert(ctx, body)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
