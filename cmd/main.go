package main

import (
	"go-todo-app"
	"go-todo-app/pkg/handler"
	"go-todo-app/pkg/repository"
	"go-todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
