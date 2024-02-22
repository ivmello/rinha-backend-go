package database

import (
	"database/sql"

	"rinha-backend-go/internal/core"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) core.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Create(transaction *core.Transaction) (*core.Account, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var accountResponse core.Account
	rows := tx.QueryRow("SELECT account_limit, balance FROM accounts WHERE id = $1 FOR UPDATE;", transaction.AccountID)
	err = rows.Scan(&accountResponse.Limit, &accountResponse.Balance)
	if err == sql.ErrNoRows {
		tx.Rollback()
		return nil, core.ErrAccountNotFound
	}
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var newBalance int
	if transaction.Operation == "d" {
		newBalance = accountResponse.Balance - transaction.Amount
	}
	if transaction.Operation == "c" {
		newBalance = accountResponse.Balance + transaction.Amount
	}
	if newBalance < (accountResponse.Limit * -1) {
		tx.Rollback()
		return nil, core.ErrInsufficientFunds
	}
	_, err = tx.Exec("INSERT INTO transactions (account_id, amount, operation, description) VALUES ($1, $2, $3, $4);", transaction.AccountID, transaction.Amount, transaction.Operation, transaction.Description)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = tx.Exec("UPDATE accounts SET balance = $1 WHERE id = $2;", newBalance, transaction.AccountID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &core.Account{
		Limit:   accountResponse.Limit,
		Balance: newBalance,
	}, nil
}

func (r *transactionRepository) GetByAccountID(accountID int) ([]core.Transaction, error) {
	rows, err := r.db.Query("SELECT amount, operation, description, created_at FROM transactions WHERE account_id = $1 ORDER BY created_at DESC LIMIT 10;", accountID)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, core.ErrAccountNotFound
	}
	var transactions []core.Transaction
	for rows.Next() {
		var transaction core.Transaction
		err = rows.Scan(&transaction.Amount, &transaction.Operation, &transaction.Description, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
