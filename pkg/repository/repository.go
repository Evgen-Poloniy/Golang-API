package repository

import (
	"API/pkg/attribute"
	"database/sql"

	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateUser(user *Users) (uint32, error)
	GetPasswordHashByUsername(username string) (string, error)
}

type Action interface {
	GetUserByID(user_id uint32) (*attribute.ActionField, error)
	GetUserByUsername(username string) (*attribute.ActionField, error)
	GetUserIDByUsername(username string) (uint32, error)
	GetUserByAttributes(attributes map[string]string) (*attribute.ActionField, error)
	GetUserBalance(user_id uint32) (float64, error)
}

type Transaction interface {
	MakeTransaction(senderID uint32, recipientID uint32, amount float64) error
	CreateRecordOfTransaction(transaction *Transactions) (uint32, error)
	GetTransactionByID(transaction_id uint32) (*Transactions, error)
}

type Repository struct {
	Authorization
	Action
	Transaction
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Action:        NewActionRepository(db),
		Transaction:   NewTransactionRepository(db),
	}
}
