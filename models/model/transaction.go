package model

import "github.com/hibakun/nova-store/models"

type Transaction struct {
	models.Model
	UUID        string `json:"uuid" form:"-" gorm:"not null;unique"`
	UserID      uint   `json:"user_id" form:"-" gorm:"not null"`
	GameID      uint   `json:"game_id" form:"game_id" gorm:"not null" validate:"required"`
	ItemID      uint   `json:"item_id" form:"item_id" gorm:"not null" validate:"required"`
	PlayerID    string `json:"player_id" form:"player_id" gorm:"not null" validate:"required"`
	ZoneID      string `json:"zone_id" form:"zone_id"`
	Amount      int    `json:"amount" form:"amount" gorm:"not null" validate:"required"`
	Total       int    `json:"total" form:"-" gorm:"not null"`
	PaymentID   uint   `json:"payment_id" form:"payment_id" gorm:"not null" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" gorm:"not null" validate:"required"`
}
