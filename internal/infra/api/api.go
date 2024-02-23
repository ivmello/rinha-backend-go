package api

import (
	"log"
	"os"

	"rinha-backend-go/internal/core"
	"rinha-backend-go/internal/infra/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(conn *pgxpool.Pool) {
	accountRepository := database.NewAccountRepository(conn)
	transactionRepository := database.NewTransactionRepository(conn)
	service := core.NewService(accountRepository, transactionRepository)
	handler := NewHandler(service)
	app := fiber.New()
	app.Get("/clientes/:id/extrato", handler.GetBalance)
	app.Post("/clientes/:id/transacoes", handler.CreateTransaction)
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
