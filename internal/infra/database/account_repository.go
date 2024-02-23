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

func (r *accountRepository) GetBalance(ctx context.Context, id int) (map[string]interface{}, error) {
	var jsonResult map[string]interface{}
	err := r.db.QueryRow(ctx, "SELECT * FROM get_balance($1);", id).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}
	return jsonResult, nil
}
