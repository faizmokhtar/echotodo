package handler

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.todoStore.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todo)
}

func (h Handler) UpdateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.todoStore.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	p := &UpdateRequestPayload{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	todo, err = h.todoStore.Update(todo, p.Title)

	return c.JSON(http.StatusOK, todo)
}

type UpdateRequestPayload struct {
	Title string `json:"title"`
}
