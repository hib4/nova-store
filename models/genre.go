package models

type Genre struct {
	Model
	Name string `json:"name" form:"name" gorm:"not null" validate:"required"`
}

type GenreResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (GenreResponse) TableName() string {
	return "genres"
}
