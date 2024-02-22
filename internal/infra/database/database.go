package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	db.SetMaxOpenConns(200)
	return db, nil
}
