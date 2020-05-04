package store

import (
	"github.com/faizmokhtar/echotodo/model"
	"github.com/jinzhu/gorm"
)

type TodoStore struct {
	db *gorm.DB
}

func NewTodoStore(db *gorm.DB) *TodoStore {
	return &TodoStore{
		db: db,
	}
}

func (ts *TodoStore) Create(t *model.Todo) error {
	return ts.db.Create(t).Error
}
