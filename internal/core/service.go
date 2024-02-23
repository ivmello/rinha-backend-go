package core

import (
	"context"
)

type Service interface {
	GetBalance(ctx context.Context, id int) (map[string]interface{}, error)
	CreateTransaction(ctx context.Context, accountID int, input CreateTransactionInput) (map[string]interface{}, error)
}

type service struct {
	accountRepository     AccountRepository
	transactionRepository TransactionRepository
}

func NewService(accountRepository AccountRepository, transactionRepository TransactionRepository) Service {
	return &service{accountRepository, transactionRepository}
}

func (s *service) GetBalance(ctx context.Context, id int) (map[string]interface{}, error) {
	output, err := s.accountRepository.GetBalance(ctx, id)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *service) CreateTransaction(ctx context.Context, id int, input CreateTransactionInput) (map[string]interface{}, error) {
	accountResponse, err := s.transactionRepository.Create(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return accountResponse, nil
}
