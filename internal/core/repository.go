package core

import "context"

type AccountRepository interface {
	GetByID(ctx context.Context, id int) ([]byte, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction Transaction) (Account, error)
}
