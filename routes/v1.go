package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibakun/nova-store/handlers"
	"github.com/hibakun/nova-store/middleware"
)

func V1(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", handlers.Login)
	v1.Post("/register", handlers.CreateUser)
	v1.Get("/logout", middleware.Protected, handlers.Logout)
}
