package usecase

import (
	"bankapp/internal/domain"
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Save(ctx context.Context, u *domain.User) error
	Update(ctx context.Context, u *domain.User) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) (*[]domain.User, error)
	Transfer(ctx context.Context, senderID, receiverID int, amount int64) error
}

type TransactionRepository interface {
	GetAllByUserID(ctx context.Context, userID int) (*[]domain.Transaction, error)
}
