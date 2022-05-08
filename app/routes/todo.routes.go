package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pranavraagz/gofiber-todo/app/handlers"
	"github.com/pranavraagz/gofiber-todo/utils/middleware"
)

func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo").Use(middleware.Auth)

	r.Post("/create", handlers.CreateTodo)
}
