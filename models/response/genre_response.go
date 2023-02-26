package response

type GenreResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (GenreResponse) TableName() string {
	return "genres"
}
