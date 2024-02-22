package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New() (*pgxpool.Pool, error) {
	connStr := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return pool, nil
}
