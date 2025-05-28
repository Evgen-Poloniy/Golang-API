package main

import (
	"API/pkg/handler"
	"API/pkg/repository"
	serverHTTP "API/pkg/server"
	"API/pkg/service"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error of loading .env file: %s", err.Error())
	}

	var cfg repository.DBConfig
	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Database connection error: %s\n", err.Error())
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	hand := handler.NewHendler(serv)

	srv := serverHTTP.NewServer(os.Getenv("API_HOST"), os.Getenv("API_PORT"), 1<<20, hand.Handle(), 10*time.Second, 10*time.Second)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("Server error: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Gracefully shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
