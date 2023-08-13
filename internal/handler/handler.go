package handler

import "github.com/gin-gonic/gin"

func SetupHandlers(
	todoHandler *TodoHandler,
) *gin.Engine {
	router := gin.Default()

	todoRouterGroup := router.Group("/todos")
	todoRouterGroup.GET("/", todoHandler.List)
	todoRouterGroup.POST("/", todoHandler.Create)

	return router
}
