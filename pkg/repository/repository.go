package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Authorization interface {
	InsertUser(user *Users) error
	GetPasswordByUsername(nikname *string) (string, error)
}

type Messenger interface {
	GetUserIDByUsername(username *string) (uint64, error)
	//GetOrCreateChat(senderID uint64, recipientID uint64)
	InsertMessage()
}

type Transaction interface {
}

type Repository struct {
	db *sql.DB
	Authorization
	Messenger
	Transaction
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) InsertUser(user *Users) error {
	var query string = "INSERT INTO Users (nikname, name, surname, password, coins) VALUE ($1, $2, $3, $4, $5)"

	if _, err := r.db.Exec(query, user.Username, user.Name, user.Surname, user.Password, user.Coins); err != nil {
		return err
	}

	return nil
}

func (r *Repository) TestInsertUser(user *Users) error {
	fmt.Println(*user)

	return nil
}

func (r *Repository) GetPasswordByNikname(nikname *string) (string, error) {
	query := "SELECT password FROM users WHERE username = $1"

	var password string

	if err := r.db.QueryRow(query, *nikname).Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}

func (r *Repository) GetUserIDByUsername(username *string) (uint64, error) {
	query := "SELECT user_id FROM users WHERE username = $1"

	var ID uint64
	if err := r.db.QueryRow(query, *username).Scan(&ID); err != nil {
		return 0, nil
	}

	return ID, nil
}
