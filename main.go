package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pranavraagz/gofiber-todo/app/dal"
	"github.com/pranavraagz/gofiber-todo/app/routes"
	"github.com/pranavraagz/gofiber-todo/config"
	"github.com/pranavraagz/gofiber-todo/config/database"
	"github.com/pranavraagz/gofiber-todo/utils"
)

func main() {
	database.Connect()
	database.Migrate(&dal.Todo{}, &dal.User{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	routes.TodoRoutes(app)
	routes.AuthRoutes(app)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
