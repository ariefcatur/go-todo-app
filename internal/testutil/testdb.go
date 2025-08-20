package testutil

import (
	"go-todo-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		panic(err)
	}
	return db
}
