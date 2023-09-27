//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hienvl125/todo-api/internal/handler"
	"github.com/hienvl125/todo-api/internal/repository"
	"github.com/hienvl125/todo-api/internal/service"
	"github.com/hienvl125/todo-api/internal/util/config"
	"github.com/hienvl125/todo-api/internal/util/db"
)

func NewGinServer(conf *config.Config) (*gin.Engine, error) {
	wire.Build(
		handler.SetupHandlers,
		handler.NewTodoHandler,
		service.NewTodoService,
		repository.NewTodoRepository,
		db.NewPostgresDB,
	)

	return nil, nil
}
