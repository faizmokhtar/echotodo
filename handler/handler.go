package handler

import "github.com/faizmokhtar/echotodo/model"

type Handler struct {
	todoStore model.TodoStore
}

func NewHandler(ts model.TodoStore) *Handler {
	return &Handler{
		todoStore: ts,
	}
}
