package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title string `json:"title"`
}

type TodoStore interface {
	Create(*Todo) error
}
