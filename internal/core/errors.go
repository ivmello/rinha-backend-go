package core

import "errors"

var (
	ErrInvalidTransaction = errors.New("invalid transaction")
	ErrAccountNotFound    = errors.New("account not found")
	ErrInsufficientFunds  = errors.New("insufficient funds")
	ErrInvalidOperation   = errors.New("invalid operation")
)
