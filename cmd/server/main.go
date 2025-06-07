package main

import (
	"log"

	"github.com/example/go-todo-app/internal/entity"
	"github.com/example/go-todo-app/internal/infrastructure/mysql"
	"github.com/example/go-todo-app/internal/interface/controller"
	"github.com/example/go-todo-app/internal/router"
	"github.com/example/go-todo-app/internal/usecase"
)

func main() {
	db, err := mysql.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&entity.Todo{}); err != nil {
		log.Fatal(err)
	}

	todoRepo := mysql.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoController := controller.NewTodoController(todoUsecase)

	e := router.NewRouter(todoController)
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
