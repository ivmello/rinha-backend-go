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

func (r *accountRepository) GetByID(ctx context.Context, id int) ([]byte, error) {
	var json_response []byte
	err := r.db.QueryRow(ctx, "SELECT * FROM get_balance($1);", id).Scan(&json_response)
	if err != nil {
		return nil, err
	}
	return json_response, nil
}
