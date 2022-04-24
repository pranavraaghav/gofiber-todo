package main

import "github.com/pranavraagz/gofiber-todo/config/database"

func main() {
	database.Connect()
	database.Migrate()
}
