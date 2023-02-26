package models

type Item struct {
	Model
	GameID uint   `json:"game_id" form:"game_id" gorm:"not null" validate:"required"`
	Name   string `json:"name" form:"name" gorm:"not null" validate:"required"`
	Price  int    `json:"price" form:"price" gorm:"not null" validate:"required"`
	Image  string `json:"image" form:"image" gorm:"not null"`
}

type ItemResponse struct {
	GameID uint   `json:"-"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
}

func (ItemResponse) TableName() string {
	return "items"
}
