package main

import (
	"log"
	"todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
