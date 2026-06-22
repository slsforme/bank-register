package usecase

import (
	"bankapp/internal/domain"
	"context"
	"errors"
)

type AccountUseCase struct {
	repo   UserRepository
	txRepo TransactionRepository
}

func (ac *AccountUseCase) CreateAccount(ctx context.Context, ownerName string, balance int64) (*domain.User, error) {
	u := &domain.User{OwnerName: ownerName, Balance: balance}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	return u, ac.repo.Save(ctx, u)
}

func (ac *AccountUseCase) GetAccount(ctx context.Context, userID int) (*domain.User, error) {
	u, err := ac.repo.GetByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return u, err
}

func (ac *AccountUseCase) UpdateAccount(ctx context.Context, userID int, ownerName *string, balance *int64) (*domain.User, error) {
	u, err := ac.repo.GetByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	if ownerName != nil {
		u.OwnerName = *ownerName
	}

	if balance != nil {
		u.Balance = *balance
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	return u, ac.repo.Update(ctx, u)
}

func (ac *AccountUseCase) DeleteAccount(ctx context.Context, userID int) (*domain.User, error) {
	u, err := ac.repo.GetByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	if u.Balance > 0 {
		return nil, errors.New("cannot delete account with balance above 0")
	}

	return u, ac.repo.Delete(ctx, userID)
}

func (ac *AccountUseCase) GetHistory(ctx context.Context, userID int) (*[]domain.Transaction, error) {
	txs, err := ac.txRepo.GetAllByUserID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return txs, nil
}
