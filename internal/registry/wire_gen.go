// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/hienvl125/todo-api/internal/handler"
	"github.com/hienvl125/todo-api/internal/repository"
	"github.com/hienvl125/todo-api/internal/service"
	"github.com/hienvl125/todo-api/internal/util/config"
	"github.com/hienvl125/todo-api/internal/util/db"
)

// Injectors from wire.go:

func NewGinServer(conf *config.Config) (*gin.Engine, error) {
	sqlxDB, err := db.NewPostgresDB(conf)
	if err != nil {
		return nil, err
	}
	todoRepository := repository.NewTodoRepository(sqlxDB)
	todoUsecase := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoUsecase)
	engine := handler.SetupHandlers(todoHandler)
	return engine, nil
}
