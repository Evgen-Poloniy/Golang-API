package repository

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	DriverName string
	DBName     string
	Host       string
	Port       string
	Username   string
	Password   string
	SSLMode    string
}

func NewDatabase(cfg DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
	)

	db, err := sql.Open(cfg.DriverName, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
