package routes

import (
	"github.com/ericprd/todo-api-go/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/todos", controller.GetTodos)
	r.POST("/todos", controller.AddTodo)

	return r
}
