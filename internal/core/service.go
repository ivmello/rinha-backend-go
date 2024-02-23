package core

import (
	"context"
	"time"
)

type Service interface {
	GetBalance(ctx context.Context, id int) (GetBalanceOutput, error)
	CreateTransaction(ctx context.Context, accountID int, input CreateTransactionInput) (CreateTransactionOutput, error)
}

type service struct {
	accountRepository     AccountRepository
	transactionRepository TransactionRepository
}

func NewService(accountRepository AccountRepository, transactionRepository TransactionRepository) Service {
	return &service{accountRepository, transactionRepository}
}

func (s *service) GetBalance(ctx context.Context, id int) (GetBalanceOutput, error) {
	account, err := s.accountRepository.GetByID(ctx, id)
	if err != nil {
		return GetBalanceOutput{}, err
	}
	transactionsOutput := make([]TransactionOutput, 0)
	for _, transaction := range account.Transactions {
		transactionsOutput = append(transactionsOutput, TransactionOutput{
			Amount:      transaction.Amount,
			Operation:   string(transaction.Operation),
			Description: transaction.Description,
			Date:        transaction.CreatedAt.Format(time.RFC3339),
		})
	}
	output := GetBalanceOutput{
		Balance: Balance{
			Total: account.Balance,
			Date:  account.Date.UTC().Format(time.RFC3339),
			Limit: account.Limit,
		},
		Transactions: transactionsOutput,
	}
	return output, nil
}

func (s *service) CreateTransaction(ctx context.Context, accountID int, input CreateTransactionInput) (CreateTransactionOutput, error) {
	transaction := Transaction{
		AccountID:   accountID,
		Amount:      input.Amount,
		Operation:   Operation(input.Operation),
		Description: input.Description,
	}
	accountResponse, err := s.transactionRepository.Create(ctx, transaction)
	if err != nil {
		return CreateTransactionOutput{}, ErrInvalidTransaction
	}
	output := CreateTransactionOutput{
		Balance: accountResponse.Balance,
		Limit:   accountResponse.Limit,
	}
	return output, nil
}
