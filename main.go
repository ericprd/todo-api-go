package main

import (
	"github.com/ericprd/todo-api-go/routes"
)

func main() {
	r := routes.SetupRoutes()

	r.Run()
}
