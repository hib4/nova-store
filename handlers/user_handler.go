package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models"
	"github.com/hibakun/nova-store/utils"
	"log"
	"net/http"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "check your input",
			"data":    err,
		})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if _, err := utils.GetUserByEmail(user.Email); err != false {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "email already exist",
		})
	}

	if _, err := utils.GetUserByPhoneNumber(user.PhoneNumber); err != false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "phone number already exist",
		})
	}

	hash, errHash := utils.HashPassword(user.Password)
	if errHash != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "couldn't hash password",
			"data":    errHash,
		})
	}

	user.Password = hash
	user.Avatar = "avatar.jpg"

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "couldn't create user",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user created",
		"data":    user,
	})
}
