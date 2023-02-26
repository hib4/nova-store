package response

type ItemResponse struct {
	ID     uint   `json:"id"`
	GameID uint   `json:"-"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
}

type ItemGameResponse struct {
	ID     uint             `json:"id"`
	GameID uint             `json:"-"`
	Game   GameNameResponse `json:"game"`
	Name   string           `json:"name"`
	Price  int              `json:"price"`
	Image  string           `json:"image"`
}

func (ItemResponse) TableName() string {
	return "items"
}

func (ItemGameResponse) TableName() string {
	return "items"
}
