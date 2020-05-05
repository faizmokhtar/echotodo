package main

import (
	"net/http"

	"github.com/faizmokhtar/echotodo/handler"
	"github.com/faizmokhtar/echotodo/store"

	"github.com/faizmokhtar/echotodo/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!")
	})

	d := db.New()
	db.AutoMigrate(d)

	ts := store.NewTodoStore(d)
	h := handler.NewHandler(ts)

	e.POST("/todos", h.CreateTodo)
	e.GET("/todos", h.Todos)
	e.GET("/todos/:id", h.GetTodo)
	e.PUT("/todos/:id", h.UpdateTodo)
	e.DELETE("/todos/:id", h.DeleteTodo)

	e.Logger.Fatal(e.Start(":8000"))
}

func deleteTodo(c echo.Context) error {
	return c.JSON(http.StatusBadGateway, "DELETE /todos/:id not implemented")
}
