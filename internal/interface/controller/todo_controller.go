package controller

import (
	"net/http"
	"strconv"

	"github.com/example/go-todo-app/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	u usecase.TodoUsecase
}

func NewTodoController(u usecase.TodoUsecase) *TodoController {
	return &TodoController{u: u}
}

func (c *TodoController) Register(g *echo.Group) {
	g.POST("/todos", c.create)
	g.GET("/todos", c.list)
	g.PUT("/todos/:id", c.complete)
}

func (c *TodoController) create(ctx echo.Context) error {
	var req struct {
		Title string `json:"title"`
	}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todo, err := c.u.Add(req.Title)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) list(ctx echo.Context) error {
	todos, err := c.u.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) complete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := c.u.Complete(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusOK)
}
