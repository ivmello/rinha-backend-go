package core

import "time"

type Service interface {
	GetBalance(id int) (*GetBalanceOutput, error)
	CreateTransaction(accountID int, input CreateTransactionInput) (*CreateTransactionOutput, error)
}

type service struct {
	accountRepository     AccountRepository
	transactionRepository TransactionRepository
}

func NewService(accountRepository AccountRepository, transactionRepository TransactionRepository) Service {
	return &service{accountRepository, transactionRepository}
}

func (s *service) GetBalance(id int) (*GetBalanceOutput, error) {
	account, err := s.accountRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	transactions, err := s.transactionRepository.GetByAccountID(id)
	if err != nil {
		return nil, err
	}
	var transactionsOutput []TransactionOutput
	for _, transaction := range transactions {
		transactionsOutput = append(transactionsOutput, TransactionOutput{
			Amount:      transaction.Amount,
			Operation:   transaction.Operation,
			Description: transaction.Description,
			Date:        transaction.CreatedAt.Format(time.RFC3339Nano),
		})
	}
	output := &GetBalanceOutput{
		Balance: Balance{
			Total: account.Balance,
			Date:  account.Date,
			Limit: account.Limit,
		},
		Transactions: transactionsOutput,
	}
	return output, nil
}

func (s *service) CreateTransaction(accountID int, input CreateTransactionInput) (*CreateTransactionOutput, error) {
	if input.Operation != "d" && input.Operation != "c" {
		return nil, ErrInvalidOperation
	}
	transaction := &Transaction{
		AccountID:   accountID,
		Amount:      input.Amount,
		Operation:   input.Operation,
		Description: input.Description,
	}
	accountResponse, err := s.transactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}
	output := &CreateTransactionOutput{
		Balance: accountResponse.Balance,
		Limit:   accountResponse.Limit,
	}
	return output, nil
}
