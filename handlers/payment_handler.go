package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models/model"
	"github.com/hibakun/nova-store/models/response"
)

func CreatePayment(c *fiber.Ctx) error {
	payment := new(model.Payment)
	if err := c.BodyParser(payment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "check your input",
			"data":    err,
		})
	}

	validate := validator.New()
	if err := validate.Struct(payment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&payment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to create payment",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "success create payment",
		"data":    payment,
	})
}

func GetAllPayments(c *fiber.Ctx) error {
	var payments []response.Payment

	database.DB.Find(&payments)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "payments found",
		"data":    payments,
	})
}

func GetPaymentById(c *fiber.Ctx) error {
	var payment response.Payment

	id := c.Params("id")
	if err := database.DB.First(&payment, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "payment not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "payment found",
		"data":    payment,
	})
}
