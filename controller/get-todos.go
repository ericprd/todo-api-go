package controller

import (
	"net/http"

	"github.com/ericprd/todo-api-go/mock_database"
	"github.com/ericprd/todo-api-go/model"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []model.Todo

	todos = mock_database.GetData()

	c.JSON(http.StatusOK, gin.H{ "error": 0, "message": "success", "data": todos })
}
