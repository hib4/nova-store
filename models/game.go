package models

type Game struct {
	Model
	Name      string          `json:"name" form:"name" gorm:"not null" validate:"required"`
	Developer string          `json:"developer" form:"developer" gorm:"not null" validate:"required"`
	Publisher string          `json:"publisher" form:"publisher" gorm:"not null" validate:"required"`
	GenreID   []uint          `json:"genre_id" form:"genre_id" gorm:"-" validate:"required"`
	Genres    []GenreResponse `json:"-" gorm:"many2many:game_genres;ForeignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GenreID"`
	Image     string          `json:"image" form:"image" gorm:"not null"`
}

type GameResponse struct {
	ID        uint            `json:"id" form:"id"`
	Name      string          `json:"name" form:"name"`
	Developer string          `json:"developer" form:"developer"`
	Publisher string          `json:"publisher" form:"publisher"`
	Genres    []GenreResponse `json:"genres" form:"genres" gorm:"many2many:game_genres;ForeignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GenreID"`
	Image     string          `json:"image" form:"image"`
}

func (GameResponse) TableName() string {
	return "games"
}
