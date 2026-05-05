package service

import (
	"errors"
	"finance-api/internal/model"
	"finance-api/internal/repository"
)

type TransactionService struct {
	transactionsRepo *repository.TransactionRepository
}

func NewTransactionService(transactionsRepo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{transactionsRepo: transactionsRepo}
}

func (s *TransactionService) CreateTransaction(userID int, req model.TransactionRequest) error {
	if req.Type != "income" && req.Type != "expense" {
		return errors.New("tipo de despesa inválido")
	}

	if req.Amount <= 0 {
		return errors.New("valor deve ser maior que zero")
	}

	transactions := &model.Transaction{
		UserID:      userID,
		Type:        req.Type,
		Amount:      req.Amount,
		Description: req.Description,
		Category:    req.Category,
	}
	return s.transactionsRepo.CreateTransation(transactions)
}

func (s *TransactionService) GetTransactions(userID int) ([]model.Transaction, error) {
	return s.transactionsRepo.GetTransactionsByUserID(userID)
}

func (s *TransactionService) DeleteTransaction(transactionID, userID int) error {
	return s.transactionsRepo.DeleteTransaction(transactionID, userID)
}
