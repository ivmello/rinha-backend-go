package core

import "time"

type Transaction struct {
	ID          int
	AccountID   int
	Amount      int
	Operation   string
	Description string
	CreatedAt   time.Time
}
