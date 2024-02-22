package core

type AccountRepository interface {
	GetByID(id int) (*Account, error)
}

type TransactionRepository interface {
	GetByAccountID(accountID int) ([]Transaction, error)
	Create(transaction *Transaction) (*Account, error)
}
