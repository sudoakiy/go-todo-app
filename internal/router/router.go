package router

import (
	"github.com/example/go-todo-app/internal/interface/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(todoController *controller.TodoController) *echo.Echo {
	e := echo.New()
	api := e.Group("/api")
	todoController.Register(api)
	return e
}
