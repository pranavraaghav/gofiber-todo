package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pranavraagz/gofiber-todo/app/handlers"
)

func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo")

	r.Post("/create", handlers.CreateTodo)
}
