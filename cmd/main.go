package main

import (
	//"database/sql"
	serverHTTP "API/pkg/server"
	"log"
	"time"
)

/*
var db *sql.DB
var dbConfig DBConfig

func initDB() {
	var err error
	db, err = sql.Open("mysql", dbConfig.Username+":"+dbConfig.Password+"@tcp("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Database)
	if err != nil {
		log.Fatalf()
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
*/

func main() {
	srv := serverHTTP.NewServer("localhost", "3505", 1<<20, 10*time.Second, 10*time.Second)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
