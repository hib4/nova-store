package response

type ItemResponse struct {
	GameID uint   `json:"-"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
}

func (ItemResponse) TableName() string {
	return "items"
}
