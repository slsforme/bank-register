package domain

import "time"

type Transaction struct {
	ID         int
	SenderID   int
	ReceiverID int
	Amount     int64
	Status     string
	CreatedAt  time.Time
}
