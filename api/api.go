package api

import (
	"fmt"

	"github.com/destinio/go-stack/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetTodos(limit int) []models.Todo {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if limit > 0 {
		db.Limit(limit)
	}

	var todos []models.Todo
	db.Find(&todos)

	return todos
}

func NewTodo(title string) string {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create
	id := db.Create(&models.Todo{Title: title})

	return fmt.Sprint(id)
}
