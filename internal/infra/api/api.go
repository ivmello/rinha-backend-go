package api

import (
	"log"

	"rinha-backend-go/internal/core"
	"rinha-backend-go/internal/infra/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(conn *pgxpool.Pool) {
	app := fiber.New()
	accountRepository := database.NewAccountRepository(conn)
	transactionRepository := database.NewTransactionRepository(conn)
	service := core.NewService(accountRepository, transactionRepository)
	handler := NewHandler(service)
	app.Get("/clientes/:id/extrato", handler.GetBalance)
	app.Post("/clientes/:id/transacoes", handler.CreateTransaction)
	log.Fatal(app.Listen(":8080"))
}