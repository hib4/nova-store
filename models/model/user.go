package model

import "github.com/hibakun/nova-store/models"

type User struct {
	models.Model
	Name        string `json:"name" form:"name" gorm:"not null" validate:"required"`
	Avatar      string `json:"avatar" form:"avatar" gorm:"not null"`
	Email       string `json:"email" form:"email" gorm:"unique;not null" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" gorm:"unique;not null" validate:"required,min=8"`
	Password    string `json:"-" form:"password" gorm:"not null" validate:"required,min=6"`
}
