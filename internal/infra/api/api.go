package api

import (
	"database/sql"
	"log"

	"rinha-backend-go/internal/core"
	"rinha-backend-go/internal/infra/database"

	"github.com/gofiber/fiber/v2"
)

func Run(conn *sql.DB) {
	app := fiber.New()
	accountRepository := database.NewAccountRepository(conn)
	transactionRepository := database.NewTransactionRepository(conn)
	service := core.NewService(accountRepository, transactionRepository)
	handler := NewHandler(service)
	app.Get("/clientes/:id/extrato", handler.GetBalance)
	app.Post("/clientes/:id/transacoes", handler.CreateTransaction)
	log.Fatal(app.Listen(":8080"))
}
