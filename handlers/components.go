package handlers

import (
	"github.com/destinio/go-stack/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TodosListHandler(c *fiber.Ctx) error {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var todos []models.Todo
	db.Find(&todos)

	return c.Render("todo-list", fiber.Map{
		"Todos": todos,
	})
}
