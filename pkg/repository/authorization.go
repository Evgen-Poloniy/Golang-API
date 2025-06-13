package repository

import (
	"database/sql"
	"fmt"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user *Users) (uint32, error) {
	var query string = "INSERT INTO Users (username, name, surname, password_hash, coins) VALUES ($1, $2, $3, $4, $5) RETURNING user_id"
	var user_id uint32

	row := r.db.QueryRow(query, user.Username, user.Name, user.Surname, user.PasswordHash, user.Coins)
	if err := row.Scan(&user_id); err != nil {
		return 0, err
	}

	return user_id, nil
}

func (r *AuthRepository) GetPasswordHashByUsername(username string) (string, error) {
	var query string = "SELECT password_hash FROM Users WHERE username = $1"
	var passwordHash string

	row := r.db.QueryRow(query, username)

	if err := row.Scan(&passwordHash); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}

	return passwordHash, nil
}
