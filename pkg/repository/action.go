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

func (r *ActionRepository) GetUserByID(user_id int) (*Users, error) {
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

func (r *ActionRepository) GetUserByUsername(username string) (*Users, error) {
	var query string = "SELECT * FROM Users WHERE username = $1"

	row := r.db.QueryRow(query, username)

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

func (r *ActionRepository) GetUserIDByUsername(username string) (uint32, error) {
	var query string = "SELECT user_id FROM Users WHERE username = $1"

	var user_id uint32 = 0
	err := r.db.QueryRow(query, username).Scan(&user_id)
	if err != nil {
		return 0, err
	}

	return user_id, nil

}

func (r *ActionRepository) GetUserByAttributes(attributes map[string]string) (*Users, error) {
	var query string = "SELECT * FROM Users WHERE"

	var i int = len(attributes)
	for key, value := range attributes {
		query += fmt.Sprintf(" %s = '%s'", key, value)
		i--
		if i != 0 {
			query += " AND"
		}
	}

	fmt.Printf("%s\n", query)

	row := r.db.QueryRow(query)

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

func (r *ActionRepository) GetUserBalance(user_id uint32) (float64, error) {
	var query string = "SELECT coins FROM Users WHERE user_id = $1"
	var balance float64

	if err := r.db.QueryRow(query, user_id).Scan(&balance); err != nil {
		return 0, err
	}

	return balance, nil
}
