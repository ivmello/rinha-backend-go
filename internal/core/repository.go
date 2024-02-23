package core

import "context"

type AccountRepository interface {
	GetBalance(ctx context.Context, id int) (map[string]interface{}, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, id int, input CreateTransactionInput) (map[string]interface{}, error)
}
