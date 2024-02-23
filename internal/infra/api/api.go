package api

import (
	"net/http"

	"rinha-backend-go/internal/core"
	"rinha-backend-go/internal/infra/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(conn *pgxpool.Pool) {
	accountRepository := database.NewAccountRepository(conn)
	transactionRepository := database.NewTransactionRepository(conn)
	service := core.NewService(accountRepository, transactionRepository)
	handler := NewHandler(service)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /clientes/{id}/extrato", handler.GetBalance)
	mux.HandleFunc("POST /clientes/{id}/transacoes", handler.CreateTransaction)
	http.ListenAndServe(":8080", mux)
}
