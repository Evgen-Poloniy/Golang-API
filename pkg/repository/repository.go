package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateUser(user *Users) (int, error)
}

type Action interface {
	GetUserByID(user_id int) (*Users, error)
	GetUserByUsername(username string) (*Users, error)
	GetUserIDByUsername(username string) (uint32, error)
	GetUserByAttributes(attributes map[string]string) (*Users, error)
	GetUserBalance(user_id uint32) (float64, error)
}

type Transaction interface {
	MakeTransaction(senderID uint32, recipientID uint32, amount float64) error
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
