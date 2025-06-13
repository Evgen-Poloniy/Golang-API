package repository

import (
	"API/pkg/attribute"
	"database/sql"
	"fmt"
)

type ActionRepository struct {
	db *sql.DB
}

func NewActionRepository(db *sql.DB) *ActionRepository {
	return &ActionRepository{db: db}
}

func (r *ActionRepository) GetUserByID(user_id uint32) (*attribute.ActionField, error) {
	var query string = "SELECT user_id, username, name, surname, coins FROM Users WHERE user_id = $1"
	var user attribute.ActionField

	row := r.db.QueryRow(query, user_id)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Surname, &user.Coins)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *ActionRepository) GetUserByUsername(username string) (*attribute.ActionField, error) {
	var query string = "SELECT user_id, username, name, surname, coins FROM Users WHERE username = $1"
	var user attribute.ActionField

	row := r.db.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Surname, &user.Coins)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *ActionRepository) GetUserIDByUsername(username string) (uint32, error) {
	var query string = "SELECT user_id FROM Users WHERE username = $1"

	var user_id uint32 = 0
	err := r.db.QueryRow(query, username).Scan(&user_id)
	if err != nil {
		return 0, err
	}

	return user_id, nil

}

func (r *ActionRepository) GetUserByAttributes(attributes map[string]string) (*attribute.ActionField, error) {
	var query string = "SELECT user_id, username, name, surname, coins FROM Users WHERE"

	var i int = len(attributes)
	for key, value := range attributes {
		query += fmt.Sprintf(" %s = '%s'", key, value)
		i--
		if i != 0 {
			query += " AND"
		}
	}

	row := r.db.QueryRow(query)

	var user attribute.ActionField
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Surname, &user.Coins)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *ActionRepository) GetUserBalance(user_id uint32) (float64, error) {
	var query string = "SELECT coins FROM Users WHERE user_id = $1"
	var balance float64

	if err := r.db.QueryRow(query, user_id).Scan(&balance); err != nil {
		return 0, err
	}

	return balance, nil
}
