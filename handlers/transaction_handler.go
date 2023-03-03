package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models/model"
	"github.com/hibakun/nova-store/models/response"
	"github.com/hibakun/nova-store/utils"
)

func CreateTransaction(c *fiber.Ctx) error {
	transaction := new(model.Transaction)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "check your input",
			"data":    err,
		})
	}

	validate := validator.New()
	if err := validate.Struct(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	_, errGame := utils.CheckGameExist(transaction.GameID)
	if errGame != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "game doesn't exist",
		})
	}

	item, errItem := utils.CheckItemExist(transaction.ItemID)
	if errItem != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "item doesn't exist",
		})
	}

	_, errPayment := utils.CheckPaymentExist(transaction.PaymentID)
	if errPayment != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "payment doesn't exist",
		})
	}

	uuid := utils.GenerateUUID()
	if check := utils.CheckUUID(uuid); check {
		for next := true; next; next = utils.CheckUUID(uuid) {
			uuid = utils.GenerateUUID()
		}
	}

	userId := c.Locals("id").(float64)
	transaction.UserID = uint(userId)
	transaction.UUID = uuid
	transaction.Total = item.Price * transaction.Amount

	if err := database.DB.Create(&transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to create transaction",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "success create transaction",
		"data":    transaction,
	})
}

func GetAllTransactions(c *fiber.Ctx) error {
	var transactions []response.TransactionResponse

	database.DB.Preload("User").Preload("Game").Preload("Item").Preload("Payment").Find(&transactions)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "transactions found",
		"data":    transactions,
	})
}

func GetTransactionByUuid(c *fiber.Ctx) error {
	var transaction response.TransactionResponse

	id := c.Params("id")
	if err := database.DB.Preload("User").Preload("Game").Preload("Item").Preload("Payment").First(&transaction, "uuid = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "transaction not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "transaction found",
		"data":    transaction,
	})
}

func GetTransactionByUserId(c *fiber.Ctx) error {
	var transactions []response.TransactionResponse

	id := c.Locals("id")
	if err := database.DB.Preload("User").Preload("Game").Preload("Item").Preload("Payment").Find(&transactions, "user_id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "transaction not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "transactions found",
		"data":    transactions,
	})
}
