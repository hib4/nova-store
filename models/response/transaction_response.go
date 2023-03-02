package response

import "github.com/hibakun/nova-store/models"

type TransactionResponse struct {
	models.Model
	UUID        string           `json:"uuid"`
	UserID      uint             `json:"user_id"`
	User        UserResponse     `json:"user"`
	GameID      uint             `json:"-"`
	Game        GameNameResponse `json:"game"`
	ItemID      uint             `json:"-"`
	Item        ItemResponse     `json:"item"`
	PlayerID    string           `json:"player_id"`
	ZoneID      string           `json:"zone_id"`
	Amount      int              `json:"amount"`
	Total       int              `json:"total"`
	PaymentID   uint             `json:"-"`
	Payment     PaymentResponse  `json:"payment"`
	PhoneNumber string           `json:"phone_number"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
