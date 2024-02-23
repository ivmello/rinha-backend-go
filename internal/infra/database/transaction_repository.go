package database

import (
	"context"

	"rinha-backend-go/internal/core"

	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) core.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Create(ctx context.Context, transaction core.Transaction) (core.Account, error) {
	amount := transaction.Amount
	if transaction.Operation == core.Debit {
		amount = amount * -1
	}
	var updatedBalance, updatedLimit *int
	err := r.db.QueryRow(ctx, "CALL create_transaction($1, $2, $3, $4)", transaction.AccountID, amount, transaction.Operation, transaction.Description).Scan(&updatedBalance, &updatedLimit)
	if err != nil {
		return core.Account{}, err
	}
	if updatedBalance == nil {
		return core.Account{}, core.ErrInvalidTransaction
	}
	return core.Account{
		Balance: *updatedBalance,
		Limit:   *updatedLimit,
	}, nil
}
