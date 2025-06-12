package repository

import (
	"context"
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
	tx, err := r.db.BeginTx(context.TODO(), &sql.TxOptions{
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
		return fmt.Errorf("failed to add coins: %v", err)
	}

	return tx.Commit()
}

func (r *TransactionRepository) CreateRecordOfTransaction(transaction *Transactions) (uint32, error) {
	var query string = "INSERT INTO Transactions (sender_id, recipient_id, amount) VALUES ($1, $2, $3) RETURNING transaction_id"
	var transaction_id uint32

	row := r.db.QueryRow(query, transaction.SenderID, transaction.RecipientID, transaction.Amount)
	if err := row.Scan(&transaction_id); err != nil {
		return 0, err
	}

	return transaction_id, nil
}

func (r *TransactionRepository) GetTransactionByID(transaction_id uint32) (*Transactions, error) {
	var query string = "SELECT * FROM Transactions WHERE transaction_id = $1"

	row := r.db.QueryRow(query, transaction_id)

	var transaction Transactions
	err := row.Scan(&transaction.ID, &transaction.SenderID, &transaction.RecipientID, &transaction.Amount, &transaction.TransactionTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}

	return &transaction, nil
}
