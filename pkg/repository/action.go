package repository

import (
	"database/sql"
	"fmt"
)

type ActionRepository struct {
	db *sql.DB
}

func NewActionRepository(db *sql.DB) *ActionRepository {
	return &ActionRepository{db: db}
}

func (r *ActionRepository) GetUserById(user_id int) (*Users, error) {
	var query string = "SELECT * FROM Users WHERE user_id = $1"

	row := r.db.QueryRow(query, user_id)

	var user Users
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Surname, &user.Password, &user.Coins)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}
