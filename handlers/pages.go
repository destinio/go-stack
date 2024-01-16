package handlers

import (
	"github.com/destinio/go-stack/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func IndexHandler(c *fiber.Ctx) error {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var todos []models.Todo
	db.Limit(5).Find(&todos)

	return c.Render("todo-list", fiber.Map{
		"Todos": todos,
	})
}
