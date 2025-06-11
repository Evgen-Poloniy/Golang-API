package repository

import (
	"database/sql"
	"fmt"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) MakeTransaction(senderID uint32, recipientID uint32, amount float64) error {
	tx, err := r.db.BeginTx(nil, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		"UPDATE users SET coins = coins - $1 WHERE user_id = $2",
		amount, senderID,
	)
	if err != nil {
		return fmt.Errorf("failed to deduct coins: %v", err)
	}

	_, err = tx.Exec(
		"UPDATE users SET coins = coins + $1 WHERE user_id = $2",
		amount, recipientID,
	)
	if err != nil {
		return fmt.Errorf("failed to add coins: %w", err)
	}

	return tx.Commit()
}
