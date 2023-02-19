package models

type Game struct {
	Model
	Name       string `json:"name" form:"name" gorm:"not null" validate:"required"`
	Developer  string `json:"developer" form:"developer" gorm:"not null" validate:"required"`
	Publisher  string `json:"publisher" form:"publisher" gorm:"not null" validate:"required"`
	GenreID    []uint `json:"genre_id" form:"genre_id" gorm:"-" validate:"required"`
	PlatformID []uint `json:"platform_id" form:"platform_id" gorm:"-" validate:"required"`
	Image      string `json:"image" form:"image" gorm:"not null" validate:"required"`
}
