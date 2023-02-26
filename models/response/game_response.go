package response

type GameResponse struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	Developer string          `json:"developer"`
	Publisher string          `json:"publisher"`
	Genres    []GenreResponse `json:"genres" gorm:"many2many:game_genres;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GenreID"`
	Platform  string          `json:"platform"`
	Image     string          `json:"image"`
	Items     []ItemResponse  `json:"items" gorm:"foreignKey:GameID"`
}

func (GameResponse) TableName() string {
	return "games"
}
