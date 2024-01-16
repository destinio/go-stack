package handlers

import (
	"github.com/destinio/go-stack/api"
	"github.com/gofiber/fiber/v2"
)

func NewTodoHandler(c *fiber.Ctx) error {
	title := c.FormValue("title")
	todoId := api.NewTodo(title)

	return c.Render("partials/todo-list-item", fiber.Map{
		"Id":    todoId,
		"Title": title,
	})
}
