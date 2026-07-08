package dto

type GetAllTransactionsRequest struct {
	Limit  int `form:"limit" binding:"omitempty,min=1,max=1000"`
	Offset int `form:"offset" binding:"omitempty,min=0"`
}

type GetTransactionsByUserIDRequest struct {
	ID int `json:"id" binding:"required"`
}

type CreateTransaction struct {
	SenderID   int    `json:"sender_id" binding:"required"`
	ReceiverID int    `json:"receiver_id" binding:"required"`
	Amount     int64  `json:"amount" binding:"required"`
	Status     string `json:"status" binding:"required"`
}
