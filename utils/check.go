package utils

import (
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models/model"
)

func CheckGameExist(id uint) (*model.Game, error) {
	var model model.Game

	if err := database.DB.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func CheckItemExist(id uint) (*model.Item, error) {
	var model model.Item

	if err := database.DB.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func CheckPaymentExist(id uint) (*model.Payment, error) {
	var model model.Payment

	if err := database.DB.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func CheckUUID(uuid string) bool {
	var transaction model.Transaction

	if err := database.DB.First(&transaction, "uuid = ?", uuid).Error; err != nil {
		return false
	}

	return true
}
