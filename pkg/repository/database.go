package repository

import (
	"database/sql"
	"fmt"
	"log"
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
		log.Println("Database Configuration:")
		log.Printf("Driver:   %s\n", cfg.DriverName)
		log.Printf("DB Name:  %s\n", cfg.DBName)
		log.Printf("Host:     %s\n", cfg.Host)
		log.Printf("Port:     %s\n", cfg.Port)
		log.Printf("User:     %s\n", cfg.Username)
		log.Printf("Password: [REDACTED]\n")
		log.Printf("SSL Mode: %s\n", cfg.SSLMode)
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
