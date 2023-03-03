package model

import "github.com/hibakun/nova-store/models"

type Item struct {
	models.Model
	GameID uint   `json:"game_id" form:"game_id" gorm:"not null" validate:"required"`
	Name   string `json:"name" form:"name" gorm:"not null" validate:"required"`
	Price  int    `json:"price" form:"price" gorm:"not null" validate:"required"`
	Image  string `json:"image" form:"image" gorm:"not null"`
}
