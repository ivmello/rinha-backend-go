package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"rinha-backend-go/internal/core"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetBalance(c *fiber.Ctx) error
	CreateTransaction(c *fiber.Ctx) error
}

type handler struct {
	service core.Service
}

func NewHandler(service core.Service) Handler {
	return &handler{service}
}

func (h *handler) GetBalance(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	if id < 1 || id > 5 {
		return c.SendStatus(http.StatusNotFound)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	output, err := h.service.GetBalance(ctx, id)
	if err != nil {
		if err == core.ErrAccountNotFound {
			return c.SendStatus(http.StatusNotFound)
		}
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).JSON(output)
}

func (h *handler) CreateTransaction(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	if id < 1 || id > 5 {
		return c.SendStatus(http.StatusNotFound)
	}
	var input core.CreateTransactionInput
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	if input.Amount < 1 {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	if input.Operation != string(core.Debit) && input.Operation != string(core.Credit) {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	if len(input.Description) < 1 || len(input.Description) > 10 {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	output, err := h.service.CreateTransaction(ctx, id, input)
	if err != nil {
		if err == core.ErrAccountNotFound {
			return c.SendStatus(http.StatusNotFound)
		}
		if err == core.ErrInsufficientFunds || err == core.ErrInvalidOperation || err == core.ErrInvalidTransaction {
			return c.SendStatus(http.StatusUnprocessableEntity)
		}
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).JSON(output)
}
