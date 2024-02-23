package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"rinha-backend-go/internal/core"
)

type Handler interface {
	GetBalance(w http.ResponseWriter, r *http.Request)
	CreateTransaction(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service core.Service
}

func NewHandler(service core.Service) Handler {
	return &handler{service}
}

func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if id < 1 || id > 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	output, err := h.service.GetBalance(ctx, id)
	if err != nil {
		if err == core.ErrAccountNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&output)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(response))
}

func (h *handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if id < 1 || id > 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var input core.CreateTransactionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if input.Amount < 1 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if input.Operation != string(core.Debit) && input.Operation != string(core.Credit) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if len(input.Description) < 1 || len(input.Description) > 10 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	output, err := h.service.CreateTransaction(ctx, id, input)
	if err != nil {
		if err == core.ErrAccountNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err == core.ErrInsufficientFunds || err == core.ErrInvalidOperation || err == core.ErrInvalidTransaction {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(response))
}
