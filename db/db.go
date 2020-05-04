package db

import (
	"fmt"

	"github.com/faizmokhtar/echotodo/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "echotodo_dev"
)

func New() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Todo{},
	)
}
