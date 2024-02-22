package database

import (
	"database/sql"
	"time"

	"rinha-backend-go/internal/core"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) core.AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) GetByID(id int) (*core.Account, error) {
	rows := r.db.QueryRow("SELECT account_limit, balance FROM accounts WHERE id = $1", id)
	var accountLimit int
	var balance int
	err := rows.Scan(&accountLimit, &balance)
	if err == sql.ErrNoRows {
		return nil, core.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}
	return &core.Account{
		ID:      id,
		Date:    time.Now().Format(time.RFC3339Nano),
		Limit:   accountLimit,
		Balance: balance,
	}, nil
}
