package domain

import (
	"errors"
	"time"
)

type User struct {
	ID        int
	OwnerName string
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Validate() error {
	if u.Balance < 0 {
		return errors.New("balance cannot be below 0")
	}

	if !ownerNameRegex.MatchString(u.OwnerName) {
		return errors.New("owner name should be shorter than 255 symbols, longer than 1 symbol and contain only latin")
	}

	return nil
}

func (u *User) Withdraw(amount int64) error {
	if amount < 0 {
		return errors.New("amount cannot be below 0")
	}

	if u.Balance < amount {
		return errors.New("insufficient funds")
	}

	u.Balance -= amount

	return nil
}

func (u *User) Deposit(amount int64) error {
	if amount < 0 {
		return errors.New("amount cannot be below 0")
	}

	u.Balance += amount

	return nil
}
