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

	e.POST("/todos", createTodo)
	e.GET("/todos/:id", getTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	d := db.New()
	db.AutoMigrate(d)

	ts := store.NewTodoStore(d)
	h := handler.NewHandler(ts)

	e.POST("/todos", h.CreateTodo)
	e.GET("/todos", h.Todos)

	e.Logger.Fatal(e.Start(":8000"))
}

func createTodo(c echo.Context) error {
	return c.JSON(http.StatusBadGateway, "POST /todos not implemented")
}

func getTodo(c echo.Context) error {
	return c.JSON(http.StatusBadGateway, "GET /todos/:id not implemented")
}

func updateTodo(c echo.Context) error {
	return c.JSON(http.StatusBadGateway, "PUT /todos/:id not implemented")
}

func deleteTodo(c echo.Context) error {
	return c.JSON(http.StatusBadGateway, "DELETE /todos/:id not implemented")
}
