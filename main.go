package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hibakun/nova-store/config"
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/routes"
	"log"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	routes.V1(app)

	log.Fatal(app.Listen("0.0.0.0:" + config.Config("PORT")))
}
