package web

type CreateTransactionInput struct {
	Amount   int `json:"amount" binding:"required"`
	CourseId int `json:"course_id" binding:"required"`
	User     UserResponse
}

type MidtransTransactionResponse struct {
	Id         string `json:"id"`
	UserId     int    `json:"user_id"`
	CourseId   int    `json:"course_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	PaymentUrl string `json:"payment_url"`
}

type TransactionNotificationFromMidtrans struct {
	OrderId           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
}
