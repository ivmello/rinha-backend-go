package core

import "errors"

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidOperation  = errors.New("invalid operation")
)
