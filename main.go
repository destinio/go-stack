package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

func main() {
	engine := html.New("./views", ".html")
	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// static files
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Todos!",
			"Todos": []Todo{
				{1, "a todo", false},
				{2, "Crh blaheate a todo list", false},
				{3, "and this one", false},
			},
		})
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString("<h1>Hello, World!</h1>")
	})

	app.Listen(":1337")
}
