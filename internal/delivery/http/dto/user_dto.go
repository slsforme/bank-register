package dto

type CreateUserRequest struct {
	OwnerName string `json:"owner_name" binding:"required,min=2,max=255"`
	Balance   int64  `json:"balance" binding:"required, min=0"`
}

type UpdateUserRequest struct {
	OwnerName *string `json:"owner_name,omitempty" binding:"omitempty,min=2,max=255"`
	Balance   *int64  `json:"balance,omitempty" binding:"omitempty,min=0"`
}

type DeleteUserRequest struct {
	ID int `json:"id" binding:"omitempty,min=1"`
}

type GetUserByIDRequest struct {
	ID int `json:"id" binding:"required,min=1"`
}

type GetAllUsersRequest struct {
	Limit  int `form:"limit,omitempty" binding:"omitempty,min=1,max=1000"`
	Offset int `form:"offset,omitempty" binding:"omitempty,min=0"`
}

type TransferRequest struct {
	SenderID   int   `json:"sender_id" binding:"required,min=1"`
	ReceiverID int   `json:"receiver_id" binding:"required,min=1"`
	Amount     int64 `json:"amount" binding:"required"`
}
