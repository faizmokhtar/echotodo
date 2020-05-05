package model

import (
	"time"
)

type Todo struct {
	Title     string `json:"title"`
	ID        uint   `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoStore interface {
	Create(*Todo) error
	List() ([]Todo, error)
	GetByID(todoID int) (*Todo, error)
	Update(*Todo, string) (*Todo, error)
	Delete(*Todo) error
}
