package repository

import "database/sql"

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user *Users) (int, error) {
	var query string = "INSERT INTO Users (username, name, surname, password, coins) VALUES ($1, $2, $3, $4, $5) RETURNING user_id"
	var user_id int

	row := r.db.QueryRow(query, user.Username, user.Name, user.Surname, user.Password, user.Coins)
	if err := row.Scan(&user_id); err != nil {
		return 0, err
	}

	return user_id, nil
}
