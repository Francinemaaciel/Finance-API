package repository

import (
	"finance-api/internal/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransation(tr *model.Transaction) error {
	query := `INSERT INTO transactions (user_id, description, amount, type, category, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(
		query,
		tr.UserID,
		tr.Description,
		tr.Amount,
		tr.Type,
		tr.Category,
		tr.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) GetTransactionsByUserID(userID int) ([]model.Transaction, error) {
	var transactions []model.Transaction
	query := `SELECT id, user_id, description, amount, type, category, created_at 
	FROM transactions WHERE user_id = $1`

	err := r.db.Select(&transactions, query, userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) DeleteTransaction(transactionID, userID int) error {
	query := `DELETE FROM transactions WHERE id = $1 AND user_id = $2`

	_, err := r.db.Exec(query, transactionID, userID)
	if err != nil {
		return err
	}
	return nil
}
