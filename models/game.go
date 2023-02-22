package models

type Game struct {
	Model
	Name      string          `json:"name" form:"name" gorm:"not null" validate:"required"`
	Developer string          `json:"developer" form:"developer" gorm:"not null" validate:"required"`
	Publisher string          `json:"publisher" form:"publisher" gorm:"not null" validate:"required"`
	GenreID   []uint          `json:"genre_id" form:"genre_id" gorm:"-" validate:"required"`
	Genres    []GenreResponse `json:"-" gorm:"foreignKey:ID"`
	Platform  string          `json:"platform" form:"platform" gorm:"not null" validate:"required"`
	Image     string          `json:"image" form:"image" gorm:"not null"`
}

type GameResponse struct {
	ID        uint            `json:"id" form:"id"`
	Name      string          `json:"name" form:"name"`
	Developer string          `json:"developer" form:"developer"`
	Publisher string          `json:"publisher" form:"publisher"`
	Genres    []GenreResponse `json:"genres" form:"genres" gorm:"foreignKey:ID"`
	Platform  string          `json:"platform" form:"platform"`
	Image     string          `json:"image" form:"image"`
}

func (GameResponse) TableName() string {
	return "games"
}
