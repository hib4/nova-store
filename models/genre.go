package models

type Genre struct {
	Model
	Name  string         `json:"name" form:"name" gorm:"not null" validate:"required"`
	Games []GameResponse `json:"-" gorm:"many2many:game_genres;ForeignKey:ID;joinForeignKey:GenreID;References:ID;joinReferences:GameID"`
}

type GenreResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (GenreResponse) TableName() string {
	return "genres"
}
