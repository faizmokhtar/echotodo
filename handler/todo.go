package handler

import (
	"net/http"

	"github.com/faizmokhtar/echotodo/model"
	"github.com/faizmokhtar/echotodo/utils"
	"github.com/labstack/echo"
)

func (h *Handler) CreateTodo(c echo.Context) error {
	todo := &model.Todo{}
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	err := h.todoStore.Create(todo)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))

	}

	return c.JSON(http.StatusCreated, todo)
}

func (h *Handler) Todos(c echo.Context) error {
	var todos []model.Todo
	todos, err := h.todoStore.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todos)
}
