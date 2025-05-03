package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Authorization interface {
	InsertUser(user *Users) error
	CheckUser(user *Users) error
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

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) InsertUser(user *Users, db *sql.DB) error {
	var quary string = "INSERT INTO Users (nikname, name, surname, password, coins) VALUE ($1, $2, $3, $4, $5)"

	if _, err := db.Exec(quary, user.Niсkname, user.Name, user.Surname, user.Password, user.Coins); err != nil {
		return err
	}

	return nil
}
