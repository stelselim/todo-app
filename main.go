package main

import (
	"todo-app/app"
)

func main() {
	router := app.SetupRouter()
	router.Run(":8080")
}
