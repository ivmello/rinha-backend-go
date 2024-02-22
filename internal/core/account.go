package core

import "time"

type Account struct {
	ID           int
	Limit        int
	Balance      int
	Transactions []Transaction
	Date         time.Time
}
