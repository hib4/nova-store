package response

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Avatar      string `json:"avatar"`
}

func (UserResponse) TableName() string {
	return "users"
}
