package core

import "context"

type AccountRepository interface {
	GetByID(ctx context.Context, id int) (Account, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction Transaction) (Account, error)
}
