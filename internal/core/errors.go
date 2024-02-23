package core

import "errors"

var (
	ErrInvalidTransaction = errors.New("INVALID_TRANSACTION")
	ErrAccountNotFound    = errors.New("ACCOUNT_NOT_FOUND")
	ErrInsufficientFunds  = errors.New("INSUFFICIENT_FUNDS")
	ErrInvalidOperation   = errors.New("INVALID_OPERATION")
)
