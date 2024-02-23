package core

import (
	"context"
	"encoding/json"
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
	var output GetBalanceOutput
	if err := json.Unmarshal(account, &output); err != nil {
		return GetBalanceOutput{}, err
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
