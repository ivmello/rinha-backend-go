package core

import "time"

type Operation string

const (
	Debit  Operation = "d"
	Credit Operation = "c"
)

type Transaction struct {
	ID          int
	AccountID   int
	Amount      int
	Operation   Operation
	Description string
	CreatedAt   time.Time
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return ErrInvalidTransaction
	}
	if t.Operation != Debit && t.Operation != Credit {
		return ErrInvalidTransaction
	}
	if len(t.Description) < 1 || len(t.Description) > 10 {
		return ErrInvalidTransaction
	}
	return nil
}
