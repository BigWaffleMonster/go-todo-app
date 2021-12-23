package main

import (
	"go-todo-app"
	"go-todo-app/pkg/handler"
	"go-todo-app/pkg/repository"
  "github.com/spf13/viper"
	"go-todo-app/pkg/service"
	"log"
)

func main() {
  if err := initConfig(); err != nil {
    log.Fatalf("Error with initializing configs: %s", err.Error())
  }

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func initConfig() error {
  viper.AddConfigPath("configs")
  viper.SetConfigName("config")
  return viper.ReadInConfig() 
}
