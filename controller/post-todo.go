package controller

import (
	"net/http"

	"github.com/ericprd/todo-api-go/mock_database"
	"github.com/ericprd/todo-api-go/model"
	"github.com/gin-gonic/gin"
)

type InputTodo struct {
	Title string `json:"title"`
}

func AddTodo(c *gin.Context) {
	var todo InputTodo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": 1, "message": err.Error(), "data": nil })
		return
	}

	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{ "error": 1, "message": "title cannot be empty", "data": nil })
		return
	}

	data := mock_database.GetData()

	newTodo := model.Todo {
		ID: len(data) + 1,
		Title: todo.Title,
		Completed: false,
	}

	mock_database.DATA = append(mock_database.DATA, newTodo)
	c.JSON(http.StatusCreated, gin.H{ "error": 0, "message": "success", "data": newTodo })
}
