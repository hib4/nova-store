package response

type PaymentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (PaymentResponse) TableName() string {
	return "payments"
}
