package model

import (
	"github.com/hibakun/nova-store/models"
	"github.com/hibakun/nova-store/models/response"
)

type Game struct {
	models.Model
	Name      string                   `json:"name" form:"name" gorm:"not null" validate:"required"`
	Developer string                   `json:"developer" form:"developer" gorm:"not null" validate:"required"`
	Publisher string                   `json:"publisher" form:"publisher" gorm:"not null" validate:"required"`
	GenreID   []uint                   `json:"genre_id" form:"genre_id" gorm:"-" validate:"required"`
	Genres    []response.GenreResponse `json:"-" gorm:"many2many:game_genres;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GenreID"`
	Platform  string                   `json:"platform" form:"platform" gorm:"not null" validate:"required"`
	Image     string                   `json:"image" form:"image" gorm:"not null"`
}
