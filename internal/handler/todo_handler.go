package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hienvl125/todo-api/internal/domain"
)

type TodoHandler struct {
	todoService domain.TodoUsecase
}

func NewTodoHandler(todoService domain.TodoUsecase) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h TodoHandler) List(ctx *gin.Context) {
	todos, err := h.todoService.List(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func (h TodoHandler) Create(ctx *gin.Context) {
	var todoInput domain.Todo
	if err := ctx.BindJSON(&todoInput); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if todoInput.Body == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "body can't be blank",
		})
		return
	}

	todo, err := h.todoService.Create(ctx, todoInput.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}
