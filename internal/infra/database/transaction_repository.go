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

func (r *transactionRepository) Create(ctx context.Context, accountID int, input core.CreateTransactionInput) (map[string]interface{}, error) {
	var json_result map[string]interface{}
	err := r.db.QueryRow(ctx, "SELECT * FROM create_transaction($1, $2, $3, $4)", accountID, input.Amount, input.Operation, input.Description).Scan(&json_result)
	if err != nil {
		return nil, core.ErrInvalidTransaction
	}
	return json_result, nil
}
