package repository

import (
	"database/sql"

	"github.com/avalokitasharma/transaction_service/transaction_service/internal/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionoRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(t *models.Transactions) error {
	query := `INSERT INTO transactions (id, amount, type, parent_id) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, t.ID, t.Amount, t.Type, t.ParentID)
	return err
}

func (r *TransactionRepository) Get(id int64) (*models.Transactions, error) {
	query := `SELECT id, amount, type, parent_id FROM transactions WHERE id=$1`
	t := &models.Transactions{}
	err := r.db.QueryRow(query, id).Scan(&t.ID, &t.Amount, &t.Type, &t.ParentID)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TransactionRepository) GetByType(transactionType string) ([]int64, error) {
	query := `SELECT id from transactions WHERE type = $1`
	rows, err := r.db.Query(query, transactionType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (r *TransactionRepository) GetSum(id int64) (float64, error) {
	query := `
	WITH RECURSIVE transaction_tree AS (
		SELECT id, amount, parent_id
		FROM transactions
		WHERE id = $1
		UNION ALL
		SELECT id, amount, parent_id
		FROM transactions t
		JOIN transaction_tree tt ON t.parent_id = tt.id
	)
	SELECT SUM(amount) FROM transaction_tree`

	var sum float64
	err := r.db.QueryRow(query, id).Scan(&sum)
	if err != nil {
		return 0, err
	}
	return sum, nil
}
