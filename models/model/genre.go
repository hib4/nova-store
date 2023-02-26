package model

import "github.com/hibakun/nova-store/models"

type Genre struct {
	models.Model
	Name string `json:"name" form:"name" gorm:"not null" validate:"required"`
}
