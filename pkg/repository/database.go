package repository

import (
	"database/sql"
	"fmt"
	"os"
	"time"
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

func (cfg *DBConfig) getEnv() {
	cfg.DriverName = os.Getenv("DB_DRIVER_NAME")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.Username = os.Getenv("DB_USER")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.SSLMode = os.Getenv("SSL_MODE")
}

func (cfg *DBConfig) printConfig() {
	var mode string = os.Getenv("MODE")

	switch mode {
	case "debug":
		fmt.Printf("Database Configuration:\n")
		fmt.Printf("\tDriver:   %s\n", cfg.DriverName)
		fmt.Printf("\tDB Name:  %s\n", cfg.DBName)
		fmt.Printf("\tHost:     %s\n", cfg.Host)
		fmt.Printf("\tPort:     %s\n", cfg.Port)
		fmt.Printf("\tUser:     %s\n", cfg.Username)
		fmt.Printf("\tPassword: [REDACTED]\n")
		fmt.Printf("\tSSL Mode: %s\n", cfg.SSLMode)
	}
}

func NewDatabase(cfg DBConfig) (*sql.DB, error) {
	cfg.getEnv()

	cfg.printConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
	)

	var db *sql.DB
	var err error

	for i := 0; i < 3; i++ {
		db, err = sql.Open(cfg.DriverName, dsn)
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			err = nil
			break
		}
	}

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
