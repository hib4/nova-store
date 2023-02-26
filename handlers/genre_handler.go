package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models/model"
	"github.com/hibakun/nova-store/models/response"
)

func CreateGenre(c *fiber.Ctx) error {
	genre := new(model.Genre)

	if err := c.BodyParser(genre); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "check your input",
			"data":    err,
		})
	}

	validate := validator.New()
	if err := validate.Struct(genre); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&genre).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to create game",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "success create genre",
		"data":    genre,
	})
}

func GetAllGenres(c *fiber.Ctx) error {
	var genres []response.GenreResponse

	database.DB.Find(&genres)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "genres found",
		"data":    genres,
	})
}

func GetGenreById(c *fiber.Ctx) error {
	var genre response.GenreResponse

	id := c.Params("id")
	if err := database.DB.First(&genre, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "game not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "genre found",
		"data":    genre,
	})
}
