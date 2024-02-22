package api

import (
	"net/http"
	"strconv"

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
	id, _ := strconv.Atoi(c.Params("id"))
	output, err := h.service.GetBalance(id)
	if err != nil {
		if err == core.ErrAccountNotFound {
			return c.Status(http.StatusNotFound).JSON(err.Error())
		}
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(output)
}

func (h *handler) CreateTransaction(c *fiber.Ctx) error {
	var input core.CreateTransactionInput
	accountID, _ := strconv.Atoi(c.Params("id"))
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	output, err := h.service.CreateTransaction(accountID, input)
	if err != nil {
		if err == core.ErrAccountNotFound {
			return c.Status(http.StatusNotFound).JSON(err.Error())
		}
		if err == core.ErrInsufficientFunds || err == core.ErrInvalidOperation {
			return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
		}
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(output)
}
