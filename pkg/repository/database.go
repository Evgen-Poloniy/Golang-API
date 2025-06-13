package repository

import (
	"API/pkg/constant"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func NewDatabase(cfg *DBConfig) (*sql.DB, error) {
	cfg.getEnv()

	var mode string = os.Getenv("MODE")
	var db *sql.DB
	var err error

	switch mode {
	case constant.DEBUG:
		db, err = makeConnectionDB(cfg)

		fmt.Print("Database Configuration:\n")
		fmt.Printf("\tDriver:		%s\n", cfg.DriverName)
		fmt.Printf("\tDB Name:  	%s\n", cfg.DBName)
		fmt.Printf("\tHost:     	%s\n", cfg.Host)
		fmt.Printf("\tPort:     	%s\n", cfg.Port)
		fmt.Printf("\tUser:     	%s\n", cfg.Username)
		fmt.Print("\tPassword: 	[REDACTED]\n")
		fmt.Printf("\tSSL Mode: 	%s\n", cfg.SSLMode)

		return db, err

	case constant.DEBUG_WITHOUT_DB:
		fmt.Printf("Database has'nt started, because of chosed mode: \"%s\"\n", mode)
		return nil, nil

	case constant.PRODUCTION:
		db, err = makeConnectionDB(cfg)

		return db, err

	default:
		var errorString string = fmt.Sprintf(
			"uncorrect mode of API. Check .env file and set API mode:\n\t\"%s\"\n\t\"%s\"\n\t\"%s\"\n",
			constant.DEBUG,
			constant.DEBUG_WITHOUT_DB,
			constant.PRODUCTION,
		)
		return nil, errors.New(errorString)
	}

}

func makeConnectionDB(cfg *DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
	)

	delay, err := strconv.Atoi(os.Getenv("API_DELAY_BD_CONNECTION"))
	if err != nil {
		return nil, fmt.Errorf("failed to convert 'string' to 'int': %v", err)
	}

	time.Sleep(time.Duration(delay) * time.Second)

	db, err := sql.Open(cfg.DriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}
