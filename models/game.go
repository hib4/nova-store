package models

type Game struct {
	Model
	Name      string          `json:"name" form:"name" gorm:"not null" validate:"required"`
	Developer string          `json:"developer" form:"developer" gorm:"not null" validate:"required"`
	Publisher string          `json:"publisher" form:"publisher" gorm:"not null" validate:"required"`
	GenreID   []uint          `json:"genre_id" form:"genre_id" gorm:"-" validate:"required"`
	Genres    []GenreResponse `json:"-" gorm:"many2many:game_genres;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GenreID"`
	Platform  string          `json:"platform" form:"platform" gorm:"not null" validate:"required"`
	Image     string          `json:"image" form:"image" gorm:"not null"`
}

type GameResponse struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	Developer string          `json:"developer"`
	Publisher string          `json:"publisher"`
	Genres    []GenreResponse `json:"genres" gorm:"many2many:game_genres;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GenreID"`
	Platform  string          `json:"platform"`
	Image     string          `json:"image"`
}

func (GameResponse) TableName() string {
	return "games"
}
