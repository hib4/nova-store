package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models"
	"github.com/hibakun/nova-store/utils"
	"net/http"
	"strconv"
)

func CreateGame(c *fiber.Ctx) error {
	game := new(models.Game)
	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "check your input",
			"data":    err,
		})
	}

	validate := validator.New()
	if err := validate.Struct(game); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "error",
			"error":   err.Error(),
		})
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "error on input image",
		})
	}

	if err := utils.CheckContentType(file, "jpg", "jpeg", "png"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	fileName := strconv.FormatInt(carbon.Now().Timestamp(), 10) + file.Filename

	game.Image = fileName

	if err := database.DB.Create(&game).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to create game",
		})
	}

	if err := c.SaveFile(file, fmt.Sprintf("./public/images/%s", fileName)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to store image",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "success create game",
		"data":    game,
	})
}
