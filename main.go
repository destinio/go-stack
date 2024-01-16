package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString("<h1>Hello, World!</h1>")
	})
}
