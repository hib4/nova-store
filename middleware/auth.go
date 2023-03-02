package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibakun/nova-store/utils"
)

func Protected(c *fiber.Ctx) error {
	token := c.Cookies("USER_SESSION")

	claims, err := utils.DecodeJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "unauthorized",
		})
	}

	c.Locals("id", claims["id"])

	return c.Next()
}
