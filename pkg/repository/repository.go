package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateUser(user *Users) (int, error)
}

type Messenger interface {
}

type Transaction interface {
}

type Repository struct {
	Authorization
	Messenger
	Transaction
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
