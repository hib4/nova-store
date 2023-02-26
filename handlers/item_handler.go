package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models/model"
	"github.com/hibakun/nova-store/models/response"
	"github.com/hibakun/nova-store/utils"
	"strconv"
)

func CreateItem(c *fiber.Ctx) error {
	item := new(model.Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "check your input",
			"data":    err,
		})
	}

	validate := validator.New()
	if err := validate.Struct(item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	var gameID model.User
	if err := database.DB.Where("id = ?", item.GameID).Find(&gameID); err != nil {
		if err.RowsAffected < 1 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "game id not found",
			})
		}
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

	item.Image = fileName

	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to create item",
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
		"message": "success create item",
		"data":    item,
	})
}

func GetItemById(c *fiber.Ctx) error {
	var item response.ItemGameResponse

	id := c.Params("id")
	if err := database.DB.Preload("Game").First(&item, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "item not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "item found",
		"data":    item,
	})
}

func GetItemsByGameId(c *fiber.Ctx) error {
	var items []response.ItemGameResponse

	id := c.Params("id")
	if err := database.DB.Preload("Game").Find(&items, "game_id= ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "items not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "items found",
		"data":    items,
	})
}
