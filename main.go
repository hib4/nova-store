package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibakun/nova-store/database"
	"log"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()

	log.Fatal(app.Listen(":8080"))
}
