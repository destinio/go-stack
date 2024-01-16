package main

import (
	"github.com/destinio/go-stack/handlers"
	"github.com/destinio/go-stack/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Todo{})

	engine := html.New("./views", ".html")
	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// static files
	app.Static("/", "./public")

	app.Get("/", handlers.IndexHandler)
	app.Get("/todos", handlers.TodosListHandler)

	app.Post("/new", handlers.NewTodoHandler)

	app.Listen("localhost:1337")
}
