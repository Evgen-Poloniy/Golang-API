package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateUser(user *Users) (int, error)
}

type Action interface {
	GetUserById(user_id int) (*Users, error)
}

type Transaction interface {
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
	}
}
