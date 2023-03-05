package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hibakun/nova-store/config"
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models/model"
	"github.com/hibakun/nova-store/utils"
	"net/http"
	"time"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

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
	user.Avatar = "avatar.png"

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "couldn't create user",
			"data":    nil,
		})
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"identity": user.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	token, err := t.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to create token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "USER_SESSION",
		Value:    token,
		MaxAge:   3600 * 24 * 7,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		SameSite: "lax",
	})

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user created",
		"data":    user,
		"token":   token,
	})
}
