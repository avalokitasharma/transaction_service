package service

import (
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/models"
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(t *models.Transactions) error {
	return s.repo.Create(t)
}

func (s *TransactionService) GetTransaction(id int64) (*models.Transactions, error) {
	t, err := s.repo.Get(id)
	return t, err
}

func (s *TransactionService) GetTransactionsByType(transactionType string) ([]int64, error) {
	ids, err := s.repo.GetByType(transactionType)
	return ids, err
}
func (s *TransactionService) GetTransactionSum(id int64) (float64, error) {
	return s.repo.GetSum(id)
}
