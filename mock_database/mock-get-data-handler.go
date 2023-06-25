package mock_database

import "github.com/ericprd/todo-api-go/model"

func GetData() []model.Todo {
	return DATA
}
