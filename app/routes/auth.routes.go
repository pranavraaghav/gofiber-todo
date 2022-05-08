package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pranavraagz/gofiber-todo/app/handlers"
)

func AuthRoutes(app *fiber.App) {
	r := app.Group("/auth")

	r.Post("/signup", handlers.Signup)
	r.Post("/login", handlers.Login)
}
