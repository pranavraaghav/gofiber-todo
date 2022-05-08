package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pranavraagz/gofiber-todo/config"
	"github.com/pranavraagz/gofiber-todo/config/database"
	"github.com/pranavraagz/gofiber-todo/utils"
)

func main() {
	database.Connect()
	database.Migrate()

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
