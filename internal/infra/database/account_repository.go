package database

import (
	"context"

	"rinha-backend-go/internal/core"

	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	db *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) core.AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) GetByID(ctx context.Context, id int) (core.Account, error) {
	var account core.Account
	err := r.db.QueryRow(ctx, "SELECT account_limit, balance FROM accounts WHERE id = $1;", id).Scan(&account.Limit, &account.Balance)
	if err != nil {
		return core.Account{}, err
	}
	transactions := make([]core.Transaction, 0)
	rows, err := r.db.Query(ctx, "SELECT amount, operation, description, created_at FROM transactions WHERE account_id = $1 ORDER BY created_at DESC LIMIT 10;", id)
	if err != nil {
		return core.Account{}, err
	}
	for rows.Next() {
		var transaction core.Transaction
		err := rows.Scan(&transaction.Amount, &transaction.Operation, &transaction.Description, &transaction.CreatedAt)
		if err != nil {
			return core.Account{}, err
		}
		transactions = append(transactions, transaction)
	}
	account.Transactions = transactions
	return account, nil
}
