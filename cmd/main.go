package main

import (
	"API/pkg/handler"
	"API/pkg/repository"
	serverHTTP "API/pkg/server"
	"API/pkg/service"
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := getEnv(envPath); err != nil {
		log.Fatalf("Error of loading .env file: %s", err.Error())
	}

	var db *sql.DB
	var cfg repository.DBConfig
	var err error

	db, err = repository.NewDatabase(&cfg)
	if err != nil {
		log.Fatalf("Database connection error: %s\n", err.Error())
	}
	defer repository.CloseDB(db)

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	hand := handler.NewHendler(serv)

	srv := serverHTTP.NewServer(
		os.Getenv("API_HOST"),
		os.Getenv("API_PORT"),
		maxHeaderBytes,
		hand.Handle(),
		writeTimeout,
		readTimeout,
	)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("Server error: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
