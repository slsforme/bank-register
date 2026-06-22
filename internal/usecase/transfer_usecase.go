package usecase

import (
	"context"
	"errors"
)

type TransferUseCase struct {
	transactionRepo TransactionRepository
	userRepo        UserRepository
}

func (tc *TransferUseCase) Withdraw(ctx context.Context, amount int64, userID int) (int64, error) {
	u, err := tc.userRepo.GetByID(ctx, userID)

	if err != nil {
		return 0, err
	}

	if err := u.Withdraw(amount); err != nil {
		return 0, err
	}

	return u.Balance, tc.userRepo.Update(ctx, u)
}

func (tc *TransferUseCase) Deposit(ctx context.Context, amount int64, userID int) (int64, error) {
	u, err := tc.userRepo.GetByID(ctx, userID)

	if err != nil {
		return 0, err
	}

	if err := u.Deposit(amount); err != nil {
		return 0, err
	}

	return u.Balance, tc.userRepo.Update(ctx, u)
}

func (tc *TransferUseCase) Transfer(ctx context.Context, senderID, receiverID int, amount int64) error {
	if amount <= 0 {
		return errors.New("amount cannot be below 0")
	}

	_, err := tc.userRepo.GetByID(ctx, senderID)
	if err != nil {
		return errors.New("sender not found")
	}

	_, err = tc.userRepo.GetByID(ctx, receiverID)
	if err != nil {
		return errors.New("receiver not found")
	}

	return tc.userRepo.Transfer(ctx, senderID, receiverID, amount)
}
