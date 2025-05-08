package main

import (
	"API/pkg/handler"
	"API/pkg/repository"
	serverHTTP "API/pkg/server"
	"API/pkg/service"
	"log"
	"os"
	"time"
	//"github.com/joho/godotenv"
)

func main() {
	/*
		err := godotenv.Load("./config/.env")
		if err != nil {
			log.Fatalf("Error of loading .env file: %s", err.Error())
		}
	*/

	time.Sleep(20 * time.Second)

	var cfg repository.DBConfig
	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Database connection error: %s\n", err.Error())
	}
	defer db.Close()

	repos := repository.NewRepository()
	serv := service.NewService(repos)
	hand := handler.NewHendler(serv)

	srv := serverHTTP.NewServer(os.Getenv("API_HOST"), os.Getenv("API_PORT"), 1<<20, hand.Handle(), 10*time.Second, 10*time.Second)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %s\n", err.Error())
	}
}
